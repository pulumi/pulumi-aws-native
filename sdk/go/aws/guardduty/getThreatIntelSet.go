// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::ThreatIntelSet
func LookupThreatIntelSet(ctx *pulumi.Context, args *LookupThreatIntelSetArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupThreatIntelSetResult
	err := ctx.Invoke("aws-native:guardduty:getThreatIntelSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelSetArgs struct {
	// The unique ID of the detector of the GuardDuty account for which you want to create a `ThreatIntelSet` .
	//
	// To find the `detectorId` in the current Region, see the
	// Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
	DetectorId string `pulumi:"detectorId"`
	// The unique ID of the `threatIntelSet` .
	Id string `pulumi:"id"`
}

type LookupThreatIntelSetResult struct {
	// The unique ID of the `threatIntelSet` .
	Id *string `pulumi:"id"`
	// The URI of the file that contains the ThreatIntelSet.
	Location *string `pulumi:"location"`
	// A user-friendly ThreatIntelSet name displayed in all findings that are generated by activity that involves IP addresses included in this ThreatIntelSet.
	Name *string `pulumi:"name"`
	// The tags to be added to a new threat list resource. Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupThreatIntelSetOutput(ctx *pulumi.Context, args LookupThreatIntelSetOutputArgs, opts ...pulumi.InvokeOption) LookupThreatIntelSetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupThreatIntelSetResultOutput, error) {
			args := v.(LookupThreatIntelSetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:guardduty:getThreatIntelSet", args, LookupThreatIntelSetResultOutput{}, options).(LookupThreatIntelSetResultOutput), nil
		}).(LookupThreatIntelSetResultOutput)
}

type LookupThreatIntelSetOutputArgs struct {
	// The unique ID of the detector of the GuardDuty account for which you want to create a `ThreatIntelSet` .
	//
	// To find the `detectorId` in the current Region, see the
	// Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
	DetectorId pulumi.StringInput `pulumi:"detectorId"`
	// The unique ID of the `threatIntelSet` .
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupThreatIntelSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelSetArgs)(nil)).Elem()
}

type LookupThreatIntelSetResultOutput struct{ *pulumi.OutputState }

func (LookupThreatIntelSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelSetResult)(nil)).Elem()
}

func (o LookupThreatIntelSetResultOutput) ToLookupThreatIntelSetResultOutput() LookupThreatIntelSetResultOutput {
	return o
}

func (o LookupThreatIntelSetResultOutput) ToLookupThreatIntelSetResultOutputWithContext(ctx context.Context) LookupThreatIntelSetResultOutput {
	return o
}

// The unique ID of the `threatIntelSet` .
func (o LookupThreatIntelSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The URI of the file that contains the ThreatIntelSet.
func (o LookupThreatIntelSetResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// A user-friendly ThreatIntelSet name displayed in all findings that are generated by activity that involves IP addresses included in this ThreatIntelSet.
func (o LookupThreatIntelSetResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The tags to be added to a new threat list resource. Each tag consists of a key and an optional value, both of which you define.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
func (o LookupThreatIntelSetResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupThreatIntelSetResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThreatIntelSetResultOutput{})
}
