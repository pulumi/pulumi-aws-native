// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for Activity
 */
export function getActivity(args: GetActivityArgs, opts?: pulumi.InvokeOptions): Promise<GetActivityResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:stepfunctions:getActivity", {
        "arn": args.arn,
    }, opts);
}

export interface GetActivityArgs {
    arn: string;
}

export interface GetActivityResult {
    readonly arn?: string;
    readonly tags?: outputs.stepfunctions.ActivityTagsEntry[];
}

export function getActivityOutput(args: GetActivityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetActivityResult> {
    return pulumi.output(args).apply(a => getActivity(a, opts))
}

export interface GetActivityOutputArgs {
    arn: pulumi.Input<string>;
}