// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TrafficMirrorFilter
 */
export function getTrafficMirrorFilter(args: GetTrafficMirrorFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficMirrorFilterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getTrafficMirrorFilter", {
        "id": args.id,
    }, opts);
}

export interface GetTrafficMirrorFilterArgs {
    id: string;
}

export interface GetTrafficMirrorFilterResult {
    readonly id?: string;
    readonly networkServices?: string[];
    readonly tags?: outputs.ec2.TrafficMirrorFilterTag[];
}

export function getTrafficMirrorFilterOutput(args: GetTrafficMirrorFilterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficMirrorFilterResult> {
    return pulumi.output(args).apply(a => getTrafficMirrorFilter(a, opts))
}

export interface GetTrafficMirrorFilterOutputArgs {
    id: pulumi.Input<string>;
}