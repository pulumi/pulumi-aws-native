// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS , an AWS Marketplace Partner, or another AWS accounts in your organization. For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
 *
 * An endpoint of type `Interface` establishes connections between the subnets in your VPC and an AWS service , your own service, or a service hosted by another AWS account . With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
 *
 * An endpoint of type `gateway` serves as a target for a route in your route table for traffic destined for Amazon S3 or DynamoDB . You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to Amazon S3 , see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint)
 *
 * An endpoint of type `GatewayLoadBalancer` provides private connectivity between your VPC and virtual appliances from a service provider.
 */
export class VpcEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VpcEndpoint {
        return new VpcEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:VpcEndpoint';

    /**
     * Returns true if the given object is an instance of VpcEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpoint.__pulumiType;
    }

    /**
     * The ID of the VPC endpoint.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The date and time the VPC endpoint was created. For example: `Fri Sep 28 23:34:36 UTC 2018.`
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * (Interface endpoints) The DNS entries for the endpoint. Each entry is a combination of the hosted zone ID and the DNS name. The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS. This order is not enforced for AWS Marketplace services.
     *
     * The following is an example. In the first entry, the hosted zone ID is Z1HUB23UULQXV and the DNS name is vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com.
     *
     * ["Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com", "Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3-us-east-1a.ec2.us-east-1.vpce.amazonaws.com", "Z1C12344VYDITB0:ec2.us-east-1.amazonaws.com"]
     *
     * If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the DNS entries in the list will change.
     */
    public /*out*/ readonly dnsEntries!: pulumi.Output<string[]>;
    /**
     * Describes the DNS options for an endpoint.
     */
    public readonly dnsOptions!: pulumi.Output<outputs.ec2.VpcEndpointDnsOptionsSpecification | undefined>;
    /**
     * The supported IP address types.
     */
    public readonly ipAddressType!: pulumi.Output<enums.ec2.VpcEndpointIpAddressType | undefined>;
    /**
     * (Interface endpoints) The network interface IDs. If you update the `PrivateDnsEnabled` or `SubnetIds` properties, the items in this list might change.
     */
    public /*out*/ readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
     *  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:
     *  ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ "Version":"2012-10-17", "Statement": [{ "Effect":"Allow", "Principal":"*", "Action":["logs:Describe*","logs:Get*","logs:List*","logs:FilterLogEvents"], "Resource":"*" }] }'``
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
     */
    public readonly policyDocument!: pulumi.Output<any | undefined>;
    /**
     * Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
     *  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
     *  This property is supported only for interface endpoints.
     *  Default: ``false``
     */
    public readonly privateDnsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the resource configuration.
     */
    public readonly resourceConfigurationArn!: pulumi.Output<string | undefined>;
    /**
     * The IDs of the route tables. Routing is supported only for gateway endpoints.
     */
    public readonly routeTableIds!: pulumi.Output<string[] | undefined>;
    /**
     * The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the endpoint service.
     */
    public readonly serviceName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the service network.
     */
    public readonly serviceNetworkArn!: pulumi.Output<string | undefined>;
    /**
     * Describes a Region.
     */
    public readonly serviceRegion!: pulumi.Output<string | undefined>;
    /**
     * The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * The tags to associate with the endpoint.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The type of endpoint.
     *  Default: Gateway
     */
    public readonly vpcEndpointType!: pulumi.Output<enums.ec2.VpcEndpointType | undefined>;
    /**
     * The ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["dnsOptions"] = args ? args.dnsOptions : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["privateDnsEnabled"] = args ? args.privateDnsEnabled : undefined;
            resourceInputs["resourceConfigurationArn"] = args ? args.resourceConfigurationArn : undefined;
            resourceInputs["routeTableIds"] = args ? args.routeTableIds : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["serviceNetworkArn"] = args ? args.serviceNetworkArn : undefined;
            resourceInputs["serviceRegion"] = args ? args.serviceRegion : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcEndpointType"] = args ? args.vpcEndpointType : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["dnsEntries"] = undefined /*out*/;
            resourceInputs["networkInterfaceIds"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["dnsEntries"] = undefined /*out*/;
            resourceInputs["dnsOptions"] = undefined /*out*/;
            resourceInputs["ipAddressType"] = undefined /*out*/;
            resourceInputs["networkInterfaceIds"] = undefined /*out*/;
            resourceInputs["policyDocument"] = undefined /*out*/;
            resourceInputs["privateDnsEnabled"] = undefined /*out*/;
            resourceInputs["resourceConfigurationArn"] = undefined /*out*/;
            resourceInputs["routeTableIds"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["serviceNetworkArn"] = undefined /*out*/;
            resourceInputs["serviceRegion"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcEndpointType"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["resourceConfigurationArn", "serviceName", "serviceNetworkArn", "serviceRegion", "vpcEndpointType", "vpcId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(VpcEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VpcEndpoint resource.
 */
export interface VpcEndpointArgs {
    /**
     * Describes the DNS options for an endpoint.
     */
    dnsOptions?: pulumi.Input<inputs.ec2.VpcEndpointDnsOptionsSpecificationArgs>;
    /**
     * The supported IP address types.
     */
    ipAddressType?: pulumi.Input<enums.ec2.VpcEndpointIpAddressType>;
    /**
     * An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
     *  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:
     *  ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ "Version":"2012-10-17", "Statement": [{ "Effect":"Allow", "Principal":"*", "Action":["logs:Describe*","logs:Get*","logs:List*","logs:FilterLogEvents"], "Resource":"*" }] }'``
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
     */
    policyDocument?: any;
    /**
     * Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
     *  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
     *  This property is supported only for interface endpoints.
     *  Default: ``false``
     */
    privateDnsEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the resource configuration.
     */
    resourceConfigurationArn?: pulumi.Input<string>;
    /**
     * The IDs of the route tables. Routing is supported only for gateway endpoints.
     */
    routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the endpoint service.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the service network.
     */
    serviceNetworkArn?: pulumi.Input<string>;
    /**
     * Describes a Region.
     */
    serviceRegion?: pulumi.Input<string>;
    /**
     * The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags to associate with the endpoint.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The type of endpoint.
     *  Default: Gateway
     */
    vpcEndpointType?: pulumi.Input<enums.ec2.VpcEndpointType>;
    /**
     * The ID of the VPC.
     */
    vpcId: pulumi.Input<string>;
}
