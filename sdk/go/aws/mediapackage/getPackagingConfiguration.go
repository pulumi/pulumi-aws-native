// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaPackage::PackagingConfiguration
func LookupPackagingConfiguration(ctx *pulumi.Context, args *LookupPackagingConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupPackagingConfigurationResult, error) {
	var rv LookupPackagingConfigurationResult
	err := ctx.Invoke("aws-native:mediapackage:getPackagingConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPackagingConfigurationArgs struct {
	// The ID of the PackagingConfiguration.
	Id string `pulumi:"id"`
}

type LookupPackagingConfigurationResult struct {
	// The ARN of the PackagingConfiguration.
	Arn *string `pulumi:"arn"`
	// A CMAF packaging configuration.
	CmafPackage *PackagingConfigurationCmafPackage `pulumi:"cmafPackage"`
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *PackagingConfigurationDashPackage `pulumi:"dashPackage"`
	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *PackagingConfigurationHlsPackage `pulumi:"hlsPackage"`
	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *PackagingConfigurationMssPackage `pulumi:"mssPackage"`
	// The ID of a PackagingGroup.
	PackagingGroupId *string `pulumi:"packagingGroupId"`
	// A collection of tags associated with a resource
	Tags []PackagingConfigurationTag `pulumi:"tags"`
}

func LookupPackagingConfigurationOutput(ctx *pulumi.Context, args LookupPackagingConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupPackagingConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPackagingConfigurationResult, error) {
			args := v.(LookupPackagingConfigurationArgs)
			r, err := LookupPackagingConfiguration(ctx, &args, opts...)
			var s LookupPackagingConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPackagingConfigurationResultOutput)
}

type LookupPackagingConfigurationOutputArgs struct {
	// The ID of the PackagingConfiguration.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPackagingConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackagingConfigurationArgs)(nil)).Elem()
}

type LookupPackagingConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupPackagingConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackagingConfigurationResult)(nil)).Elem()
}

func (o LookupPackagingConfigurationResultOutput) ToLookupPackagingConfigurationResultOutput() LookupPackagingConfigurationResultOutput {
	return o
}

func (o LookupPackagingConfigurationResultOutput) ToLookupPackagingConfigurationResultOutputWithContext(ctx context.Context) LookupPackagingConfigurationResultOutput {
	return o
}

// The ARN of the PackagingConfiguration.
func (o LookupPackagingConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A CMAF packaging configuration.
func (o LookupPackagingConfigurationResultOutput) CmafPackage() PackagingConfigurationCmafPackagePtrOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) *PackagingConfigurationCmafPackage { return v.CmafPackage }).(PackagingConfigurationCmafPackagePtrOutput)
}

// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
func (o LookupPackagingConfigurationResultOutput) DashPackage() PackagingConfigurationDashPackagePtrOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) *PackagingConfigurationDashPackage { return v.DashPackage }).(PackagingConfigurationDashPackagePtrOutput)
}

// An HTTP Live Streaming (HLS) packaging configuration.
func (o LookupPackagingConfigurationResultOutput) HlsPackage() PackagingConfigurationHlsPackagePtrOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) *PackagingConfigurationHlsPackage { return v.HlsPackage }).(PackagingConfigurationHlsPackagePtrOutput)
}

// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
func (o LookupPackagingConfigurationResultOutput) MssPackage() PackagingConfigurationMssPackagePtrOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) *PackagingConfigurationMssPackage { return v.MssPackage }).(PackagingConfigurationMssPackagePtrOutput)
}

// The ID of a PackagingGroup.
func (o LookupPackagingConfigurationResultOutput) PackagingGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) *string { return v.PackagingGroupId }).(pulumi.StringPtrOutput)
}

// A collection of tags associated with a resource
func (o LookupPackagingConfigurationResultOutput) Tags() PackagingConfigurationTagArrayOutput {
	return o.ApplyT(func(v LookupPackagingConfigurationResult) []PackagingConfigurationTag { return v.Tags }).(PackagingConfigurationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPackagingConfigurationResultOutput{})
}