// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::VPCConnection Resource Type.
func LookupVpcConnection(ctx *pulumi.Context, args *LookupVpcConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpcConnectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcConnectionResult
	err := ctx.Invoke("aws-native:quicksight:getVpcConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcConnectionArgs struct {
	// The AWS account ID of the account where you want to create a new VPC connection.
	AwsAccountId string `pulumi:"awsAccountId"`
	// The ID of the VPC connection that you're creating. This ID is a unique identifier for each AWS Region in an AWS account.
	VpcConnectionId string `pulumi:"vpcConnectionId"`
}

type LookupVpcConnectionResult struct {
	// <p>The Amazon Resource Name (ARN) of the VPC connection.</p>
	Arn *string `pulumi:"arn"`
	// The availability status of the VPC connection.
	AvailabilityStatus *VpcConnectionVpcConnectionAvailabilityStatus `pulumi:"availabilityStatus"`
	// <p>The time that the VPC connection was created.</p>
	CreatedTime *string `pulumi:"createdTime"`
	// A list of IP addresses of DNS resolver endpoints for the VPC connection.
	DnsResolvers []string `pulumi:"dnsResolvers"`
	// <p>The time that the VPC connection was last updated.</p>
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// The display name for the VPC connection.
	Name *string `pulumi:"name"`
	// <p>A list of network interfaces.</p>
	NetworkInterfaces []VpcConnectionNetworkInterface `pulumi:"networkInterfaces"`
	// The ARN of the IAM role associated with the VPC connection.
	RoleArn *string `pulumi:"roleArn"`
	// The Amazon EC2 security group IDs associated with the VPC connection.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The HTTP status of the request.
	Status *VpcConnectionVpcConnectionResourceStatus `pulumi:"status"`
	// A map of the key-value pairs for the resource tag or tags assigned to the VPC connection.
	Tags []aws.Tag `pulumi:"tags"`
	// <p>The Amazon EC2 VPC ID associated with the VPC connection.</p>
	VpcId *string `pulumi:"vpcId"`
}

func LookupVpcConnectionOutput(ctx *pulumi.Context, args LookupVpcConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpcConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcConnectionResultOutput, error) {
			args := v.(LookupVpcConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:quicksight:getVpcConnection", args, LookupVpcConnectionResultOutput{}, options).(LookupVpcConnectionResultOutput), nil
		}).(LookupVpcConnectionResultOutput)
}

type LookupVpcConnectionOutputArgs struct {
	// The AWS account ID of the account where you want to create a new VPC connection.
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	// The ID of the VPC connection that you're creating. This ID is a unique identifier for each AWS Region in an AWS account.
	VpcConnectionId pulumi.StringInput `pulumi:"vpcConnectionId"`
}

func (LookupVpcConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcConnectionArgs)(nil)).Elem()
}

type LookupVpcConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpcConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcConnectionResult)(nil)).Elem()
}

func (o LookupVpcConnectionResultOutput) ToLookupVpcConnectionResultOutput() LookupVpcConnectionResultOutput {
	return o
}

func (o LookupVpcConnectionResultOutput) ToLookupVpcConnectionResultOutputWithContext(ctx context.Context) LookupVpcConnectionResultOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) of the VPC connection.</p>
func (o LookupVpcConnectionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The availability status of the VPC connection.
func (o LookupVpcConnectionResultOutput) AvailabilityStatus() VpcConnectionVpcConnectionAvailabilityStatusPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *VpcConnectionVpcConnectionAvailabilityStatus {
		return v.AvailabilityStatus
	}).(VpcConnectionVpcConnectionAvailabilityStatusPtrOutput)
}

// <p>The time that the VPC connection was created.</p>
func (o LookupVpcConnectionResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// A list of IP addresses of DNS resolver endpoints for the VPC connection.
func (o LookupVpcConnectionResultOutput) DnsResolvers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) []string { return v.DnsResolvers }).(pulumi.StringArrayOutput)
}

// <p>The time that the VPC connection was last updated.</p>
func (o LookupVpcConnectionResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// The display name for the VPC connection.
func (o LookupVpcConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// <p>A list of network interfaces.</p>
func (o LookupVpcConnectionResultOutput) NetworkInterfaces() VpcConnectionNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) []VpcConnectionNetworkInterface { return v.NetworkInterfaces }).(VpcConnectionNetworkInterfaceArrayOutput)
}

// The ARN of the IAM role associated with the VPC connection.
func (o LookupVpcConnectionResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// The Amazon EC2 security group IDs associated with the VPC connection.
func (o LookupVpcConnectionResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The HTTP status of the request.
func (o LookupVpcConnectionResultOutput) Status() VpcConnectionVpcConnectionResourceStatusPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *VpcConnectionVpcConnectionResourceStatus { return v.Status }).(VpcConnectionVpcConnectionResourceStatusPtrOutput)
}

// A map of the key-value pairs for the resource tag or tags assigned to the VPC connection.
func (o LookupVpcConnectionResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// <p>The Amazon EC2 VPC ID associated with the VPC connection.</p>
func (o LookupVpcConnectionResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcConnectionResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcConnectionResultOutput{})
}
