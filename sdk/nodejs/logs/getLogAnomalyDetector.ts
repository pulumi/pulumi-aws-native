// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Logs::LogAnomalyDetector resource specifies a CloudWatch Logs LogAnomalyDetector.
 */
export function getLogAnomalyDetector(args: GetLogAnomalyDetectorArgs, opts?: pulumi.InvokeOptions): Promise<GetLogAnomalyDetectorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:logs:getLogAnomalyDetector", {
        "anomalyDetectorArn": args.anomalyDetectorArn,
    }, opts);
}

export interface GetLogAnomalyDetectorArgs {
    /**
     * ARN of LogAnomalyDetector
     */
    anomalyDetectorArn: string;
}

export interface GetLogAnomalyDetectorResult {
    /**
     * ARN of LogAnomalyDetector
     */
    readonly anomalyDetectorArn?: string;
    /**
     * Current status of detector.
     */
    readonly anomalyDetectorStatus?: string;
    readonly anomalyVisibilityTime?: number;
    /**
     * When detector was created.
     */
    readonly creationTimeStamp?: number;
    /**
     * Name of detector
     */
    readonly detectorName?: string;
    /**
     * How often log group is evaluated
     */
    readonly evaluationFrequency?: enums.logs.LogAnomalyDetectorEvaluationFrequency;
    readonly filterPattern?: string;
    /**
     * The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
     */
    readonly kmsKeyId?: string;
    /**
     * When detector was lsat modified.
     */
    readonly lastModifiedTimeStamp?: number;
    /**
     * List of Arns for the given log group
     */
    readonly logGroupArnList?: string[];
}
/**
 * The AWS::Logs::LogAnomalyDetector resource specifies a CloudWatch Logs LogAnomalyDetector.
 */
export function getLogAnomalyDetectorOutput(args: GetLogAnomalyDetectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogAnomalyDetectorResult> {
    return pulumi.output(args).apply((a: any) => getLogAnomalyDetector(a, opts))
}

export interface GetLogAnomalyDetectorOutputArgs {
    /**
     * ARN of LogAnomalyDetector
     */
    anomalyDetectorArn: pulumi.Input<string>;
}