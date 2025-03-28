// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::InferenceComponent
func LookupInferenceComponent(ctx *pulumi.Context, args *LookupInferenceComponentArgs, opts ...pulumi.InvokeOption) (*LookupInferenceComponentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInferenceComponentResult
	err := ctx.Invoke("aws-native:sagemaker:getInferenceComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInferenceComponentArgs struct {
	// The Amazon Resource Name (ARN) of the inference component.
	InferenceComponentArn string `pulumi:"inferenceComponentArn"`
}

type LookupInferenceComponentResult struct {
	// The time when the inference component was created.
	CreationTime *string `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) of the endpoint that hosts the inference component.
	EndpointArn *string `pulumi:"endpointArn"`
	// The name of the endpoint that hosts the inference component.
	EndpointName  *string `pulumi:"endpointName"`
	FailureReason *string `pulumi:"failureReason"`
	// The Amazon Resource Name (ARN) of the inference component.
	InferenceComponentArn *string `pulumi:"inferenceComponentArn"`
	// The name of the inference component.
	InferenceComponentName *string `pulumi:"inferenceComponentName"`
	// The status of the inference component.
	InferenceComponentStatus *InferenceComponentStatus `pulumi:"inferenceComponentStatus"`
	// The time when the inference component was last updated.
	LastModifiedTime *string                          `pulumi:"lastModifiedTime"`
	RuntimeConfig    *InferenceComponentRuntimeConfig `pulumi:"runtimeConfig"`
	Specification    *InferenceComponentSpecification `pulumi:"specification"`
	Tags             []aws.Tag                        `pulumi:"tags"`
	// The name of the production variant that hosts the inference component.
	VariantName *string `pulumi:"variantName"`
}

func LookupInferenceComponentOutput(ctx *pulumi.Context, args LookupInferenceComponentOutputArgs, opts ...pulumi.InvokeOption) LookupInferenceComponentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInferenceComponentResultOutput, error) {
			args := v.(LookupInferenceComponentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:sagemaker:getInferenceComponent", args, LookupInferenceComponentResultOutput{}, options).(LookupInferenceComponentResultOutput), nil
		}).(LookupInferenceComponentResultOutput)
}

type LookupInferenceComponentOutputArgs struct {
	// The Amazon Resource Name (ARN) of the inference component.
	InferenceComponentArn pulumi.StringInput `pulumi:"inferenceComponentArn"`
}

func (LookupInferenceComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceComponentArgs)(nil)).Elem()
}

type LookupInferenceComponentResultOutput struct{ *pulumi.OutputState }

func (LookupInferenceComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceComponentResult)(nil)).Elem()
}

func (o LookupInferenceComponentResultOutput) ToLookupInferenceComponentResultOutput() LookupInferenceComponentResultOutput {
	return o
}

func (o LookupInferenceComponentResultOutput) ToLookupInferenceComponentResultOutputWithContext(ctx context.Context) LookupInferenceComponentResultOutput {
	return o
}

// The time when the inference component was created.
func (o LookupInferenceComponentResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the endpoint that hosts the inference component.
func (o LookupInferenceComponentResultOutput) EndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.EndpointArn }).(pulumi.StringPtrOutput)
}

// The name of the endpoint that hosts the inference component.
func (o LookupInferenceComponentResultOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the inference component.
func (o LookupInferenceComponentResultOutput) InferenceComponentArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.InferenceComponentArn }).(pulumi.StringPtrOutput)
}

// The name of the inference component.
func (o LookupInferenceComponentResultOutput) InferenceComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.InferenceComponentName }).(pulumi.StringPtrOutput)
}

// The status of the inference component.
func (o LookupInferenceComponentResultOutput) InferenceComponentStatus() InferenceComponentStatusPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *InferenceComponentStatus { return v.InferenceComponentStatus }).(InferenceComponentStatusPtrOutput)
}

// The time when the inference component was last updated.
func (o LookupInferenceComponentResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) RuntimeConfig() InferenceComponentRuntimeConfigPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *InferenceComponentRuntimeConfig { return v.RuntimeConfig }).(InferenceComponentRuntimeConfigPtrOutput)
}

func (o LookupInferenceComponentResultOutput) Specification() InferenceComponentSpecificationPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *InferenceComponentSpecification { return v.Specification }).(InferenceComponentSpecificationPtrOutput)
}

func (o LookupInferenceComponentResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The name of the production variant that hosts the inference component.
func (o LookupInferenceComponentResultOutput) VariantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.VariantName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInferenceComponentResultOutput{})
}
