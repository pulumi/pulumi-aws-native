// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an authorizer.
type Authorizer struct {
	pulumi.CustomResourceState

	Arn                    pulumi.StringOutput       `pulumi:"arn"`
	AuthorizerFunctionArn  pulumi.StringOutput       `pulumi:"authorizerFunctionArn"`
	AuthorizerName         pulumi.StringPtrOutput    `pulumi:"authorizerName"`
	EnableCachingForHttp   pulumi.BoolPtrOutput      `pulumi:"enableCachingForHttp"`
	SigningDisabled        pulumi.BoolPtrOutput      `pulumi:"signingDisabled"`
	Status                 AuthorizerStatusPtrOutput `pulumi:"status"`
	Tags                   AuthorizerTagArrayOutput  `pulumi:"tags"`
	TokenKeyName           pulumi.StringPtrOutput    `pulumi:"tokenKeyName"`
	TokenSigningPublicKeys pulumi.AnyOutput          `pulumi:"tokenSigningPublicKeys"`
}

// NewAuthorizer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizer(ctx *pulumi.Context,
	name string, args *AuthorizerArgs, opts ...pulumi.ResourceOption) (*Authorizer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizerFunctionArn == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizerFunctionArn'")
	}
	var resource Authorizer
	err := ctx.RegisterResource("aws-native:iot:Authorizer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizer gets an existing Authorizer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizerState, opts ...pulumi.ResourceOption) (*Authorizer, error) {
	var resource Authorizer
	err := ctx.ReadResource("aws-native:iot:Authorizer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Authorizer resources.
type authorizerState struct {
}

type AuthorizerState struct {
}

func (AuthorizerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizerState)(nil)).Elem()
}

type authorizerArgs struct {
	AuthorizerFunctionArn  string            `pulumi:"authorizerFunctionArn"`
	AuthorizerName         *string           `pulumi:"authorizerName"`
	EnableCachingForHttp   *bool             `pulumi:"enableCachingForHttp"`
	SigningDisabled        *bool             `pulumi:"signingDisabled"`
	Status                 *AuthorizerStatus `pulumi:"status"`
	Tags                   []AuthorizerTag   `pulumi:"tags"`
	TokenKeyName           *string           `pulumi:"tokenKeyName"`
	TokenSigningPublicKeys interface{}       `pulumi:"tokenSigningPublicKeys"`
}

// The set of arguments for constructing a Authorizer resource.
type AuthorizerArgs struct {
	AuthorizerFunctionArn  pulumi.StringInput
	AuthorizerName         pulumi.StringPtrInput
	EnableCachingForHttp   pulumi.BoolPtrInput
	SigningDisabled        pulumi.BoolPtrInput
	Status                 AuthorizerStatusPtrInput
	Tags                   AuthorizerTagArrayInput
	TokenKeyName           pulumi.StringPtrInput
	TokenSigningPublicKeys pulumi.Input
}

func (AuthorizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizerArgs)(nil)).Elem()
}

type AuthorizerInput interface {
	pulumi.Input

	ToAuthorizerOutput() AuthorizerOutput
	ToAuthorizerOutputWithContext(ctx context.Context) AuthorizerOutput
}

func (*Authorizer) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorizer)(nil)).Elem()
}

func (i *Authorizer) ToAuthorizerOutput() AuthorizerOutput {
	return i.ToAuthorizerOutputWithContext(context.Background())
}

func (i *Authorizer) ToAuthorizerOutputWithContext(ctx context.Context) AuthorizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerOutput)
}

type AuthorizerOutput struct{ *pulumi.OutputState }

func (AuthorizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorizer)(nil)).Elem()
}

func (o AuthorizerOutput) ToAuthorizerOutput() AuthorizerOutput {
	return o
}

func (o AuthorizerOutput) ToAuthorizerOutputWithContext(ctx context.Context) AuthorizerOutput {
	return o
}

func (o AuthorizerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o AuthorizerOutput) AuthorizerFunctionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringOutput { return v.AuthorizerFunctionArn }).(pulumi.StringOutput)
}

func (o AuthorizerOutput) AuthorizerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringPtrOutput { return v.AuthorizerName }).(pulumi.StringPtrOutput)
}

func (o AuthorizerOutput) EnableCachingForHttp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.BoolPtrOutput { return v.EnableCachingForHttp }).(pulumi.BoolPtrOutput)
}

func (o AuthorizerOutput) SigningDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.BoolPtrOutput { return v.SigningDisabled }).(pulumi.BoolPtrOutput)
}

func (o AuthorizerOutput) Status() AuthorizerStatusPtrOutput {
	return o.ApplyT(func(v *Authorizer) AuthorizerStatusPtrOutput { return v.Status }).(AuthorizerStatusPtrOutput)
}

func (o AuthorizerOutput) Tags() AuthorizerTagArrayOutput {
	return o.ApplyT(func(v *Authorizer) AuthorizerTagArrayOutput { return v.Tags }).(AuthorizerTagArrayOutput)
}

func (o AuthorizerOutput) TokenKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringPtrOutput { return v.TokenKeyName }).(pulumi.StringPtrOutput)
}

func (o AuthorizerOutput) TokenSigningPublicKeys() pulumi.AnyOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.AnyOutput { return v.TokenSigningPublicKeys }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizerInput)(nil)).Elem(), &Authorizer{})
	pulumi.RegisterOutputType(AuthorizerOutput{})
}