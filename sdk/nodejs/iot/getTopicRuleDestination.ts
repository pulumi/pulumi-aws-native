// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT::TopicRuleDestination
 */
export function getTopicRuleDestination(args: GetTopicRuleDestinationArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicRuleDestinationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iot:getTopicRuleDestination", {
        "arn": args.arn,
    }, opts);
}

export interface GetTopicRuleDestinationArgs {
    /**
     * Amazon Resource Name (ARN).
     */
    arn: string;
}

export interface GetTopicRuleDestinationResult {
    /**
     * Amazon Resource Name (ARN).
     */
    readonly arn?: string;
    /**
     * The status of the TopicRuleDestination.
     */
    readonly status?: enums.iot.TopicRuleDestinationStatus;
    /**
     * The reasoning for the current status of the TopicRuleDestination.
     */
    readonly statusReason?: string;
}

export function getTopicRuleDestinationOutput(args: GetTopicRuleDestinationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicRuleDestinationResult> {
    return pulumi.output(args).apply(a => getTopicRuleDestination(a, opts))
}

export interface GetTopicRuleDestinationOutputArgs {
    /**
     * Amazon Resource Name (ARN).
     */
    arn: pulumi.Input<string>;
}