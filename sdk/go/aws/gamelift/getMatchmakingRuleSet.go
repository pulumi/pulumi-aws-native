// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GameLift::MatchmakingRuleSet
func LookupMatchmakingRuleSet(ctx *pulumi.Context, args *LookupMatchmakingRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupMatchmakingRuleSetResult, error) {
	var rv LookupMatchmakingRuleSetResult
	err := ctx.Invoke("aws-native:gamelift:getMatchmakingRuleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMatchmakingRuleSetArgs struct {
	Id string `pulumi:"id"`
}

type LookupMatchmakingRuleSetResult struct {
	Arn  *string                 `pulumi:"arn"`
	Id   *string                 `pulumi:"id"`
	Tags []MatchmakingRuleSetTag `pulumi:"tags"`
}

func LookupMatchmakingRuleSetOutput(ctx *pulumi.Context, args LookupMatchmakingRuleSetOutputArgs, opts ...pulumi.InvokeOption) LookupMatchmakingRuleSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMatchmakingRuleSetResult, error) {
			args := v.(LookupMatchmakingRuleSetArgs)
			r, err := LookupMatchmakingRuleSet(ctx, &args, opts...)
			var s LookupMatchmakingRuleSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMatchmakingRuleSetResultOutput)
}

type LookupMatchmakingRuleSetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMatchmakingRuleSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMatchmakingRuleSetArgs)(nil)).Elem()
}

type LookupMatchmakingRuleSetResultOutput struct{ *pulumi.OutputState }

func (LookupMatchmakingRuleSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMatchmakingRuleSetResult)(nil)).Elem()
}

func (o LookupMatchmakingRuleSetResultOutput) ToLookupMatchmakingRuleSetResultOutput() LookupMatchmakingRuleSetResultOutput {
	return o
}

func (o LookupMatchmakingRuleSetResultOutput) ToLookupMatchmakingRuleSetResultOutputWithContext(ctx context.Context) LookupMatchmakingRuleSetResultOutput {
	return o
}

func (o LookupMatchmakingRuleSetResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingRuleSetResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingRuleSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMatchmakingRuleSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMatchmakingRuleSetResultOutput) Tags() MatchmakingRuleSetTagArrayOutput {
	return o.ApplyT(func(v LookupMatchmakingRuleSetResult) []MatchmakingRuleSetTag { return v.Tags }).(MatchmakingRuleSetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMatchmakingRuleSetResultOutput{})
}