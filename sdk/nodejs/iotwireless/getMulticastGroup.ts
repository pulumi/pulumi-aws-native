// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create and manage Multicast groups.
 */
export function getMulticastGroup(args: GetMulticastGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetMulticastGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotwireless:getMulticastGroup", {
        "id": args.id,
    }, opts);
}

export interface GetMulticastGroupArgs {
    /**
     * Multicast group id. Returned after successful create.
     */
    id: string;
}

export interface GetMulticastGroupResult {
    /**
     * Multicast group arn. Returned after successful create.
     */
    readonly arn?: string;
    /**
     * Wireless device to associate. Only for update request.
     */
    readonly associateWirelessDevice?: string;
    /**
     * Multicast group description
     */
    readonly description?: string;
    /**
     * Wireless device to disassociate. Only for update request.
     */
    readonly disassociateWirelessDevice?: string;
    /**
     * Multicast group id. Returned after successful create.
     */
    readonly id?: string;
    /**
     * Multicast group LoRaWAN
     */
    readonly loRaWAN?: outputs.iotwireless.MulticastGroupLoRaWAN;
    /**
     * Name of Multicast group
     */
    readonly name?: string;
    /**
     * Multicast group status. Returned after successful read.
     */
    readonly status?: string;
    /**
     * A list of key-value pairs that contain metadata for the Multicast group.
     */
    readonly tags?: outputs.iotwireless.MulticastGroupTag[];
}

export function getMulticastGroupOutput(args: GetMulticastGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMulticastGroupResult> {
    return pulumi.output(args).apply(a => getMulticastGroup(a, opts))
}

export interface GetMulticastGroupOutputArgs {
    /**
     * Multicast group id. Returned after successful create.
     */
    id: pulumi.Input<string>;
}