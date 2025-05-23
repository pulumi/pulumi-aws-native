// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Schedule.
 */
export function getSchedule(args: GetScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:databrew:getSchedule", {
        "name": args.name,
    }, opts);
}

export interface GetScheduleArgs {
    /**
     * Schedule Name
     */
    name: string;
}

export interface GetScheduleResult {
    /**
     * Schedule cron
     */
    readonly cronExpression?: string;
    /**
     * A list of jobs to be run, according to the schedule.
     */
    readonly jobNames?: string[];
    /**
     * Metadata tags that have been applied to the schedule.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::DataBrew::Schedule.
 */
export function getScheduleOutput(args: GetScheduleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetScheduleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:databrew:getSchedule", {
        "name": args.name,
    }, opts);
}

export interface GetScheduleOutputArgs {
    /**
     * Schedule Name
     */
    name: pulumi.Input<string>;
}
