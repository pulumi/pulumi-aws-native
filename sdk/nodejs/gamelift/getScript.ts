// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GameLift::Script
 */
export function getScript(args: GetScriptArgs, opts?: pulumi.InvokeOptions): Promise<GetScriptResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:gamelift:getScript", {
        "id": args.id,
    }, opts);
}

export interface GetScriptArgs {
    id: string;
}

export interface GetScriptResult {
    readonly arn?: string;
    readonly id?: string;
    readonly name?: string;
    readonly storageLocation?: outputs.gamelift.ScriptS3Location;
    readonly tags?: outputs.gamelift.ScriptTag[];
    readonly version?: string;
}

export function getScriptOutput(args: GetScriptOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScriptResult> {
    return pulumi.output(args).apply(a => getScript(a, opts))
}

export interface GetScriptOutputArgs {
    id: pulumi.Input<string>;
}