// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::NimbleStudio::StreamingImage
 */
export class StreamingImage extends pulumi.CustomResource {
    /**
     * Get an existing StreamingImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StreamingImage {
        return new StreamingImage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:nimblestudio:StreamingImage';

    /**
     * Returns true if the given object is an instance of StreamingImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamingImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamingImage.__pulumiType;
    }

    /**
     * A human-readable description of the streaming image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of an EC2 machine image with which to create the streaming image.
     */
    public readonly ec2ImageId!: pulumi.Output<string>;
    public /*out*/ readonly encryptionConfiguration!: pulumi.Output<outputs.nimblestudio.StreamingImageEncryptionConfiguration>;
    public readonly encryptionConfigurationKeyArn!: pulumi.Output<string | undefined>;
    public readonly encryptionConfigurationKeyType!: pulumi.Output<string | undefined>;
    /**
     * The list of IDs of EULAs that must be accepted before a streaming session can be started using this streaming image.
     */
    public /*out*/ readonly eulaIds!: pulumi.Output<string[]>;
    /**
     * A friendly name for a streaming image resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The owner of the streaming image, either the studioId that contains the streaming image or 'amazon' for images that are provided by  .
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * The platform of the streaming image, either WINDOWS or LINUX.
     */
    public /*out*/ readonly platform!: pulumi.Output<string>;
    /**
     * The unique identifier for the streaming image resource.
     */
    public /*out*/ readonly streamingImageId!: pulumi.Output<string>;
    /**
     * The unique identifier for a studio resource. In Nimble Studio, all other resources are contained in a studio resource.
     */
    public readonly studioId!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a StreamingImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamingImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.ec2ImageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ec2ImageId'");
            }
            if ((!args || args.studioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'studioId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ec2ImageId"] = args ? args.ec2ImageId : undefined;
            resourceInputs["encryptionConfigurationKeyArn"] = args ? args.encryptionConfigurationKeyArn : undefined;
            resourceInputs["encryptionConfigurationKeyType"] = args ? args.encryptionConfigurationKeyType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["studioId"] = args ? args.studioId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["encryptionConfiguration"] = undefined /*out*/;
            resourceInputs["eulaIds"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["platform"] = undefined /*out*/;
            resourceInputs["streamingImageId"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["ec2ImageId"] = undefined /*out*/;
            resourceInputs["encryptionConfiguration"] = undefined /*out*/;
            resourceInputs["encryptionConfigurationKeyArn"] = undefined /*out*/;
            resourceInputs["encryptionConfigurationKeyType"] = undefined /*out*/;
            resourceInputs["eulaIds"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["platform"] = undefined /*out*/;
            resourceInputs["streamingImageId"] = undefined /*out*/;
            resourceInputs["studioId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["ec2ImageId", "studioId", "tags.*"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(StreamingImage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StreamingImage resource.
 */
export interface StreamingImageArgs {
    /**
     * A human-readable description of the streaming image.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of an EC2 machine image with which to create the streaming image.
     */
    ec2ImageId: pulumi.Input<string>;
    encryptionConfigurationKeyArn?: pulumi.Input<string>;
    encryptionConfigurationKeyType?: pulumi.Input<string>;
    /**
     * A friendly name for a streaming image resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique identifier for a studio resource. In Nimble Studio, all other resources are contained in a studio resource.
     */
    studioId: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
