// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::WorkSpacesWeb::NetworkSettings Resource Type
func LookupNetworkSettings(ctx *pulumi.Context, args *LookupNetworkSettingsArgs, opts ...pulumi.InvokeOption) (*LookupNetworkSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkSettingsResult
	err := ctx.Invoke("aws-native:workspacesweb:getNetworkSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkSettingsArgs struct {
	// The ARN of the network settings.
	NetworkSettingsArn string `pulumi:"networkSettingsArn"`
}

type LookupNetworkSettingsResult struct {
	// A list of web portal ARNs that this network settings is associated with.
	AssociatedPortalArns []string `pulumi:"associatedPortalArns"`
	// The ARN of the network settings.
	NetworkSettingsArn *string `pulumi:"networkSettingsArn"`
	// One or more security groups used to control access from streaming instances to your VPC.
	//
	// *Pattern* : `^[\w+\-]+$`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The subnets in which network interfaces are created to connect streaming instances to your VPC. At least two of these subnets must be in different availability zones.
	//
	// *Pattern* : `^subnet-([0-9a-f]{8}|[0-9a-f]{17})$`
	SubnetIds []string `pulumi:"subnetIds"`
	// The tags to add to the network settings resource. A tag is a key-value pair.
	Tags []aws.Tag `pulumi:"tags"`
	// The VPC that streaming instances will connect to.
	//
	// *Pattern* : `^vpc-[0-9a-z]*$`
	VpcId *string `pulumi:"vpcId"`
}

func LookupNetworkSettingsOutput(ctx *pulumi.Context, args LookupNetworkSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkSettingsResultOutput, error) {
			args := v.(LookupNetworkSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:workspacesweb:getNetworkSettings", args, LookupNetworkSettingsResultOutput{}, options).(LookupNetworkSettingsResultOutput), nil
		}).(LookupNetworkSettingsResultOutput)
}

type LookupNetworkSettingsOutputArgs struct {
	// The ARN of the network settings.
	NetworkSettingsArn pulumi.StringInput `pulumi:"networkSettingsArn"`
}

func (LookupNetworkSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSettingsArgs)(nil)).Elem()
}

type LookupNetworkSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSettingsResult)(nil)).Elem()
}

func (o LookupNetworkSettingsResultOutput) ToLookupNetworkSettingsResultOutput() LookupNetworkSettingsResultOutput {
	return o
}

func (o LookupNetworkSettingsResultOutput) ToLookupNetworkSettingsResultOutputWithContext(ctx context.Context) LookupNetworkSettingsResultOutput {
	return o
}

// A list of web portal ARNs that this network settings is associated with.
func (o LookupNetworkSettingsResultOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkSettingsResult) []string { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

// The ARN of the network settings.
func (o LookupNetworkSettingsResultOutput) NetworkSettingsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSettingsResult) *string { return v.NetworkSettingsArn }).(pulumi.StringPtrOutput)
}

// One or more security groups used to control access from streaming instances to your VPC.
//
// *Pattern* : `^[\w+\-]+$`
func (o LookupNetworkSettingsResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkSettingsResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The subnets in which network interfaces are created to connect streaming instances to your VPC. At least two of these subnets must be in different availability zones.
//
// *Pattern* : `^subnet-([0-9a-f]{8}|[0-9a-f]{17})$`
func (o LookupNetworkSettingsResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkSettingsResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The tags to add to the network settings resource. A tag is a key-value pair.
func (o LookupNetworkSettingsResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupNetworkSettingsResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The VPC that streaming instances will connect to.
//
// *Pattern* : `^vpc-[0-9a-z]*$`
func (o LookupNetworkSettingsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSettingsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkSettingsResultOutput{})
}
