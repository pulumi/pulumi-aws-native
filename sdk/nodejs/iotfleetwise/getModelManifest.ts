// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::IoTFleetWise::ModelManifest Resource Type
 */
export function getModelManifest(args: GetModelManifestArgs, opts?: pulumi.InvokeOptions): Promise<GetModelManifestResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotfleetwise:getModelManifest", {
        "name": args.name,
    }, opts);
}

export interface GetModelManifestArgs {
    name: string;
}

export interface GetModelManifestResult {
    readonly arn?: string;
    readonly creationTime?: string;
    readonly description?: string;
    readonly lastModificationTime?: string;
    readonly nodes?: string[];
    readonly signalCatalogArn?: string;
    readonly status?: enums.iotfleetwise.ModelManifestManifestStatus;
    readonly tags?: outputs.iotfleetwise.ModelManifestTag[];
}

export function getModelManifestOutput(args: GetModelManifestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelManifestResult> {
    return pulumi.output(args).apply(a => getModelManifest(a, opts))
}

export interface GetModelManifestOutputArgs {
    name: pulumi.Input<string>;
}