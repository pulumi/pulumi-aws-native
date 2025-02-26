// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enable a resource that has been published in the CloudFormation Registry.
func LookupTypeActivation(ctx *pulumi.Context, args *LookupTypeActivationArgs, opts ...pulumi.InvokeOption) (*LookupTypeActivationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTypeActivationResult
	err := ctx.Invoke("aws-native:cloudformation:getTypeActivation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTypeActivationArgs struct {
	// The Amazon Resource Name (ARN) of the extension.
	Arn string `pulumi:"arn"`
}

type LookupTypeActivationResult struct {
	// The Amazon Resource Name (ARN) of the extension.
	Arn *string `pulumi:"arn"`
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication
	PublicTypeArn *string `pulumi:"publicTypeArn"`
	// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
	PublisherId *string `pulumi:"publisherId"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName *string `pulumi:"typeName"`
	// An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
	TypeNameAlias *string `pulumi:"typeNameAlias"`
}

func LookupTypeActivationOutput(ctx *pulumi.Context, args LookupTypeActivationOutputArgs, opts ...pulumi.InvokeOption) LookupTypeActivationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTypeActivationResultOutput, error) {
			args := v.(LookupTypeActivationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cloudformation:getTypeActivation", args, LookupTypeActivationResultOutput{}, options).(LookupTypeActivationResultOutput), nil
		}).(LookupTypeActivationResultOutput)
}

type LookupTypeActivationOutputArgs struct {
	// The Amazon Resource Name (ARN) of the extension.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupTypeActivationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTypeActivationArgs)(nil)).Elem()
}

type LookupTypeActivationResultOutput struct{ *pulumi.OutputState }

func (LookupTypeActivationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTypeActivationResult)(nil)).Elem()
}

func (o LookupTypeActivationResultOutput) ToLookupTypeActivationResultOutput() LookupTypeActivationResultOutput {
	return o
}

func (o LookupTypeActivationResultOutput) ToLookupTypeActivationResultOutputWithContext(ctx context.Context) LookupTypeActivationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the extension.
func (o LookupTypeActivationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTypeActivationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Number (ARN) assigned to the public extension upon publication
func (o LookupTypeActivationResultOutput) PublicTypeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTypeActivationResult) *string { return v.PublicTypeArn }).(pulumi.StringPtrOutput)
}

// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
func (o LookupTypeActivationResultOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTypeActivationResult) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o LookupTypeActivationResultOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTypeActivationResult) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

// An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
func (o LookupTypeActivationResultOutput) TypeNameAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTypeActivationResult) *string { return v.TypeNameAlias }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTypeActivationResultOutput{})
}
