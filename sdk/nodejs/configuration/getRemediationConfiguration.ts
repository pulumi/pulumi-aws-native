// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Config::RemediationConfiguration
 */
export function getRemediationConfiguration(args: GetRemediationConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetRemediationConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:configuration:getRemediationConfiguration", {
        "id": args.id,
    }, opts);
}

export interface GetRemediationConfigurationArgs {
    id: string;
}

export interface GetRemediationConfigurationResult {
    readonly automatic?: boolean;
    readonly executionControls?: outputs.configuration.RemediationConfigurationExecutionControls;
    readonly id?: string;
    readonly maximumAutomaticAttempts?: number;
    readonly parameters?: any;
    readonly resourceType?: string;
    readonly retryAttemptSeconds?: number;
    readonly targetId?: string;
    readonly targetType?: string;
    readonly targetVersion?: string;
}

export function getRemediationConfigurationOutput(args: GetRemediationConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRemediationConfigurationResult> {
    return pulumi.output(args).apply(a => getRemediationConfiguration(a, opts))
}

export interface GetRemediationConfigurationOutputArgs {
    id: pulumi.Input<string>;
}