// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::TopicRule
func LookupTopicRule(ctx *pulumi.Context, args *LookupTopicRuleArgs, opts ...pulumi.InvokeOption) (*LookupTopicRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTopicRuleResult
	err := ctx.Invoke("aws-native:iot:getTopicRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicRuleArgs struct {
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

type LookupTopicRuleResult struct {
	// The Amazon Resource Name (ARN) of the AWS IoT rule, such as `arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule` .
	Arn *string `pulumi:"arn"`
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []aws.Tag `pulumi:"tags"`
	// The rule payload.
	TopicRulePayload *TopicRulePayload `pulumi:"topicRulePayload"`
}

func LookupTopicRuleOutput(ctx *pulumi.Context, args LookupTopicRuleOutputArgs, opts ...pulumi.InvokeOption) LookupTopicRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTopicRuleResultOutput, error) {
			args := v.(LookupTopicRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:iot:getTopicRule", args, LookupTopicRuleResultOutput{}, options).(LookupTopicRuleResultOutput), nil
		}).(LookupTopicRuleResultOutput)
}

type LookupTopicRuleOutputArgs struct {
	// The name of the rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupTopicRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicRuleArgs)(nil)).Elem()
}

type LookupTopicRuleResultOutput struct{ *pulumi.OutputState }

func (LookupTopicRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicRuleResult)(nil)).Elem()
}

func (o LookupTopicRuleResultOutput) ToLookupTopicRuleResultOutput() LookupTopicRuleResultOutput {
	return o
}

func (o LookupTopicRuleResultOutput) ToLookupTopicRuleResultOutputWithContext(ctx context.Context) LookupTopicRuleResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the AWS IoT rule, such as `arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule` .
func (o LookupTopicRuleResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Metadata which can be used to manage the topic rule.
//
// > For URI Request parameters use format: ...key1=value1&key2=value2...
// >
// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
// >
// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
func (o LookupTopicRuleResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupTopicRuleResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The rule payload.
func (o LookupTopicRuleResultOutput) TopicRulePayload() TopicRulePayloadPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleResult) *TopicRulePayload { return v.TopicRulePayload }).(TopicRulePayloadPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicRuleResultOutput{})
}
