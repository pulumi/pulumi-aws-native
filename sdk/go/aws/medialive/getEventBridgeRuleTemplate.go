// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaLive::EventBridgeRuleTemplate Resource Type
func LookupEventBridgeRuleTemplate(ctx *pulumi.Context, args *LookupEventBridgeRuleTemplateArgs, opts ...pulumi.InvokeOption) (*LookupEventBridgeRuleTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEventBridgeRuleTemplateResult
	err := ctx.Invoke("aws-native:medialive:getEventBridgeRuleTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventBridgeRuleTemplateArgs struct {
	// Placeholder documentation for __string
	Identifier string `pulumi:"identifier"`
}

type LookupEventBridgeRuleTemplateResult struct {
	// An eventbridge rule template's ARN (Amazon Resource Name)
	Arn *string `pulumi:"arn"`
	// Placeholder documentation for __timestampIso8601
	CreatedAt *string `pulumi:"createdAt"`
	// A resource's optional description.
	Description *string `pulumi:"description"`
	// Placeholder documentation for __listOfEventBridgeRuleTemplateTarget
	EventTargets []EventBridgeRuleTemplateTarget `pulumi:"eventTargets"`
	// The type of event to match with the rule.
	EventType *EventBridgeRuleTemplateEventType `pulumi:"eventType"`
	// An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`
	GroupId *string `pulumi:"groupId"`
	// An eventbridge rule template's id. AWS provided templates have ids that start with `aws-`
	Id *string `pulumi:"id"`
	// Placeholder documentation for __string
	Identifier *string `pulumi:"identifier"`
	// Placeholder documentation for __timestampIso8601
	ModifiedAt *string `pulumi:"modifiedAt"`
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	Name *string `pulumi:"name"`
}

func LookupEventBridgeRuleTemplateOutput(ctx *pulumi.Context, args LookupEventBridgeRuleTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupEventBridgeRuleTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEventBridgeRuleTemplateResultOutput, error) {
			args := v.(LookupEventBridgeRuleTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:medialive:getEventBridgeRuleTemplate", args, LookupEventBridgeRuleTemplateResultOutput{}, options).(LookupEventBridgeRuleTemplateResultOutput), nil
		}).(LookupEventBridgeRuleTemplateResultOutput)
}

type LookupEventBridgeRuleTemplateOutputArgs struct {
	// Placeholder documentation for __string
	Identifier pulumi.StringInput `pulumi:"identifier"`
}

func (LookupEventBridgeRuleTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBridgeRuleTemplateArgs)(nil)).Elem()
}

type LookupEventBridgeRuleTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupEventBridgeRuleTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBridgeRuleTemplateResult)(nil)).Elem()
}

func (o LookupEventBridgeRuleTemplateResultOutput) ToLookupEventBridgeRuleTemplateResultOutput() LookupEventBridgeRuleTemplateResultOutput {
	return o
}

func (o LookupEventBridgeRuleTemplateResultOutput) ToLookupEventBridgeRuleTemplateResultOutputWithContext(ctx context.Context) LookupEventBridgeRuleTemplateResultOutput {
	return o
}

// An eventbridge rule template's ARN (Amazon Resource Name)
func (o LookupEventBridgeRuleTemplateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Placeholder documentation for __timestampIso8601
func (o LookupEventBridgeRuleTemplateResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// A resource's optional description.
func (o LookupEventBridgeRuleTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Placeholder documentation for __listOfEventBridgeRuleTemplateTarget
func (o LookupEventBridgeRuleTemplateResultOutput) EventTargets() EventBridgeRuleTemplateTargetArrayOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) []EventBridgeRuleTemplateTarget { return v.EventTargets }).(EventBridgeRuleTemplateTargetArrayOutput)
}

// The type of event to match with the rule.
func (o LookupEventBridgeRuleTemplateResultOutput) EventType() EventBridgeRuleTemplateEventTypePtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *EventBridgeRuleTemplateEventType { return v.EventType }).(EventBridgeRuleTemplateEventTypePtrOutput)
}

// An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`
func (o LookupEventBridgeRuleTemplateResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// An eventbridge rule template's id. AWS provided templates have ids that start with `aws-`
func (o LookupEventBridgeRuleTemplateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Placeholder documentation for __string
func (o LookupEventBridgeRuleTemplateResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Placeholder documentation for __timestampIso8601
func (o LookupEventBridgeRuleTemplateResultOutput) ModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.ModifiedAt }).(pulumi.StringPtrOutput)
}

// A resource's name. Names must be unique within the scope of a resource type in a specific region.
func (o LookupEventBridgeRuleTemplateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventBridgeRuleTemplateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventBridgeRuleTemplateResultOutput{})
}
