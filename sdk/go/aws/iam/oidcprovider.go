// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::OIDCProvider
type OIDCProvider struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the OIDC provider
	Arn            pulumi.StringOutput        `pulumi:"arn"`
	ClientIdList   pulumi.StringArrayOutput   `pulumi:"clientIdList"`
	Tags           OIDCProviderTagArrayOutput `pulumi:"tags"`
	ThumbprintList pulumi.StringArrayOutput   `pulumi:"thumbprintList"`
	Url            pulumi.StringPtrOutput     `pulumi:"url"`
}

// NewOIDCProvider registers a new resource with the given unique name, arguments, and options.
func NewOIDCProvider(ctx *pulumi.Context,
	name string, args *OIDCProviderArgs, opts ...pulumi.ResourceOption) (*OIDCProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ThumbprintList == nil {
		return nil, errors.New("invalid value for required argument 'ThumbprintList'")
	}
	var resource OIDCProvider
	err := ctx.RegisterResource("aws-native:iam:OIDCProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOIDCProvider gets an existing OIDCProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOIDCProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OIDCProviderState, opts ...pulumi.ResourceOption) (*OIDCProvider, error) {
	var resource OIDCProvider
	err := ctx.ReadResource("aws-native:iam:OIDCProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OIDCProvider resources.
type oidcproviderState struct {
}

type OIDCProviderState struct {
}

func (OIDCProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcproviderState)(nil)).Elem()
}

type oidcproviderArgs struct {
	ClientIdList   []string          `pulumi:"clientIdList"`
	Tags           []OIDCProviderTag `pulumi:"tags"`
	ThumbprintList []string          `pulumi:"thumbprintList"`
	Url            *string           `pulumi:"url"`
}

// The set of arguments for constructing a OIDCProvider resource.
type OIDCProviderArgs struct {
	ClientIdList   pulumi.StringArrayInput
	Tags           OIDCProviderTagArrayInput
	ThumbprintList pulumi.StringArrayInput
	Url            pulumi.StringPtrInput
}

func (OIDCProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcproviderArgs)(nil)).Elem()
}

type OIDCProviderInput interface {
	pulumi.Input

	ToOIDCProviderOutput() OIDCProviderOutput
	ToOIDCProviderOutputWithContext(ctx context.Context) OIDCProviderOutput
}

func (*OIDCProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**OIDCProvider)(nil)).Elem()
}

func (i *OIDCProvider) ToOIDCProviderOutput() OIDCProviderOutput {
	return i.ToOIDCProviderOutputWithContext(context.Background())
}

func (i *OIDCProvider) ToOIDCProviderOutputWithContext(ctx context.Context) OIDCProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OIDCProviderOutput)
}

type OIDCProviderOutput struct{ *pulumi.OutputState }

func (OIDCProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OIDCProvider)(nil)).Elem()
}

func (o OIDCProviderOutput) ToOIDCProviderOutput() OIDCProviderOutput {
	return o
}

func (o OIDCProviderOutput) ToOIDCProviderOutputWithContext(ctx context.Context) OIDCProviderOutput {
	return o
}

// Amazon Resource Name (ARN) of the OIDC provider
func (o OIDCProviderOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *OIDCProvider) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o OIDCProviderOutput) ClientIdList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OIDCProvider) pulumi.StringArrayOutput { return v.ClientIdList }).(pulumi.StringArrayOutput)
}

func (o OIDCProviderOutput) Tags() OIDCProviderTagArrayOutput {
	return o.ApplyT(func(v *OIDCProvider) OIDCProviderTagArrayOutput { return v.Tags }).(OIDCProviderTagArrayOutput)
}

func (o OIDCProviderOutput) ThumbprintList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OIDCProvider) pulumi.StringArrayOutput { return v.ThumbprintList }).(pulumi.StringArrayOutput)
}

func (o OIDCProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OIDCProvider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OIDCProviderInput)(nil)).Elem(), &OIDCProvider{})
	pulumi.RegisterOutputType(OIDCProviderOutput{})
}