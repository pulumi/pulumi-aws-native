// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Logs::DeliverySource.
func LookupDeliverySource(ctx *pulumi.Context, args *LookupDeliverySourceArgs, opts ...pulumi.InvokeOption) (*LookupDeliverySourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeliverySourceResult
	err := ctx.Invoke("aws-native:logs:getDeliverySource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeliverySourceArgs struct {
	// The unique name of the Log source.
	Name string `pulumi:"name"`
}

type LookupDeliverySourceResult struct {
	// The ARN of the Aqueduct Source.
	Arn *string `pulumi:"arn"`
	// The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.
	LogType *string `pulumi:"logType"`
	// List of ARN of the resource that will be sending the logs
	ResourceArns []string `pulumi:"resourceArns"`
	// The service generating the log
	Service *string `pulumi:"service"`
	// An array of key-value pairs to apply to this resource.
	Tags []DeliverySourceTag `pulumi:"tags"`
}

func LookupDeliverySourceOutput(ctx *pulumi.Context, args LookupDeliverySourceOutputArgs, opts ...pulumi.InvokeOption) LookupDeliverySourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeliverySourceResult, error) {
			args := v.(LookupDeliverySourceArgs)
			r, err := LookupDeliverySource(ctx, &args, opts...)
			var s LookupDeliverySourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeliverySourceResultOutput)
}

type LookupDeliverySourceOutputArgs struct {
	// The unique name of the Log source.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDeliverySourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliverySourceArgs)(nil)).Elem()
}

type LookupDeliverySourceResultOutput struct{ *pulumi.OutputState }

func (LookupDeliverySourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliverySourceResult)(nil)).Elem()
}

func (o LookupDeliverySourceResultOutput) ToLookupDeliverySourceResultOutput() LookupDeliverySourceResultOutput {
	return o
}

func (o LookupDeliverySourceResultOutput) ToLookupDeliverySourceResultOutputWithContext(ctx context.Context) LookupDeliverySourceResultOutput {
	return o
}

func (o LookupDeliverySourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDeliverySourceResult] {
	return pulumix.Output[LookupDeliverySourceResult]{
		OutputState: o.OutputState,
	}
}

// The ARN of the Aqueduct Source.
func (o LookupDeliverySourceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeliverySourceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.
func (o LookupDeliverySourceResultOutput) LogType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeliverySourceResult) *string { return v.LogType }).(pulumi.StringPtrOutput)
}

// List of ARN of the resource that will be sending the logs
func (o LookupDeliverySourceResultOutput) ResourceArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeliverySourceResult) []string { return v.ResourceArns }).(pulumi.StringArrayOutput)
}

// The service generating the log
func (o LookupDeliverySourceResultOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeliverySourceResult) *string { return v.Service }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupDeliverySourceResultOutput) Tags() DeliverySourceTagArrayOutput {
	return o.ApplyT(func(v LookupDeliverySourceResult) []DeliverySourceTag { return v.Tags }).(DeliverySourceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeliverySourceResultOutput{})
}