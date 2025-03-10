// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cleanrooms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an association between an ID namespace and a collaboration
func LookupIdNamespaceAssociation(ctx *pulumi.Context, args *LookupIdNamespaceAssociationArgs, opts ...pulumi.InvokeOption) (*LookupIdNamespaceAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdNamespaceAssociationResult
	err := ctx.Invoke("aws-native:cleanrooms:getIdNamespaceAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdNamespaceAssociationArgs struct {
	// The unique identifier of the ID namespace association that you want to retrieve.
	IdNamespaceAssociationIdentifier string `pulumi:"idNamespaceAssociationIdentifier"`
	// The unique identifier of the membership that contains the ID namespace association.
	MembershipIdentifier string `pulumi:"membershipIdentifier"`
}

type LookupIdNamespaceAssociationResult struct {
	// The Amazon Resource Name (ARN) of the ID namespace association.
	Arn *string `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the collaboration that contains this ID namespace association.
	CollaborationArn *string `pulumi:"collaborationArn"`
	// The unique identifier of the collaboration that contains this ID namespace association.
	CollaborationIdentifier *string `pulumi:"collaborationIdentifier"`
	// The description of the ID namespace association.
	Description *string `pulumi:"description"`
	// The configuration settings for the ID mapping table.
	IdMappingConfig *IdNamespaceAssociationIdMappingConfig `pulumi:"idMappingConfig"`
	// The unique identifier of the ID namespace association that you want to retrieve.
	IdNamespaceAssociationIdentifier *string                                         `pulumi:"idNamespaceAssociationIdentifier"`
	InputReferenceProperties         *IdNamespaceAssociationInputReferenceProperties `pulumi:"inputReferenceProperties"`
	// The Amazon Resource Name (ARN) of the membership resource for this ID namespace association.
	MembershipArn *string `pulumi:"membershipArn"`
	// The name of this ID namespace association.
	Name *string `pulumi:"name"`
	// An optional label that you can assign to a resource when you create it. Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupIdNamespaceAssociationOutput(ctx *pulumi.Context, args LookupIdNamespaceAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupIdNamespaceAssociationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIdNamespaceAssociationResultOutput, error) {
			args := v.(LookupIdNamespaceAssociationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cleanrooms:getIdNamespaceAssociation", args, LookupIdNamespaceAssociationResultOutput{}, options).(LookupIdNamespaceAssociationResultOutput), nil
		}).(LookupIdNamespaceAssociationResultOutput)
}

type LookupIdNamespaceAssociationOutputArgs struct {
	// The unique identifier of the ID namespace association that you want to retrieve.
	IdNamespaceAssociationIdentifier pulumi.StringInput `pulumi:"idNamespaceAssociationIdentifier"`
	// The unique identifier of the membership that contains the ID namespace association.
	MembershipIdentifier pulumi.StringInput `pulumi:"membershipIdentifier"`
}

func (LookupIdNamespaceAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdNamespaceAssociationArgs)(nil)).Elem()
}

type LookupIdNamespaceAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupIdNamespaceAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdNamespaceAssociationResult)(nil)).Elem()
}

func (o LookupIdNamespaceAssociationResultOutput) ToLookupIdNamespaceAssociationResultOutput() LookupIdNamespaceAssociationResultOutput {
	return o
}

func (o LookupIdNamespaceAssociationResultOutput) ToLookupIdNamespaceAssociationResultOutputWithContext(ctx context.Context) LookupIdNamespaceAssociationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the ID namespace association.
func (o LookupIdNamespaceAssociationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the collaboration that contains this ID namespace association.
func (o LookupIdNamespaceAssociationResultOutput) CollaborationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.CollaborationArn }).(pulumi.StringPtrOutput)
}

// The unique identifier of the collaboration that contains this ID namespace association.
func (o LookupIdNamespaceAssociationResultOutput) CollaborationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.CollaborationIdentifier }).(pulumi.StringPtrOutput)
}

// The description of the ID namespace association.
func (o LookupIdNamespaceAssociationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The configuration settings for the ID mapping table.
func (o LookupIdNamespaceAssociationResultOutput) IdMappingConfig() IdNamespaceAssociationIdMappingConfigPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *IdNamespaceAssociationIdMappingConfig {
		return v.IdMappingConfig
	}).(IdNamespaceAssociationIdMappingConfigPtrOutput)
}

// The unique identifier of the ID namespace association that you want to retrieve.
func (o LookupIdNamespaceAssociationResultOutput) IdNamespaceAssociationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.IdNamespaceAssociationIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupIdNamespaceAssociationResultOutput) InputReferenceProperties() IdNamespaceAssociationInputReferencePropertiesPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *IdNamespaceAssociationInputReferenceProperties {
		return v.InputReferenceProperties
	}).(IdNamespaceAssociationInputReferencePropertiesPtrOutput)
}

// The Amazon Resource Name (ARN) of the membership resource for this ID namespace association.
func (o LookupIdNamespaceAssociationResultOutput) MembershipArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.MembershipArn }).(pulumi.StringPtrOutput)
}

// The name of this ID namespace association.
func (o LookupIdNamespaceAssociationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// An optional label that you can assign to a resource when you create it. Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
func (o LookupIdNamespaceAssociationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupIdNamespaceAssociationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdNamespaceAssociationResultOutput{})
}
