// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EMR::Step
 */
export function getStep(args: GetStepArgs, opts?: pulumi.InvokeOptions): Promise<GetStepResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:emr:getStep", {
        "id": args.id,
    }, opts);
}

export interface GetStepArgs {
    id: string;
}

export interface GetStepResult {
    readonly id?: string;
}

export function getStepOutput(args: GetStepOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStepResult> {
    return pulumi.output(args).apply(a => getStep(a, opts))
}

export interface GetStepOutputArgs {
    id: pulumi.Input<string>;
}