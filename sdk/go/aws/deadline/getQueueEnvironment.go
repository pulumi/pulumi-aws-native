// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deadline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Deadline::QueueEnvironment Resource Type
func LookupQueueEnvironment(ctx *pulumi.Context, args *LookupQueueEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupQueueEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQueueEnvironmentResult
	err := ctx.Invoke("aws-native:deadline:getQueueEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueEnvironmentArgs struct {
	FarmId             string `pulumi:"farmId"`
	QueueEnvironmentId string `pulumi:"queueEnvironmentId"`
	QueueId            string `pulumi:"queueId"`
}

type LookupQueueEnvironmentResult struct {
	Name               *string                                  `pulumi:"name"`
	Priority           *int                                     `pulumi:"priority"`
	QueueEnvironmentId *string                                  `pulumi:"queueEnvironmentId"`
	Template           *string                                  `pulumi:"template"`
	TemplateType       *QueueEnvironmentEnvironmentTemplateType `pulumi:"templateType"`
}

func LookupQueueEnvironmentOutput(ctx *pulumi.Context, args LookupQueueEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupQueueEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueEnvironmentResult, error) {
			args := v.(LookupQueueEnvironmentArgs)
			r, err := LookupQueueEnvironment(ctx, &args, opts...)
			var s LookupQueueEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueEnvironmentResultOutput)
}

type LookupQueueEnvironmentOutputArgs struct {
	FarmId             pulumi.StringInput `pulumi:"farmId"`
	QueueEnvironmentId pulumi.StringInput `pulumi:"queueEnvironmentId"`
	QueueId            pulumi.StringInput `pulumi:"queueId"`
}

func (LookupQueueEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueEnvironmentArgs)(nil)).Elem()
}

type LookupQueueEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupQueueEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueEnvironmentResult)(nil)).Elem()
}

func (o LookupQueueEnvironmentResultOutput) ToLookupQueueEnvironmentResultOutput() LookupQueueEnvironmentResultOutput {
	return o
}

func (o LookupQueueEnvironmentResultOutput) ToLookupQueueEnvironmentResultOutputWithContext(ctx context.Context) LookupQueueEnvironmentResultOutput {
	return o
}

func (o LookupQueueEnvironmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueEnvironmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupQueueEnvironmentResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupQueueEnvironmentResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupQueueEnvironmentResultOutput) QueueEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueEnvironmentResult) *string { return v.QueueEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupQueueEnvironmentResultOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueEnvironmentResult) *string { return v.Template }).(pulumi.StringPtrOutput)
}

func (o LookupQueueEnvironmentResultOutput) TemplateType() QueueEnvironmentEnvironmentTemplateTypePtrOutput {
	return o.ApplyT(func(v LookupQueueEnvironmentResult) *QueueEnvironmentEnvironmentTemplateType { return v.TemplateType }).(QueueEnvironmentEnvironmentTemplateTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueEnvironmentResultOutput{})
}