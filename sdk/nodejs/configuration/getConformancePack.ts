// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.
 */
export function getConformancePack(args: GetConformancePackArgs, opts?: pulumi.InvokeOptions): Promise<GetConformancePackResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:configuration:getConformancePack", {
        "conformancePackName": args.conformancePackName,
    }, opts);
}

export interface GetConformancePackArgs {
    /**
     * Name of the conformance pack which will be assigned as the unique identifier.
     */
    conformancePackName: string;
}

export interface GetConformancePackResult {
    /**
     * A list of ConformancePackInputParameter objects.
     */
    readonly conformancePackInputParameters?: outputs.configuration.ConformancePackInputParameter[];
    /**
     * AWS Config stores intermediate files while processing conformance pack template.
     */
    readonly deliveryS3Bucket?: string;
    /**
     * The prefix for delivery S3 bucket.
     */
    readonly deliveryS3KeyPrefix?: string;
}

export function getConformancePackOutput(args: GetConformancePackOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConformancePackResult> {
    return pulumi.output(args).apply(a => getConformancePack(a, opts))
}

export interface GetConformancePackOutputArgs {
    /**
     * Name of the conformance pack which will be assigned as the unique identifier.
     */
    conformancePackName: pulumi.Input<string>;
}