// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::CloudFrontOriginAccessIdentity
func LookupCloudFrontOriginAccessIdentity(ctx *pulumi.Context, args *LookupCloudFrontOriginAccessIdentityArgs, opts ...pulumi.InvokeOption) (*LookupCloudFrontOriginAccessIdentityResult, error) {
	var rv LookupCloudFrontOriginAccessIdentityResult
	err := ctx.Invoke("aws-native:cloudfront:getCloudFrontOriginAccessIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudFrontOriginAccessIdentityArgs struct {
	Id string `pulumi:"id"`
}

type LookupCloudFrontOriginAccessIdentityResult struct {
	CloudFrontOriginAccessIdentityConfig *CloudFrontOriginAccessIdentityConfig `pulumi:"cloudFrontOriginAccessIdentityConfig"`
	Id                                   *string                               `pulumi:"id"`
	S3CanonicalUserId                    *string                               `pulumi:"s3CanonicalUserId"`
}

func LookupCloudFrontOriginAccessIdentityOutput(ctx *pulumi.Context, args LookupCloudFrontOriginAccessIdentityOutputArgs, opts ...pulumi.InvokeOption) LookupCloudFrontOriginAccessIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudFrontOriginAccessIdentityResult, error) {
			args := v.(LookupCloudFrontOriginAccessIdentityArgs)
			r, err := LookupCloudFrontOriginAccessIdentity(ctx, &args, opts...)
			var s LookupCloudFrontOriginAccessIdentityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudFrontOriginAccessIdentityResultOutput)
}

type LookupCloudFrontOriginAccessIdentityOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCloudFrontOriginAccessIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudFrontOriginAccessIdentityArgs)(nil)).Elem()
}

type LookupCloudFrontOriginAccessIdentityResultOutput struct{ *pulumi.OutputState }

func (LookupCloudFrontOriginAccessIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudFrontOriginAccessIdentityResult)(nil)).Elem()
}

func (o LookupCloudFrontOriginAccessIdentityResultOutput) ToLookupCloudFrontOriginAccessIdentityResultOutput() LookupCloudFrontOriginAccessIdentityResultOutput {
	return o
}

func (o LookupCloudFrontOriginAccessIdentityResultOutput) ToLookupCloudFrontOriginAccessIdentityResultOutputWithContext(ctx context.Context) LookupCloudFrontOriginAccessIdentityResultOutput {
	return o
}

func (o LookupCloudFrontOriginAccessIdentityResultOutput) CloudFrontOriginAccessIdentityConfig() CloudFrontOriginAccessIdentityConfigPtrOutput {
	return o.ApplyT(func(v LookupCloudFrontOriginAccessIdentityResult) *CloudFrontOriginAccessIdentityConfig {
		return v.CloudFrontOriginAccessIdentityConfig
	}).(CloudFrontOriginAccessIdentityConfigPtrOutput)
}

func (o LookupCloudFrontOriginAccessIdentityResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudFrontOriginAccessIdentityResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCloudFrontOriginAccessIdentityResultOutput) S3CanonicalUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudFrontOriginAccessIdentityResult) *string { return v.S3CanonicalUserId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudFrontOriginAccessIdentityResultOutput{})
}