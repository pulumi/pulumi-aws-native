// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53profiles

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Profiles::ProfileResourceAssociation
func LookupProfileResourceAssociation(ctx *pulumi.Context, args *LookupProfileResourceAssociationArgs, opts ...pulumi.InvokeOption) (*LookupProfileResourceAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProfileResourceAssociationResult
	err := ctx.Invoke("aws-native:route53profiles:getProfileResourceAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileResourceAssociationArgs struct {
	// Primary Identifier for  Profile Resource Association
	Id string `pulumi:"id"`
}

type LookupProfileResourceAssociationResult struct {
	// Primary Identifier for  Profile Resource Association
	Id *string `pulumi:"id"`
	// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
	ResourceProperties *string `pulumi:"resourceProperties"`
	// The type of the resource associated to the  Profile.
	ResourceType *string `pulumi:"resourceType"`
}

func LookupProfileResourceAssociationOutput(ctx *pulumi.Context, args LookupProfileResourceAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResourceAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResourceAssociationResult, error) {
			args := v.(LookupProfileResourceAssociationArgs)
			r, err := LookupProfileResourceAssociation(ctx, &args, opts...)
			var s LookupProfileResourceAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResourceAssociationResultOutput)
}

type LookupProfileResourceAssociationOutputArgs struct {
	// Primary Identifier for  Profile Resource Association
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupProfileResourceAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResourceAssociationArgs)(nil)).Elem()
}

type LookupProfileResourceAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResourceAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResourceAssociationResult)(nil)).Elem()
}

func (o LookupProfileResourceAssociationResultOutput) ToLookupProfileResourceAssociationResultOutput() LookupProfileResourceAssociationResultOutput {
	return o
}

func (o LookupProfileResourceAssociationResultOutput) ToLookupProfileResourceAssociationResultOutputWithContext(ctx context.Context) LookupProfileResourceAssociationResultOutput {
	return o
}

// Primary Identifier for  Profile Resource Association
func (o LookupProfileResourceAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResourceAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
func (o LookupProfileResourceAssociationResultOutput) ResourceProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResourceAssociationResult) *string { return v.ResourceProperties }).(pulumi.StringPtrOutput)
}

// The type of the resource associated to the  Profile.
func (o LookupProfileResourceAssociationResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResourceAssociationResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResourceAssociationResultOutput{})
}