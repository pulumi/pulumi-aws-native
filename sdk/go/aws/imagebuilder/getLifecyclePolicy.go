// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::LifecyclePolicy
func LookupLifecyclePolicy(ctx *pulumi.Context, args *LookupLifecyclePolicyArgs, opts ...pulumi.InvokeOption) (*LookupLifecyclePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLifecyclePolicyResult
	err := ctx.Invoke("aws-native:imagebuilder:getLifecyclePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLifecyclePolicyArgs struct {
	// The Amazon Resource Name (ARN) of the lifecycle policy.
	Arn string `pulumi:"arn"`
}

type LookupLifecyclePolicyResult struct {
	// The Amazon Resource Name (ARN) of the lifecycle policy.
	Arn *string `pulumi:"arn"`
	// The description of the lifecycle policy.
	Description *string `pulumi:"description"`
	// The execution role of the lifecycle policy.
	ExecutionRole *string `pulumi:"executionRole"`
	// The policy details of the lifecycle policy.
	PolicyDetails []LifecyclePolicyPolicyDetail `pulumi:"policyDetails"`
	// The resource selection of the lifecycle policy.
	ResourceSelection *LifecyclePolicyResourceSelection `pulumi:"resourceSelection"`
	// The resource type of the lifecycle policy.
	ResourceType *LifecyclePolicyResourceType `pulumi:"resourceType"`
	// The status of the lifecycle policy.
	Status *LifecyclePolicyStatus `pulumi:"status"`
	// The tags associated with the lifecycle policy.
	Tags map[string]string `pulumi:"tags"`
}

func LookupLifecyclePolicyOutput(ctx *pulumi.Context, args LookupLifecyclePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLifecyclePolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLifecyclePolicyResultOutput, error) {
			args := v.(LookupLifecyclePolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:imagebuilder:getLifecyclePolicy", args, LookupLifecyclePolicyResultOutput{}, options).(LookupLifecyclePolicyResultOutput), nil
		}).(LookupLifecyclePolicyResultOutput)
}

type LookupLifecyclePolicyOutputArgs struct {
	// The Amazon Resource Name (ARN) of the lifecycle policy.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupLifecyclePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLifecyclePolicyArgs)(nil)).Elem()
}

type LookupLifecyclePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLifecyclePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLifecyclePolicyResult)(nil)).Elem()
}

func (o LookupLifecyclePolicyResultOutput) ToLookupLifecyclePolicyResultOutput() LookupLifecyclePolicyResultOutput {
	return o
}

func (o LookupLifecyclePolicyResultOutput) ToLookupLifecyclePolicyResultOutputWithContext(ctx context.Context) LookupLifecyclePolicyResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The description of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The execution role of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) ExecutionRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.ExecutionRole }).(pulumi.StringPtrOutput)
}

// The policy details of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) PolicyDetails() LifecyclePolicyPolicyDetailArrayOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) []LifecyclePolicyPolicyDetail { return v.PolicyDetails }).(LifecyclePolicyPolicyDetailArrayOutput)
}

// The resource selection of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) ResourceSelection() LifecyclePolicyResourceSelectionPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *LifecyclePolicyResourceSelection { return v.ResourceSelection }).(LifecyclePolicyResourceSelectionPtrOutput)
}

// The resource type of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) ResourceType() LifecyclePolicyResourceTypePtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *LifecyclePolicyResourceType { return v.ResourceType }).(LifecyclePolicyResourceTypePtrOutput)
}

// The status of the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) Status() LifecyclePolicyStatusPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *LifecyclePolicyStatus { return v.Status }).(LifecyclePolicyStatusPtrOutput)
}

// The tags associated with the lifecycle policy.
func (o LookupLifecyclePolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLifecyclePolicyResultOutput{})
}
