// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Launch.
 */
export function getLaunch(args: GetLaunchArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:evidently:getLaunch", {
        "arn": args.arn,
    }, opts);
}

export interface GetLaunchArgs {
    /**
     * The ARN of the launch. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`
     */
    arn: string;
}

export interface GetLaunchResult {
    /**
     * The ARN of the launch. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`
     */
    readonly arn?: string;
    /**
     * An optional description for the launch.
     */
    readonly description?: string;
    /**
     * Start or Stop Launch Launch. Default is not started.
     */
    readonly executionStatus?: outputs.evidently.LaunchExecutionStatusObject;
    /**
     * An array of structures that contains the feature and variations that are to be used for the launch. You can up to five launch groups in a launch.
     */
    readonly groups?: outputs.evidently.LaunchGroupObject[];
    /**
     * An array of structures that define the metrics that will be used to monitor the launch performance. You can have up to three metric monitors in the array.
     */
    readonly metricMonitors?: outputs.evidently.LaunchMetricDefinitionObject[];
    /**
     * When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the launch name as the `randomizationsSalt` .
     */
    readonly randomizationSalt?: string;
    /**
     * An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.
     */
    readonly scheduledSplitsConfig?: outputs.evidently.LaunchStepConfig[];
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::Evidently::Launch.
 */
export function getLaunchOutput(args: GetLaunchOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLaunchResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:evidently:getLaunch", {
        "arn": args.arn,
    }, opts);
}

export interface GetLaunchOutputArgs {
    /**
     * The ARN of the launch. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`
     */
    arn: pulumi.Input<string>;
}
