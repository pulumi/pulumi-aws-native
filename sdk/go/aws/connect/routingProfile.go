// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::RoutingProfile
type RoutingProfile struct {
	pulumi.CustomResourceState

	// The identifier of the default outbound queue for this routing profile.
	DefaultOutboundQueueArn pulumi.StringOutput `pulumi:"defaultOutboundQueueArn"`
	// The description of the routing profile.
	Description pulumi.StringOutput `pulumi:"description"`
	// The identifier of the Amazon Connect instance.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
	MediaConcurrencies RoutingProfileMediaConcurrencyArrayOutput `pulumi:"mediaConcurrencies"`
	// The name of the routing profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// The queues to associate with this routing profile.
	QueueConfigs RoutingProfileQueueConfigArrayOutput `pulumi:"queueConfigs"`
	// The Amazon Resource Name (ARN) of the routing profile.
	RoutingProfileArn pulumi.StringOutput `pulumi:"routingProfileArn"`
	// An array of key-value pairs to apply to this resource.
	Tags RoutingProfileTagArrayOutput `pulumi:"tags"`
}

// NewRoutingProfile registers a new resource with the given unique name, arguments, and options.
func NewRoutingProfile(ctx *pulumi.Context,
	name string, args *RoutingProfileArgs, opts ...pulumi.ResourceOption) (*RoutingProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultOutboundQueueArn == nil {
		return nil, errors.New("invalid value for required argument 'DefaultOutboundQueueArn'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	if args.MediaConcurrencies == nil {
		return nil, errors.New("invalid value for required argument 'MediaConcurrencies'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RoutingProfile
	err := ctx.RegisterResource("aws-native:connect:RoutingProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoutingProfile gets an existing RoutingProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoutingProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutingProfileState, opts ...pulumi.ResourceOption) (*RoutingProfile, error) {
	var resource RoutingProfile
	err := ctx.ReadResource("aws-native:connect:RoutingProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoutingProfile resources.
type routingProfileState struct {
}

type RoutingProfileState struct {
}

func (RoutingProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*routingProfileState)(nil)).Elem()
}

type routingProfileArgs struct {
	// The identifier of the default outbound queue for this routing profile.
	DefaultOutboundQueueArn string `pulumi:"defaultOutboundQueueArn"`
	// The description of the routing profile.
	Description string `pulumi:"description"`
	// The identifier of the Amazon Connect instance.
	InstanceArn string `pulumi:"instanceArn"`
	// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
	MediaConcurrencies []RoutingProfileMediaConcurrency `pulumi:"mediaConcurrencies"`
	// The name of the routing profile.
	Name *string `pulumi:"name"`
	// The queues to associate with this routing profile.
	QueueConfigs []RoutingProfileQueueConfig `pulumi:"queueConfigs"`
	// An array of key-value pairs to apply to this resource.
	Tags []RoutingProfileTag `pulumi:"tags"`
}

// The set of arguments for constructing a RoutingProfile resource.
type RoutingProfileArgs struct {
	// The identifier of the default outbound queue for this routing profile.
	DefaultOutboundQueueArn pulumi.StringInput
	// The description of the routing profile.
	Description pulumi.StringInput
	// The identifier of the Amazon Connect instance.
	InstanceArn pulumi.StringInput
	// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
	MediaConcurrencies RoutingProfileMediaConcurrencyArrayInput
	// The name of the routing profile.
	Name pulumi.StringPtrInput
	// The queues to associate with this routing profile.
	QueueConfigs RoutingProfileQueueConfigArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags RoutingProfileTagArrayInput
}

func (RoutingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routingProfileArgs)(nil)).Elem()
}

type RoutingProfileInput interface {
	pulumi.Input

	ToRoutingProfileOutput() RoutingProfileOutput
	ToRoutingProfileOutputWithContext(ctx context.Context) RoutingProfileOutput
}

func (*RoutingProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProfile)(nil)).Elem()
}

func (i *RoutingProfile) ToRoutingProfileOutput() RoutingProfileOutput {
	return i.ToRoutingProfileOutputWithContext(context.Background())
}

func (i *RoutingProfile) ToRoutingProfileOutputWithContext(ctx context.Context) RoutingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingProfileOutput)
}

type RoutingProfileOutput struct{ *pulumi.OutputState }

func (RoutingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProfile)(nil)).Elem()
}

func (o RoutingProfileOutput) ToRoutingProfileOutput() RoutingProfileOutput {
	return o
}

func (o RoutingProfileOutput) ToRoutingProfileOutputWithContext(ctx context.Context) RoutingProfileOutput {
	return o
}

// The identifier of the default outbound queue for this routing profile.
func (o RoutingProfileOutput) DefaultOutboundQueueArn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProfile) pulumi.StringOutput { return v.DefaultOutboundQueueArn }).(pulumi.StringOutput)
}

// The description of the routing profile.
func (o RoutingProfileOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProfile) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The identifier of the Amazon Connect instance.
func (o RoutingProfileOutput) InstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProfile) pulumi.StringOutput { return v.InstanceArn }).(pulumi.StringOutput)
}

// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
func (o RoutingProfileOutput) MediaConcurrencies() RoutingProfileMediaConcurrencyArrayOutput {
	return o.ApplyT(func(v *RoutingProfile) RoutingProfileMediaConcurrencyArrayOutput { return v.MediaConcurrencies }).(RoutingProfileMediaConcurrencyArrayOutput)
}

// The name of the routing profile.
func (o RoutingProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The queues to associate with this routing profile.
func (o RoutingProfileOutput) QueueConfigs() RoutingProfileQueueConfigArrayOutput {
	return o.ApplyT(func(v *RoutingProfile) RoutingProfileQueueConfigArrayOutput { return v.QueueConfigs }).(RoutingProfileQueueConfigArrayOutput)
}

// The Amazon Resource Name (ARN) of the routing profile.
func (o RoutingProfileOutput) RoutingProfileArn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingProfile) pulumi.StringOutput { return v.RoutingProfileArn }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o RoutingProfileOutput) Tags() RoutingProfileTagArrayOutput {
	return o.ApplyT(func(v *RoutingProfile) RoutingProfileTagArrayOutput { return v.Tags }).(RoutingProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingProfileInput)(nil)).Elem(), &RoutingProfile{})
	pulumi.RegisterOutputType(RoutingProfileOutput{})
}