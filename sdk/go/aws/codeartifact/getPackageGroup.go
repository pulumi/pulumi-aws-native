// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource schema to create a CodeArtifact package group.
func LookupPackageGroup(ctx *pulumi.Context, args *LookupPackageGroupArgs, opts ...pulumi.InvokeOption) (*LookupPackageGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPackageGroupResult
	err := ctx.Invoke("aws-native:codeartifact:getPackageGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPackageGroupArgs struct {
	// The ARN of the package group.
	Arn string `pulumi:"arn"`
}

type LookupPackageGroupResult struct {
	// The ARN of the package group.
	Arn *string `pulumi:"arn"`
	// The contact info of the package group.
	ContactInfo *string `pulumi:"contactInfo"`
	// The text description of the package group.
	Description *string `pulumi:"description"`
	// The 12-digit account ID of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// The package origin configuration of the package group.
	OriginConfiguration *PackageGroupOriginConfiguration `pulumi:"originConfiguration"`
	// An array of key-value pairs to apply to the package group.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupPackageGroupOutput(ctx *pulumi.Context, args LookupPackageGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPackageGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPackageGroupResult, error) {
			args := v.(LookupPackageGroupArgs)
			r, err := LookupPackageGroup(ctx, &args, opts...)
			var s LookupPackageGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPackageGroupResultOutput)
}

type LookupPackageGroupOutputArgs struct {
	// The ARN of the package group.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupPackageGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackageGroupArgs)(nil)).Elem()
}

type LookupPackageGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPackageGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackageGroupResult)(nil)).Elem()
}

func (o LookupPackageGroupResultOutput) ToLookupPackageGroupResultOutput() LookupPackageGroupResultOutput {
	return o
}

func (o LookupPackageGroupResultOutput) ToLookupPackageGroupResultOutputWithContext(ctx context.Context) LookupPackageGroupResultOutput {
	return o
}

// The ARN of the package group.
func (o LookupPackageGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackageGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The contact info of the package group.
func (o LookupPackageGroupResultOutput) ContactInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackageGroupResult) *string { return v.ContactInfo }).(pulumi.StringPtrOutput)
}

// The text description of the package group.
func (o LookupPackageGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackageGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The 12-digit account ID of the AWS account that owns the domain.
func (o LookupPackageGroupResultOutput) DomainOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPackageGroupResult) *string { return v.DomainOwner }).(pulumi.StringPtrOutput)
}

// The package origin configuration of the package group.
func (o LookupPackageGroupResultOutput) OriginConfiguration() PackageGroupOriginConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPackageGroupResult) *PackageGroupOriginConfiguration { return v.OriginConfiguration }).(PackageGroupOriginConfigurationPtrOutput)
}

// An array of key-value pairs to apply to the package group.
func (o LookupPackageGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupPackageGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPackageGroupResultOutput{})
}