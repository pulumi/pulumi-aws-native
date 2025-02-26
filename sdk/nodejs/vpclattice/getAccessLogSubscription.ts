// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.
 */
export function getAccessLogSubscription(args: GetAccessLogSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessLogSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:vpclattice:getAccessLogSubscription", {
        "arn": args.arn,
    }, opts);
}

export interface GetAccessLogSubscriptionArgs {
    /**
     * The Amazon Resource Name (ARN) of the access log subscription.
     */
    arn: string;
}

export interface GetAccessLogSubscriptionResult {
    /**
     * The Amazon Resource Name (ARN) of the access log subscription.
     */
    readonly arn?: string;
    /**
     * The Amazon Resource Name (ARN) of the destination. The supported destination types are CloudWatch Log groups, Kinesis Data Firehose delivery streams, and Amazon S3 buckets.
     */
    readonly destinationArn?: string;
    /**
     * The ID of the access log subscription.
     */
    readonly id?: string;
    /**
     * The Amazon Resource Name (ARN) of the access log subscription.
     */
    readonly resourceArn?: string;
    /**
     * The ID of the service network or service.
     */
    readonly resourceId?: string;
    /**
     * Log type of the service network.
     */
    readonly serviceNetworkLogType?: enums.vpclattice.AccessLogSubscriptionServiceNetworkLogType;
    /**
     * The tags for the access log subscription.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.
 */
export function getAccessLogSubscriptionOutput(args: GetAccessLogSubscriptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccessLogSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:vpclattice:getAccessLogSubscription", {
        "arn": args.arn,
    }, opts);
}

export interface GetAccessLogSubscriptionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the access log subscription.
     */
    arn: pulumi.Input<string>;
}
