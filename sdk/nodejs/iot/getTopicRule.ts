// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT::TopicRule
 */
export function getTopicRule(args: GetTopicRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iot:getTopicRule", {
        "ruleName": args.ruleName,
    }, opts);
}

export interface GetTopicRuleArgs {
    ruleName: string;
}

export interface GetTopicRuleResult {
    readonly arn?: string;
    readonly tags?: outputs.iot.TopicRuleTag[];
    readonly topicRulePayload?: outputs.iot.TopicRulePayload;
}

export function getTopicRuleOutput(args: GetTopicRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicRuleResult> {
    return pulumi.output(args).apply(a => getTopicRule(a, opts))
}

export interface GetTopicRuleOutputArgs {
    ruleName: pulumi.Input<string>;
}