// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationFSxONTAP.
 */
export function getLocationFSxONTAP(args: GetLocationFSxONTAPArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationFSxONTAPResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getLocationFSxONTAP", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationFSxONTAPArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.
     */
    locationArn: string;
}

export interface GetLocationFSxONTAPResult {
    /**
     * The Amazon Resource Name (ARN) for the FSx ONAP file system.
     */
    readonly fsxFilesystemArn?: string;
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.
     */
    readonly locationArn?: string;
    /**
     * The URL of the FSx ONTAP file system that was described.
     */
    readonly locationUri?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationFSxONTAPTag[];
}

export function getLocationFSxONTAPOutput(args: GetLocationFSxONTAPOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationFSxONTAPResult> {
    return pulumi.output(args).apply(a => getLocationFSxONTAP(a, opts))
}

export interface GetLocationFSxONTAPOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.
     */
    locationArn: pulumi.Input<string>;
}