// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
type UserPoolUiCustomizationAttachment struct {
	pulumi.CustomResourceState

	// The app client ID for your UI customization. When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client. To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
	Css pulumi.StringPtrOutput `pulumi:"css"`
	// The ID of the user pool where you want to apply branding to the classic hosted UI.
	UserPoolId pulumi.StringOutput `pulumi:"userPoolId"`
}

// NewUserPoolUiCustomizationAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPoolUiCustomizationAttachment(ctx *pulumi.Context,
	name string, args *UserPoolUiCustomizationAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPoolUiCustomizationAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clientId",
		"userPoolId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPoolUiCustomizationAttachment
	err := ctx.RegisterResource("aws-native:cognito:UserPoolUiCustomizationAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPoolUiCustomizationAttachment gets an existing UserPoolUiCustomizationAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPoolUiCustomizationAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPoolUiCustomizationAttachmentState, opts ...pulumi.ResourceOption) (*UserPoolUiCustomizationAttachment, error) {
	var resource UserPoolUiCustomizationAttachment
	err := ctx.ReadResource("aws-native:cognito:UserPoolUiCustomizationAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPoolUiCustomizationAttachment resources.
type userPoolUiCustomizationAttachmentState struct {
}

type UserPoolUiCustomizationAttachmentState struct {
}

func (UserPoolUiCustomizationAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolUiCustomizationAttachmentState)(nil)).Elem()
}

type userPoolUiCustomizationAttachmentArgs struct {
	// The app client ID for your UI customization. When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
	ClientId string `pulumi:"clientId"`
	// A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client. To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
	Css *string `pulumi:"css"`
	// The ID of the user pool where you want to apply branding to the classic hosted UI.
	UserPoolId string `pulumi:"userPoolId"`
}

// The set of arguments for constructing a UserPoolUiCustomizationAttachment resource.
type UserPoolUiCustomizationAttachmentArgs struct {
	// The app client ID for your UI customization. When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
	ClientId pulumi.StringInput
	// A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client. To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
	Css pulumi.StringPtrInput
	// The ID of the user pool where you want to apply branding to the classic hosted UI.
	UserPoolId pulumi.StringInput
}

func (UserPoolUiCustomizationAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolUiCustomizationAttachmentArgs)(nil)).Elem()
}

type UserPoolUiCustomizationAttachmentInput interface {
	pulumi.Input

	ToUserPoolUiCustomizationAttachmentOutput() UserPoolUiCustomizationAttachmentOutput
	ToUserPoolUiCustomizationAttachmentOutputWithContext(ctx context.Context) UserPoolUiCustomizationAttachmentOutput
}

func (*UserPoolUiCustomizationAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPoolUiCustomizationAttachment)(nil)).Elem()
}

func (i *UserPoolUiCustomizationAttachment) ToUserPoolUiCustomizationAttachmentOutput() UserPoolUiCustomizationAttachmentOutput {
	return i.ToUserPoolUiCustomizationAttachmentOutputWithContext(context.Background())
}

func (i *UserPoolUiCustomizationAttachment) ToUserPoolUiCustomizationAttachmentOutputWithContext(ctx context.Context) UserPoolUiCustomizationAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolUiCustomizationAttachmentOutput)
}

type UserPoolUiCustomizationAttachmentOutput struct{ *pulumi.OutputState }

func (UserPoolUiCustomizationAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPoolUiCustomizationAttachment)(nil)).Elem()
}

func (o UserPoolUiCustomizationAttachmentOutput) ToUserPoolUiCustomizationAttachmentOutput() UserPoolUiCustomizationAttachmentOutput {
	return o
}

func (o UserPoolUiCustomizationAttachmentOutput) ToUserPoolUiCustomizationAttachmentOutputWithContext(ctx context.Context) UserPoolUiCustomizationAttachmentOutput {
	return o
}

// The app client ID for your UI customization. When this value isn't present, the customization applies to all user pool app clients that don't have client-level settings..
func (o UserPoolUiCustomizationAttachmentOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPoolUiCustomizationAttachment) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// A plaintext CSS file that contains the custom fields that you want to apply to your user pool or app client. To download a template, go to the Amazon Cognito console. Navigate to your user pool *App clients* tab, select *Login pages* , edit *Hosted UI (classic) style* , and select the link to `CSS template.css` .
func (o UserPoolUiCustomizationAttachmentOutput) Css() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPoolUiCustomizationAttachment) pulumi.StringPtrOutput { return v.Css }).(pulumi.StringPtrOutput)
}

// The ID of the user pool where you want to apply branding to the classic hosted UI.
func (o UserPoolUiCustomizationAttachmentOutput) UserPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPoolUiCustomizationAttachment) pulumi.StringOutput { return v.UserPoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPoolUiCustomizationAttachmentInput)(nil)).Elem(), &UserPoolUiCustomizationAttachment{})
	pulumi.RegisterOutputType(UserPoolUiCustomizationAttachmentOutput{})
}
