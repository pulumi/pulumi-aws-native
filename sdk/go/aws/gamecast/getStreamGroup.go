// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamecast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::GameCast::StreamGroup Resource Type
func LookupStreamGroup(ctx *pulumi.Context, args *LookupStreamGroupArgs, opts ...pulumi.InvokeOption) (*LookupStreamGroupResult, error) {
	var rv LookupStreamGroupResult
	err := ctx.Invoke("aws-native:gamecast:getStreamGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamGroupArgs struct {
	// ARN of the resource.
	Arn string `pulumi:"arn"`
}

type LookupStreamGroupResult struct {
	// ARN of the resource.
	Arn                *string                        `pulumi:"arn"`
	DefaultApplication *StreamGroupDefaultApplication `pulumi:"defaultApplication"`
	// Descriptive label for the resource, not a unique ID.
	Description *string `pulumi:"description"`
	// DesiredCapacity is the target number of stream sessions customer sets.
	DesiredCapacity *int             `pulumi:"desiredCapacity"`
	Tags            *StreamGroupTags `pulumi:"tags"`
}

func LookupStreamGroupOutput(ctx *pulumi.Context, args LookupStreamGroupOutputArgs, opts ...pulumi.InvokeOption) LookupStreamGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamGroupResult, error) {
			args := v.(LookupStreamGroupArgs)
			r, err := LookupStreamGroup(ctx, &args, opts...)
			var s LookupStreamGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamGroupResultOutput)
}

type LookupStreamGroupOutputArgs struct {
	// ARN of the resource.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupStreamGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamGroupArgs)(nil)).Elem()
}

type LookupStreamGroupResultOutput struct{ *pulumi.OutputState }

func (LookupStreamGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamGroupResult)(nil)).Elem()
}

func (o LookupStreamGroupResultOutput) ToLookupStreamGroupResultOutput() LookupStreamGroupResultOutput {
	return o
}

func (o LookupStreamGroupResultOutput) ToLookupStreamGroupResultOutputWithContext(ctx context.Context) LookupStreamGroupResultOutput {
	return o
}

// ARN of the resource.
func (o LookupStreamGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupStreamGroupResultOutput) DefaultApplication() StreamGroupDefaultApplicationPtrOutput {
	return o.ApplyT(func(v LookupStreamGroupResult) *StreamGroupDefaultApplication { return v.DefaultApplication }).(StreamGroupDefaultApplicationPtrOutput)
}

// Descriptive label for the resource, not a unique ID.
func (o LookupStreamGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// DesiredCapacity is the target number of stream sessions customer sets.
func (o LookupStreamGroupResultOutput) DesiredCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamGroupResult) *int { return v.DesiredCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupStreamGroupResultOutput) Tags() StreamGroupTagsPtrOutput {
	return o.ApplyT(func(v LookupStreamGroupResult) *StreamGroupTags { return v.Tags }).(StreamGroupTagsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamGroupResultOutput{})
}