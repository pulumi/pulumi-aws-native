// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SecurityLake::SubscriberNotification
 */
export function getSubscriberNotification(args: GetSubscriberNotificationArgs, opts?: pulumi.InvokeOptions): Promise<GetSubscriberNotificationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:securitylake:getSubscriberNotification", {
        "subscriberArn": args.subscriberArn,
    }, opts);
}

export interface GetSubscriberNotificationArgs {
    /**
     * The ARN for the subscriber
     */
    subscriberArn: string;
}

export interface GetSubscriberNotificationResult {
    /**
     * Specify the configurations you want to use for subscriber notification. The subscriber is notified when new data is written to the data lake for sources that the subscriber consumes in Security Lake .
     */
    readonly notificationConfiguration?: outputs.securitylake.SubscriberNotificationNotificationConfiguration;
    /**
     * The endpoint the subscriber should listen to for notifications
     */
    readonly subscriberEndpoint?: string;
}
/**
 * Resource Type definition for AWS::SecurityLake::SubscriberNotification
 */
export function getSubscriberNotificationOutput(args: GetSubscriberNotificationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubscriberNotificationResult> {
    return pulumi.output(args).apply((a: any) => getSubscriberNotification(a, opts))
}

export interface GetSubscriberNotificationOutputArgs {
    /**
     * The ARN for the subscriber
     */
    subscriberArn: pulumi.Input<string>;
}