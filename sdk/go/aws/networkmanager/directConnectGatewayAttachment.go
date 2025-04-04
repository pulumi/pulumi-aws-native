// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::NetworkManager::DirectConnectGatewayAttachment Resource Type
type DirectConnectGatewayAttachment struct {
	pulumi.CustomResourceState

	// Id of the attachment.
	AttachmentId pulumi.StringOutput `pulumi:"attachmentId"`
	// The policy rule number associated with the attachment.
	AttachmentPolicyRuleNumber pulumi.IntOutput `pulumi:"attachmentPolicyRuleNumber"`
	// Attachment type.
	AttachmentType pulumi.StringOutput `pulumi:"attachmentType"`
	// The ARN of a core network for the Direct Connect Gateway attachment.
	CoreNetworkArn pulumi.StringOutput `pulumi:"coreNetworkArn"`
	// The ID of a core network for the Direct Connect Gateway attachment.
	CoreNetworkId pulumi.StringOutput `pulumi:"coreNetworkId"`
	// Creation time of the attachment.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ARN of the Direct Connect Gateway.
	DirectConnectGatewayArn pulumi.StringOutput `pulumi:"directConnectGatewayArn"`
	// The Regions where the edges are located.
	EdgeLocations pulumi.StringArrayOutput `pulumi:"edgeLocations"`
	// The name of the network function group attachment.
	NetworkFunctionGroupName pulumi.StringOutput `pulumi:"networkFunctionGroupName"`
	// Owner account of the attachment.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The attachment to move from one network function group to another.
	ProposedNetworkFunctionGroupChange DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangePtrOutput `pulumi:"proposedNetworkFunctionGroupChange"`
	// The attachment to move from one segment to another.
	ProposedSegmentChange DirectConnectGatewayAttachmentProposedSegmentChangePtrOutput `pulumi:"proposedSegmentChange"`
	// The ARN of the Resource.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// The name of the segment attachment..
	SegmentName pulumi.StringOutput `pulumi:"segmentName"`
	// State of the attachment.
	State pulumi.StringOutput `pulumi:"state"`
	// Tags for the attachment.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// Last update time of the attachment.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDirectConnectGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewDirectConnectGatewayAttachment(ctx *pulumi.Context,
	name string, args *DirectConnectGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*DirectConnectGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CoreNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'CoreNetworkId'")
	}
	if args.DirectConnectGatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'DirectConnectGatewayArn'")
	}
	if args.EdgeLocations == nil {
		return nil, errors.New("invalid value for required argument 'EdgeLocations'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"coreNetworkId",
		"directConnectGatewayArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectConnectGatewayAttachment
	err := ctx.RegisterResource("aws-native:networkmanager:DirectConnectGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectConnectGatewayAttachment gets an existing DirectConnectGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectConnectGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectConnectGatewayAttachmentState, opts ...pulumi.ResourceOption) (*DirectConnectGatewayAttachment, error) {
	var resource DirectConnectGatewayAttachment
	err := ctx.ReadResource("aws-native:networkmanager:DirectConnectGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectConnectGatewayAttachment resources.
type directConnectGatewayAttachmentState struct {
}

type DirectConnectGatewayAttachmentState struct {
}

func (DirectConnectGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*directConnectGatewayAttachmentState)(nil)).Elem()
}

type directConnectGatewayAttachmentArgs struct {
	// The ID of a core network for the Direct Connect Gateway attachment.
	CoreNetworkId string `pulumi:"coreNetworkId"`
	// The ARN of the Direct Connect Gateway.
	DirectConnectGatewayArn string `pulumi:"directConnectGatewayArn"`
	// The Regions where the edges are located.
	EdgeLocations []string `pulumi:"edgeLocations"`
	// The attachment to move from one network function group to another.
	ProposedNetworkFunctionGroupChange *DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChange `pulumi:"proposedNetworkFunctionGroupChange"`
	// The attachment to move from one segment to another.
	ProposedSegmentChange *DirectConnectGatewayAttachmentProposedSegmentChange `pulumi:"proposedSegmentChange"`
	// Tags for the attachment.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DirectConnectGatewayAttachment resource.
type DirectConnectGatewayAttachmentArgs struct {
	// The ID of a core network for the Direct Connect Gateway attachment.
	CoreNetworkId pulumi.StringInput
	// The ARN of the Direct Connect Gateway.
	DirectConnectGatewayArn pulumi.StringInput
	// The Regions where the edges are located.
	EdgeLocations pulumi.StringArrayInput
	// The attachment to move from one network function group to another.
	ProposedNetworkFunctionGroupChange DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangePtrInput
	// The attachment to move from one segment to another.
	ProposedSegmentChange DirectConnectGatewayAttachmentProposedSegmentChangePtrInput
	// Tags for the attachment.
	Tags aws.TagArrayInput
}

func (DirectConnectGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directConnectGatewayAttachmentArgs)(nil)).Elem()
}

type DirectConnectGatewayAttachmentInput interface {
	pulumi.Input

	ToDirectConnectGatewayAttachmentOutput() DirectConnectGatewayAttachmentOutput
	ToDirectConnectGatewayAttachmentOutputWithContext(ctx context.Context) DirectConnectGatewayAttachmentOutput
}

func (*DirectConnectGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectConnectGatewayAttachment)(nil)).Elem()
}

func (i *DirectConnectGatewayAttachment) ToDirectConnectGatewayAttachmentOutput() DirectConnectGatewayAttachmentOutput {
	return i.ToDirectConnectGatewayAttachmentOutputWithContext(context.Background())
}

func (i *DirectConnectGatewayAttachment) ToDirectConnectGatewayAttachmentOutputWithContext(ctx context.Context) DirectConnectGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectConnectGatewayAttachmentOutput)
}

type DirectConnectGatewayAttachmentOutput struct{ *pulumi.OutputState }

func (DirectConnectGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectConnectGatewayAttachment)(nil)).Elem()
}

func (o DirectConnectGatewayAttachmentOutput) ToDirectConnectGatewayAttachmentOutput() DirectConnectGatewayAttachmentOutput {
	return o
}

func (o DirectConnectGatewayAttachmentOutput) ToDirectConnectGatewayAttachmentOutputWithContext(ctx context.Context) DirectConnectGatewayAttachmentOutput {
	return o
}

// Id of the attachment.
func (o DirectConnectGatewayAttachmentOutput) AttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.AttachmentId }).(pulumi.StringOutput)
}

// The policy rule number associated with the attachment.
func (o DirectConnectGatewayAttachmentOutput) AttachmentPolicyRuleNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.IntOutput { return v.AttachmentPolicyRuleNumber }).(pulumi.IntOutput)
}

// Attachment type.
func (o DirectConnectGatewayAttachmentOutput) AttachmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.AttachmentType }).(pulumi.StringOutput)
}

// The ARN of a core network for the Direct Connect Gateway attachment.
func (o DirectConnectGatewayAttachmentOutput) CoreNetworkArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.CoreNetworkArn }).(pulumi.StringOutput)
}

// The ID of a core network for the Direct Connect Gateway attachment.
func (o DirectConnectGatewayAttachmentOutput) CoreNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.CoreNetworkId }).(pulumi.StringOutput)
}

// Creation time of the attachment.
func (o DirectConnectGatewayAttachmentOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ARN of the Direct Connect Gateway.
func (o DirectConnectGatewayAttachmentOutput) DirectConnectGatewayArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.DirectConnectGatewayArn }).(pulumi.StringOutput)
}

// The Regions where the edges are located.
func (o DirectConnectGatewayAttachmentOutput) EdgeLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringArrayOutput { return v.EdgeLocations }).(pulumi.StringArrayOutput)
}

// The name of the network function group attachment.
func (o DirectConnectGatewayAttachmentOutput) NetworkFunctionGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.NetworkFunctionGroupName }).(pulumi.StringOutput)
}

// Owner account of the attachment.
func (o DirectConnectGatewayAttachmentOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// The attachment to move from one network function group to another.
func (o DirectConnectGatewayAttachmentOutput) ProposedNetworkFunctionGroupChange() DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangePtrOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangePtrOutput {
		return v.ProposedNetworkFunctionGroupChange
	}).(DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChangePtrOutput)
}

// The attachment to move from one segment to another.
func (o DirectConnectGatewayAttachmentOutput) ProposedSegmentChange() DirectConnectGatewayAttachmentProposedSegmentChangePtrOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) DirectConnectGatewayAttachmentProposedSegmentChangePtrOutput {
		return v.ProposedSegmentChange
	}).(DirectConnectGatewayAttachmentProposedSegmentChangePtrOutput)
}

// The ARN of the Resource.
func (o DirectConnectGatewayAttachmentOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

// The name of the segment attachment..
func (o DirectConnectGatewayAttachmentOutput) SegmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.SegmentName }).(pulumi.StringOutput)
}

// State of the attachment.
func (o DirectConnectGatewayAttachmentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Tags for the attachment.
func (o DirectConnectGatewayAttachmentOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// Last update time of the attachment.
func (o DirectConnectGatewayAttachmentOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectConnectGatewayAttachment) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectConnectGatewayAttachmentInput)(nil)).Elem(), &DirectConnectGatewayAttachment{})
	pulumi.RegisterOutputType(DirectConnectGatewayAttachmentOutput{})
}
