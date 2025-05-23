// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.
 *  You can allocate an Elastic IP address from an address pool owned by AWS or from an address pool created from a public IPv4 address range that you have brought to AWS for use with your AWS resources using bring your own IP addresses (BYOIP). For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the *Amazon EC2 User Guide*.
 *  For more information, see [Elastic IP Addresses](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) in the *Amazon EC2 User Guide*.
 */
export function getEip(args: GetEipArgs, opts?: pulumi.InvokeOptions): Promise<GetEipResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getEip", {
        "allocationId": args.allocationId,
        "publicIp": args.publicIp,
    }, opts);
}

export interface GetEipArgs {
    /**
     * The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
     */
    allocationId: string;
    /**
     * The Elastic IP address.
     */
    publicIp: string;
}

export interface GetEipResult {
    /**
     * The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
     */
    readonly allocationId?: string;
    /**
     * The network (``vpc``).
     *  If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.
     */
    readonly domain?: string;
    /**
     * The ID of the instance.
     *   Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
     */
    readonly instanceId?: string;
    /**
     * The Elastic IP address.
     */
    readonly publicIp?: string;
    /**
     * The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.
     *   Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
     */
    readonly publicIpv4Pool?: string;
    /**
     * Any tags assigned to the Elastic IP address.
     *   Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.
 *  You can allocate an Elastic IP address from an address pool owned by AWS or from an address pool created from a public IPv4 address range that you have brought to AWS for use with your AWS resources using bring your own IP addresses (BYOIP). For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the *Amazon EC2 User Guide*.
 *  For more information, see [Elastic IP Addresses](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) in the *Amazon EC2 User Guide*.
 */
export function getEipOutput(args: GetEipOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEipResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getEip", {
        "allocationId": args.allocationId,
        "publicIp": args.publicIp,
    }, opts);
}

export interface GetEipOutputArgs {
    /**
     * The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
     */
    allocationId: pulumi.Input<string>;
    /**
     * The Elastic IP address.
     */
    publicIp: pulumi.Input<string>;
}
