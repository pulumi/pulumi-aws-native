// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codegurureviewer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.
func LookupRepositoryAssociation(ctx *pulumi.Context, args *LookupRepositoryAssociationArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryAssociationResult
	err := ctx.Invoke("aws-native:codegurureviewer:getRepositoryAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRepositoryAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the repository association.
	AssociationArn string `pulumi:"associationArn"`
}

type LookupRepositoryAssociationResult struct {
	// The Amazon Resource Name (ARN) of the repository association.
	AssociationArn *string `pulumi:"associationArn"`
}

func LookupRepositoryAssociationOutput(ctx *pulumi.Context, args LookupRepositoryAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryAssociationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRepositoryAssociationResultOutput, error) {
			args := v.(LookupRepositoryAssociationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:codegurureviewer:getRepositoryAssociation", args, LookupRepositoryAssociationResultOutput{}, options).(LookupRepositoryAssociationResultOutput), nil
		}).(LookupRepositoryAssociationResultOutput)
}

type LookupRepositoryAssociationOutputArgs struct {
	// The Amazon Resource Name (ARN) of the repository association.
	AssociationArn pulumi.StringInput `pulumi:"associationArn"`
}

func (LookupRepositoryAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryAssociationArgs)(nil)).Elem()
}

type LookupRepositoryAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryAssociationResult)(nil)).Elem()
}

func (o LookupRepositoryAssociationResultOutput) ToLookupRepositoryAssociationResultOutput() LookupRepositoryAssociationResultOutput {
	return o
}

func (o LookupRepositoryAssociationResultOutput) ToLookupRepositoryAssociationResultOutputWithContext(ctx context.Context) LookupRepositoryAssociationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the repository association.
func (o LookupRepositoryAssociationResultOutput) AssociationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryAssociationResult) *string { return v.AssociationArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryAssociationResultOutput{})
}
