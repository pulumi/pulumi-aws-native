// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * <p>Represents a channel group that facilitates the grouping of multiple channels.</p>
 */
export function getChannelGroup(args: GetChannelGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:mediapackagev2:getChannelGroup", {
        "arn": args.arn,
    }, opts);
}

export interface GetChannelGroupArgs {
    /**
     * <p>The Amazon Resource Name (ARN) associated with the resource.</p>
     */
    arn: string;
}

export interface GetChannelGroupResult {
    /**
     * <p>The Amazon Resource Name (ARN) associated with the resource.</p>
     */
    readonly arn?: string;
    /**
     * <p>The date and time the channel group was created.</p>
     */
    readonly createdAt?: string;
    /**
     * <p>Enter any descriptive text that helps you to identify the channel group.</p>
     */
    readonly description?: string;
    /**
     * <p>The output domain where the source stream should be sent. Integrate the domain with a downstream CDN (such as Amazon CloudFront) or playback device.</p>
     */
    readonly egressDomain?: string;
    /**
     * <p>The date and time the channel group was modified.</p>
     */
    readonly modifiedAt?: string;
    /**
     * The tags associated with the channel group.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * <p>Represents a channel group that facilitates the grouping of multiple channels.</p>
 */
export function getChannelGroupOutput(args: GetChannelGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetChannelGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:mediapackagev2:getChannelGroup", {
        "arn": args.arn,
    }, opts);
}

export interface GetChannelGroupOutputArgs {
    /**
     * <p>The Amazon Resource Name (ARN) associated with the resource.</p>
     */
    arn: pulumi.Input<string>;
}
