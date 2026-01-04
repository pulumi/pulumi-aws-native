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
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/ratelimit"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	ststypes "github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/client"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	pOutputs "github.com/pulumi/pulumi-aws-native/provider/pkg/outputs"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/resources"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi-go-provider/resourcex"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil/rpcerror"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
)

// The APN 1.1 AWS Marketplace identifier to should be used in the User-Agent header.
var PulumiAWSMarketplaceCode string = "c7qiae2l6usvzoynupds6v7hf"

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
	pulumirpc.UnimplementedResourceProviderServer

	host     *provider.HostClient
	name     string
	canceler *cancellationContext

	configured          bool
	version             string
	accountID           string
	region              string
	partition           partition
	resourceMap         *metadata.CloudAPIMetadata
	roleArn             *string
	autoNamingConfig    *autonaming.ProviderAutoNamingConfig
	allowedAccountIds   []string
	forbiddenAccountIds []string
	defaultTags         map[string]string

	pulumiSchema []byte

	cfn    *cloudformation.Client
	ccc    client.CloudControlClient
	ec2    *ec2.Client
	ssm    *ssm.Client
	sts    *sts.Client
	s3     *s3.Client
	lambda *lambda.Client

	customResources map[string]resources.CustomResource
}

var _ pulumirpc.ResourceProviderServer = (*cfnProvider)(nil)

func NewAwsNativeProvider(host *provider.HostClient, name, version string,
	pulumiSchema, cloudAPIResourcesBytes []byte) (pulumirpc.ResourceProviderServer, error) {
	resourceMap, err := LoadMetadata(cloudAPIResourcesBytes)
	if err != nil {
		return nil, err
	}

	return &cfnProvider{
		host:         host,
		canceler:     makeCancellationContext(),
		name:         name,
		version:      version,
		resourceMap:  resourceMap,
		pulumiSchema: pulumiSchema,
	}, nil
}

// LoadMetadata deserializes the provided compressed json byte array into a CloudAPIMetadata struct.
func LoadMetadata(metadataBytes []byte) (*metadata.CloudAPIMetadata, error) {
	var resourceMap metadata.CloudAPIMetadata

	uncompressed, err := gzip.NewReader(bytes.NewReader(metadataBytes))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed metadata")
	}
	if err = json.NewDecoder(uncompressed).Decode(&resourceMap); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	if err = uncompressed.Close(); err != nil {
		return nil, errors.Wrap(err, "closing uncompress stream for metadata")
	}
	return &resourceMap, nil
}

func (p *cfnProvider) Attach(context context.Context, req *pulumirpc.PluginAttach) (*emptypb.Empty, error) {
	host, err := provider.NewHostClient(req.GetAddress())
	if err != nil {
		return nil, err
	}
	p.host = host
	return &pbempty.Empty{}, nil
}

func (p *cfnProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	decompressed, err := schema.DecompressSchema(p.pulumiSchema)
	if err != nil {
		return nil, errors.New("failure loading schema")
	}
	return &pulumirpc.GetSchemaResponse{Schema: string(decompressed)}, nil
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

	truthyValue := func(argName resource.PropertyKey, props resource.PropertyMap) bool {
		if arg := props[argName]; arg.HasValue() {
			switch {
			case arg.IsString() && len(arg.StringValue()) > 0:
				return true
			case arg.IsBool() && arg.BoolValue():
				return true
			}
		}
		return false
	}

	// These provider values are not yet implemented, so return an error and link to the relevant issue(s).
	var failures []*pulumirpc.CheckFailure
	for k := range news {
		switch k {
		case "ignoreTags":
			failures = append(failures, &pulumirpc.CheckFailure{Property: string(k),
				Reason: fmt.Sprintf("not yet implemented. See https://github.com/pulumi/pulumi-aws-native/issues/110")})
		case "insecure":
			failures = append(failures, &pulumirpc.CheckFailure{Property: string(k),
				Reason: fmt.Sprintf("not yet implemented. See https://github.com/pulumi/pulumi-aws-native/issues/111")})
		case "skipGetEc2Platforms":
			if !truthyValue(k, news) {
				failures = append(failures, &pulumirpc.CheckFailure{Property: string(k),
					Reason: fmt.Sprintf("not yet implemented. See https://github.com/pulumi/pulumi-aws-native/issues/115")})
			}
		case "skipMetadataApiCheck":
			if !truthyValue(k, news) {
				failures = append(failures, &pulumirpc.CheckFailure{Property: string(k),
					Reason: fmt.Sprintf("not yet implemented. See https://github.com/pulumi/pulumi-aws-native/issues/116")})
			}
		case "skipRegionValidation":
			if !truthyValue(k, news) {
				failures = append(failures, &pulumirpc.CheckFailure{Property: string(k),
					Reason: fmt.Sprintf("not yet implemented. See https://github.com/pulumi/pulumi-aws-native/issues/117")})
			}
		case "skipRequestingAccountId":
			failures = append(failures, &pulumirpc.CheckFailure{Property: string(k),
				Reason: fmt.Sprintf("not yet implemented. See https://github.com/pulumi/pulumi-aws-native/issues/118")})
		}
	}
	if failures != nil {
		return &pulumirpc.CheckResponse{Failures: failures}, nil
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

	var diffs []string
	for _, k := range diff.Keys() {
		diffs = append(diffs, string(k))
	}

	return &pulumirpc.DiffResponse{
		Changes: pulumirpc.DiffResponse_DIFF_SOME,
		Diffs:   diffs,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *cfnProvider) Configure(ctx context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	vars := req.GetVariables()

	// loadOptions are used to override default config loading behavior.
	var loadOptions []func(*config.LoadOptions) error

	if region, ok := varsOrEnv(vars, "aws-native:config:region", "AWS_REGION", "AWS_DEFAULT_REGION"); ok {
		glog.V(4).Infof("using AWS region: %q", region)
		loadOptions = append(loadOptions, config.WithRegion(region))
		p.region, p.partition = region, getPartition(region)
	} else {
		return nil, errors.New("missing required configuration key \"aws-native:region\":" +
			"The region where AWS operations will take place. Examples are eu-east-1, eu-west-2, etc.\n" +
			"\tSet a value using the command `pulumi config set aws-native:region <value>`, or by setting the environment variables `AWS_REGION` or `AWS_DEFAULT_REGION`.")
	}

	if profile, ok := varsOrEnv(vars, "aws-native:config:profile", "AWS_PROFILE"); ok {
		glog.V(4).Infof("using AWS profile: %q", profile)
		loadOptions = append(loadOptions, config.WithSharedConfigProfile(profile))
	} else {
		glog.V(4).Infof(`using AWS profile: "default"`)
	}

	if sharedCredentialsFilePath, ok := varsOrEnv(vars, "aws-native:config:sharedCredentialsFile", "AWS_SHARED_CREDENTIALS_FILE"); ok {
		glog.V(4).Infof("using AWS shared credentials file at path: %q", sharedCredentialsFilePath)
		normalizedPath, err := normalizeFilePath(sharedCredentialsFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load shared credentials file: %w", err)
		}
		loadOptions = append(loadOptions, config.WithSharedCredentialsFiles([]string{normalizedPath}))
	} else {
		if runtime.GOOS == "windows" {
			windowsPath := `%USERPROFILE%\.aws\credentials`
			glog.V(4).Infof(`using AWS shared credentials file at path: %q`, windowsPath)
		} else {
			glog.V(4).Info(`using AWS shared credentials file at path: "~/.aws/credentials"`)
		}
	}

	// Environment variables are checked by the AWS SDK by default as a fallback after explicitly defined config.
	// See https://github.com/pulumi/pulumi-aws-native/pull/1378
	var accessKey, secretKey, token string

	if v, ok := varsOrEnv(vars, "aws-native:config:accessKey"); ok {
		accessKey = v
	}
	if v, ok := varsOrEnv(vars, "aws-native:config:secretKey"); ok {
		secretKey = v
	}
	if v, ok := varsOrEnv(vars, "aws-native:config:token"); ok {
		token = v
	}

	if len(accessKey) > 0 || len(secretKey) > 0 || len(token) > 0 {
		// If all required values are not present/valid, the client will return an appropriate error.
		credsProvider := credentials.NewStaticCredentialsProvider(accessKey, secretKey, token)
		loadOptions = append(loadOptions, config.WithCredentialsProvider(credsProvider))
	}

	if endpointsString, ok := vars["aws-native:config:endpoints"]; ok {
		var endpoints map[string]string
		err := json.Unmarshal([]byte(endpointsString), &endpoints)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal 'endpoints' config: %w", err)
		}
		glog.V(4).Infof("using AWS endpoints: %v", endpoints)
		resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if endpoint, ok := endpoints[strings.ToLower(service)]; ok {
				return aws.Endpoint{URL: endpoint}, nil
			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})
		loadOptions = append(loadOptions, config.WithEndpointResolverWithOptions(resolver))
	}

	// Attach custom middleware to the client.
	loadOptions = append(loadOptions, config.WithAPIOptions(func() (v []func(stack *middleware.Stack) error) {
		v = append(v, attachCustomMiddleware)
		return v
	}()))

	if glog.V(9) {
		loadOptions = append(loadOptions, config.WithClientLogMode(aws.LogRequestWithBody|aws.LogResponseWithBody))
	}

	// Load configuration from the environment, overriding with any config that was explicitly set on the Provider.
	cfg, err := config.LoadDefaultConfig(ctx, loadOptions...)
	if err != nil {
		return nil, errors.Wrapf(err, "could not load AWS config")
	}

	if assumeRoleJson, ok := vars["aws-native:config:assumeRole"]; ok {
		var assumeRole ProviderAssumeRole
		err := json.Unmarshal([]byte(assumeRoleJson), &assumeRole)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal 'assumeRole' config: %w", err)
		}

		creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(cfg), "roleArn",
			func(o *stscreds.AssumeRoleOptions) {
				if assumeRole.DurationSeconds != nil {
					o.Duration = time.Duration(*(assumeRole.DurationSeconds)) * time.Second
				}
				o.ExternalID = assumeRole.ExternalId
				o.Policy = assumeRole.Policy

				for _, arn := range assumeRole.PolicyArns {
					o.PolicyARNs = append(o.PolicyARNs, ststypes.PolicyDescriptorType{Arn: aws.String(arn)})
				}
				if assumeRole.RoleArn != nil {
					o.RoleARN = *assumeRole.RoleArn
				}
				if assumeRole.SessionName != nil {
					o.RoleSessionName = *assumeRole.SessionName
				}
				for k, v := range assumeRole.Tags {
					o.Tags = append(o.Tags, ststypes.Tag{Key: aws.String(k), Value: aws.String(v)})
				}
				if assumeRole.TransitiveTagKeys != nil {
					o.TransitiveTagKeys = assumeRole.TransitiveTagKeys
				}
			})
		cfg.Credentials = aws.NewCredentialsCache(creds)
	}

	if roleArn, ok := vars["aws-native:config:roleArn"]; ok {
		p.roleArn = &roleArn
	}

	maxRetries := 25
	if maxRetriesConf, ok := vars["aws-native:config:maxRetries"]; ok {
		maxRetries, err = strconv.Atoi(maxRetriesConf)
		if err != nil {
			return nil, fmt.Errorf("invalid config value for 'maxRetries': %q: %w", maxRetriesConf, err)
		}
	}
	glog.V(4).Infof("using Max Retry Attempts: %d", maxRetries)

	maxRetryRateTokens := -1
	if maxRetryRateTokensConf, ok := vars["aws-native:config:maxRetryRateTokens"]; ok {
		maxRetryRateTokens, err = strconv.Atoi(maxRetryRateTokensConf)
		if err != nil {
			return nil, fmt.Errorf("invalid config value for 'maxRetryRateTokens': %q: %w", maxRetryRateTokensConf, err)
		}
		glog.V(4).Infof("using Retry Token Rate Limit: %d", maxRetryRateTokens)
	} else {
		glog.V(4).Infof("using Retry Token Rate Limit: unlimited")
	}

	cfg.Retryer = func() (r aws.Retryer) {
		rateLimiter := func(o *retry.StandardOptions) {
			// We don't use a token bucket by default because that will raise
			// errors when the bucket is empty. Instead, the adaptive retryer
			// will sleep for us when the server responds with Throttling
			// errors. The default backoff strategy also introduces delays when
			// retrying.
			o.RateLimiter = ratelimit.None
			if maxRetryRateTokens > 0 {
				o.RateLimiter = ratelimit.NewTokenRateLimit(uint(maxRetryRateTokens))
			}
		}

		// Increase the maximum number of retry attempts from the default of 3.
		// Note this must be applied after AddWithErrorCodes for some reason.
		retries := func(o *retry.StandardOptions) {
			o.MaxAttempts = maxRetries
		}

		// AdaptiveMode is similar to the default Standard retryer but sleeps
		// when it gets a Throttling response.
		retryer := retry.NewAdaptiveMode(func(o *retry.AdaptiveModeOptions) {
			o.StandardOptions = append(o.StandardOptions,
				rateLimiter,
				retries,
			)
		})
		return retryer
	}

	if allowedAccountIdsJson, ok := vars["aws-native:config:allowedAccountIds"]; ok {
		var allowedAccountIds []string
		err := json.Unmarshal([]byte(allowedAccountIdsJson), &allowedAccountIds)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal 'allowedAccountIds' config: %w", err)
		}
		p.allowedAccountIds = allowedAccountIds
	}

	if forbiddenAccountIdsJson, ok := vars["aws-native:config:forbiddenAccountIds"]; ok {
		var forbiddenAccountIds []string
		err := json.Unmarshal([]byte(forbiddenAccountIdsJson), &forbiddenAccountIds)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal 'forbiddenAccountIds' config: %w", err)
		}
		p.forbiddenAccountIds = forbiddenAccountIds
	}

	skipCredentialsValidation := false
	if skipCredentialsValidationVar, ok := varsOrEnv(vars, "aws-native:config:skipCredentialsValidation", "AWS_SKIP_CREDENTIALS_VALIDATION"); ok {
		if skipCredentialsValidationVar == "true" {
			skipCredentialsValidation = true
		}
	}

	if defaultTagsJson, ok := vars["aws-native:config:defaultTags"]; ok {
		type DefaultTags struct {
			Tags map[string]string `json:"tags"`
		}
		var defaultTags DefaultTags
		err := json.Unmarshal([]byte(defaultTagsJson), &defaultTags)
		if err == nil {
			p.defaultTags = defaultTags.Tags
		} else {
			// As a fallback, we also try to unmarshal the default tags as a simple map[string]string
			// as this was originally supported when implementing defaultTags.
			fallbackErr := json.Unmarshal([]byte(defaultTagsJson), &p.defaultTags)
			if fallbackErr != nil {
				return nil, fmt.Errorf("failed to unmarshal 'defaultTags' config: %w", err)
			}
		}
	} else {
		p.defaultTags = nil
	}

	s3ForcePathStyle := false
	if s3ForcePathStyleVar, ok := vars["aws-native:config:s3ForcePathStyle"]; ok {
		if s3ForcePathStyleVar == "true" {
			s3ForcePathStyle = true
			glog.V(4).Infof("using S3 path-style addressing")
		}
	}

	p.cfn = cloudformation.NewFromConfig(cfg)
	p.ccc = client.NewCloudControlClient(cloudcontrol.NewFromConfig(cfg), p.roleArn, cfg.Retryer, p.host)
	p.ec2 = ec2.NewFromConfig(cfg)
	p.ssm = ssm.NewFromConfig(cfg)
	p.sts = sts.NewFromConfig(cfg)
	p.s3 = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = s3ForcePathStyle
	})
	p.lambda = lambda.NewFromConfig(cfg)

	if !skipCredentialsValidation {
		callerIdentityResp, err := p.sts.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
		if err != nil {
			return nil, errors.Wrapf(err, "could not get AWS account ID")
		}
		if callerIdentityResp.Account == nil {
			return nil, errors.New("could not get AWS account ID: nil account")
		}
		p.accountID = *callerIdentityResp.Account

		err = p.validateAccountId()
		if err != nil {
			return nil, err
		}
	}

	lambdaClient := client.NewLambdaClient(p.lambda)
	s3Client := client.NewS3Client(p.s3, s3.NewPresignClient(p.s3))

	p.customResources = map[string]resources.CustomResource{
		metadata.ExtensionResourceToken: resources.NewExtensionResource(p.ccc),
		metadata.CfnCustomResourceToken: resources.NewCfnCustomResource(p.name, s3Client, lambdaClient),
	}

	if autoNaming, ok := vars["aws-native:config:autoNaming"]; ok {
		var autoNamingConfig autonaming.ProviderAutoNamingConfig
		if err := json.Unmarshal([]byte(autoNaming), &autoNamingConfig); err != nil {
			return nil, fmt.Errorf("failed to unmarshal 'autoNaming' config: %w", err)
		}
		p.autoNamingConfig = &autoNamingConfig
	}

	p.configured = true

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets:                   true,
		SupportsPreview:                 true,
		SupportsAutonamingConfiguration: true,
	}, nil
}

func varsOrEnv(vars map[string]string, key string, env ...string) (string, bool) {
	if val, ok := vars[key]; ok {
		return val, true
	}
	for _, e := range env {
		if val, ok := os.LookupEnv(e); ok {
			return val, true
		}
	}
	return "", false
}

// ValidateAccountId returns a context-specific error if the configured account
// id is explicitly forbidden or not authorised; and nil if it is authorised.
func (p *cfnProvider) validateAccountId() error {
	if len(p.allowedAccountIds) == 0 && len(p.forbiddenAccountIds) == 0 {
		return nil
	}
	if p.forbiddenAccountIds != nil {
		for _, id := range p.forbiddenAccountIds {
			if id == p.accountID {
				return fmt.Errorf("forbidden account ID (%s)", id)
			}
		}
	}
	if p.allowedAccountIds != nil {
		for _, id := range p.allowedAccountIds {
			if id == p.accountID {
				return nil
			}
		}
		return fmt.Errorf("account ID not allowed (%s)", p.accountID)
	}
	return nil
}

var functions = map[string]func(*cfnProvider, context.Context, resource.PropertyMap) (resource.PropertyMap, error){
	"aws-native:index:getAccountId":          (*cfnProvider).getAccountID,
	"aws-native:index:getAzs":                (*cfnProvider).getAZs,
	"aws-native:index:getRegion":             (*cfnProvider).getRegion,
	"aws-native:index:getUrlSuffix":          (*cfnProvider).getURLSuffix,
	"aws-native:index:cidr":                  (*cfnProvider).cidr,
	"aws-native:index:getSsmParameterString": (*cfnProvider).getSSMParameterString,
	"aws-native:index:getSsmParameterList":   (*cfnProvider).getSSMParameterList,
	"aws-native:index:importValue":           (*cfnProvider).importValue,
	"aws-native:index:getPartition":          (*cfnProvider).getPartition,
}

type invokeFunc func(p *cfnProvider, ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error)

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
	var result resource.PropertyMap
	fn, ok := functions[tok]
	if !ok {
		fn, ok = p.getInvokeFunc(ctx, tok)
	}
	if !ok {
		return nil, fmt.Errorf("unknown function '%s'", tok)
	}
	result, err = fn(p, ctx, inputs)
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

func (p *cfnProvider) getInvokeFunc(ctx context.Context, tok string) (invokeFunc, bool) {
	cf, ok := p.resourceMap.Functions[tok]
	if !ok {
		return nil, false
	}
	return func(p *cfnProvider, ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
		idParts := make([]string, len(cf.Identifiers))
		for i, v := range cf.Identifiers {
			pv, ok := inputs[resource.PropertyKey(v)]
			if !ok {
				return nil, errors.Errorf("missing identifier property %q", v)
			}
			if !pv.IsString() {
				return nil, errors.Errorf("identifier property %q, expected type string, found %q", v, pv.TypeString())
			}
			idParts[i] = pv.StringValue()
		}
		identifier := strings.Join(idParts, "|")
		glog.V(9).Infof("%s invoking", cf.CfType)
		outputs, exists, err := p.ccc.Read(ctx, cf.CfType, identifier)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("resource not found")
		}
		sdkOutput := naming.CfnToSdk(outputs)

		return resource.NewPropertyMapFromMap(sdkOutput), nil
	}, true
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

func (p *cfnProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Check(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label: fmt.Sprintf("%s.olds", label), SkipNulls: true,
	})
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

	var failures []resources.ValidationFailure
	resourceToken := string(urn.Type())
	if customResource, ok := p.customResources[resourceToken]; ok {
		newInputs, failures, err = customResource.Check(ctx, urn, engineAutonaming(req), newInputs, olds, p.defaultTags)
		if err != nil {
			return nil, fmt.Errorf("failed to check custom resource %q: %w", resourceToken, err)
		}
	} else {
		spec, ok := p.resourceMap.Resources[resourceToken]
		if !ok {
			return nil, errors.Errorf("resource type %s not found", resourceToken)
		}

		if err := autonaming.ApplyAutoNaming(spec.AutoNamingSpec, urn, engineAutonaming(req), p.autoNamingConfig,
			olds, newInputs); err != nil {
			return nil, fmt.Errorf("failed to apply auto-naming: %w", err)
		}

		// Merge default tags into the inputs if the resource supports tags and the user has not overridden them.
		if len(p.defaultTags) > 0 && spec.TagsProperty != "" && spec.TagsStyle != default_tags.TagsStyleUnknown {
			tagsKey := resource.PropertyKey(spec.TagsProperty)
			val, err := default_tags.MergeDefaultTags(newInputs[tagsKey], p.defaultTags, spec.TagsStyle)
			if err != nil {
				return nil, err
			}
			newInputs[tagsKey] = val
		}

		failures, err = resources.ValidateResource(&spec, p.resourceMap.Types, newInputs)
		if err != nil {
			return nil, err
		}
	}
	var checkFailures []*pulumirpc.CheckFailure
	for _, f := range failures {
		checkFailures = append(checkFailures, &pulumirpc.CheckFailure{Property: f.Path, Reason: f.Reason})
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
	return &pulumirpc.CheckResponse{Failures: checkFailures}, nil
}

func engineAutonaming(req *pulumirpc.CheckRequest) autonaming.EngineAutoNamingConfig {
	engineAutonaming := autonaming.EngineAutoNamingConfig{
		RandomSeed: req.RandomSeed,
	}
	if req.Autonaming != nil {
		mode := autonaming.EngineAutonamingMode(req.Autonaming.Mode)
		engineAutonaming.AutonamingMode = &mode
		engineAutonaming.ProposedName = req.Autonaming.ProposedName
	}
	return engineAutonaming
}

func (p *cfnProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse inputs for update")
	}

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.oldState", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := resources.ParseCheckpointObject(oldState)

	diff := oldInputs.Diff(newInputs)

	if diff == nil {
		return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_NONE}, nil
	}

	return &pulumirpc.DiffResponse{
		Changes:             pulumirpc.DiffResponse_DIFF_UNKNOWN,
		DeleteBeforeReplace: true,
	}, nil
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

	resourceToken := string(urn.Type())

	var outputs resource.PropertyMap
	if req.GetPreview() {
		if customResource, ok := p.customResources[resourceToken]; ok {
			outputs = customResource.PreviewCustomResourceOutputs()
		} else {
			spec, hasSpec := p.resourceMap.Resources[resourceToken]
			if !hasSpec {
				return nil, errors.Errorf("Resource type %s not found", resourceToken)
			}
			outputs = pOutputs.PreviewOutputs(inputs, p.resourceMap.Types, spec.Outputs, spec.ReadOnly, urn.Type().Name(), nil)
		}

		previewState, err := plugin.MarshalProperties(
			outputs,
			plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
		)
		if err != nil {
			return nil, err
		}
		return &pulumirpc.CreateResponse{
			Properties: previewState,
		}, nil
	}

	var id *string
	var createErr error
	timeout := time.Duration(req.GetTimeout()) * time.Second
	if customResource, ok := p.customResources[resourceToken]; ok {
		id, outputs, createErr = customResource.Create(ctx, urn, inputs, timeout)
	} else {
		// Standard resource
		spec, hasSpec := p.resourceMap.Resources[resourceToken]
		if !hasSpec {
			return nil, errors.Errorf("Resource type %s not found", resourceToken)
		}
		cfType := spec.CfType

		// Convert SDK inputs to CFN payload.
		payload, err := naming.SdkToCfn(&spec, p.resourceMap.Types, resourcex.Decode(inputs))
		if err != nil {
			return nil, fmt.Errorf("Failed to convert value for %s: %w", resourceToken, err)
		}

		// Create the resource with Cloud API.
		glog.V(9).Infof("%s.CreateResource %q", label, cfType)
		var resourceState map[string]interface{}
		id, resourceState, createErr = p.ccc.Create(ctx, urn, cfType, payload)
		if createErr != nil && (id == nil || resourceState == nil) {
			return nil, errors.Wrapf(createErr, "creating resource")
		}

		rawOutputs := naming.CfnToSdk(resourceState)
		// Write-only properties are not returned in the outputs, so we assume they should have the same value we sent from the inputs.
		if hasSpec && len(spec.WriteOnly) > 0 {
			inputsMap := inputs.Mappable()
			for _, writeOnlyProp := range spec.WriteOnly {
				if _, ok := rawOutputs[writeOnlyProp]; !ok {
					inputValue, ok := inputsMap[writeOnlyProp]
					if ok {
						rawOutputs[writeOnlyProp] = inputValue
					}
				}
			}
		}
		outputs = resources.CheckpointObject(inputs, rawOutputs)
	}

	if createErr != nil {
		if id == nil {
			return nil, fmt.Errorf("creating resource: %w", createErr)
		}
		// Resource was created but failed to fully initialize.
		// It has some state, so we return a partial error.
		checkpoint, err := plugin.MarshalProperties(
			outputs,
			plugin.MarshalOptions{
				Label:        "currentResourceStateCheckpoint.checkpoint",
				KeepSecrets:  true,
				KeepUnknowns: true,
				SkipNulls:    true,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("marshalling currentResourceStateCheckpoint: %w", err)
		}
		return nil, partialError(*id, createErr, checkpoint, req.GetProperties())
	}

	// Store both outputs and inputs into the state.
	checkpoint, err := plugin.MarshalProperties(
		outputs,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         *id,
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
	// Extract old inputs from the `__inputs` field of the old state.
	inputs := resources.ParseCheckpointObject(oldState)

	// Read the resource state from AWS.
	resourceToken := string(urn.Type())
	var newInputs resource.PropertyMap
	var newState resource.PropertyMap
	var exists bool
	if customResource, ok := p.customResources[resourceToken]; ok {
		newState, newInputs, exists, err = customResource.Read(ctx, urn, id, inputs, oldState)
		if err != nil {
			return nil, err
		}
		if !exists {
			// Not Exists means that the resource was deleted.
			return &pulumirpc.ReadResponse{Id: ""}, nil
		}
	} else {
		// Standard resource
		spec, ok := p.resourceMap.Resources[resourceToken]
		if !ok {
			return nil, errors.Errorf("Resource type %s not found", resourceToken)
		}
		var resourceState map[string]interface{}
		resourceState, exists, err = p.ccc.Read(p.canceler.context, spec.CfType, id)
		if err != nil {
			return nil, err
		}
		if !exists {
			// Not Exists means that the resource was deleted.
			return &pulumirpc.ReadResponse{Id: ""}, nil
		}
		rawState := naming.CfnToSdk(resourceState)

		if inputs == nil {
			// There may be no old state (i.e., importing a new resource).
			// Extract inputs from the response body.
			newStateProps := resource.NewPropertyMapFromMap(rawState)
			newInputs, err = schema.GetInputsFromState(&spec, newStateProps)
			if err != nil {
				return nil, err
			}
			if len(spec.WriteOnly) > 0 {
				p.host.Log(ctx, diag.Warning, urn, fmt.Sprintf("Can't import write-only properties: %s", strings.Join(spec.WriteOnly, ", ")))
			}
		} else {
			// It's hard to infer the changes in the inputs shape based on the outputs without false positives.
			// The current approach is complicated but it's aimed to minimize the noise while refreshing:
			// 0. We have "old" inputs and outputs before refresh and "new" outputs read from AWS.
			// 1. Project old outputs to their corresponding input shape (exclude attributes).
			oldInputProjection, err := schema.GetInputsFromState(&spec, oldState)
			if err != nil {
				return nil, err
			}
			oldStateMap := oldState.Mappable()
			// Fill in the write-only properties from the old state as they won't included when reading.
			if len(spec.WriteOnly) > 0 {
				missingProps := make([]string, 0, len(spec.WriteOnly))
				for _, writeOnlyProp := range spec.WriteOnly {
					if _, ok := rawState[writeOnlyProp]; !ok {
						oldValue, ok := oldStateMap[writeOnlyProp]
						missingProps = append(missingProps, writeOnlyProp)
						if ok {
							rawState[writeOnlyProp] = oldValue
						}
					}
				}
				p.host.Log(ctx, diag.Warning, urn, fmt.Sprintf("Can't refresh write-only properties: %s", strings.Join(missingProps, ", ")))
			}
			// 2. Project new outputs to their corresponding input shape (exclude attributes).
			newStateProps := resource.NewPropertyMapFromMap(rawState)
			newInputProjection, err := schema.GetInputsFromState(&spec, newStateProps)
			if err != nil {
				return nil, err
			}
			// 3. Calculate the difference between two projections. This should give us actual significant changes
			// that happened in AWS between the last resource update and its current state.
			diff := oldInputProjection.Diff(newInputProjection)
			// 4. Suppress AWS-managed changes from the diff. This removes:
			//    - aws:* prefixed tags that AWS adds automatically (users cannot manage these)
			//    - Resource-specific state transitions (e.g., EFS replication protection)
			diff = resources.SuppressAWSManagedDiffs(resourceToken, &spec, diff, inputs)
			// 5. Apply this difference to the actual inputs (not a projection) that we have in state.
			newInputs = resources.ApplyDiff(inputs, diff)
		}

		newState = resources.CheckpointObject(newInputs, rawState)
	}
	// Store both outputs and inputs into the state checkpoint.
	checkpoint, err := plugin.MarshalProperties(
		newState,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	// Serialize and return the calculated inputs.
	inputsRecord, err := plugin.MarshalProperties(
		newInputs,
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

	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.properties", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse inputs for update")
	}

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.oldState", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}
	oldInputs := resources.ParseCheckpointObject(oldState)

	var outputs resource.PropertyMap
	id := req.GetId()
	resourceToken := string(urn.Type())

	if req.GetPreview() {
		if customResource, ok := p.customResources[resourceToken]; ok {
			outputs = customResource.PreviewCustomResourceOutputs()
		} else {
			spec, hasSpec := p.resourceMap.Resources[resourceToken]
			if !hasSpec {
				return nil, errors.Errorf("Resource type %s not found", resourceToken)
			}
			resourceTypeName := urn.Type().Name()
			outputs = pOutputs.PreviewOutputs(newInputs, p.resourceMap.Types, spec.Outputs, spec.ReadOnly, resourceTypeName, &oldState)
		}

		previewState, err := plugin.MarshalProperties(
			outputs,
			plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
		)
		if err != nil {
			return nil, err
		}
		return &pulumirpc.UpdateResponse{Properties: previewState}, nil
	}

	if customResource, ok := p.customResources[resourceToken]; ok {
		timeout := time.Duration(req.GetTimeout()) * time.Second
		// Custom resource
		outputs, err = customResource.Update(ctx, urn, id, newInputs, oldInputs, oldState, timeout)
		if err != nil {
			return nil, err
		}
	} else {
		// Standard resource
		spec, ok := p.resourceMap.Resources[resourceToken]
		if !ok {
			return nil, errors.Errorf("Resource type %s not found", resourceToken)
		}

		ops, err := resources.CalcPatch(oldInputs, newInputs, spec, p.resourceMap.Types)
		if err != nil {
			return nil, err
		}

		glog.V(9).Infof("%s.UpdateResource %q id %q state %+v", label, spec.CfType, id, ops)
		resourceState, err := p.ccc.Update(p.canceler.context, urn, spec.CfType, id, ops)
		if err != nil {
			return nil, err
		}
		rawOutputs := naming.CfnToSdk(resourceState)

		// Write-only properties are not returned in the outputs, so we assume they should have the same value we sent from the inputs.
		if len(spec.WriteOnly) > 0 {
			inputsMap := newInputs.Mappable()
			for _, writeOnlyProp := range spec.WriteOnly {
				if _, ok := rawOutputs[writeOnlyProp]; !ok {
					inputValue, ok := inputsMap[writeOnlyProp]
					if ok {
						rawOutputs[writeOnlyProp] = inputValue
					}
				}
			}
		}
		outputs = resources.CheckpointObject(newInputs, rawOutputs)
	}
	// Store both outputs and inputs into the state and return RPC checkpoint.
	checkpoint, err := plugin.MarshalProperties(
		outputs,
		plugin.MarshalOptions{Label: fmt.Sprintf("%s.checkpoint", label), KeepSecrets: true, KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{Properties: checkpoint}, nil
}

func (p *cfnProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Delete(%s)", p.name, urn)
	glog.V(9).Infof("%s executing", label)

	id := req.GetId()
	resourceToken := string(urn.Type())
	if customResource, ok := p.customResources[resourceToken]; ok {
		oldInputs, err := plugin.UnmarshalProperties(req.OldInputs, plugin.MarshalOptions{
			Label:        fmt.Sprintf("%s.properties", label),
			KeepUnknowns: true,
			RejectAssets: true,
			KeepSecrets:  true,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse inputs for delete")
		}

		// Retrieve the state.
		state, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
			Label: fmt.Sprintf("%s.state", label), KeepUnknowns: true, SkipNulls: true, KeepSecrets: true,
		})

		timeout := time.Duration(req.GetTimeout()) * time.Second
		err = customResource.Delete(ctx, urn, id, oldInputs, state, timeout)
		if err != nil {
			return nil, err
		}
	} else {
		// Standard resource
		spec, ok := p.resourceMap.Resources[resourceToken]
		if !ok {
			return nil, errors.Errorf("Resource type %s not found", resourceToken)
		}

		glog.V(9).Infof("%s.DeleteResource %q id %q", label, spec.CfType, id)
		err := p.ccc.Delete(p.canceler.context, urn, spec.CfType, id)
		if err != nil {
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

// pulumiUserAgentMiddleware adds a Pulumi-specific user-agent to the request middleware.
// Example: APN/1.1 (marketplace-identifier)
var pulumiUserAgentMiddleware = middleware.BuildMiddlewareFunc("PulumiUserAgent", func(
	ctx context.Context, input middleware.BuildInput, next middleware.BuildHandler,
) (
	out middleware.BuildOutput, middlewareMetadata middleware.Metadata, err error,
) {
	request, ok := input.Request.(*smithyhttp.Request)
	if !ok {
		return out, middlewareMetadata, fmt.Errorf("unknown transport type %T", input.Request)
	}

	const userAgentKey = "User-Agent"

	// Replace the whole user-agent string to include only the marketplace identifier
	newValue := fmt.Sprintf("APN/1.1 (%s)", PulumiAWSMarketplaceCode)

	request.Header.Set(userAgentKey, newValue)

	return next.HandleBuild(ctx, input)
})

func attachCustomMiddleware(stack *middleware.Stack) error {
	return stack.Build.Add(pulumiUserAgentMiddleware, middleware.After)
}

// normalizeFilePath expands home directory prefixes to a canonical path.
func normalizeFilePath(filePath string) (string, error) {
	if strings.HasPrefix(filePath, `~/`) {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		dir := usr.HomeDir
		filePath = filepath.Join(dir, filePath[2:])
	} else if strings.HasPrefix(filePath, `%USERPROFILE%\`) {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		dir := usr.HomeDir
		filePath = filepath.Join(dir, filePath[14:])
	}
	return filePath, nil
}

// partialError creates an error for resources that did not complete an operation in progress.
// The last known state of the object is included in the error so that it can be checkpointed.
func partialError(id string, err error, state *structpb.Struct, inputs *structpb.Struct) error {
	detail := pulumirpc.ErrorResourceInitFailed{
		Id:         id,
		Properties: state,
		Reasons:    []string{err.Error()},
		Inputs:     inputs,
	}
	return rpcerror.WithDetails(rpcerror.New(codes.Unknown, err.Error()), &detail)
}
