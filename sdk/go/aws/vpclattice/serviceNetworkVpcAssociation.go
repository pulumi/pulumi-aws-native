// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates a VPC with a service network.
type ServiceNetworkVpcAssociation struct {
	pulumi.CustomResourceState

	Arn                      pulumi.StringOutput                        `pulumi:"arn"`
	CreatedAt                pulumi.StringOutput                        `pulumi:"createdAt"`
	SecurityGroupIds         pulumi.StringArrayOutput                   `pulumi:"securityGroupIds"`
	ServiceNetworkArn        pulumi.StringOutput                        `pulumi:"serviceNetworkArn"`
	ServiceNetworkId         pulumi.StringOutput                        `pulumi:"serviceNetworkId"`
	ServiceNetworkIdentifier pulumi.StringPtrOutput                     `pulumi:"serviceNetworkIdentifier"`
	ServiceNetworkName       pulumi.StringOutput                        `pulumi:"serviceNetworkName"`
	Status                   ServiceNetworkVpcAssociationStatusOutput   `pulumi:"status"`
	Tags                     ServiceNetworkVpcAssociationTagArrayOutput `pulumi:"tags"`
	VpcId                    pulumi.StringOutput                        `pulumi:"vpcId"`
	VpcIdentifier            pulumi.StringPtrOutput                     `pulumi:"vpcIdentifier"`
}

// NewServiceNetworkVpcAssociation registers a new resource with the given unique name, arguments, and options.
func NewServiceNetworkVpcAssociation(ctx *pulumi.Context,
	name string, args *ServiceNetworkVpcAssociationArgs, opts ...pulumi.ResourceOption) (*ServiceNetworkVpcAssociation, error) {
	if args == nil {
		args = &ServiceNetworkVpcAssociationArgs{}
	}

	var resource ServiceNetworkVpcAssociation
	err := ctx.RegisterResource("aws-native:vpclattice:ServiceNetworkVpcAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceNetworkVpcAssociation gets an existing ServiceNetworkVpcAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceNetworkVpcAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceNetworkVpcAssociationState, opts ...pulumi.ResourceOption) (*ServiceNetworkVpcAssociation, error) {
	var resource ServiceNetworkVpcAssociation
	err := ctx.ReadResource("aws-native:vpclattice:ServiceNetworkVpcAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceNetworkVpcAssociation resources.
type serviceNetworkVpcAssociationState struct {
}

type ServiceNetworkVpcAssociationState struct {
}

func (ServiceNetworkVpcAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceNetworkVpcAssociationState)(nil)).Elem()
}

type serviceNetworkVpcAssociationArgs struct {
	SecurityGroupIds         []string                          `pulumi:"securityGroupIds"`
	ServiceNetworkIdentifier *string                           `pulumi:"serviceNetworkIdentifier"`
	Tags                     []ServiceNetworkVpcAssociationTag `pulumi:"tags"`
	VpcIdentifier            *string                           `pulumi:"vpcIdentifier"`
}

// The set of arguments for constructing a ServiceNetworkVpcAssociation resource.
type ServiceNetworkVpcAssociationArgs struct {
	SecurityGroupIds         pulumi.StringArrayInput
	ServiceNetworkIdentifier pulumi.StringPtrInput
	Tags                     ServiceNetworkVpcAssociationTagArrayInput
	VpcIdentifier            pulumi.StringPtrInput
}

func (ServiceNetworkVpcAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceNetworkVpcAssociationArgs)(nil)).Elem()
}

type ServiceNetworkVpcAssociationInput interface {
	pulumi.Input

	ToServiceNetworkVpcAssociationOutput() ServiceNetworkVpcAssociationOutput
	ToServiceNetworkVpcAssociationOutputWithContext(ctx context.Context) ServiceNetworkVpcAssociationOutput
}

func (*ServiceNetworkVpcAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNetworkVpcAssociation)(nil)).Elem()
}

func (i *ServiceNetworkVpcAssociation) ToServiceNetworkVpcAssociationOutput() ServiceNetworkVpcAssociationOutput {
	return i.ToServiceNetworkVpcAssociationOutputWithContext(context.Background())
}

func (i *ServiceNetworkVpcAssociation) ToServiceNetworkVpcAssociationOutputWithContext(ctx context.Context) ServiceNetworkVpcAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceNetworkVpcAssociationOutput)
}

type ServiceNetworkVpcAssociationOutput struct{ *pulumi.OutputState }

func (ServiceNetworkVpcAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNetworkVpcAssociation)(nil)).Elem()
}

func (o ServiceNetworkVpcAssociationOutput) ToServiceNetworkVpcAssociationOutput() ServiceNetworkVpcAssociationOutput {
	return o
}

func (o ServiceNetworkVpcAssociationOutput) ToServiceNetworkVpcAssociationOutputWithContext(ctx context.Context) ServiceNetworkVpcAssociationOutput {
	return o
}

func (o ServiceNetworkVpcAssociationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ServiceNetworkVpcAssociationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ServiceNetworkVpcAssociationOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o ServiceNetworkVpcAssociationOutput) ServiceNetworkArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringOutput { return v.ServiceNetworkArn }).(pulumi.StringOutput)
}

func (o ServiceNetworkVpcAssociationOutput) ServiceNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringOutput { return v.ServiceNetworkId }).(pulumi.StringOutput)
}

func (o ServiceNetworkVpcAssociationOutput) ServiceNetworkIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringPtrOutput { return v.ServiceNetworkIdentifier }).(pulumi.StringPtrOutput)
}

func (o ServiceNetworkVpcAssociationOutput) ServiceNetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringOutput { return v.ServiceNetworkName }).(pulumi.StringOutput)
}

func (o ServiceNetworkVpcAssociationOutput) Status() ServiceNetworkVpcAssociationStatusOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) ServiceNetworkVpcAssociationStatusOutput { return v.Status }).(ServiceNetworkVpcAssociationStatusOutput)
}

func (o ServiceNetworkVpcAssociationOutput) Tags() ServiceNetworkVpcAssociationTagArrayOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) ServiceNetworkVpcAssociationTagArrayOutput { return v.Tags }).(ServiceNetworkVpcAssociationTagArrayOutput)
}

func (o ServiceNetworkVpcAssociationOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func (o ServiceNetworkVpcAssociationOutput) VpcIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceNetworkVpcAssociation) pulumi.StringPtrOutput { return v.VpcIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceNetworkVpcAssociationInput)(nil)).Elem(), &ServiceNetworkVpcAssociation{})
	pulumi.RegisterOutputType(ServiceNetworkVpcAssociationOutput{})
}