// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::VirtualMFADevice
 */
export function getVirtualMFADevice(args: GetVirtualMFADeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMFADeviceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iam:getVirtualMFADevice", {
        "serialNumber": args.serialNumber,
    }, opts);
}

export interface GetVirtualMFADeviceArgs {
    serialNumber: string;
}

export interface GetVirtualMFADeviceResult {
    readonly serialNumber?: string;
    readonly tags?: outputs.iam.VirtualMFADeviceTag[];
    readonly users?: string[];
}

export function getVirtualMFADeviceOutput(args: GetVirtualMFADeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualMFADeviceResult> {
    return pulumi.output(args).apply(a => getVirtualMFADevice(a, opts))
}

export interface GetVirtualMFADeviceOutputArgs {
    serialNumber: pulumi.Input<string>;
}