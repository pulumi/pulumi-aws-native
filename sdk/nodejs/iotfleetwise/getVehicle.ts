// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::IoTFleetWise::Vehicle Resource Type
 */
export function getVehicle(args: GetVehicleArgs, opts?: pulumi.InvokeOptions): Promise<GetVehicleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotfleetwise:getVehicle", {
        "name": args.name,
    }, opts);
}

export interface GetVehicleArgs {
    name: string;
}

export interface GetVehicleResult {
    readonly arn?: string;
    readonly attributes?: outputs.iotfleetwise.VehicleattributesMap;
    readonly creationTime?: string;
    readonly decoderManifestArn?: string;
    readonly lastModificationTime?: string;
    readonly modelManifestArn?: string;
    readonly tags?: outputs.iotfleetwise.VehicleTag[];
}

export function getVehicleOutput(args: GetVehicleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVehicleResult> {
    return pulumi.output(args).apply(a => getVehicle(a, opts))
}

export interface GetVehicleOutputArgs {
    name: pulumi.Input<string>;
}