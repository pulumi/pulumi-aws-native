// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SQS::Queue
 */
export function getQueue(args: GetQueueArgs, opts?: pulumi.InvokeOptions): Promise<GetQueueResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sqs:getQueue", {
        "queueUrl": args.queueUrl,
    }, opts);
}

export interface GetQueueArgs {
    /**
     * URL of the source queue.
     */
    queueUrl: string;
}

export interface GetQueueResult {
    /**
     * Amazon Resource Name (ARN) of the queue.
     */
    readonly arn?: string;
    /**
     * For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
     */
    readonly contentBasedDeduplication?: boolean;
    /**
     * Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
     */
    readonly deduplicationScope?: string;
    /**
     * The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
     */
    readonly delaySeconds?: number;
    /**
     * Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
     */
    readonly fifoThroughputLimit?: string;
    /**
     * The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
     */
    readonly kmsDataKeyReusePeriodSeconds?: number;
    /**
     * The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
     */
    readonly kmsMasterKeyId?: string;
    /**
     * The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
     */
    readonly maximumMessageSize?: number;
    /**
     * The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
     */
    readonly messageRetentionPeriod?: number;
    /**
     * URL of the source queue.
     */
    readonly queueUrl?: string;
    /**
     * Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
     */
    readonly receiveMessageWaitTimeSeconds?: number;
    /**
     * The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
     */
    readonly redriveAllowPolicy?: any;
    /**
     * A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
     */
    readonly redrivePolicy?: any;
    /**
     * Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).
     */
    readonly sqsManagedSseEnabled?: boolean;
    /**
     * The tags that you attach to this queue.
     */
    readonly tags?: outputs.sqs.QueueTag[];
    /**
     * The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
     */
    readonly visibilityTimeout?: number;
}

export function getQueueOutput(args: GetQueueOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueueResult> {
    return pulumi.output(args).apply(a => getQueue(a, opts))
}

export interface GetQueueOutputArgs {
    /**
     * URL of the source queue.
     */
    queueUrl: pulumi.Input<string>;
}