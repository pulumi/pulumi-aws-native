// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppIntegrations::EventIntegration
func LookupEventIntegration(ctx *pulumi.Context, args *LookupEventIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupEventIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEventIntegrationResult
	err := ctx.Invoke("aws-native:appintegrations:getEventIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventIntegrationArgs struct {
	// The name of the event integration.
	Name string `pulumi:"name"`
}

type LookupEventIntegrationResult struct {
	// The event integration description.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the event integration.
	EventIntegrationArn *string `pulumi:"eventIntegrationArn"`
	// The tags (keys and values) associated with the event integration.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupEventIntegrationOutput(ctx *pulumi.Context, args LookupEventIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupEventIntegrationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEventIntegrationResultOutput, error) {
			args := v.(LookupEventIntegrationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:appintegrations:getEventIntegration", args, LookupEventIntegrationResultOutput{}, options).(LookupEventIntegrationResultOutput), nil
		}).(LookupEventIntegrationResultOutput)
}

type LookupEventIntegrationOutputArgs struct {
	// The name of the event integration.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupEventIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventIntegrationArgs)(nil)).Elem()
}

type LookupEventIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupEventIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventIntegrationResult)(nil)).Elem()
}

func (o LookupEventIntegrationResultOutput) ToLookupEventIntegrationResultOutput() LookupEventIntegrationResultOutput {
	return o
}

func (o LookupEventIntegrationResultOutput) ToLookupEventIntegrationResultOutputWithContext(ctx context.Context) LookupEventIntegrationResultOutput {
	return o
}

// The event integration description.
func (o LookupEventIntegrationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the event integration.
func (o LookupEventIntegrationResultOutput) EventIntegrationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) *string { return v.EventIntegrationArn }).(pulumi.StringPtrOutput)
}

// The tags (keys and values) associated with the event integration.
func (o LookupEventIntegrationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventIntegrationResultOutput{})
}
