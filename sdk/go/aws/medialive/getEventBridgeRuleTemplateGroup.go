// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaLive::EventBridgeRuleTemplateGroup Resource Type
func LookupEventBridgeRuleTemplateGroup(ctx *pulumi.Context, args *LookupEventBridgeRuleTemplateGroupArgs, opts ...pulumi.InvokeOption) (*LookupEventBridgeRuleTemplateGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEventBridgeRuleTemplateGroupResult
	err := ctx.Invoke("aws-native:medialive:getEventBridgeRuleTemplateGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventBridgeRuleTemplateGroupArgs struct {
	Identifier string `pulumi:"identifier"`
}

type LookupEventBridgeRuleTemplateGroupResult struct {
	// An eventbridge rule template group's ARN (Amazon Resource Name)
	Arn *string `pulumi:"arn"`
	// The date and time of resource creation.
	CreatedAt *string `pulumi:"createdAt"`
	// A resource's optional description.
	Description *string `pulumi:"description"`
	// An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`
	Id         *string `pulumi:"id"`
	Identifier *string `pulumi:"identifier"`
	// The date and time of latest resource modification.
	ModifiedAt *string `pulumi:"modifiedAt"`
}

func LookupEventBridgeRuleTemplateGroupOutput(ctx *pulumi.Context, args LookupEventBridgeRuleTemplateGroupOutputArgs, opts ...pulumi.InvokeOption) LookupEventBridgeRuleTemplateGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventBridgeRuleTemplateGroupResult, error) {
			args := v.(LookupEventBridgeRuleTemplateGroupArgs)
			r, err := LookupEventBridgeRuleTemplateGroup(ctx, &args, opts...)
			var s LookupEventBridgeRuleTemplateGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventBridgeRuleTemplateGroupResultOutput)
}

type LookupEventBridgeRuleTemplateGroupOutputArgs struct {
	Identifier pulumi.StringInput `pulumi:"identifier"`
}

func (LookupEventBridgeRuleTemplateGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBridgeRuleTemplateGroupArgs)(nil)).Elem()
}

type LookupEventBridgeRuleTemplateGroupResultOutput struct{ *pulumi.OutputState }

func (LookupEventBridgeRuleTemplateGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBridgeRuleTemplateGroupResult)(nil)).Elem()
}

func (o LookupEventBridgeRuleTemplateGroupResultOutput) ToLookupEventBridgeRuleTemplateGroupResultOutput() LookupEventBridgeRuleTemplateGroupResultOutput {
	return o
}

func (o LookupEventBridgeRuleTemplateGroupResultOutput) ToLookupEventBridgeRuleTemplateGroupResultOutputWithContext(ctx context.Context) LookupEventBridgeRuleTemplateGroupResultOutput {
	return o
}

// An eventbridge rule template group's ARN (Amazon Resource Name)
func (o LookupEventBridgeRuleTemplateGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The date and time of resource creation.
func (o LookupEventBridgeRuleTemplateGroupResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateGroupResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// A resource's optional description.
func (o LookupEventBridgeRuleTemplateGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`
func (o LookupEventBridgeRuleTemplateGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEventBridgeRuleTemplateGroupResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateGroupResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// The date and time of latest resource modification.
func (o LookupEventBridgeRuleTemplateGroupResultOutput) ModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateGroupResult) *string { return v.ModifiedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventBridgeRuleTemplateGroupResultOutput{})
}