// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a listener for a service. Before you start using your Amazon VPC Lattice service, you must add one or more listeners. A listener is a process that checks for connection requests to your services.
type Listener struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the listener.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the listener.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The action for the default rule. Each listener has a default rule. The default rule is used if no other rules match.
	DefaultAction ListenerDefaultActionOutput `pulumi:"defaultAction"`
	// The name of the listener. A listener name must be unique within a service. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The listener port. You can specify a value from 1 to 65535. For HTTP, the default is 80. For HTTPS, the default is 443.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The listener protocol.
	Protocol ListenerProtocolOutput `pulumi:"protocol"`
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn pulumi.StringOutput `pulumi:"serviceArn"`
	// The ID of the service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The ID or ARN of the service.
	ServiceIdentifier pulumi.StringPtrOutput `pulumi:"serviceIdentifier"`
	// The tags for the listener.
	Tags aws.TagArrayOutput `pulumi:"tags"`
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
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"port",
		"protocol",
		"serviceIdentifier",
	})
	opts = append(opts, replaceOnChanges)
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
	// The action for the default rule. Each listener has a default rule. The default rule is used if no other rules match.
	DefaultAction ListenerDefaultAction `pulumi:"defaultAction"`
	// The name of the listener. A listener name must be unique within a service. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name *string `pulumi:"name"`
	// The listener port. You can specify a value from 1 to 65535. For HTTP, the default is 80. For HTTPS, the default is 443.
	Port *int `pulumi:"port"`
	// The listener protocol.
	Protocol ListenerProtocol `pulumi:"protocol"`
	// The ID or ARN of the service.
	ServiceIdentifier *string `pulumi:"serviceIdentifier"`
	// The tags for the listener.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The action for the default rule. Each listener has a default rule. The default rule is used if no other rules match.
	DefaultAction ListenerDefaultActionInput
	// The name of the listener. A listener name must be unique within a service. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name pulumi.StringPtrInput
	// The listener port. You can specify a value from 1 to 65535. For HTTP, the default is 80. For HTTPS, the default is 443.
	Port pulumi.IntPtrInput
	// The listener protocol.
	Protocol ListenerProtocolInput
	// The ID or ARN of the service.
	ServiceIdentifier pulumi.StringPtrInput
	// The tags for the listener.
	Tags aws.TagArrayInput
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

// The Amazon Resource Name (ARN) of the listener.
func (o ListenerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the listener.
func (o ListenerOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The action for the default rule. Each listener has a default rule. The default rule is used if no other rules match.
func (o ListenerOutput) DefaultAction() ListenerDefaultActionOutput {
	return o.ApplyT(func(v *Listener) ListenerDefaultActionOutput { return v.DefaultAction }).(ListenerDefaultActionOutput)
}

// The name of the listener. A listener name must be unique within a service. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
//
// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
func (o ListenerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The listener port. You can specify a value from 1 to 65535. For HTTP, the default is 80. For HTTPS, the default is 443.
func (o ListenerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The listener protocol.
func (o ListenerOutput) Protocol() ListenerProtocolOutput {
	return o.ApplyT(func(v *Listener) ListenerProtocolOutput { return v.Protocol }).(ListenerProtocolOutput)
}

// The Amazon Resource Name (ARN) of the service.
func (o ListenerOutput) ServiceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ServiceArn }).(pulumi.StringOutput)
}

// The ID of the service.
func (o ListenerOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The ID or ARN of the service.
func (o ListenerOutput) ServiceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.ServiceIdentifier }).(pulumi.StringPtrOutput)
}

// The tags for the listener.
func (o ListenerOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Listener) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterOutputType(ListenerOutput{})
}
