// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
 */
export function getRobot(args: GetRobotArgs, opts?: pulumi.InvokeOptions): Promise<GetRobotResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:robomaker:getRobot", {
        "arn": args.arn,
    }, opts);
}

export interface GetRobotArgs {
    arn: string;
}

export interface GetRobotResult {
    readonly arn?: string;
    readonly tags?: outputs.robomaker.RobotTags;
}

export function getRobotOutput(args: GetRobotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRobotResult> {
    return pulumi.output(args).apply(a => getRobot(a, opts))
}

export interface GetRobotOutputArgs {
    arn: pulumi.Input<string>;
}