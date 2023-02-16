// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaPackage::PackagingGroup
func LookupPackagingGroup(ctx *pulumi.Context, args *LookupPackagingGroupArgs, opts ...pulumi.InvokeOption) (*LookupPackagingGroupResult, error) {
	var rv LookupPackagingGroupResult
	err := ctx.Invoke("aws-native:mediapackage:getPackagingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPackagingGroupArgs struct {
	// The ID of the PackagingGroup.
	Id string `pulumi:"id"`
}

type LookupPackagingGroupResult struct {
	// The ARN of the PackagingGroup.
	Arn *string `pulumi:"arn"`
	// CDN Authorization
	Authorization *PackagingGroupAuthorization `pulumi:"authorization"`
	// The fully qualified domain name for Assets in the PackagingGroup.
	DomainName *string `pulumi:"domainName"`
	// The configuration parameters for egress access logging.
	EgressAccessLogs *PackagingGroupLogConfiguration `pulumi:"egressAccessLogs"`
}

func LookupPackagingGroupOutput(ctx *pulumi.Context, args LookupPackagingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPackagingGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPackagingGroupResult, error) {
			args := v.(LookupPackagingGroupArgs)
			r, err := LookupPackagingGroup(ctx, &args, opts...)
			var s LookupPackagingGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPackagingGroupResultOutput)
}

type LookupPackagingGroupOutputArgs struct {
	// The ID of the PackagingGroup.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPackagingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackagingGroupArgs)(nil)).Elem()
}

type LookupPackagingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPackagingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackagingGroupResult)(nil)).Elem()
}

func (o LookupPackagingGroupResultOutput) ToLookupPackagingGroupResultOutput() LookupPackagingGroupResultOutput {
	return o
}

func (o LookupPackagingGroupResultOutput) ToLookupPackagingGroupResultOutputWithContext(ctx context.Context) LookupPackagingGroupResultOutput {
	return o
}

// The ARN of the PackagingGroup.
func (o LookupPackagingGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackagingGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// CDN Authorization
func (o LookupPackagingGroupResultOutput) Authorization() PackagingGroupAuthorizationPtrOutput {
	return o.ApplyT(func(v LookupPackagingGroupResult) *PackagingGroupAuthorization { return v.Authorization }).(PackagingGroupAuthorizationPtrOutput)
}

// The fully qualified domain name for Assets in the PackagingGroup.
func (o LookupPackagingGroupResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackagingGroupResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The configuration parameters for egress access logging.
func (o LookupPackagingGroupResultOutput) EgressAccessLogs() PackagingGroupLogConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPackagingGroupResult) *PackagingGroupLogConfiguration { return v.EgressAccessLogs }).(PackagingGroupLogConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPackagingGroupResultOutput{})
}