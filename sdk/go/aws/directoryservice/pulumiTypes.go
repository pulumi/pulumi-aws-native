// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type SimpleAdVpcSettings struct {
	// The identifiers of the subnets for the directory servers. The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.
	SubnetIds []string `pulumi:"subnetIds"`
	// The identifier of the VPC in which to create the directory.
	VpcId string `pulumi:"vpcId"`
}

// SimpleAdVpcSettingsInput is an input type that accepts SimpleAdVpcSettingsArgs and SimpleAdVpcSettingsOutput values.
// You can construct a concrete instance of `SimpleAdVpcSettingsInput` via:
//
//	SimpleAdVpcSettingsArgs{...}
type SimpleAdVpcSettingsInput interface {
	pulumi.Input

	ToSimpleAdVpcSettingsOutput() SimpleAdVpcSettingsOutput
	ToSimpleAdVpcSettingsOutputWithContext(context.Context) SimpleAdVpcSettingsOutput
}

type SimpleAdVpcSettingsArgs struct {
	// The identifiers of the subnets for the directory servers. The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
	// The identifier of the VPC in which to create the directory.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (SimpleAdVpcSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleAdVpcSettings)(nil)).Elem()
}

func (i SimpleAdVpcSettingsArgs) ToSimpleAdVpcSettingsOutput() SimpleAdVpcSettingsOutput {
	return i.ToSimpleAdVpcSettingsOutputWithContext(context.Background())
}

func (i SimpleAdVpcSettingsArgs) ToSimpleAdVpcSettingsOutputWithContext(ctx context.Context) SimpleAdVpcSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleAdVpcSettingsOutput)
}

type SimpleAdVpcSettingsOutput struct{ *pulumi.OutputState }

func (SimpleAdVpcSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleAdVpcSettings)(nil)).Elem()
}

func (o SimpleAdVpcSettingsOutput) ToSimpleAdVpcSettingsOutput() SimpleAdVpcSettingsOutput {
	return o
}

func (o SimpleAdVpcSettingsOutput) ToSimpleAdVpcSettingsOutputWithContext(ctx context.Context) SimpleAdVpcSettingsOutput {
	return o
}

// The identifiers of the subnets for the directory servers. The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.
func (o SimpleAdVpcSettingsOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleAdVpcSettings) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The identifier of the VPC in which to create the directory.
func (o SimpleAdVpcSettingsOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v SimpleAdVpcSettings) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleAdVpcSettingsInput)(nil)).Elem(), SimpleAdVpcSettingsArgs{})
	pulumi.RegisterOutputType(SimpleAdVpcSettingsOutput{})
}
