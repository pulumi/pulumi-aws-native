// Copyright 2016-2021, Pulumi Corporation.
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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	oldaws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	pbstruct "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/uuid"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/update"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	name     string
	canceler *cancellationContext

	configured   bool
	version      string
	accountID    string
	region       string
	schema       schema.CloudFormationSchema
	cfTypes      map[string]string
	pulumiSchema []byte

	cfn *cloudformation.Client
	ec2 *ec2.Client
	ssm *ssm.Client

	cfg aws.Config
	update *update.StackUpdate
	supportedResourceTypes codegen.StringSet
}

var _ pulumirpc.ResourceProviderServer = (*cfnProvider)(nil)

func newAwsNativeProvider(host *provider.HostClient, name, version string, pulumiSchema []byte) (pulumirpc.ResourceProviderServer, error) {
	return &cfnProvider{
		host:         host,
		canceler:     makeCancellationContext(),
		name:         name,
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
		Label:        fmt.Sprintf("%s.CheckConfig.news", p.name),
		KeepUnknowns: true,
	})
	if err != nil {
		return nil, err
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
		if k != "region" && k != "version" {
			failures := []*pulumirpc.CheckFailure{{Property: string(k), Reason: fmt.Sprintf("unknown property '%v'", k)}}
			return &pulumirpc.CheckResponse{Failures: failures}, nil
		}
	}

	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *cfnProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.DiffConfig(%s)", p.name, urn)
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
	for _, k := range diff.Keys() {
		diffs = append(diffs, string(k))
	}

	return &pulumirpc.DiffResponse{
		Changes:  pulumirpc.DiffResponse_DIFF_SOME,
		Diffs:    diffs,
		Replaces: replaces,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *cfnProvider) Configure(ctx context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	region, ok := req.Variables["aws-native:config:region"]
	if !ok {
		return nil, errors.New("missing required property 'region'")
	}
	p.region = region

	cfg, err := config.LoadDefaultConfig(ctx)
	cfg.Region = p.region
	cfg.APIOptions = append(cfg.APIOptions, pulumiUserAgent()...)
	if err != nil {
		return nil, errors.Wrapf(err, "could not load AWS config")
	}

	p.cfn, p.ec2, p.ssm = cloudformation.NewFromConfig(cfg), ec2.NewFromConfig(cfg), ssm.NewFromConfig(cfg)
	p.cfg = cfg

	callerIdentityResp, err := sts.NewFromConfig(cfg).GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, errors.Wrapf(err, "could not get AWS account ID")
	}
	if callerIdentityResp.Account == nil {
		return nil, errors.New("could not get AWS account ID: nil account")
	}
	p.accountID = *callerIdentityResp.Account

	schema, err := getCloudFormationSchema(p.canceler.context, region)
	if err != nil {
		return nil, errors.Wrapf(err, "could not fetch CloudFormation schema")
	}

	// Map to convert module names to CF modules.
	cfTypes := map[string]string{}
	for key := range schema.ResourceTypes {
		parts := strings.Split(key, "::")
		if len(parts) != 3 {
			return nil, errors.Errorf("invalid resource type %s", key)
		}
		resourceKey := fmt.Sprintf("%s::%s::%s", parts[0], strings.ToLower(parts[1]), parts[2])
		cfTypes[resourceKey] = key
	}
	p.schema = *schema
	p.cfTypes = cfTypes

	p.supportedResourceTypes = readSupportedResourceTypes()
	p.configured = true

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

var functions = map[string]func(*cfnProvider, context.Context, resource.PropertyMap) (resource.PropertyMap, error){
	"aws-native:index:getAccountId":          (*cfnProvider).getAccountID,
	"aws-native:index:getAzs":                (*cfnProvider).getAZs,
	"aws-native:index:getRegion":             (*cfnProvider).getRegion,
	"aws-native:index:cidr":                  (*cfnProvider).cidr,
	"aws-native:index:getSsmParameterString": (*cfnProvider).getSSMParameterString,
	"aws-native:index:getSsmParameterList":   (*cfnProvider).getSSMParameterList,
	"aws-native:index:importValue":           (*cfnProvider).importValue,
}

// Invoke dynamically executes a built-in function in the provider.
func (p *cfnProvider) Invoke(ctx context.Context,
	req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {

	// Unmarshal arguments.
	tok := req.GetTok()

	inputs, err := plugin.UnmarshalProperties(req.GetArgs(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.Invoke(%s).inputs", p.name, tok),
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
		Label:        fmt.Sprintf("%s.Invoke(%s).outputs", p.name, tok),
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

	failures, err := schema.ValidateResource(sch, resourceType, inputs)
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
	label := fmt.Sprintf("%s.Check(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	_, resourceType, err := p.resourceInfoFromURN(urn)
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

func (p *cfnProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	_, resourceType, err := p.resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}

	diff, err := p.diffState(req.GetOlds(), req.GetNews(), label)
	if err != nil {
		return nil, err
	}
	if diff == nil {
		return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_NONE}, nil
	}

	var replaces []string
	rs, err := schema.GetResourceReplaces(p.schema, resourceType, diff)
	if err != nil {
		return nil, err
	}
	replaces = rs

	return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_SOME, Replaces: replaces}, nil
}

func (p *cfnProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs.
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "malformed resource inputs")
	}

	logicalID, resourceType, err := p.resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}

	// Convert SDK inputs to CFN payload.
	payload, err := schema.SdkToCfn(p.schema, resourceType, inputs.MapRepl(nil, mapReplStripSecrets))
	if err != nil {
		return nil, err
	}

	var resourceState map[string]interface{}
	var id string
	// Serialize inputs as a desired state JSON.
	if p.supportedResourceTypes.Has(resourceType) {
		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal as JSON")
		}
		desiredState := string(jsonBytes)

		// Create the resource with Cloud API.
		clientToken := uuid.New().String()
		glog.V(9).Infof("%s.CreateResource %q token %q state %q", label, resourceType, clientToken, desiredState)
		res, err := p.cfn.CreateResource(ctx, &cloudformation.CreateResourceInput{
			ClientToken:  aws.String(clientToken),
			TypeName:     aws.String(resourceType),
			DesiredState: aws.String(desiredState),
		})
		if err != nil {
			return nil, errors.Wrapf(err, "creating resource")
		}
		pi, err := p.waitForResourceOpCompletion(ctx, res.ProgressEvent)
		if err != nil {
			return nil, err
		}
		if pi.Identifier == nil {
			return nil, errors.New("received nil identifier while reading resource state")
		}

		// Retrieve the resource state from AWS.
		id = *pi.Identifier
		glog.V(9).Infof("%s.GetResource %q id %q", label, resourceType, id)
		rs, err := p.readResourceState(ctx, resourceType, id)
		if err != nil {
			return nil, errors.Wrapf(err, "reading resource state")
		}
		resourceState = rs
	} else {
		p.host.Log(ctx, diag.Warning, urn, "resource is deployed with CFN")
		if p.update == nil {
			sess, err := session.NewSession(&oldaws.Config{
				Region: oldaws.String(p.region),
			})
			if err != nil {
				return nil, errors.Errorf("could not create AWS session: %v", err)
			}
			update, err := update.StartStackUpdate(p.canceler.context, sess, "test-stack-name")
			if err != nil {
				return nil, errors.Errorf("could not start stack updater: %v", err)
			}
			p.update = update
		}

		data, err := p.update.CreateResource(p.canceler.context, logicalID, resourceType, payload) //, options...)
		if err != nil {
			return nil, errors.Wrap(err, "failed to CreateResource")
		}
		resourceState = cleanTopLevelAttributeNames(data.Attributes)
		id = data.PhysicalResourceID
	}

	outputs := schema.CfnToSdk(resourceState)

	// Store both outputs and inputs into the state.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(inputs, outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         id,
		Properties: checkpoint,
	}, nil
}

func (p *cfnProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)
	id := req.GetId()

	// Retrieve the old state.
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
	})
	if err != nil {
		return nil, err
	}

	// Read the resource state from AWS.
	_, resourceType, err := p.resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}
	resourceState, err := p.readResourceState(ctx, resourceType, id)
	if err != nil {
		var oe *smithy.OperationError
		if errors.As(err, &oe) {
			if re, ok := oe.Unwrap().(*awshttp.ResponseError); ok && re.HTTPStatusCode() == 404 {
				// ResourceNotFound means that the resource was deleted.
				return &pulumirpc.ReadResponse{Id: ""}, nil
			}
		}
		return nil, err
	}
	newState := schema.CfnToSdk(resourceState)

	// Extract old inputs from the `__inputs` field of the old state.
	inputs := parseCheckpointObject(oldState)
	if inputs == nil {
		// There may be no old state (i.e., importing a new resource).
		// Extract inputs from the response body.
		newStateProps := resource.NewPropertyMapFromMap(newState)
		inputs, err = schema.GetInputsFromState(p.schema, resourceType, newStateProps)
		if err != nil {
			return nil, err
		}
	} else {
		// It's hard to infer the changes in the inputs shape based on the outputs without false positives.
		// The current approach is complicated but it's aimed to minimize the noise while refreshing:
		// 0. We have "old" inputs and outputs before refresh and "new" outputs read from AWS.
		// 1. Project old outputs to their corresponding input shape (exclude attributes).
		oldInputProjection, err := schema.GetInputsFromState(p.schema, resourceType, oldState)
		if err != nil {
			return nil, err
		}
		// 2. Project new outputs to their corresponding input shape (exclude attributes).
		newStateProps := resource.NewPropertyMapFromMap(newState)
		newInputProjection, err := schema.GetInputsFromState(p.schema, resourceType, newStateProps)
		if err != nil {
			return nil, err
		}
		// 3. Calculate the difference between two projections. This should give us actual significant changes
		// that happened in AWS between the last resource update and its current state.
		diff := oldInputProjection.Diff(newInputProjection)
		// 4. Apply this difference to the actual inputs (not a projection) that we have in state.
		inputs = applyDiff(inputs, diff)
	}

	// Store both outputs and inputs into the state checkpoint.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(inputs, newState),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	// Serialize and return the calculated inputs.
	inputsRecord, err := plugin.MarshalProperties(
		inputs,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.inputs", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	return &pulumirpc.ReadResponse{
		Id:         id,
		Inputs:     inputsRecord,
		Properties: checkpoint,
	}, nil
}

func (p *cfnProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	id := req.GetId()
	_, resourceType, err := p.resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}

	diff, err := p.diffState(req.GetOlds(), req.GetNews(), label)
	if err != nil {
		return nil, err
	}

	ops, err := schema.DiffToPatch(p.schema, resourceType, diff)
	if err != nil {
		return nil, errors.Wrapf(err, "calculating diff patch")
	}

	clientToken := uuid.New().String()
	glog.V(9).Infof("%s.UpdateResource %q id %q token %q state %+v", label, resourceType, id, clientToken, ops)
	res, err := p.cfn.UpdateResource(ctx, &cloudformation.UpdateResourceInput{
		ClientToken:     aws.String(clientToken),
		TypeName:        aws.String(resourceType),
		Identifier:      aws.String(id),
		PatchOperations: ops,
	})
	if err != nil {
		return nil, err
	}
	if _, err = p.waitForResourceOpCompletion(ctx, res.ProgressEvent); err != nil {
		return nil, err
	}

	resourceState, err := p.readResourceState(ctx, resourceType, id)
	if err != nil {
		return nil, errors.Wrapf(err, "reading resource state")
	}
	outputs := schema.CfnToSdk(resourceState)

	// Read the inputs to persist them into state.
	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.newInputs", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	// Store both outputs and inputs into the state and return RPC checkpoint.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(newInputs, outputs),
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{Properties: checkpoint}, nil
}

func readSupportedResourceTypes() codegen.StringSet {
	outDir := "/Users/mikhailshilkov/go/src/github.com/pulumi/pulumi-aws-native/provider/cmd/pulumi-gen-aws-native"
	path := filepath.Join(outDir, "supported-types.txt")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return codegen.NewStringSet(lines...)
}

func (p *cfnProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	logicalID, resourceType, err := p.resourceInfoFromURN(urn)
	if err != nil {
		return nil, err
	}
	id := req.GetId()

	// Serialize inputs as a desired state JSON.
	if p.supportedResourceTypes.Has(resourceType) {
		clientToken := uuid.New().String()
		glog.V(9).Infof("%s.DeleteResource %q id %q token %q", label, resourceType, id, clientToken)
		res, err := p.cfn.DeleteResource(ctx, &cloudformation.DeleteResourceInput{
			ClientToken: aws.String(clientToken),
			TypeName:    aws.String(resourceType),
			Identifier:  aws.String(id),
		})
		if err != nil {
			return nil, err
		}
		if pi, err := p.waitForResourceOpCompletion(ctx, res.ProgressEvent); err != nil {
			if pi != nil {
				errorCode := pi.ErrorCode
				if errorCode == types.HandlerErrorCodeNotFound {
					// NotFound means that the resource was already deleted, so the operation can succeed.
					return &pbempty.Empty{}, nil
				}
			}
		}
	} else {
		p.host.Log(ctx, diag.Warning, urn, "resource is deleted with CFN")
		if p.update == nil {
			sess, err := session.NewSession(&oldaws.Config{
				Region: oldaws.String(p.region),
			})
			if err != nil {
				return nil, errors.Errorf("could not create AWS session: %v", err)
			}
			update, err := update.StartStackUpdate(p.canceler.context, sess, "test-stack-name")
			if err != nil {
				return nil, errors.Errorf("could not start stack updater: %v", err)
			}
			p.update = update
		}

		if err := p.update.DeleteResource(p.canceler.context, logicalID, id); err != nil {
			return nil, err
		}
	}

	return &pbempty.Empty{}, nil
}

// Construct creates a new component resource.
func (p *cfnProvider) Construct(_ context.Context, _ *pulumirpc.ConstructRequest) (*pulumirpc.ConstructResponse, error) {
	panic("Construct not implemented")
}

// Call dynamically executes a method in the provider associated with a component resource.
func (p *cfnProvider) Call(_ context.Context, _ *pulumirpc.CallRequest) (*pulumirpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
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

func (p *cfnProvider) resourceInfoFromURN(urn resource.URN) (string, string, error) {
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
	cfType, has := p.cfTypes[typ]
	if !has {
		return "", "", errors.Errorf("resource type %s not found", typ)
	}

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

	return name + hash, cfType, nil
}

func (p *cfnProvider) readResourceState(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
	getRes, err := p.cfn.GetResource(ctx, &cloudformation.GetResourceInput{
		TypeName:   aws.String(typeName),
		Identifier: aws.String(identifier),
	})
	if err != nil {
		return nil, err
	}
	if getRes.ResourceDescription.ResourceModel == nil {
		return nil, errors.New("received nil resource model")
	}

	resourceModel := *getRes.ResourceDescription.ResourceModel
	var outputs map[string]interface{}
	if err = json.Unmarshal([]byte(resourceModel), &outputs); err != nil {
		return nil, err
	}

	return outputs, nil
}

func (p *cfnProvider) waitForResourceOpCompletion(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
	// TODO: Consider using ExponentialJitterBackoff from the AWS Go SDK v2 when we switch over.
	retryPolicy := backoff.Backoff{
		Min:    1 * time.Second,
		Max:    30 * time.Second,
		Factor: 1.5,
		Jitter: true,
	}
	for {
		status := pi.OperationStatus
		identifier := ""
		if pi.Identifier != nil {
			identifier = *pi.Identifier
		}
		glog.V(9).Infof("waiting for resource %q: status %q", identifier, status)
		switch status {
		case "SUCCESS":
			return pi, nil
		case "FAILED":
			statusMessage := ""
			if pi.StatusMessage != nil {
				statusMessage = *pi.StatusMessage
			}
			return pi, errors.Errorf("operation %s failed with %q: %s", pi.Operation, pi.ErrorCode, statusMessage)
		case "IN_PROGRESS":
			var pause time.Duration
			if pi.RetryAfter != nil {
				pause = pi.RetryAfter.Sub(time.Now())
			} else {
				pause = retryPolicy.Duration()
			}
			time.Sleep(pause)
		default:
			return nil, errors.Errorf("unknown status %q: %+v", status, pi)
		}

		output, err := p.cfn.GetResourceRequestStatus(ctx, &cloudformation.GetResourceRequestStatusInput{
			RequestToken: pi.RequestToken,
		})
		if err != nil {
			return nil, errors.Wrap(err, "getting resource request status")
		}
		pi = output.ProgressEvent
	}
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

// diffState extracts old and new inputs and calculates a diff between them.
func (p *cfnProvider) diffState(olds *pbstruct.Struct, news *pbstruct.Struct, label string) (*resource.ObjectDiff, error) {
	oldState, err := plugin.UnmarshalProperties(olds, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.oldState", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)

	newInputs, err := plugin.UnmarshalProperties(news, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.newInputs", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	return oldInputs.Diff(newInputs), nil
}

// applyDiff produces a new map as a merge of a calculated diff into an existing map of values.
func applyDiff(values resource.PropertyMap, diff *resource.ObjectDiff) resource.PropertyMap {
	if diff == nil {
		return values
	}

	result := resource.PropertyMap{}
	for name, value := range values {
		if !diff.Deleted(name) {
			result[name] = value
		}
	}
	for key, value := range diff.Adds {
		result[key] = value
	}
	for key, value := range diff.Updates {
		result[key] = value.New
	}
	return result
}

// checkpointObject puts inputs in the `__inputs` field of the state.
func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

// parseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
func parseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.SecretValue().Element.ObjectValue()
	}

	return nil
}

// getPulumiVersion parses the version of the pulumi SDK used in the running program.
func getPulumiVersion() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range bi.Deps {
			if strings.HasPrefix(dep.Path, "github.com/pulumi/pulumi/sdk") {
				return strings.TrimPrefix(dep.Version, "v")
			}
		}
	}
	// We should never get here but let's not panic and return something sensible if we do.
	logging.V(4).Info("No Pulumi package version found, using '3.0' as the default version for user-agent")
	return "3.0"
}

// pulumiUserAgent adds a Pulumi-specific user-agent to the request middleware.
// The resulting string looks like this: `APN/1.0 Pulumi/1.0 PulumiAwsNative/0.0.2`
func pulumiUserAgent() []func(*middleware.Stack) error {
	return []func(*middleware.Stack) error{
		awsmiddleware.AddUserAgentKeyValue("APN", "1.0"),
		awsmiddleware.AddUserAgentKeyValue("Pulumi", getPulumiVersion()),
		awsmiddleware.AddUserAgentKeyValue("PulumiAwsNative", version.Version),
	}
}
