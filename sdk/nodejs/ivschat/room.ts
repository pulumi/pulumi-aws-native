// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::IVSChat::Room.
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const room = new aws_native.ivschat.Room("room", {
 *     name: "MyRoom",
 *     tags: [{
 *         key: "MyKey",
 *         value: "MyValue",
 *     }],
 * });
 * export const roomArn = room.id;
 * export const roomId = room.id;
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const room = new aws_native.ivschat.Room("room", {
 *     name: "MyRoom",
 *     tags: [{
 *         key: "MyKey",
 *         value: "MyValue",
 *     }],
 * });
 * export const roomArn = room.id;
 * export const roomId = room.id;
 *
 * ```
 */
export class Room extends pulumi.CustomResource {
    /**
     * Get an existing Room resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Room {
        return new Room(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ivschat:Room';

    /**
     * Returns true if the given object is an instance of Room.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Room {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Room.__pulumiType;
    }

    /**
     * Room ARN is automatically generated on creation and assigned as the unique identifier.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The system-generated ID of the room.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Array of logging configuration identifiers attached to the room.
     */
    public readonly loggingConfigurationIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * The maximum number of characters in a single message.
     */
    public readonly maximumMessageLength!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of messages per second that can be sent to the room.
     */
    public readonly maximumMessageRatePerSecond!: pulumi.Output<number | undefined>;
    /**
     * Configuration information for optional review of messages.
     */
    public readonly messageReviewHandler!: pulumi.Output<outputs.ivschat.RoomMessageReviewHandler | undefined>;
    /**
     * The name of the room. The value does not need to be unique.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Room resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RoomArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["loggingConfigurationIdentifiers"] = args ? args.loggingConfigurationIdentifiers : undefined;
            resourceInputs["maximumMessageLength"] = args ? args.maximumMessageLength : undefined;
            resourceInputs["maximumMessageRatePerSecond"] = args ? args.maximumMessageRatePerSecond : undefined;
            resourceInputs["messageReviewHandler"] = args ? args.messageReviewHandler : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["loggingConfigurationIdentifiers"] = undefined /*out*/;
            resourceInputs["maximumMessageLength"] = undefined /*out*/;
            resourceInputs["maximumMessageRatePerSecond"] = undefined /*out*/;
            resourceInputs["messageReviewHandler"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Room.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Room resource.
 */
export interface RoomArgs {
    /**
     * Array of logging configuration identifiers attached to the room.
     */
    loggingConfigurationIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of characters in a single message.
     */
    maximumMessageLength?: pulumi.Input<number>;
    /**
     * The maximum number of messages per second that can be sent to the room.
     */
    maximumMessageRatePerSecond?: pulumi.Input<number>;
    /**
     * Configuration information for optional review of messages.
     */
    messageReviewHandler?: pulumi.Input<inputs.ivschat.RoomMessageReviewHandlerArgs>;
    /**
     * The name of the room. The value does not need to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
