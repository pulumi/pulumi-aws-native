// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
func LookupResolverDnssecConfig(ctx *pulumi.Context, args *LookupResolverDnssecConfigArgs, opts ...pulumi.InvokeOption) (*LookupResolverDnssecConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResolverDnssecConfigResult
	err := ctx.Invoke("aws-native:route53resolver:getResolverDnssecConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResolverDnssecConfigArgs struct {
	// Id
	Id string `pulumi:"id"`
}

type LookupResolverDnssecConfigResult struct {
	// Id
	Id *string `pulumi:"id"`
	// AccountId
	OwnerId *string `pulumi:"ownerId"`
	// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
	ValidationStatus *ResolverDnssecConfigValidationStatus `pulumi:"validationStatus"`
}

func LookupResolverDnssecConfigOutput(ctx *pulumi.Context, args LookupResolverDnssecConfigOutputArgs, opts ...pulumi.InvokeOption) LookupResolverDnssecConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResolverDnssecConfigResult, error) {
			args := v.(LookupResolverDnssecConfigArgs)
			r, err := LookupResolverDnssecConfig(ctx, &args, opts...)
			var s LookupResolverDnssecConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResolverDnssecConfigResultOutput)
}

type LookupResolverDnssecConfigOutputArgs struct {
	// Id
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResolverDnssecConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverDnssecConfigArgs)(nil)).Elem()
}

type LookupResolverDnssecConfigResultOutput struct{ *pulumi.OutputState }

func (LookupResolverDnssecConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverDnssecConfigResult)(nil)).Elem()
}

func (o LookupResolverDnssecConfigResultOutput) ToLookupResolverDnssecConfigResultOutput() LookupResolverDnssecConfigResultOutput {
	return o
}

func (o LookupResolverDnssecConfigResultOutput) ToLookupResolverDnssecConfigResultOutputWithContext(ctx context.Context) LookupResolverDnssecConfigResultOutput {
	return o
}

// Id
func (o LookupResolverDnssecConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverDnssecConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// AccountId
func (o LookupResolverDnssecConfigResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverDnssecConfigResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
func (o LookupResolverDnssecConfigResultOutput) ValidationStatus() ResolverDnssecConfigValidationStatusPtrOutput {
	return o.ApplyT(func(v LookupResolverDnssecConfigResult) *ResolverDnssecConfigValidationStatus {
		return v.ValidationStatus
	}).(ResolverDnssecConfigValidationStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverDnssecConfigResultOutput{})
}