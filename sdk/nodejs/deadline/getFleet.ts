// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Deadline::Fleet Resource Type
 */
export function getFleet(args: GetFleetArgs, opts?: pulumi.InvokeOptions): Promise<GetFleetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:deadline:getFleet", {
        "arn": args.arn,
    }, opts);
}

export interface GetFleetArgs {
    arn: string;
}

export interface GetFleetResult {
    readonly arn?: string;
    readonly capabilities?: outputs.deadline.FleetCapabilities;
    readonly configuration?: outputs.deadline.FleetConfiguration0Properties | outputs.deadline.FleetConfiguration1Properties;
    readonly description?: string;
    readonly displayName?: string;
    readonly fleetId?: string;
    readonly maxWorkerCount?: number;
    readonly minWorkerCount?: number;
    readonly roleArn?: string;
    readonly status?: enums.deadline.FleetStatus;
    readonly workerCount?: number;
}
/**
 * Definition of AWS::Deadline::Fleet Resource Type
 */
export function getFleetOutput(args: GetFleetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFleetResult> {
    return pulumi.output(args).apply((a: any) => getFleet(a, opts))
}

export interface GetFleetOutputArgs {
    arn: pulumi.Input<string>;
}