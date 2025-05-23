// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetVpcEndpoint
    {
        /// <summary>
        /// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS , an AWS Marketplace Partner, or another AWS accounts in your organization. For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
        /// 
        /// An endpoint of type `Interface` establishes connections between the subnets in your VPC and an AWS service , your own service, or a service hosted by another AWS account . With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
        /// 
        /// An endpoint of type `gateway` serves as a target for a route in your route table for traffic destined for Amazon S3 or DynamoDB . You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to Amazon S3 , see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
        /// 
        /// An endpoint of type `GatewayLoadBalancer` provides private connectivity between your VPC and virtual appliances from a service provider.
        /// </summary>
        public static Task<GetVpcEndpointResult> InvokeAsync(GetVpcEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcEndpointResult>("aws-native:ec2:getVpcEndpoint", args ?? new GetVpcEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS , an AWS Marketplace Partner, or another AWS accounts in your organization. For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
        /// 
        /// An endpoint of type `Interface` establishes connections between the subnets in your VPC and an AWS service , your own service, or a service hosted by another AWS account . With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
        /// 
        /// An endpoint of type `gateway` serves as a target for a route in your route table for traffic destined for Amazon S3 or DynamoDB . You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to Amazon S3 , see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
        /// 
        /// An endpoint of type `GatewayLoadBalancer` provides private connectivity between your VPC and virtual appliances from a service provider.
        /// </summary>
        public static Output<GetVpcEndpointResult> Invoke(GetVpcEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcEndpointResult>("aws-native:ec2:getVpcEndpoint", args ?? new GetVpcEndpointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS , an AWS Marketplace Partner, or another AWS accounts in your organization. For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
        /// 
        /// An endpoint of type `Interface` establishes connections between the subnets in your VPC and an AWS service , your own service, or a service hosted by another AWS account . With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
        /// 
        /// An endpoint of type `gateway` serves as a target for a route in your route table for traffic destined for Amazon S3 or DynamoDB . You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to Amazon S3 , see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
        /// 
        /// An endpoint of type `GatewayLoadBalancer` provides private connectivity between your VPC and virtual appliances from a service provider.
        /// </summary>
        public static Output<GetVpcEndpointResult> Invoke(GetVpcEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcEndpointResult>("aws-native:ec2:getVpcEndpoint", args ?? new GetVpcEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the VPC endpoint.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVpcEndpointArgs()
        {
        }
        public static new GetVpcEndpointArgs Empty => new GetVpcEndpointArgs();
    }

    public sealed class GetVpcEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the VPC endpoint.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVpcEndpointInvokeArgs()
        {
        }
        public static new GetVpcEndpointInvokeArgs Empty => new GetVpcEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcEndpointResult
    {
        /// <summary>
        /// The date and time the VPC endpoint was created. For example: `Fri Sep 28 23:34:36 UTC 2018.`
        /// </summary>
        public readonly string? CreationTimestamp;
        /// <summary>
        /// (Interface endpoints) The DNS entries for the endpoint. Each entry is a combination of the hosted zone ID and the DNS name. The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS. This order is not enforced for AWS Marketplace services.
        /// 
        /// The following is an example. In the first entry, the hosted zone ID is Z1HUB23UULQXV and the DNS name is vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com.
        /// 
        /// ["Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com", "Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3-us-east-1a.ec2.us-east-1.vpce.amazonaws.com", "Z1C12344VYDITB0:ec2.us-east-1.amazonaws.com"]
        /// 
        /// If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the DNS entries in the list will change.
        /// </summary>
        public readonly ImmutableArray<string> DnsEntries;
        /// <summary>
        /// Describes the DNS options for an endpoint.
        /// </summary>
        public readonly Outputs.VpcEndpointDnsOptionsSpecification? DnsOptions;
        /// <summary>
        /// The ID of the VPC endpoint.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The supported IP address types.
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.VpcEndpointIpAddressType? IpAddressType;
        /// <summary>
        /// (Interface endpoints) The network interface IDs. If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the items in this list might change.
        /// </summary>
        public readonly ImmutableArray<string> NetworkInterfaceIds;
        /// <summary>
        /// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
        ///  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:
        ///  ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ "Version":"2012-10-17", "Statement": [{ "Effect":"Allow", "Principal":"*", "Action":["logs:Describe*","logs:Get*","logs:List*","logs:FilterLogEvents"], "Resource":"*" }] }'``
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? PolicyDocument;
        /// <summary>
        /// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
        ///  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
        ///  This property is supported only for interface endpoints.
        ///  Default: ``false``
        /// </summary>
        public readonly bool? PrivateDnsEnabled;
        /// <summary>
        /// The IDs of the route tables. Routing is supported only for gateway endpoints.
        /// </summary>
        public readonly ImmutableArray<string> RouteTableIds;
        /// <summary>
        /// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The tags to associate with the endpoint.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetVpcEndpointResult(
            string? creationTimestamp,

            ImmutableArray<string> dnsEntries,

            Outputs.VpcEndpointDnsOptionsSpecification? dnsOptions,

            string? id,

            Pulumi.AwsNative.Ec2.VpcEndpointIpAddressType? ipAddressType,

            ImmutableArray<string> networkInterfaceIds,

            object? policyDocument,

            bool? privateDnsEnabled,

            ImmutableArray<string> routeTableIds,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            CreationTimestamp = creationTimestamp;
            DnsEntries = dnsEntries;
            DnsOptions = dnsOptions;
            Id = id;
            IpAddressType = ipAddressType;
            NetworkInterfaceIds = networkInterfaceIds;
            PolicyDocument = policyDocument;
            PrivateDnsEnabled = privateDnsEnabled;
            RouteTableIds = routeTableIds;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            Tags = tags;
        }
    }
}
