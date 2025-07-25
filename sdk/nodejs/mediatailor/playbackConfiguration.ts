// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaTailor::PlaybackConfiguration
 */
export class PlaybackConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing PlaybackConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PlaybackConfiguration {
        return new PlaybackConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:mediatailor:PlaybackConfiguration';

    /**
     * Returns true if the given object is an instance of PlaybackConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PlaybackConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PlaybackConfiguration.__pulumiType;
    }

    /**
     * The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns, and what priority MediaTailor uses when inserting ads.
     */
    public readonly adConditioningConfiguration!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationAdConditioningConfiguration | undefined>;
    /**
     * The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
     */
    public readonly adDecisionServerUrl!: pulumi.Output<string>;
    /**
     * The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
     */
    public readonly availSuppression!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationAvailSuppression | undefined>;
    /**
     * The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
     */
    public readonly bumper!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationBumper | undefined>;
    /**
     * The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
     */
    public readonly cdnConfiguration!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationCdnConfiguration | undefined>;
    /**
     * The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
     */
    public readonly configurationAliases!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The configuration for DASH content.
     */
    public readonly dashConfiguration!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationDashConfiguration | undefined>;
    /**
     * The configuration for HLS content.
     */
    public readonly hlsConfiguration!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationHlsConfiguration | undefined>;
    /**
     * The setting that controls whether players can use stitched or guided ad insertion. The default, STITCHED_ONLY, forces all player sessions to use stitched (server-side) ad insertion. Choosing PLAYER_SELECT allows players to select either stitched or guided ad insertion at session-initialization time. The default for players that do not specify an insertion mode is stitched.
     */
    public readonly insertionMode!: pulumi.Output<enums.mediatailor.PlaybackConfigurationInsertionMode | undefined>;
    /**
     * The configuration for pre-roll ad insertion.
     */
    public readonly livePreRollConfiguration!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationLivePreRollConfiguration | undefined>;
    /**
     * The configuration that defines where AWS Elemental MediaTailor sends logs for the playback configuration.
     */
    public readonly logConfiguration!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationLogConfiguration | undefined>;
    /**
     * The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
     */
    public readonly manifestProcessingRules!: pulumi.Output<outputs.mediatailor.PlaybackConfigurationManifestProcessingRules | undefined>;
    /**
     * The identifier for the playback configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
     */
    public readonly personalizationThresholdSeconds!: pulumi.Output<number | undefined>;
    /**
     * The Amazon Resource Name (ARN) for the playback configuration.
     */
    public /*out*/ readonly playbackConfigurationArn!: pulumi.Output<string>;
    /**
     * The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
     */
    public /*out*/ readonly playbackEndpointPrefix!: pulumi.Output<string>;
    /**
     * The URL that the player uses to initialize a session that uses client-side reporting.
     */
    public /*out*/ readonly sessionInitializationEndpointPrefix!: pulumi.Output<string>;
    /**
     * The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
     */
    public readonly slateAdUrl!: pulumi.Output<string | undefined>;
    /**
     * The tags to assign to the playback configuration.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
     */
    public readonly transcodeProfileName!: pulumi.Output<string | undefined>;
    /**
     * The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
     */
    public readonly videoContentSourceUrl!: pulumi.Output<string>;

    /**
     * Create a PlaybackConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlaybackConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.adDecisionServerUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adDecisionServerUrl'");
            }
            if ((!args || args.videoContentSourceUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'videoContentSourceUrl'");
            }
            resourceInputs["adConditioningConfiguration"] = args ? args.adConditioningConfiguration : undefined;
            resourceInputs["adDecisionServerUrl"] = args ? args.adDecisionServerUrl : undefined;
            resourceInputs["availSuppression"] = args ? args.availSuppression : undefined;
            resourceInputs["bumper"] = args ? args.bumper : undefined;
            resourceInputs["cdnConfiguration"] = args ? args.cdnConfiguration : undefined;
            resourceInputs["configurationAliases"] = args ? args.configurationAliases : undefined;
            resourceInputs["dashConfiguration"] = args ? args.dashConfiguration : undefined;
            resourceInputs["hlsConfiguration"] = args ? args.hlsConfiguration : undefined;
            resourceInputs["insertionMode"] = args ? args.insertionMode : undefined;
            resourceInputs["livePreRollConfiguration"] = args ? args.livePreRollConfiguration : undefined;
            resourceInputs["logConfiguration"] = args ? args.logConfiguration : undefined;
            resourceInputs["manifestProcessingRules"] = args ? args.manifestProcessingRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["personalizationThresholdSeconds"] = args ? args.personalizationThresholdSeconds : undefined;
            resourceInputs["slateAdUrl"] = args ? args.slateAdUrl : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transcodeProfileName"] = args ? args.transcodeProfileName : undefined;
            resourceInputs["videoContentSourceUrl"] = args ? args.videoContentSourceUrl : undefined;
            resourceInputs["playbackConfigurationArn"] = undefined /*out*/;
            resourceInputs["playbackEndpointPrefix"] = undefined /*out*/;
            resourceInputs["sessionInitializationEndpointPrefix"] = undefined /*out*/;
        } else {
            resourceInputs["adConditioningConfiguration"] = undefined /*out*/;
            resourceInputs["adDecisionServerUrl"] = undefined /*out*/;
            resourceInputs["availSuppression"] = undefined /*out*/;
            resourceInputs["bumper"] = undefined /*out*/;
            resourceInputs["cdnConfiguration"] = undefined /*out*/;
            resourceInputs["configurationAliases"] = undefined /*out*/;
            resourceInputs["dashConfiguration"] = undefined /*out*/;
            resourceInputs["hlsConfiguration"] = undefined /*out*/;
            resourceInputs["insertionMode"] = undefined /*out*/;
            resourceInputs["livePreRollConfiguration"] = undefined /*out*/;
            resourceInputs["logConfiguration"] = undefined /*out*/;
            resourceInputs["manifestProcessingRules"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["personalizationThresholdSeconds"] = undefined /*out*/;
            resourceInputs["playbackConfigurationArn"] = undefined /*out*/;
            resourceInputs["playbackEndpointPrefix"] = undefined /*out*/;
            resourceInputs["sessionInitializationEndpointPrefix"] = undefined /*out*/;
            resourceInputs["slateAdUrl"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["transcodeProfileName"] = undefined /*out*/;
            resourceInputs["videoContentSourceUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PlaybackConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PlaybackConfiguration resource.
 */
export interface PlaybackConfigurationArgs {
    /**
     * The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns, and what priority MediaTailor uses when inserting ads.
     */
    adConditioningConfiguration?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationAdConditioningConfigurationArgs>;
    /**
     * The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
     */
    adDecisionServerUrl: pulumi.Input<string>;
    /**
     * The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
     */
    availSuppression?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationAvailSuppressionArgs>;
    /**
     * The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
     */
    bumper?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationBumperArgs>;
    /**
     * The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
     */
    cdnConfiguration?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationCdnConfigurationArgs>;
    /**
     * The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
     */
    configurationAliases?: pulumi.Input<{[key: string]: any}>;
    /**
     * The configuration for DASH content.
     */
    dashConfiguration?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationDashConfigurationArgs>;
    /**
     * The configuration for HLS content.
     */
    hlsConfiguration?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationHlsConfigurationArgs>;
    /**
     * The setting that controls whether players can use stitched or guided ad insertion. The default, STITCHED_ONLY, forces all player sessions to use stitched (server-side) ad insertion. Choosing PLAYER_SELECT allows players to select either stitched or guided ad insertion at session-initialization time. The default for players that do not specify an insertion mode is stitched.
     */
    insertionMode?: pulumi.Input<enums.mediatailor.PlaybackConfigurationInsertionMode>;
    /**
     * The configuration for pre-roll ad insertion.
     */
    livePreRollConfiguration?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationLivePreRollConfigurationArgs>;
    /**
     * The configuration that defines where AWS Elemental MediaTailor sends logs for the playback configuration.
     */
    logConfiguration?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationLogConfigurationArgs>;
    /**
     * The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
     */
    manifestProcessingRules?: pulumi.Input<inputs.mediatailor.PlaybackConfigurationManifestProcessingRulesArgs>;
    /**
     * The identifier for the playback configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
     */
    personalizationThresholdSeconds?: pulumi.Input<number>;
    /**
     * The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
     */
    slateAdUrl?: pulumi.Input<string>;
    /**
     * The tags to assign to the playback configuration.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
     */
    transcodeProfileName?: pulumi.Input<string>;
    /**
     * The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
     */
    videoContentSourceUrl: pulumi.Input<string>;
}
