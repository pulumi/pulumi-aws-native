// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates an Elastic IP address with an instance or a network interface. Before you can use an Elastic IP address, you must allocate it to your account. For more information about working with Elastic IP addresses, see [Elastic IP address concepts and rules](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#vpc-eip-overview).
 *  You must specify ``AllocationId`` and either ``InstanceId``, ``NetworkInterfaceId``, or ``PrivateIpAddress``.
 */
export function getEipAssociation(args: GetEipAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetEipAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getEipAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetEipAssociationArgs {
    /**
     * The ID of the association.
     */
    id: string;
}

export interface GetEipAssociationResult {
    /**
     * The ID of the association.
     */
    readonly id?: string;
}
/**
 * Associates an Elastic IP address with an instance or a network interface. Before you can use an Elastic IP address, you must allocate it to your account. For more information about working with Elastic IP addresses, see [Elastic IP address concepts and rules](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#vpc-eip-overview).
 *  You must specify ``AllocationId`` and either ``InstanceId``, ``NetworkInterfaceId``, or ``PrivateIpAddress``.
 */
export function getEipAssociationOutput(args: GetEipAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEipAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getEipAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetEipAssociationOutputArgs {
    /**
     * The ID of the association.
     */
    id: pulumi.Input<string>;
}
