// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.
 */
export function getSlackChannelConfiguration(args: GetSlackChannelConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetSlackChannelConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:supportapp:getSlackChannelConfiguration", {
        "channelId": args.channelId,
        "teamId": args.teamId,
    }, opts);
}

export interface GetSlackChannelConfigurationArgs {
    /**
     * The channel ID in Slack, which identifies a channel within a workspace.
     */
    channelId: string;
    /**
     * The team ID in Slack, which uniquely identifies a workspace.
     */
    teamId: string;
}

export interface GetSlackChannelConfigurationResult {
    /**
     * The channel name in Slack.
     */
    readonly channelName?: string;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
     */
    readonly channelRoleArn?: string;
    /**
     * Whether to notify when a correspondence is added to a case.
     */
    readonly notifyOnAddCorrespondenceToCase?: boolean;
    /**
     * The severity level of a support case that a customer wants to get notified for.
     */
    readonly notifyOnCaseSeverity?: enums.supportapp.SlackChannelConfigurationNotifyOnCaseSeverity;
    /**
     * Whether to notify when a case is created or reopened.
     */
    readonly notifyOnCreateOrReopenCase?: boolean;
    /**
     * Whether to notify when a case is resolved.
     */
    readonly notifyOnResolveCase?: boolean;
}
/**
 * An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.
 */
export function getSlackChannelConfigurationOutput(args: GetSlackChannelConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSlackChannelConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:supportapp:getSlackChannelConfiguration", {
        "channelId": args.channelId,
        "teamId": args.teamId,
    }, opts);
}

export interface GetSlackChannelConfigurationOutputArgs {
    /**
     * The channel ID in Slack, which identifies a channel within a workspace.
     */
    channelId: pulumi.Input<string>;
    /**
     * The team ID in Slack, which uniquely identifies a workspace.
     */
    teamId: pulumi.Input<string>;
}
