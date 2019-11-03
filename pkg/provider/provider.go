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

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudformation/pkg/schema"
	"github.com/pulumi/pulumi-cloudformation/pkg/update"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/resource/plugin"
	"github.com/pulumi/pulumi/pkg/resource/provider"
	"github.com/pulumi/pulumi/pkg/util/contract"
	pulumirpc "github.com/pulumi/pulumi/sdk/proto/go"
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

	configured bool
	version    string
	stackName  string
	region     string
	schema     schema.CloudFormationSchema

	update *update.StackUpdate
}

var _ pulumirpc.ResourceProviderServer = (*cfnProvider)(nil)

func newCloudformationProvider(host *provider.HostClient, version string) (pulumirpc.ResourceProviderServer, error) {
	return &cfnProvider{
		host:     host,
		canceler: makeCancellationContext(),
		version:  version,
	}, nil
}

// CheckConfig validates the configuration for this provider.
func (p *cfnProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        "CFN.CheckConfig.news",
		KeepUnknowns: true,
		SkipNulls:    true,
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
		SkipNulls:    true,
	})
	if err != nil {
		return nil, err
	}
	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.news", label),
		KeepUnknowns: true,
		SkipNulls:    true,
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

	// TODO(pdg): get the CFN schema for the region
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

// Invoke dynamically executes a built-in function in the provider.
func (p *cfnProvider) Invoke(ctx context.Context,
	req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {

	// Unmarshal arguments.
	tok := req.GetTok()

	// Process Invoke call.
	switch tok {
	case "cloudformation:index:getStackId":
		result := resource.NewPropertyValue(map[string]interface{}{
			"stackId": p.update.StackID(),
		})
		res, err := plugin.MarshalProperties(result.ObjectValue(), plugin.MarshalOptions{
			Label:        fmt.Sprintf("CFN.Invoke.result"),
			KeepUnknowns: true,
			SkipNulls:    true,
			KeepSecrets:  true,
		})
		if err != nil {
			return nil, err
		}
		return &pulumirpc.InvokeResponse{Return: res}, nil

	default:
		return nil, fmt.Errorf("Unknown Invoke type '%s'", tok)
	}
}

func validateInputs(sch schema.CloudFormationSchema, resourceType string, inputs resource.PropertyMap) ([]*pulumirpc.CheckFailure, error) {
	var checkFailures []*pulumirpc.CheckFailure

	if metadataV, ok := inputs["metadata"]; ok {
		if !metadataV.IsObject() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "metadata", Reason: "metadata must be an object"})
		}
	}
	if creationPolicyV, ok := inputs["creationPolicy"]; ok {
		if !creationPolicyV.IsObject() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "creationPolicy", Reason: "creationPolicy must be an object"})
		}
	}
	if deletionPolicyV, ok := inputs["deletionPolicy"]; ok {
		if !deletionPolicyV.IsObject() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "deletionPolicy", Reason: "deletionPolicy must be an object"})
		}
	}
	if updatePolicyV, ok := inputs["updatePolicy"]; ok {
		if !updatePolicyV.IsObject() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "updatePolicy", Reason: "updatePolicy must be an object"})
		}
	}
	if updateReplacePolicyV, ok := inputs["updateReplacePolicy"]; ok {
		if !updateReplacePolicyV.IsObject() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "updateReplacePolicy", Reason: "updateReplacePolicy must be an object"})
		}
	}

	var properties resource.PropertyMap
	if propertiesV, ok := inputs["properties"]; ok {
		if !propertiesV.IsObject() {
			checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: "properties", Reason: "properties must be an object"})
		}
		properties = propertiesV.ObjectValue()
	}

	failures, err := schema.ValidateResource(sch, resourceType, properties)
	if err != nil {
		return nil, err
	}
	for _, f := range failures {
		checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: f.Path, Reason: f.Reason})
	}
	return checkFailures, nil
}

func (p *cfnProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("CFN.Check(%s)", urn)
	glog.V(9).Infof("%s executing", label)

	resourceType, err := resourceTypeFromURN(urn)
	if err != nil {
		return nil, err
	}

	// Parse inputs.
	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		SkipNulls:    true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	failures, err := validateInputs(p.schema, resourceType, newInputs)
	if err != nil {
		return nil, err
	}

	if len(failures) == 0 {
		return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
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

	resourceType, err := resourceTypeFromURN(urn)
	if err != nil {
		return nil, err
	}

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		SkipNulls:    true,
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
		SkipNulls:    true,
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
		SkipNulls:    true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create failed because malformed resource inputs")
	}

	inputs := newInputs.MapRepl(nil, mapReplStripSecrets)

	logicalID := string(urn.Name())
	resourceType, err := resourceTypeFromURN(urn)
	if err != nil {
		return nil, err
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
		deletionPolicy, _ := deletionPolicyV.(map[string]interface{})
		options = append(options, update.DeletionPolicy(deletionPolicy))
	}
	if updatePolicyV, ok := inputs["updatePolicy"]; ok {
		updatePolicy, _ := updatePolicyV.(map[string]interface{})
		options = append(options, update.UpdatePolicy(updatePolicy))
	}
	if updateReplacePolicyV, ok := inputs["updateReplacePolicy"]; ok {
		updateReplacePolicy, _ := updateReplacePolicyV.(map[string]interface{})
		options = append(options, update.UpdateReplacePolicy(updateReplacePolicy))
	}

	data, err := p.update.CreateResource(p.canceler.context, logicalID, resourceType, properties, options...)
	if err != nil {
		return nil, err
	}

	newInputs["attributes"] = resource.NewPropertyValue(data.Attributes)

	state, err := plugin.MarshalProperties(newInputs, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		SkipNulls:    true,
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

	logicalID := string(urn.Name())
	physicalResourceID := req.GetId()
	resourceType, err := resourceTypeFromURN(urn)
	if err != nil {
		return nil, err
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
	if data.Metadata != nil {
		stateMap["metadata"] = data.Metadata
	}
	if data.Properties != nil {
		stateMap["properties"] = data.Properties
	}
	if data.CreationPolicy != nil {
		stateMap["creationPolicy"] = data.CreationPolicy
	}
	if data.DeletionPolicy != nil {
		stateMap["deletionPolicy"] = data.DeletionPolicy
	}
	if data.UpdatePolicy != nil {
		stateMap["updatePolicy"] = data.UpdatePolicy
	}
	if data.UpdateReplacePolicy != nil {
		stateMap["updateReplacePolicy"] = data.UpdateReplacePolicy
	}
	stateMap["attributes"] = data.Attributes

	// TODO: annotate secrets
	stateProps := resource.NewPropertyMapFromMap(stateMap)
	state, err := plugin.MarshalProperties(stateProps, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	delete(stateProps, "attributes")
	inputs, err := plugin.MarshalProperties(stateProps, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		SkipNulls:    true,
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
		SkipNulls:    true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create failed because malformed resource inputs")
	}

	inputs := newInputs.MapRepl(nil, mapReplStripSecrets)

	logicalID := string(urn.Name())
	physicalResourceID := req.GetId()
	resourceType, err := resourceTypeFromURN(urn)
	if err != nil {
		return nil, err
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
		deletionPolicy, _ := deletionPolicyV.(map[string]interface{})
		options = append(options, update.DeletionPolicy(deletionPolicy))
	}
	if updatePolicyV, ok := inputs["updatePolicy"]; ok {
		updatePolicy, _ := updatePolicyV.(map[string]interface{})
		options = append(options, update.UpdatePolicy(updatePolicy))
	}
	if updateReplacePolicyV, ok := inputs["updateReplacePolicy"]; ok {
		updateReplacePolicy, _ := updateReplacePolicyV.(map[string]interface{})
		options = append(options, update.UpdateReplacePolicy(updateReplacePolicy))
	}

	data, err := p.update.UpdateResource(p.canceler.context, logicalID, physicalResourceID, resourceType, properties, options...)
	if err != nil {
		return nil, err
	}

	newInputs["attributes"] = resource.NewPropertyValue(data.Attributes)

	state, err := plugin.MarshalProperties(newInputs, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.state", label),
		KeepUnknowns: true,
		SkipNulls:    true,
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

	logicalID := string(urn.Name())
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

func resourceTypeFromURN(urn resource.URN) (string, error) {
	ns, service := "AWS", ""

	nsService := strings.Split(string(urn.Type().Module().Name()), "/")
	switch len(nsService) {
	case 1:
		service = nsService[0]
	case 2:
		ns, service = nsService[1], nsService[2]
	default:
		return "", errors.Errorf("malformed module in type %v", urn.Type())
	}

	return strings.Join([]string{ns, service, string(urn.Type().Name())}, "::"), nil
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
