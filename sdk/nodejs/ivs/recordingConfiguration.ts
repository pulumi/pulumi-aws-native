// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IVS::RecordingConfiguration
 */
export class RecordingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing RecordingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RecordingConfiguration {
        return new RecordingConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ivs:RecordingConfiguration';

    /**
     * Returns true if the given object is an instance of RecordingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordingConfiguration.__pulumiType;
    }

    /**
     * Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly destinationConfiguration!: pulumi.Output<outputs.ivs.RecordingConfigurationDestinationConfiguration>;
    /**
     * Recording Configuration Name.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Recording Reconnect Window Seconds. (0 means disabled)
     */
    public readonly recordingReconnectWindowSeconds!: pulumi.Output<number | undefined>;
    /**
     * Recording Configuration State.
     */
    public /*out*/ readonly state!: pulumi.Output<enums.ivs.RecordingConfigurationState>;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    public readonly tags!: pulumi.Output<outputs.ivs.RecordingConfigurationTag[] | undefined>;
    public readonly thumbnailConfiguration!: pulumi.Output<outputs.ivs.RecordingConfigurationThumbnailConfiguration | undefined>;

    /**
     * Create a RecordingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordingConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationConfiguration'");
            }
            resourceInputs["destinationConfiguration"] = args ? args.destinationConfiguration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recordingReconnectWindowSeconds"] = args ? args.recordingReconnectWindowSeconds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thumbnailConfiguration"] = args ? args.thumbnailConfiguration : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["destinationConfiguration"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["recordingReconnectWindowSeconds"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["thumbnailConfiguration"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RecordingConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RecordingConfiguration resource.
 */
export interface RecordingConfigurationArgs {
    destinationConfiguration: pulumi.Input<inputs.ivs.RecordingConfigurationDestinationConfigurationArgs>;
    /**
     * Recording Configuration Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Recording Reconnect Window Seconds. (0 means disabled)
     */
    recordingReconnectWindowSeconds?: pulumi.Input<number>;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ivs.RecordingConfigurationTagArgs>[]>;
    thumbnailConfiguration?: pulumi.Input<inputs.ivs.RecordingConfigurationThumbnailConfigurationArgs>;
}