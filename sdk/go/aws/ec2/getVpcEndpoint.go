// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS , an AWS Marketplace Partner, or another AWS accounts in your organization. For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
//
// An endpoint of type `Interface` establishes connections between the subnets in your VPC and an AWS service , your own service, or a service hosted by another AWS account . With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
//
// An endpoint of type `gateway` serves as a target for a route in your route table for traffic destined for Amazon S3 or DynamoDB . You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to Amazon S3 , see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
//
// An endpoint of type `GatewayLoadBalancer` provides private connectivity between your VPC and virtual appliances from a service provider.
func LookupVpcEndpoint(ctx *pulumi.Context, args *LookupVpcEndpointArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointResult
	err := ctx.Invoke("aws-native:ec2:getVpcEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcEndpointArgs struct {
	// The ID of the VPC endpoint.
	Id string `pulumi:"id"`
}

type LookupVpcEndpointResult struct {
	// The date and time the VPC endpoint was created. For example: `Fri Sep 28 23:34:36 UTC 2018.`
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// (Interface endpoints) The DNS entries for the endpoint. Each entry is a combination of the hosted zone ID and the DNS name. The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS. This order is not enforced for AWS Marketplace services.
	//
	// The following is an example. In the first entry, the hosted zone ID is Z1HUB23UULQXV and the DNS name is vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com.
	//
	// ["Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com", "Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3-us-east-1a.ec2.us-east-1.vpce.amazonaws.com", "Z1C12344VYDITB0:ec2.us-east-1.amazonaws.com"]
	//
	// If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the DNS entries in the list will change.
	DnsEntries []string `pulumi:"dnsEntries"`
	// Describes the DNS options for an endpoint.
	DnsOptions *VpcEndpointDnsOptionsSpecification `pulumi:"dnsOptions"`
	// The ID of the VPC endpoint.
	Id *string `pulumi:"id"`
	// The supported IP address types.
	IpAddressType *VpcEndpointIpAddressType `pulumi:"ipAddressType"`
	// (Interface endpoints) The network interface IDs. If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the items in this list might change.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
	//  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:
	//  ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ "Version":"2012-10-17", "Statement": [{ "Effect":"Allow", "Principal":"*", "Action":["logs:Describe*","logs:Get*","logs:List*","logs:FilterLogEvents"], "Resource":"*" }] }'``
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
	//  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
	//  This property is supported only for interface endpoints.
	//  Default: ``false``
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// The IDs of the route tables. Routing is supported only for gateway endpoints.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
	SubnetIds []string `pulumi:"subnetIds"`
	// The tags to associate with the endpoint.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupVpcEndpointOutput(ctx *pulumi.Context, args LookupVpcEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointResultOutput, error) {
			args := v.(LookupVpcEndpointArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getVpcEndpoint", args, LookupVpcEndpointResultOutput{}, options).(LookupVpcEndpointResultOutput), nil
		}).(LookupVpcEndpointResultOutput)
}

type LookupVpcEndpointOutputArgs struct {
	// The ID of the VPC endpoint.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVpcEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointArgs)(nil)).Elem()
}

type LookupVpcEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointResult)(nil)).Elem()
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutput() LookupVpcEndpointResultOutput {
	return o
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutputWithContext(ctx context.Context) LookupVpcEndpointResultOutput {
	return o
}

// The date and time the VPC endpoint was created. For example: `Fri Sep 28 23:34:36 UTC 2018.`
func (o LookupVpcEndpointResultOutput) CreationTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *string { return v.CreationTimestamp }).(pulumi.StringPtrOutput)
}

// (Interface endpoints) The DNS entries for the endpoint. Each entry is a combination of the hosted zone ID and the DNS name. The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS. This order is not enforced for AWS Marketplace services.
//
// The following is an example. In the first entry, the hosted zone ID is Z1HUB23UULQXV and the DNS name is vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com.
//
// ["Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com", "Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3-us-east-1a.ec2.us-east-1.vpce.amazonaws.com", "Z1C12344VYDITB0:ec2.us-east-1.amazonaws.com"]
//
// If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the DNS entries in the list will change.
func (o LookupVpcEndpointResultOutput) DnsEntries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.DnsEntries }).(pulumi.StringArrayOutput)
}

// Describes the DNS options for an endpoint.
func (o LookupVpcEndpointResultOutput) DnsOptions() VpcEndpointDnsOptionsSpecificationPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *VpcEndpointDnsOptionsSpecification { return v.DnsOptions }).(VpcEndpointDnsOptionsSpecificationPtrOutput)
}

// The ID of the VPC endpoint.
func (o LookupVpcEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The supported IP address types.
func (o LookupVpcEndpointResultOutput) IpAddressType() VpcEndpointIpAddressTypePtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *VpcEndpointIpAddressType { return v.IpAddressType }).(VpcEndpointIpAddressTypePtrOutput)
}

// (Interface endpoints) The network interface IDs. If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the items in this list might change.
func (o LookupVpcEndpointResultOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
//
//	For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:
//	``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ "Version":"2012-10-17", "Statement": [{ "Effect":"Allow", "Principal":"*", "Action":["logs:Describe*","logs:Get*","logs:List*","logs:FilterLogEvents"], "Resource":"*" }] }'``
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
func (o LookupVpcEndpointResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, “kinesis.us-east-1.amazonaws.com“), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
//
//	To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
//	This property is supported only for interface endpoints.
//	Default: ``false``
func (o LookupVpcEndpointResultOutput) PrivateDnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *bool { return v.PrivateDnsEnabled }).(pulumi.BoolPtrOutput)
}

// The IDs of the route tables. Routing is supported only for gateway endpoints.
func (o LookupVpcEndpointResultOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
func (o LookupVpcEndpointResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
func (o LookupVpcEndpointResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The tags to associate with the endpoint.
func (o LookupVpcEndpointResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointResultOutput{})
}
