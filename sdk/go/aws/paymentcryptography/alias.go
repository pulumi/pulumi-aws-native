// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package paymentcryptography

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::PaymentCryptography::Alias Resource Type
type Alias struct {
	pulumi.CustomResourceState

	// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
	//
	// > Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
	AliasName pulumi.StringOutput `pulumi:"aliasName"`
	// The `KeyARN` of the key associated with the alias.
	KeyArn pulumi.StringPtrOutput `pulumi:"keyArn"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		args = &AliasArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"aliasName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alias
	err := ctx.RegisterResource("aws-native:paymentcryptography:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws-native:paymentcryptography:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
	//
	// > Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
	AliasName *string `pulumi:"aliasName"`
	// The `KeyARN` of the key associated with the alias.
	KeyArn *string `pulumi:"keyArn"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
	//
	// > Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
	AliasName pulumi.StringPtrInput
	// The `KeyARN` of the key associated with the alias.
	KeyArn pulumi.StringPtrInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
//
// > Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
func (o AliasOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.AliasName }).(pulumi.StringOutput)
}

// The `KeyARN` of the key associated with the alias.
func (o AliasOutput) KeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringPtrOutput { return v.KeyArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterOutputType(AliasOutput{})
}