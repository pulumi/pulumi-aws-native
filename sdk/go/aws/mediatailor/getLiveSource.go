// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::MediaTailor::LiveSource Resource Type
func LookupLiveSource(ctx *pulumi.Context, args *LookupLiveSourceArgs, opts ...pulumi.InvokeOption) (*LookupLiveSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLiveSourceResult
	err := ctx.Invoke("aws-native:mediatailor:getLiveSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveSourceArgs struct {
	LiveSourceName     string `pulumi:"liveSourceName"`
	SourceLocationName string `pulumi:"sourceLocationName"`
}

type LookupLiveSourceResult struct {
	// <p>The ARN of the live source.</p>
	Arn *string `pulumi:"arn"`
	// <p>A list of HTTP package configuration parameters for this live source.</p>
	HttpPackageConfigurations []LiveSourceHttpPackageConfiguration `pulumi:"httpPackageConfigurations"`
	// The tags to assign to the live source.
	Tags []LiveSourceTag `pulumi:"tags"`
}

func LookupLiveSourceOutput(ctx *pulumi.Context, args LookupLiveSourceOutputArgs, opts ...pulumi.InvokeOption) LookupLiveSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLiveSourceResult, error) {
			args := v.(LookupLiveSourceArgs)
			r, err := LookupLiveSource(ctx, &args, opts...)
			var s LookupLiveSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLiveSourceResultOutput)
}

type LookupLiveSourceOutputArgs struct {
	LiveSourceName     pulumi.StringInput `pulumi:"liveSourceName"`
	SourceLocationName pulumi.StringInput `pulumi:"sourceLocationName"`
}

func (LookupLiveSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveSourceArgs)(nil)).Elem()
}

type LookupLiveSourceResultOutput struct{ *pulumi.OutputState }

func (LookupLiveSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveSourceResult)(nil)).Elem()
}

func (o LookupLiveSourceResultOutput) ToLookupLiveSourceResultOutput() LookupLiveSourceResultOutput {
	return o
}

func (o LookupLiveSourceResultOutput) ToLookupLiveSourceResultOutputWithContext(ctx context.Context) LookupLiveSourceResultOutput {
	return o
}

func (o LookupLiveSourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLiveSourceResult] {
	return pulumix.Output[LookupLiveSourceResult]{
		OutputState: o.OutputState,
	}
}

// <p>The ARN of the live source.</p>
func (o LookupLiveSourceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveSourceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>A list of HTTP package configuration parameters for this live source.</p>
func (o LookupLiveSourceResultOutput) HttpPackageConfigurations() LiveSourceHttpPackageConfigurationArrayOutput {
	return o.ApplyT(func(v LookupLiveSourceResult) []LiveSourceHttpPackageConfiguration {
		return v.HttpPackageConfigurations
	}).(LiveSourceHttpPackageConfigurationArrayOutput)
}

// The tags to assign to the live source.
func (o LookupLiveSourceResultOutput) Tags() LiveSourceTagArrayOutput {
	return o.ApplyT(func(v LookupLiveSourceResult) []LiveSourceTag { return v.Tags }).(LiveSourceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLiveSourceResultOutput{})
}