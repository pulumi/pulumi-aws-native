// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a listener for a service. Before you start using your Amazon VPC Lattice service, you must add one or more listeners. A listener is a process that checks for connection requests to your services.
type Listener struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput         `pulumi:"arn"`
	DefaultAction     ListenerDefaultActionOutput `pulumi:"defaultAction"`
	Name              pulumi.StringPtrOutput      `pulumi:"name"`
	Port              pulumi.IntPtrOutput         `pulumi:"port"`
	Protocol          ListenerProtocolOutput      `pulumi:"protocol"`
	ServiceArn        pulumi.StringOutput         `pulumi:"serviceArn"`
	ServiceId         pulumi.StringOutput         `pulumi:"serviceId"`
	ServiceIdentifier pulumi.StringPtrOutput      `pulumi:"serviceIdentifier"`
	Tags              ListenerTagArrayOutput      `pulumi:"tags"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Listener
	err := ctx.RegisterResource("aws-native:vpclattice:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("aws-native:vpclattice:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
}

type ListenerState struct {
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	DefaultAction     ListenerDefaultAction `pulumi:"defaultAction"`
	Name              *string               `pulumi:"name"`
	Port              *int                  `pulumi:"port"`
	Protocol          ListenerProtocol      `pulumi:"protocol"`
	ServiceIdentifier *string               `pulumi:"serviceIdentifier"`
	Tags              []ListenerTag         `pulumi:"tags"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	DefaultAction     ListenerDefaultActionInput
	Name              pulumi.StringPtrInput
	Port              pulumi.IntPtrInput
	Protocol          ListenerProtocolInput
	ServiceIdentifier pulumi.StringPtrInput
	Tags              ListenerTagArrayInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

func (o ListenerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ListenerOutput) DefaultAction() ListenerDefaultActionOutput {
	return o.ApplyT(func(v *Listener) ListenerDefaultActionOutput { return v.DefaultAction }).(ListenerDefaultActionOutput)
}

func (o ListenerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListenerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ListenerOutput) Protocol() ListenerProtocolOutput {
	return o.ApplyT(func(v *Listener) ListenerProtocolOutput { return v.Protocol }).(ListenerProtocolOutput)
}

func (o ListenerOutput) ServiceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ServiceArn }).(pulumi.StringOutput)
}

func (o ListenerOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func (o ListenerOutput) ServiceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.ServiceIdentifier }).(pulumi.StringPtrOutput)
}

func (o ListenerOutput) Tags() ListenerTagArrayOutput {
	return o.ApplyT(func(v *Listener) ListenerTagArrayOutput { return v.Tags }).(ListenerTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterOutputType(ListenerOutput{})
}