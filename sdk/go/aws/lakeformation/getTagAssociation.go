// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A resource schema representing a Lake Formation Tag Association. While tag associations are not explicit Lake Formation resources, this CloudFormation resource can be used to associate tags with Lake Formation entities.
func LookupTagAssociation(ctx *pulumi.Context, args *LookupTagAssociationArgs, opts ...pulumi.InvokeOption) (*LookupTagAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTagAssociationResult
	err := ctx.Invoke("aws-native:lakeformation:getTagAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagAssociationArgs struct {
	// Unique string identifying the resource. Used as primary identifier, which ideally should be a string
	ResourceIdentifier string `pulumi:"resourceIdentifier"`
	// Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string
	TagsIdentifier string `pulumi:"tagsIdentifier"`
}

type LookupTagAssociationResult struct {
	// Unique string identifying the resource. Used as primary identifier, which ideally should be a string
	ResourceIdentifier *string `pulumi:"resourceIdentifier"`
	// Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string
	TagsIdentifier *string `pulumi:"tagsIdentifier"`
}

func LookupTagAssociationOutput(ctx *pulumi.Context, args LookupTagAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupTagAssociationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTagAssociationResultOutput, error) {
			args := v.(LookupTagAssociationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:lakeformation:getTagAssociation", args, LookupTagAssociationResultOutput{}, options).(LookupTagAssociationResultOutput), nil
		}).(LookupTagAssociationResultOutput)
}

type LookupTagAssociationOutputArgs struct {
	// Unique string identifying the resource. Used as primary identifier, which ideally should be a string
	ResourceIdentifier pulumi.StringInput `pulumi:"resourceIdentifier"`
	// Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string
	TagsIdentifier pulumi.StringInput `pulumi:"tagsIdentifier"`
}

func (LookupTagAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagAssociationArgs)(nil)).Elem()
}

type LookupTagAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupTagAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagAssociationResult)(nil)).Elem()
}

func (o LookupTagAssociationResultOutput) ToLookupTagAssociationResultOutput() LookupTagAssociationResultOutput {
	return o
}

func (o LookupTagAssociationResultOutput) ToLookupTagAssociationResultOutputWithContext(ctx context.Context) LookupTagAssociationResultOutput {
	return o
}

// Unique string identifying the resource. Used as primary identifier, which ideally should be a string
func (o LookupTagAssociationResultOutput) ResourceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagAssociationResult) *string { return v.ResourceIdentifier }).(pulumi.StringPtrOutput)
}

// Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string
func (o LookupTagAssociationResultOutput) TagsIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagAssociationResult) *string { return v.TagsIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagAssociationResultOutput{})
}
