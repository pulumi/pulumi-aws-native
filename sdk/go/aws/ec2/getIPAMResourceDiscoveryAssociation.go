// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema of AWS::EC2::IPAMResourceDiscoveryAssociation Type
func LookupIPAMResourceDiscoveryAssociation(ctx *pulumi.Context, args *LookupIPAMResourceDiscoveryAssociationArgs, opts ...pulumi.InvokeOption) (*LookupIPAMResourceDiscoveryAssociationResult, error) {
	var rv LookupIPAMResourceDiscoveryAssociationResult
	err := ctx.Invoke("aws-native:ec2:getIPAMResourceDiscoveryAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPAMResourceDiscoveryAssociationArgs struct {
	// Id of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryAssociationId string `pulumi:"ipamResourceDiscoveryAssociationId"`
}

type LookupIPAMResourceDiscoveryAssociationResult struct {
	// Arn of the IPAM.
	IpamArn *string `pulumi:"ipamArn"`
	// The home region of the IPAM.
	IpamRegion *string `pulumi:"ipamRegion"`
	// The Amazon Resource Name (ARN) of the resource discovery association is a part of.
	IpamResourceDiscoveryAssociationArn *string `pulumi:"ipamResourceDiscoveryAssociationArn"`
	// Id of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryAssociationId *string `pulumi:"ipamResourceDiscoveryAssociationId"`
	// If the Resource Discovery Association exists due as part of CreateIpam.
	IsDefault *bool `pulumi:"isDefault"`
	// The AWS Account ID for the account where the shared IPAM exists.
	OwnerId *string `pulumi:"ownerId"`
	// The status of the resource discovery.
	ResourceDiscoveryStatus *string `pulumi:"resourceDiscoveryStatus"`
	// The operational state of the Resource Discovery Association. Related to Create/Delete activities.
	State *string `pulumi:"state"`
	// An array of key-value pairs to apply to this resource.
	Tags []IPAMResourceDiscoveryAssociationTag `pulumi:"tags"`
}

func LookupIPAMResourceDiscoveryAssociationOutput(ctx *pulumi.Context, args LookupIPAMResourceDiscoveryAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupIPAMResourceDiscoveryAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPAMResourceDiscoveryAssociationResult, error) {
			args := v.(LookupIPAMResourceDiscoveryAssociationArgs)
			r, err := LookupIPAMResourceDiscoveryAssociation(ctx, &args, opts...)
			var s LookupIPAMResourceDiscoveryAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPAMResourceDiscoveryAssociationResultOutput)
}

type LookupIPAMResourceDiscoveryAssociationOutputArgs struct {
	// Id of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryAssociationId pulumi.StringInput `pulumi:"ipamResourceDiscoveryAssociationId"`
}

func (LookupIPAMResourceDiscoveryAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPAMResourceDiscoveryAssociationArgs)(nil)).Elem()
}

type LookupIPAMResourceDiscoveryAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupIPAMResourceDiscoveryAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPAMResourceDiscoveryAssociationResult)(nil)).Elem()
}

func (o LookupIPAMResourceDiscoveryAssociationResultOutput) ToLookupIPAMResourceDiscoveryAssociationResultOutput() LookupIPAMResourceDiscoveryAssociationResultOutput {
	return o
}

func (o LookupIPAMResourceDiscoveryAssociationResultOutput) ToLookupIPAMResourceDiscoveryAssociationResultOutputWithContext(ctx context.Context) LookupIPAMResourceDiscoveryAssociationResultOutput {
	return o
}

// Arn of the IPAM.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) IpamArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string { return v.IpamArn }).(pulumi.StringPtrOutput)
}

// The home region of the IPAM.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) IpamRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string { return v.IpamRegion }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the resource discovery association is a part of.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) IpamResourceDiscoveryAssociationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string {
		return v.IpamResourceDiscoveryAssociationArn
	}).(pulumi.StringPtrOutput)
}

// Id of the IPAM Resource Discovery Association.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) IpamResourceDiscoveryAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string {
		return v.IpamResourceDiscoveryAssociationId
	}).(pulumi.StringPtrOutput)
}

// If the Resource Discovery Association exists due as part of CreateIpam.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The AWS Account ID for the account where the shared IPAM exists.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// The status of the resource discovery.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) ResourceDiscoveryStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string { return v.ResourceDiscoveryStatus }).(pulumi.StringPtrOutput)
}

// The operational state of the Resource Discovery Association. Related to Create/Delete activities.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupIPAMResourceDiscoveryAssociationResultOutput) Tags() IPAMResourceDiscoveryAssociationTagArrayOutput {
	return o.ApplyT(func(v LookupIPAMResourceDiscoveryAssociationResult) []IPAMResourceDiscoveryAssociationTag {
		return v.Tags
	}).(IPAMResourceDiscoveryAssociationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPAMResourceDiscoveryAssociationResultOutput{})
}