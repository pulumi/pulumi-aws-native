// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GlobalAccelerator::Accelerator
 */
export function getAccelerator(args: GetAcceleratorArgs, opts?: pulumi.InvokeOptions): Promise<GetAcceleratorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:globalaccelerator:getAccelerator", {
        "acceleratorArn": args.acceleratorArn,
    }, opts);
}

export interface GetAcceleratorArgs {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    acceleratorArn: string;
}

export interface GetAcceleratorResult {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    readonly acceleratorArn?: string;
    /**
     * The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 addresses.
     */
    readonly dnsName?: string;
    /**
     * The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 and IPv6 addresses.
     */
    readonly dualStackDnsName?: string;
    /**
     * Indicates whether an accelerator is enabled. The value is true or false.
     */
    readonly enabled?: boolean;
    /**
     * IP Address type.
     */
    readonly ipAddressType?: enums.globalaccelerator.AcceleratorIpAddressType;
    /**
     * The IP addresses from BYOIP Prefix pool.
     */
    readonly ipAddresses?: string[];
    /**
     * The IPv4 addresses assigned to the accelerator.
     */
    readonly ipv4Addresses?: string[];
    /**
     * The IPv6 addresses assigned if the accelerator is dualstack
     */
    readonly ipv6Addresses?: string[];
    /**
     * Name of accelerator.
     */
    readonly name?: string;
    readonly tags?: outputs.globalaccelerator.AcceleratorTag[];
}

export function getAcceleratorOutput(args: GetAcceleratorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAcceleratorResult> {
    return pulumi.output(args).apply(a => getAccelerator(a, opts))
}

export interface GetAcceleratorOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    acceleratorArn: pulumi.Input<string>;
}