// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This resource schema represents the ResourceCollection resource in the Amazon DevOps Guru.
 */
export function getResourceCollection(args: GetResourceCollectionArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceCollectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:devopsguru:getResourceCollection", {
        "resourceCollectionType": args.resourceCollectionType,
    }, opts);
}

export interface GetResourceCollectionArgs {
    /**
     * The type of ResourceCollection
     */
    resourceCollectionType: enums.devopsguru.ResourceCollectionType;
}

export interface GetResourceCollectionResult {
    readonly resourceCollectionFilter?: outputs.devopsguru.ResourceCollectionFilter;
    /**
     * The type of ResourceCollection
     */
    readonly resourceCollectionType?: enums.devopsguru.ResourceCollectionType;
}

export function getResourceCollectionOutput(args: GetResourceCollectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceCollectionResult> {
    return pulumi.output(args).apply(a => getResourceCollection(a, opts))
}

export interface GetResourceCollectionOutputArgs {
    /**
     * The type of ResourceCollection
     */
    resourceCollectionType: pulumi.Input<enums.devopsguru.ResourceCollectionType>;
}