// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::RefactorSpaces::Environment Resource Type
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:refactorspaces:getEnvironment", {
        "environmentIdentifier": args.environmentIdentifier,
    }, opts);
}

export interface GetEnvironmentArgs {
    environmentIdentifier: string;
}

export interface GetEnvironmentResult {
    readonly arn?: string;
    readonly environmentIdentifier?: string;
    /**
     * Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
     */
    readonly tags?: outputs.refactorspaces.EnvironmentTag[];
    readonly transitGatewayId?: string;
}

export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentResult> {
    return pulumi.output(args).apply(a => getEnvironment(a, opts))
}

export interface GetEnvironmentOutputArgs {
    environmentIdentifier: pulumi.Input<string>;
}