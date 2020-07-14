// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: gosec
package provider

import (
	"context"
	"crypto/sha1"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudformation/provider/pkg/schema"
	"github.com/pulumi/pulumi-cloudformation/provider/pkg/update"
	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	pulumirpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
)

type cancellationContext struct {
	context context.Context
	cancel  context.CancelFunc
}

func makeCancellationContext() *cancellationContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &cancellationContext{
		context: ctx,
		cancel:  cancel,
	}
}

type cfnProvider struct {
	host     *provider.HostClient
	canceler *cancellationContext

	configured   bool
	version      string
	stackName    string
	accountID    string
	region       string
	schema       schema.CloudFormationSchema
	pulumiSchema []byte

	cfn    *cloudformation.CloudFormation
	ec2    *ec2.EC2
	ssm    *ssm.SSM
	update *update.StackUpdate
}

var _ pulumirpc.ResourceProviderServer = (*cfnProvider)(nil)

func newCloudformationProvider(host *provider.HostClient, version string, pulumiSchema []byte) (pulumirpc.ResourceProviderServer, error) {
	return &cfnProvider{
		host:         host,
		canceler:     makeCancellationContext(),
		version:      version,
		pulumiSchema: pulumiSchema,
	}, nil
}

func (p *cfnProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	return &pulumirpc.GetSchemaResponse{Schema: string(p.pulumiSchema)}, nil
}

// CheckConfig validates the configuration for this provider.
func (p *cfnProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        "CFN.CheckConfig.news",
		KeepUnknowns: true,
	})
	if err != nil {
		return nil, err
	}

	stack, ok := news["stack"]
	if !ok {
		failures := []*pulumirpc.CheckFailure{{Property: "stack", Reason: "missing required property 'stack'"}}
		return &pulumirpc.CheckResponse{Failures: failures}, nil
	}
	if !stack.IsString() && !stack.IsComputed() {
		failures := []*pulumirpc.CheckFailure{{Property: "stack", Reason: "'stack' must be a string"}}
		return &pulumirpc.CheckResponse{Failures: failures}, nil
	}

	region, ok := news["region"]
	if !ok {
		failures := []*pulumirpc.CheckFailure{{Property: "region", Reason: "missing required property 'region'"}}
		return &pulumirpc.CheckResponse{Failures: failures}, nil
	}
	if !region.IsString() && !region.IsComputed() {
		failures := []*pulumirpc.CheckFailure{{Property: "region", Reason: "'region' must be a string"}}
		return &pulumirpc.CheckResponse{Failures: failures}, nil
	}

	for k := range news {
		if k != "stack" && k != "region" && k != "version" {
			failures := []*pulumirpc.CheckFailure{{Property: string(k), Reason: fmt.Sprintf("unknown property '%v'", k)}}
			return &pulumirpc.CheckResponse{Failures: failures}, nil
		}
	}

	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *cfnProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.DiffConfig(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.olds", label),
		KeepUnknowns: true,
	})
	if err != nil {
		return nil, err
	}
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.news", label),
		KeepUnknowns: true,
		RejectAssets: true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diffConfig failed because of malformed resource inputs")
	}

	diff := olds.Diff(news)
	if diff == nil {
		return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_NONE}, nil
	}

	var diffs, replaces []string
	for k := range diff.Keys() {
		diffs, replaces = append(diffs, string(k)), append(replaces, string(k))
	}

	return &pulumirpc.DiffResponse{
		Changes:  pulumirpc.DiffResponse_DIFF_SOME,
		Diffs:    diffs,
		Replaces: replaces,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *cfnProvider) Configure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	region, ok := req.Variables["cloudformation:config:region"]
	if !ok {
		return nil, errors.New("missing required property 'region'")
	}
	stackName, ok := req.Variables["cloudformation:config:stack"]
	if !ok {
		return nil, errors.New("missing required property 'stack'")
	}
	p.region, p.stackName = region, stackName

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(p.region),
	})
	if err != nil {
		return nil, errors.Errorf("could not create AWS session: %v", err)
	}
	p.cfn, p.ec2, p.ssm = cloudformation.New(sess), ec2.New(sess), ssm.New(sess)

	callerIdentityResp, err := sts.New(sess).GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, errors.Errorf("could not get AWS account ID: %v", err)
	}
	p.accountID = aws.StringValue(callerIdentityResp.Account)

	schema, err := getCloudFormationSchema(p.canceler.context, region)
	if err != nil {
		return nil, errors.Errorf("could not fetch CloudFormation schema: %v", err)
	}
	p.schema = *schema

	update, err := update.StartStackUpdate(p.canceler.context, sess, p.stackName)
	if err != nil {
		return nil, errors.Errorf("could not start stack updater: %v", err)
	}
	p.update = update

	p.configured = true

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

var functions = map[string]func(*cfnProvider, context.Context, resource.PropertyMap) (resource.PropertyMap, error){
	"cloudformation:index:getAccountId":          (*cfnProvider).getAccountID,
	"cloudformation:index:getAzs":                (*cfnProvider).getAZs,
	"cloudformation:index:getPartition":          (*cfnProvider).getPartition,
	"cloudformation:index:getRegion":             (*cfnProvider).getRegion,
	"cloudformation:index:getStackId":            (*cfnProvider).getStackID,
	"cloudformation:index:getStackName":          (*cfnProvider).getStackName,
	"cloudformation:index:getUrlSuffix":          (*cfnProvider).getURLSuffix,
	"cloudformation:index:cidr":                  (*cfnProvider).cidr,
	"cloudformation:index:getSsmParameterString": (*cfnProvider).getSSMParameterString,
	"cloudformation:index:getSsmParameterList":   (*cfnProvider).getSSMParameterList,
	"cloudformation:index:importValue":           (*cfnProvider).importValue,
}

// Invoke dynamically executes a built-in function in the provider.
func (p *cfnProvider) Invoke(ctx context.Context,
	req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {

	// Unmarshal arguments.
	tok := req.GetTok()

	inputs, err := plugin.UnmarshalProperties(req.GetArgs(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("CFN.Invoke(%s).inputs", tok),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	// Process Invoke call.
	fn, ok := functions[tok]
	if !ok {
		return nil, fmt.Errorf("unknown function '%s'", tok)
	}
	result, err := fn(p, ctx, inputs)
	if err != nil {
		return nil, err
	}

	res, err := plugin.MarshalProperties(result, plugin.MarshalOptions{
		Label:        fmt.Sprintf("CFN.Invoke(%s).outputs", tok),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}
	return &pulumirpc.InvokeResponse{Return: res}, nil
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *cfnProvider) StreamInvoke(
	req *pulumirpc.InvokeRequest, server pulumirpc.ResourceProvider_StreamInvokeServer) error {

	tok := req.GetTok()
	return errors.Errorf("unrecognized function (StreamInvoke): %s", tok)
}

// filterNullValues removes nested null values from the given property value. If all nested values in the property
// value are null, the property value itself is considered null. This allows for easier integration with CloudFormation
// templates, which use `AWS::NoValue` to remove values from lists, maps, etc. We use `null` to the same effect.
func filterNullValues(v resource.PropertyValue) resource.PropertyValue {
	switch {
	case v.IsArray():
		if len(v.ArrayValue()) == 0 {
			return v
		}

		var result []resource.PropertyValue
		for _, e := range v.ArrayValue() {
			e = filterNullValues(e)
			if !e.IsNull() {
				result = append(result, e)
			}
		}
		return resource.NewArrayProperty(result)
	case v.IsObject():
		return resource.NewObjectProperty(filterNullProperties(v.ObjectValue()))
	default:
		return v
	}
}

// filterNullValues removes nested null values from the given property map. If all nested values in the property
// map are null, the property map itself is considered null. This allows for easier integration with CloudFormation
// templates, which use `AWS::NoValue` to remove values from lists, maps, etc. We use `null` to the same effect.
func filterNullProperties(m resource.PropertyMap) resource.PropertyMap {
	if len(m) == 0 {
		return m
	}

	result := resource.PropertyMap{}
	for k, v := range m {
		e := filterNullValues(v)
		if !e.IsNull() {
			result[k] = e
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func validateInputs(sch schema.CloudFormationSchema, resourceType string, inputs resource.PropertyMap) ([]*pulumirpc.CheckFailure, error) {
	var checkFailures []*pulumirpc.CheckFailure

	if logicalID, ok := inputs["logicalId"]; ok {
		if !logicalID.IsString() && !logicalID.IsComputed() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "logicalId", Reason: "logicalId must be a string"})
		}
	}

	if metadataV, ok := inputs["metadata"]; ok {
		if !metadataV.IsObject() && !metadataV.IsComputed() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "metadata", Reason: "metadata must be an object"})
		}
	}
	if creationPolicyV, ok := inputs["creationPolicy"]; ok {
		if !creationPolicyV.IsObject() && !creationPolicyV.IsComputed() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "creationPolicy", Reason: "creationPolicy must be an object"})
		}
	}
	if deletionPolicyV, ok := inputs["deletionPolicy"]; ok {
		if !deletionPolicyV.IsString() && !deletionPolicyV.IsComputed() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "deletionPolicy", Reason: "deletionPolicy must be a string"})
		}
	}
	if updatePolicyV, ok := inputs["updatePolicy"]; ok {
		if !updatePolicyV.IsObject() && !updatePolicyV.IsComputed() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "updatePolicy", Reason: "updatePolicy must be an object"})
		}
	}
	if updateReplacePolicyV, ok := inputs["updateReplacePolicy"]; ok {
		if !updateReplacePolicyV.IsString() && !updateReplacePolicyV.IsComputed() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "updateReplacePolicy", Reason: "updateReplacePolicy must be a string"})
		}
	}

	if propertiesV, ok := inputs["properties"]; ok {
		switch {
		case propertiesV.IsObject():
			failures, err := schema.ValidateResource(sch, resourceType, propertiesV.ObjectValue())
			if err != nil {
				return nil, err
			}
			for _, f := range failures {
				checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: f.Path, Reason: f.Reason})
			}
		case propertiesV.IsComputed():
			// Nothing we can do here.
		default:
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "properties", Reason: "properties must be an object"})
		}
	}

	return checkFailures, nil
}

func (p *cfnProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Check(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	_, resourceType, err := resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}

	// Parse inputs.
	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	// Filter null properties from the inputs.
	newInputs = filterNullProperties(newInputs)

	failures, err := validateInputs(p.schema, resourceType, newInputs)
	if err != nil {
		return nil, err
	}

	if len(failures) == 0 {
		inputs, err := plugin.MarshalProperties(newInputs, plugin.MarshalOptions{
			Label:        fmt.Sprintf("%s.inputs", label),
			KeepUnknowns: true,
			KeepSecrets:  true,
		})
		if err != nil {
			return nil, err
		}
		return &pulumirpc.CheckResponse{Inputs: inputs}, nil
	}
	return &pulumirpc.CheckResponse{Failures: failures}, nil
}

func getPropertiesDiff(inputDiff *resource.ObjectDiff) (resource.ValueDiff, bool) {
	if add, ok := inputDiff.Adds["properties"]; ok {
		return resource.ValueDiff{New: add}, true
	}
	if del, ok := inputDiff.Deletes["properties"]; ok {
		return resource.ValueDiff{Old: del}, true
	}
	diff, ok := inputDiff.Updates["properties"]
	return diff, ok
}

func (p *cfnProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Diff(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	_, resourceType, err := resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create failed because malformed resource inputs")
	}
	delete(oldState, "attributes")
	oldInputs := oldState

	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create failed because malformed resource inputs")
	}

	diff := oldInputs.Diff(newInputs)
	if diff == nil {
		return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_NONE}, nil
	}

	var replaces []string
	if propertiesDiff, ok := getPropertiesDiff(diff); ok {
		rs, err := schema.GetResourceReplaces(p.schema, resourceType, propertiesDiff)
		if err != nil {
			return nil, err
		}
		replaces = rs
	}

	return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_SOME, Replaces: replaces}, nil
}

func (p *cfnProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Create(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	// Parse inputs
	newInputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create failed because malformed resource inputs")
	}

	inputs := newInputs.MapRepl(nil, mapReplStripSecrets)

	logicalID, resourceType, err := resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}
	if inputID, ok := inputs["logicalId"]; ok {
		logicalID = inputID.(string)
	}

	properties, _ := inputs["properties"].(map[string]interface{})

	var options []update.ResourceOption
	if metadataV, ok := inputs["metadata"]; ok {
		metadata, _ := metadataV.(map[string]interface{})
		options = append(options, update.Metadata(metadata))
	}
	if creationPolicyV, ok := inputs["creationPolicy"]; ok {
		creationPolicy, _ := creationPolicyV.(map[string]interface{})
		options = append(options, update.CreationPolicy(creationPolicy))
	}
	if deletionPolicyV, ok := inputs["deletionPolicy"]; ok {
		deletionPolicy, _ := deletionPolicyV.(string)
		options = append(options, update.DeletionPolicy(deletionPolicy))
	}
	if updatePolicyV, ok := inputs["updatePolicy"]; ok {
		updatePolicy, _ := updatePolicyV.(map[string]interface{})
		options = append(options, update.UpdatePolicy(updatePolicy))
	}
	if updateReplacePolicyV, ok := inputs["updateReplacePolicy"]; ok {
		updateReplacePolicy, _ := updateReplacePolicyV.(string)
		options = append(options, update.UpdateReplacePolicy(updateReplacePolicy))
	}

	data, err := p.update.CreateResource(p.canceler.context, logicalID, resourceType, properties, options...)
	if err != nil {
		return nil, err
	}

	newInputs["attributes"] = resource.NewPropertyValue(cleanTopLevelAttributeNames(data.Attributes))

	state, err := plugin.MarshalProperties(newInputs, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         data.PhysicalResourceID,
		Properties: state,
	}, nil
}

func (p *cfnProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Read(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	physicalResourceID := req.GetId()
	logicalID, resourceType, err := resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}

	// Parse old state
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	// If the resource has an explicit logical ID in the state, use it.
	if stateID, ok := oldState["logicalId"]; ok {
		logicalID = stateID.StringValue()
	}

	data, err := p.update.ReadResource(p.canceler.context, logicalID, physicalResourceID, resourceType)
	if err != nil {
		return nil, err
	}

	// This resource may have disappeared.
	if data.PhysicalResourceID == "" {
		return &pulumirpc.ReadResponse{}, nil
	}

	stateMap := map[string]interface{}{}
	if _, ok := oldState["logicalId"]; ok {
		stateMap["logicalId"] = logicalID
	}
	if data.Metadata != nil {
		stateMap["metadata"] = data.Metadata
	}
	if data.Properties != nil {
		stateMap["properties"] = data.Properties
	}
	if data.CreationPolicy != nil {
		stateMap["creationPolicy"] = data.CreationPolicy
	}
	stateMap["deletionPolicy"] = data.DeletionPolicy
	if data.UpdatePolicy != nil {
		stateMap["updatePolicy"] = data.UpdatePolicy
	}
	stateMap["updateReplacePolicy"] = data.UpdateReplacePolicy
	stateMap["attributes"] = cleanTopLevelAttributeNames(data.Attributes)

	// TODO: annotate secrets
	stateProps := resource.NewPropertyMapFromMap(stateMap)
	state, err := plugin.MarshalProperties(stateProps, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	delete(stateProps, "attributes")
	inputs, err := plugin.MarshalProperties(stateProps, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	return &pulumirpc.ReadResponse{
		Id:         data.PhysicalResourceID,
		Inputs:     inputs,
		Properties: state,
	}, nil
}

func (p *cfnProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Update(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	// Parse inputs
	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create failed because malformed resource inputs")
	}

	inputs := newInputs.MapRepl(nil, mapReplStripSecrets)

	physicalResourceID := req.GetId()
	logicalID, resourceType, err := resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}
	if inputID, ok := inputs["logicalId"]; ok {
		logicalID = inputID.(string)
	}

	properties, _ := inputs["properties"].(map[string]interface{})

	var options []update.ResourceOption
	if metadataV, ok := inputs["metadata"]; ok {
		metadata, _ := metadataV.(map[string]interface{})
		options = append(options, update.Metadata(metadata))
	}
	if creationPolicyV, ok := inputs["creationPolicy"]; ok {
		creationPolicy, _ := creationPolicyV.(map[string]interface{})
		options = append(options, update.CreationPolicy(creationPolicy))
	}
	if deletionPolicyV, ok := inputs["deletionPolicy"]; ok {
		deletionPolicy, _ := deletionPolicyV.(string)
		options = append(options, update.DeletionPolicy(deletionPolicy))
	}
	if updatePolicyV, ok := inputs["updatePolicy"]; ok {
		updatePolicy, _ := updatePolicyV.(map[string]interface{})
		options = append(options, update.UpdatePolicy(updatePolicy))
	}
	if updateReplacePolicyV, ok := inputs["updateReplacePolicy"]; ok {
		updateReplacePolicy, _ := updateReplacePolicyV.(string)
		options = append(options, update.UpdateReplacePolicy(updateReplacePolicy))
	}

	data, err := p.update.UpdateResource(p.canceler.context, logicalID, physicalResourceID, resourceType, properties, options...)
	if err != nil {
		return nil, err
	}

	newInputs["attributes"] = resource.NewPropertyValue(cleanTopLevelAttributeNames(data.Attributes))

	state, err := plugin.MarshalProperties(newInputs, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{Properties: state}, nil
}

func (p *cfnProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Delete(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	// Parse state
	state, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	logicalID, _, err := resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}
	if stateID, ok := state["logicalId"]; ok {
		logicalID = stateID.StringValue()
	}

	physicalResourceID := req.GetId()
	if err := p.update.DeleteResource(p.canceler.context, logicalID, physicalResourceID); err != nil {
		return nil, err
	}

	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *cfnProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*pulumirpc.PluginInfo, error) {
	return &pulumirpc.PluginInfo{
		Version: p.version,
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (p *cfnProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	p.canceler.cancel()
	return &pbempty.Empty{}, nil
}

func resourceInfoFromURN(urn resource.URN) (string, string, error) {
	// Transform the URN's type into its CloudFormation type.
	ns, service := "AWS", ""

	module := string(urn.Type().Module().Name())

	nsService := strings.Split(module, "/")
	switch len(nsService) {
	case 1:
		service = nsService[0]
	case 2:
		ns, service = nsService[1], nsService[2]
	default:
		return "", "", errors.Errorf("malformed module in type %v", urn.Type())
	}

	if service == "Configuration" {
		service = "Config"
	}

	typ := strings.Join([]string{ns, service, string(urn.Type().Name())}, "::")

	// Now build the resource's logical ID from the first 223 characters of its name and the base32-encoded SHA-1 hash of
	// its URN. The latter is necessary to accurately map the Pulumi URN identity model onto the CF logical ID identity
	// model. We use the aforementioned lengths because the length of a CF logical ID is limited to 255 characters.

	sum := sha1.Sum([]byte(urn))
	hash := base32.StdEncoding.EncodeToString(sum[:])
	contract.Assert(len(hash) == 32)

	const maxNameLen = 255 - 32

	name := string(urn.Name())
	if len(name) >= maxNameLen {
		name = name[:maxNameLen]
	}

	return name + hash, typ, nil
}

func cleanTopLevelAttributeNames(attrs map[string]interface{}) map[string]interface{} {
	for k, v := range attrs {
		if strings.Index(k, ".") != -1 {
			delete(attrs, k)
			attrs[strings.Replace(k, ".", "", -1)] = v
		}
	}
	return attrs
}

func mapReplStripSecrets(v resource.PropertyValue) (interface{}, bool) {
	if v.IsSecret() {
		return v.SecretValue().Element.MapRepl(nil, mapReplStripSecrets), true
	}

	return nil, false
}

func getCloudFormationSchema(ctx context.Context, region string) (*schema.CloudFormationSchema, error) {
	urlFmt := `https://cfn-resource-specifications-%[1]s-prod.s3.%[1]s.amazonaws.com/latest/gzip/CloudFormationResourceSpecification.json`
	url := fmt.Sprintf(urlFmt, region)

	// nolint: gosec
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer contract.IgnoreClose(resp.Body)

	var sch schema.CloudFormationSchema
	if err := json.NewDecoder(resp.Body).Decode(&sch); err != nil {
		return nil, err
	}

	return &sch, nil
}
