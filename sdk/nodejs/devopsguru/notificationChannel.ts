// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myNotificationChannel = new aws_native.devopsguru.NotificationChannel("myNotificationChannel", {config: {
 *     filters: {
 *         messageTypes: [
 *             aws_native.devopsguru.NotificationChannelNotificationMessageType.NewInsight,
 *             aws_native.devopsguru.NotificationChannelNotificationMessageType.ClosedInsight,
 *             aws_native.devopsguru.NotificationChannelNotificationMessageType.SeverityUpgraded,
 *         ],
 *         severities: [
 *             aws_native.devopsguru.NotificationChannelInsightSeverity.Medium,
 *             aws_native.devopsguru.NotificationChannelInsightSeverity.High,
 *         ],
 *     },
 *     sns: {
 *         topicArn: "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel",
 *     },
 * }});
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myNotificationChannel1 = new aws_native.devopsguru.NotificationChannel("myNotificationChannel1", {config: {
 *     sns: {
 *         topicArn: "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel",
 *     },
 * }});
 * const myNotificationChannel2 = new aws_native.devopsguru.NotificationChannel("myNotificationChannel2", {config: {
 *     sns: {
 *         topicArn: "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel2",
 *     },
 * }});
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myNotificationChannel1 = new aws_native.devopsguru.NotificationChannel("myNotificationChannel1", {config: {
 *     sns: {
 *         topicArn: "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel",
 *     },
 * }});
 * const myNotificationChannel2 = new aws_native.devopsguru.NotificationChannel("myNotificationChannel2", {config: {
 *     sns: {
 *         topicArn: "arn:aws:sns:us-east-1:123456789012:DefaultNotificationChannel2",
 *     },
 * }});
 *
 * ```
 */
export class NotificationChannel extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NotificationChannel {
        return new NotificationChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:devopsguru:NotificationChannel';

    /**
     * Returns true if the given object is an instance of NotificationChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannel.__pulumiType;
    }

    /**
     * The ID of a notification channel.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * A `NotificationChannelConfig` object that contains information about configured notification channels.
     */
    public readonly config!: pulumi.Output<outputs.devopsguru.NotificationChannelConfig>;

    /**
     * Create a NotificationChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["config"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["config"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(NotificationChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NotificationChannel resource.
 */
export interface NotificationChannelArgs {
    /**
     * A `NotificationChannelConfig` object that contains information about configured notification channels.
     */
    config: pulumi.Input<inputs.devopsguru.NotificationChannelConfigArgs>;
}
