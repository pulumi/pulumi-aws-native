// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaTailor::VodSource Resource Type
func LookupVodSource(ctx *pulumi.Context, args *LookupVodSourceArgs, opts ...pulumi.InvokeOption) (*LookupVodSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVodSourceResult
	err := ctx.Invoke("aws-native:mediatailor:getVodSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVodSourceArgs struct {
	SourceLocationName string `pulumi:"sourceLocationName"`
	VodSourceName      string `pulumi:"vodSourceName"`
}

type LookupVodSourceResult struct {
	// <p>The ARN of the VOD source.</p>
	Arn *string `pulumi:"arn"`
	// <p>A list of HTTP package configuration parameters for this VOD source.</p>
	HttpPackageConfigurations []VodSourceHttpPackageConfiguration `pulumi:"httpPackageConfigurations"`
	// The tags to assign to the VOD source.
	Tags []VodSourceTag `pulumi:"tags"`
}

func LookupVodSourceOutput(ctx *pulumi.Context, args LookupVodSourceOutputArgs, opts ...pulumi.InvokeOption) LookupVodSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVodSourceResult, error) {
			args := v.(LookupVodSourceArgs)
			r, err := LookupVodSource(ctx, &args, opts...)
			var s LookupVodSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVodSourceResultOutput)
}

type LookupVodSourceOutputArgs struct {
	SourceLocationName pulumi.StringInput `pulumi:"sourceLocationName"`
	VodSourceName      pulumi.StringInput `pulumi:"vodSourceName"`
}

func (LookupVodSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVodSourceArgs)(nil)).Elem()
}

type LookupVodSourceResultOutput struct{ *pulumi.OutputState }

func (LookupVodSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVodSourceResult)(nil)).Elem()
}

func (o LookupVodSourceResultOutput) ToLookupVodSourceResultOutput() LookupVodSourceResultOutput {
	return o
}

func (o LookupVodSourceResultOutput) ToLookupVodSourceResultOutputWithContext(ctx context.Context) LookupVodSourceResultOutput {
	return o
}

// <p>The ARN of the VOD source.</p>
func (o LookupVodSourceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVodSourceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>A list of HTTP package configuration parameters for this VOD source.</p>
func (o LookupVodSourceResultOutput) HttpPackageConfigurations() VodSourceHttpPackageConfigurationArrayOutput {
	return o.ApplyT(func(v LookupVodSourceResult) []VodSourceHttpPackageConfiguration { return v.HttpPackageConfigurations }).(VodSourceHttpPackageConfigurationArrayOutput)
}

// The tags to assign to the VOD source.
func (o LookupVodSourceResultOutput) Tags() VodSourceTagArrayOutput {
	return o.ApplyT(func(v LookupVodSourceResult) []VodSourceTag { return v.Tags }).(VodSourceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVodSourceResultOutput{})
}