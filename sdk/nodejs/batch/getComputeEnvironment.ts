// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Batch::ComputeEnvironment
 */
export function getComputeEnvironment(args: GetComputeEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:batch:getComputeEnvironment", {
        "computeEnvironmentArn": args.computeEnvironmentArn,
    }, opts);
}

export interface GetComputeEnvironmentArgs {
    computeEnvironmentArn: string;
}

export interface GetComputeEnvironmentResult {
    readonly computeEnvironmentArn?: string;
    readonly computeResources?: outputs.batch.ComputeEnvironmentComputeResources;
    readonly serviceRole?: string;
    readonly state?: string;
    readonly unmanagedvCpus?: number;
}

export function getComputeEnvironmentOutput(args: GetComputeEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeEnvironmentResult> {
    return pulumi.output(args).apply(a => getComputeEnvironment(a, opts))
}

export interface GetComputeEnvironmentOutputArgs {
    computeEnvironmentArn: pulumi.Input<string>;
}