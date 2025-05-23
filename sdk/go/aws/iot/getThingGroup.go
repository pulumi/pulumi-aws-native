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

// Resource Type definition for AWS::IoT::ThingGroup
func LookupThingGroup(ctx *pulumi.Context, args *LookupThingGroupArgs, opts ...pulumi.InvokeOption) (*LookupThingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupThingGroupResult
	err := ctx.Invoke("aws-native:iot:getThingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThingGroupArgs struct {
	// The thing group name.
	ThingGroupName string `pulumi:"thingGroupName"`
}

type LookupThingGroupResult struct {
	// The thing group ARN.
	Arn *string `pulumi:"arn"`
	// The thing group ID.
	Id *string `pulumi:"id"`
	// The dynamic thing group search query string.
	//
	// The `queryString` attribute *is* required for `CreateDynamicThingGroup` . The `queryString` attribute *is not* required for `CreateThingGroup` .
	QueryString *string `pulumi:"queryString"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// Thing group properties.
	ThingGroupProperties *ThingGroupPropertiesProperties `pulumi:"thingGroupProperties"`
}

func LookupThingGroupOutput(ctx *pulumi.Context, args LookupThingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupThingGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupThingGroupResultOutput, error) {
			args := v.(LookupThingGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:iot:getThingGroup", args, LookupThingGroupResultOutput{}, options).(LookupThingGroupResultOutput), nil
		}).(LookupThingGroupResultOutput)
}

type LookupThingGroupOutputArgs struct {
	// The thing group name.
	ThingGroupName pulumi.StringInput `pulumi:"thingGroupName"`
}

func (LookupThingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThingGroupArgs)(nil)).Elem()
}

type LookupThingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupThingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThingGroupResult)(nil)).Elem()
}

func (o LookupThingGroupResultOutput) ToLookupThingGroupResultOutput() LookupThingGroupResultOutput {
	return o
}

func (o LookupThingGroupResultOutput) ToLookupThingGroupResultOutputWithContext(ctx context.Context) LookupThingGroupResultOutput {
	return o
}

// The thing group ARN.
func (o LookupThingGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThingGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The thing group ID.
func (o LookupThingGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThingGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The dynamic thing group search query string.
//
// The `queryString` attribute *is* required for `CreateDynamicThingGroup` . The `queryString` attribute *is not* required for `CreateThingGroup` .
func (o LookupThingGroupResultOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThingGroupResult) *string { return v.QueryString }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupThingGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupThingGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Thing group properties.
func (o LookupThingGroupResultOutput) ThingGroupProperties() ThingGroupPropertiesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupThingGroupResult) *ThingGroupPropertiesProperties { return v.ThingGroupProperties }).(ThingGroupPropertiesPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThingGroupResultOutput{})
}
