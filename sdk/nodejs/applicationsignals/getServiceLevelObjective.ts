// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApplicationSignals::ServiceLevelObjective
 */
export function getServiceLevelObjective(args: GetServiceLevelObjectiveArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceLevelObjectiveResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:applicationsignals:getServiceLevelObjective", {
        "arn": args.arn,
    }, opts);
}

export interface GetServiceLevelObjectiveArgs {
    /**
     * The ARN of this SLO.
     */
    arn: string;
}

export interface GetServiceLevelObjectiveResult {
    /**
     * The ARN of this SLO.
     */
    readonly arn?: string;
    /**
     * Each object in this array defines the length of the look-back window used to calculate one burn rate metric for this SLO. The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO.
     */
    readonly burnRateConfigurations?: outputs.applicationsignals.ServiceLevelObjectiveBurnRateConfiguration[];
    /**
     * Epoch time in seconds of the time that this SLO was created
     */
    readonly createdTime?: number;
    /**
     * An optional description for this SLO. Default is 'No description'
     */
    readonly description?: string;
    /**
     * Displays whether this is a period-based SLO or a request-based SLO.
     */
    readonly evaluationType?: enums.applicationsignals.ServiceLevelObjectiveEvaluationType;
    /**
     * The time window to be excluded from the SLO performance metrics.
     */
    readonly exclusionWindows?: outputs.applicationsignals.ServiceLevelObjectiveExclusionWindow[];
    /**
     * This structure contains the attributes that determine the goal of an SLO. This includes the time period for evaluation and the attainment threshold.
     */
    readonly goal?: outputs.applicationsignals.ServiceLevelObjectiveGoal;
    /**
     * Epoch time in seconds of the time that this SLO was most recently updated
     */
    readonly lastUpdatedTime?: number;
    /**
     * A structure containing information about the performance metric that this SLO monitors, if this is a request-based SLO.
     */
    readonly requestBasedSli?: outputs.applicationsignals.ServiceLevelObjectiveRequestBasedSli;
    /**
     * A structure containing information about the performance metric that this SLO monitors, if this is a period-based SLO.
     */
    readonly sli?: outputs.applicationsignals.ServiceLevelObjectiveSli;
    /**
     * A list of key-value pairs to associate with the SLO. You can associate as many as 50 tags with an SLO. To be able to associate tags with the SLO when you create the SLO, you must have the cloudwatch:TagResource permission.
     *
     * Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::ApplicationSignals::ServiceLevelObjective
 */
export function getServiceLevelObjectiveOutput(args: GetServiceLevelObjectiveOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceLevelObjectiveResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:applicationsignals:getServiceLevelObjective", {
        "arn": args.arn,
    }, opts);
}

export interface GetServiceLevelObjectiveOutputArgs {
    /**
     * The ARN of this SLO.
     */
    arn: pulumi.Input<string>;
}
