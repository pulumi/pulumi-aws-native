// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolDomain
func LookupUserPoolDomain(ctx *pulumi.Context, args *LookupUserPoolDomainArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserPoolDomainResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolDomainArgs struct {
	// The resource ID.
	Id string `pulumi:"id"`
}

type LookupUserPoolDomainResult struct {
	// The Amazon CloudFront endpoint that you use as the target of the alias that you set up with your Domain Name Service (DNS) provider.
	CloudFrontDistribution *string `pulumi:"cloudFrontDistribution"`
	// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application. Use this object to specify an SSL certificate that is managed by ACM.
	CustomDomainConfig *UserPoolDomainCustomDomainConfigType `pulumi:"customDomainConfig"`
	// The resource ID.
	Id *string `pulumi:"id"`
}

func LookupUserPoolDomainOutput(ctx *pulumi.Context, args LookupUserPoolDomainOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolDomainResultOutput, error) {
			args := v.(LookupUserPoolDomainArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupUserPoolDomainResult
			secret, err := ctx.InvokePackageRaw("aws-native:cognito:getUserPoolDomain", args, &rv, "", opts...)
			if err != nil {
				return LookupUserPoolDomainResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupUserPoolDomainResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupUserPoolDomainResultOutput), nil
			}
			return output, nil
		}).(LookupUserPoolDomainResultOutput)
}

type LookupUserPoolDomainOutputArgs struct {
	// The resource ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserPoolDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolDomainArgs)(nil)).Elem()
}

type LookupUserPoolDomainResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolDomainResult)(nil)).Elem()
}

func (o LookupUserPoolDomainResultOutput) ToLookupUserPoolDomainResultOutput() LookupUserPoolDomainResultOutput {
	return o
}

func (o LookupUserPoolDomainResultOutput) ToLookupUserPoolDomainResultOutputWithContext(ctx context.Context) LookupUserPoolDomainResultOutput {
	return o
}

// The Amazon CloudFront endpoint that you use as the target of the alias that you set up with your Domain Name Service (DNS) provider.
func (o LookupUserPoolDomainResultOutput) CloudFrontDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolDomainResult) *string { return v.CloudFrontDistribution }).(pulumi.StringPtrOutput)
}

// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application. Use this object to specify an SSL certificate that is managed by ACM.
func (o LookupUserPoolDomainResultOutput) CustomDomainConfig() UserPoolDomainCustomDomainConfigTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolDomainResult) *UserPoolDomainCustomDomainConfigType { return v.CustomDomainConfig }).(UserPoolDomainCustomDomainConfigTypePtrOutput)
}

// The resource ID.
func (o LookupUserPoolDomainResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolDomainResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolDomainResultOutput{})
}
