// *** WARNING: this file was generated by pulumi-language-nodejs. ***
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
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getTopicRule", {
        "ruleName": args.ruleName,
    }, opts);
}

export interface GetTopicRuleArgs {
    /**
     * The name of the rule.
     */
    ruleName: string;
}

export interface GetTopicRuleResult {
    /**
     * The Amazon Resource Name (ARN) of the AWS IoT rule, such as `arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule` .
     */
    readonly arn?: string;
    /**
     * Metadata which can be used to manage the topic rule.
     *
     * > For URI Request parameters use format: ...key1=value1&key2=value2...
     * > 
     * > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
     * > 
     * > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
     */
    readonly tags?: outputs.Tag[];
    /**
     * The rule payload.
     */
    readonly topicRulePayload?: outputs.iot.TopicRulePayload;
}
/**
 * Resource Type definition for AWS::IoT::TopicRule
 */
export function getTopicRuleOutput(args: GetTopicRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTopicRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iot:getTopicRule", {
        "ruleName": args.ruleName,
    }, opts);
}

export interface GetTopicRuleOutputArgs {
    /**
     * The name of the rule.
     */
    ruleName: pulumi.Input<string>;
}
