// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Schema of AWS::EC2::IPAMAllocation Type
 */
export function getIpamAllocation(args: GetIpamAllocationArgs, opts?: pulumi.InvokeOptions): Promise<GetIpamAllocationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getIpamAllocation", {
        "cidr": args.cidr,
        "ipamPoolAllocationId": args.ipamPoolAllocationId,
        "ipamPoolId": args.ipamPoolId,
    }, opts);
}

export interface GetIpamAllocationArgs {
    cidr: string;
    /**
     * Id of the allocation.
     */
    ipamPoolAllocationId: string;
    /**
     * Id of the IPAM Pool.
     */
    ipamPoolId: string;
}

export interface GetIpamAllocationResult {
    /**
     * Id of the allocation.
     */
    readonly ipamPoolAllocationId?: string;
}
/**
 * Resource Schema of AWS::EC2::IPAMAllocation Type
 */
export function getIpamAllocationOutput(args: GetIpamAllocationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpamAllocationResult> {
    return pulumi.output(args).apply((a: any) => getIpamAllocation(a, opts))
}

export interface GetIpamAllocationOutputArgs {
    cidr: pulumi.Input<string>;
    /**
     * Id of the allocation.
     */
    ipamPoolAllocationId: pulumi.Input<string>;
    /**
     * Id of the IPAM Pool.
     */
    ipamPoolId: pulumi.Input<string>;
}