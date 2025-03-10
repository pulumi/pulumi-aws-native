// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synthetics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Synthetics::Canary
func LookupCanary(ctx *pulumi.Context, args *LookupCanaryArgs, opts ...pulumi.InvokeOption) (*LookupCanaryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCanaryResult
	err := ctx.Invoke("aws-native:synthetics:getCanary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCanaryArgs struct {
	// Name of the canary.
	Name string `pulumi:"name"`
}

type LookupCanaryResult struct {
	// Provide artifact configuration
	ArtifactConfig *CanaryArtifactConfig `pulumi:"artifactConfig"`
	// Provide the s3 bucket output location for test results
	ArtifactS3Location *string `pulumi:"artifactS3Location"`
	// Provide the canary script source
	Code *CanaryCode `pulumi:"code"`
	// Lambda Execution role used to run your canaries
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Retention period of failed canary runs represented in number of days
	FailureRetentionPeriod *int `pulumi:"failureRetentionPeriod"`
	// Id of the canary
	Id *string `pulumi:"id"`
	// Setting to control if provisioned resources created by Synthetics are deleted alongside the canary. Default is AUTOMATIC.
	ProvisionedResourceCleanup *CanaryProvisionedResourceCleanup `pulumi:"provisionedResourceCleanup"`
	// Provide canary run configuration
	RunConfig *CanaryRunConfig `pulumi:"runConfig"`
	// Runtime version of Synthetics Library
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// Frequency to run your canaries
	Schedule *CanarySchedule `pulumi:"schedule"`
	// State of the canary
	State *string `pulumi:"state"`
	// Retention period of successful canary runs represented in number of days
	SuccessRetentionPeriod *int `pulumi:"successRetentionPeriod"`
	// The list of key-value pairs that are associated with the canary.
	Tags []aws.Tag `pulumi:"tags"`
	// Provide VPC Configuration if enabled.
	VpcConfig *CanaryVpcConfig `pulumi:"vpcConfig"`
}

func LookupCanaryOutput(ctx *pulumi.Context, args LookupCanaryOutputArgs, opts ...pulumi.InvokeOption) LookupCanaryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCanaryResultOutput, error) {
			args := v.(LookupCanaryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:synthetics:getCanary", args, LookupCanaryResultOutput{}, options).(LookupCanaryResultOutput), nil
		}).(LookupCanaryResultOutput)
}

type LookupCanaryOutputArgs struct {
	// Name of the canary.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupCanaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCanaryArgs)(nil)).Elem()
}

type LookupCanaryResultOutput struct{ *pulumi.OutputState }

func (LookupCanaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCanaryResult)(nil)).Elem()
}

func (o LookupCanaryResultOutput) ToLookupCanaryResultOutput() LookupCanaryResultOutput {
	return o
}

func (o LookupCanaryResultOutput) ToLookupCanaryResultOutputWithContext(ctx context.Context) LookupCanaryResultOutput {
	return o
}

// Provide artifact configuration
func (o LookupCanaryResultOutput) ArtifactConfig() CanaryArtifactConfigPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryArtifactConfig { return v.ArtifactConfig }).(CanaryArtifactConfigPtrOutput)
}

// Provide the s3 bucket output location for test results
func (o LookupCanaryResultOutput) ArtifactS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.ArtifactS3Location }).(pulumi.StringPtrOutput)
}

// Provide the canary script source
func (o LookupCanaryResultOutput) Code() CanaryCodePtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryCode { return v.Code }).(CanaryCodePtrOutput)
}

// Lambda Execution role used to run your canaries
func (o LookupCanaryResultOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

// Retention period of failed canary runs represented in number of days
func (o LookupCanaryResultOutput) FailureRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *int { return v.FailureRetentionPeriod }).(pulumi.IntPtrOutput)
}

// Id of the canary
func (o LookupCanaryResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Setting to control if provisioned resources created by Synthetics are deleted alongside the canary. Default is AUTOMATIC.
func (o LookupCanaryResultOutput) ProvisionedResourceCleanup() CanaryProvisionedResourceCleanupPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryProvisionedResourceCleanup { return v.ProvisionedResourceCleanup }).(CanaryProvisionedResourceCleanupPtrOutput)
}

// Provide canary run configuration
func (o LookupCanaryResultOutput) RunConfig() CanaryRunConfigPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryRunConfig { return v.RunConfig }).(CanaryRunConfigPtrOutput)
}

// Runtime version of Synthetics Library
func (o LookupCanaryResultOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

// Frequency to run your canaries
func (o LookupCanaryResultOutput) Schedule() CanarySchedulePtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanarySchedule { return v.Schedule }).(CanarySchedulePtrOutput)
}

// State of the canary
func (o LookupCanaryResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Retention period of successful canary runs represented in number of days
func (o LookupCanaryResultOutput) SuccessRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *int { return v.SuccessRetentionPeriod }).(pulumi.IntPtrOutput)
}

// The list of key-value pairs that are associated with the canary.
func (o LookupCanaryResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupCanaryResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Provide VPC Configuration if enabled.
func (o LookupCanaryResultOutput) VpcConfig() CanaryVpcConfigPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryVpcConfig { return v.VpcConfig }).(CanaryVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCanaryResultOutput{})
}
