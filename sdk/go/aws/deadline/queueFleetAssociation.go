// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deadline

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Deadline::QueueFleetAssociation Resource Type
type QueueFleetAssociation struct {
	pulumi.CustomResourceState

	FarmId  pulumi.StringOutput `pulumi:"farmId"`
	FleetId pulumi.StringOutput `pulumi:"fleetId"`
	QueueId pulumi.StringOutput `pulumi:"queueId"`
}

// NewQueueFleetAssociation registers a new resource with the given unique name, arguments, and options.
func NewQueueFleetAssociation(ctx *pulumi.Context,
	name string, args *QueueFleetAssociationArgs, opts ...pulumi.ResourceOption) (*QueueFleetAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FarmId == nil {
		return nil, errors.New("invalid value for required argument 'FarmId'")
	}
	if args.FleetId == nil {
		return nil, errors.New("invalid value for required argument 'FleetId'")
	}
	if args.QueueId == nil {
		return nil, errors.New("invalid value for required argument 'QueueId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"farmId",
		"fleetId",
		"queueId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueueFleetAssociation
	err := ctx.RegisterResource("aws-native:deadline:QueueFleetAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueFleetAssociation gets an existing QueueFleetAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueFleetAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueFleetAssociationState, opts ...pulumi.ResourceOption) (*QueueFleetAssociation, error) {
	var resource QueueFleetAssociation
	err := ctx.ReadResource("aws-native:deadline:QueueFleetAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueFleetAssociation resources.
type queueFleetAssociationState struct {
}

type QueueFleetAssociationState struct {
}

func (QueueFleetAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueFleetAssociationState)(nil)).Elem()
}

type queueFleetAssociationArgs struct {
	FarmId  string `pulumi:"farmId"`
	FleetId string `pulumi:"fleetId"`
	QueueId string `pulumi:"queueId"`
}

// The set of arguments for constructing a QueueFleetAssociation resource.
type QueueFleetAssociationArgs struct {
	FarmId  pulumi.StringInput
	FleetId pulumi.StringInput
	QueueId pulumi.StringInput
}

func (QueueFleetAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueFleetAssociationArgs)(nil)).Elem()
}

type QueueFleetAssociationInput interface {
	pulumi.Input

	ToQueueFleetAssociationOutput() QueueFleetAssociationOutput
	ToQueueFleetAssociationOutputWithContext(ctx context.Context) QueueFleetAssociationOutput
}

func (*QueueFleetAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueFleetAssociation)(nil)).Elem()
}

func (i *QueueFleetAssociation) ToQueueFleetAssociationOutput() QueueFleetAssociationOutput {
	return i.ToQueueFleetAssociationOutputWithContext(context.Background())
}

func (i *QueueFleetAssociation) ToQueueFleetAssociationOutputWithContext(ctx context.Context) QueueFleetAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueFleetAssociationOutput)
}

type QueueFleetAssociationOutput struct{ *pulumi.OutputState }

func (QueueFleetAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueFleetAssociation)(nil)).Elem()
}

func (o QueueFleetAssociationOutput) ToQueueFleetAssociationOutput() QueueFleetAssociationOutput {
	return o
}

func (o QueueFleetAssociationOutput) ToQueueFleetAssociationOutputWithContext(ctx context.Context) QueueFleetAssociationOutput {
	return o
}

func (o QueueFleetAssociationOutput) FarmId() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueFleetAssociation) pulumi.StringOutput { return v.FarmId }).(pulumi.StringOutput)
}

func (o QueueFleetAssociationOutput) FleetId() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueFleetAssociation) pulumi.StringOutput { return v.FleetId }).(pulumi.StringOutput)
}

func (o QueueFleetAssociationOutput) QueueId() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueFleetAssociation) pulumi.StringOutput { return v.QueueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueFleetAssociationInput)(nil)).Elem(), &QueueFleetAssociation{})
	pulumi.RegisterOutputType(QueueFleetAssociationOutput{})
}