// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::SES::MailManagerAddonInstance Resource Type
type MailManagerAddonInstance struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Add On instance.
	AddonInstanceArn pulumi.StringOutput `pulumi:"addonInstanceArn"`
	// The unique ID of the Add On instance.
	AddonInstanceId pulumi.StringOutput `pulumi:"addonInstanceId"`
	// The name of the Add On for the instance.
	AddonName pulumi.StringOutput `pulumi:"addonName"`
	// The subscription ID for the instance.
	AddonSubscriptionId pulumi.StringOutput `pulumi:"addonSubscriptionId"`
	// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewMailManagerAddonInstance registers a new resource with the given unique name, arguments, and options.
func NewMailManagerAddonInstance(ctx *pulumi.Context,
	name string, args *MailManagerAddonInstanceArgs, opts ...pulumi.ResourceOption) (*MailManagerAddonInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddonSubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'AddonSubscriptionId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"addonSubscriptionId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MailManagerAddonInstance
	err := ctx.RegisterResource("aws-native:ses:MailManagerAddonInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMailManagerAddonInstance gets an existing MailManagerAddonInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMailManagerAddonInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MailManagerAddonInstanceState, opts ...pulumi.ResourceOption) (*MailManagerAddonInstance, error) {
	var resource MailManagerAddonInstance
	err := ctx.ReadResource("aws-native:ses:MailManagerAddonInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MailManagerAddonInstance resources.
type mailManagerAddonInstanceState struct {
}

type MailManagerAddonInstanceState struct {
}

func (MailManagerAddonInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*mailManagerAddonInstanceState)(nil)).Elem()
}

type mailManagerAddonInstanceArgs struct {
	// The subscription ID for the instance.
	AddonSubscriptionId string `pulumi:"addonSubscriptionId"`
	// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a MailManagerAddonInstance resource.
type MailManagerAddonInstanceArgs struct {
	// The subscription ID for the instance.
	AddonSubscriptionId pulumi.StringInput
	// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags aws.TagArrayInput
}

func (MailManagerAddonInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mailManagerAddonInstanceArgs)(nil)).Elem()
}

type MailManagerAddonInstanceInput interface {
	pulumi.Input

	ToMailManagerAddonInstanceOutput() MailManagerAddonInstanceOutput
	ToMailManagerAddonInstanceOutputWithContext(ctx context.Context) MailManagerAddonInstanceOutput
}

func (*MailManagerAddonInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**MailManagerAddonInstance)(nil)).Elem()
}

func (i *MailManagerAddonInstance) ToMailManagerAddonInstanceOutput() MailManagerAddonInstanceOutput {
	return i.ToMailManagerAddonInstanceOutputWithContext(context.Background())
}

func (i *MailManagerAddonInstance) ToMailManagerAddonInstanceOutputWithContext(ctx context.Context) MailManagerAddonInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MailManagerAddonInstanceOutput)
}

type MailManagerAddonInstanceOutput struct{ *pulumi.OutputState }

func (MailManagerAddonInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MailManagerAddonInstance)(nil)).Elem()
}

func (o MailManagerAddonInstanceOutput) ToMailManagerAddonInstanceOutput() MailManagerAddonInstanceOutput {
	return o
}

func (o MailManagerAddonInstanceOutput) ToMailManagerAddonInstanceOutputWithContext(ctx context.Context) MailManagerAddonInstanceOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Add On instance.
func (o MailManagerAddonInstanceOutput) AddonInstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *MailManagerAddonInstance) pulumi.StringOutput { return v.AddonInstanceArn }).(pulumi.StringOutput)
}

// The unique ID of the Add On instance.
func (o MailManagerAddonInstanceOutput) AddonInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MailManagerAddonInstance) pulumi.StringOutput { return v.AddonInstanceId }).(pulumi.StringOutput)
}

// The name of the Add On for the instance.
func (o MailManagerAddonInstanceOutput) AddonName() pulumi.StringOutput {
	return o.ApplyT(func(v *MailManagerAddonInstance) pulumi.StringOutput { return v.AddonName }).(pulumi.StringOutput)
}

// The subscription ID for the instance.
func (o MailManagerAddonInstanceOutput) AddonSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *MailManagerAddonInstance) pulumi.StringOutput { return v.AddonSubscriptionId }).(pulumi.StringOutput)
}

// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
func (o MailManagerAddonInstanceOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *MailManagerAddonInstance) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MailManagerAddonInstanceInput)(nil)).Elem(), &MailManagerAddonInstance{})
	pulumi.RegisterOutputType(MailManagerAddonInstanceOutput{})
}
