// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GuardDuty::Detector
 */
export class Detector extends pulumi.CustomResource {
    /**
     * Get an existing Detector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Detector {
        return new Detector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:guardduty:Detector';

    /**
     * Returns true if the given object is an instance of Detector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Detector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Detector.__pulumiType;
    }

    /**
     * The unique ID of the detector.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Describes which data sources will be enabled for the detector.
     */
    public readonly dataSources!: pulumi.Output<outputs.guardduty.DetectorCfnDataSourceConfigurations | undefined>;
    /**
     * Specifies whether the detector is to be enabled on creation.
     */
    public readonly enable!: pulumi.Output<boolean>;
    /**
     * A list of features that will be configured for the detector.
     */
    public readonly features!: pulumi.Output<outputs.guardduty.DetectorCfnFeatureConfiguration[] | undefined>;
    /**
     * Specifies how frequently updated findings are exported.
     */
    public readonly findingPublishingFrequency!: pulumi.Output<string | undefined>;
    /**
     * Specifies tags added to a new detector resource. Each tag consists of a key and an optional value, both of which you define.
     *
     * Currently, support is available only for creating and deleting a tag. No support exists for updating the tags.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Detector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DetectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.enable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enable'");
            }
            resourceInputs["dataSources"] = args ? args.dataSources : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["findingPublishingFrequency"] = args ? args.findingPublishingFrequency : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["dataSources"] = undefined /*out*/;
            resourceInputs["enable"] = undefined /*out*/;
            resourceInputs["features"] = undefined /*out*/;
            resourceInputs["findingPublishingFrequency"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Detector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Detector resource.
 */
export interface DetectorArgs {
    /**
     * Describes which data sources will be enabled for the detector.
     */
    dataSources?: pulumi.Input<inputs.guardduty.DetectorCfnDataSourceConfigurationsArgs>;
    /**
     * Specifies whether the detector is to be enabled on creation.
     */
    enable: pulumi.Input<boolean>;
    /**
     * A list of features that will be configured for the detector.
     */
    features?: pulumi.Input<pulumi.Input<inputs.guardduty.DetectorCfnFeatureConfigurationArgs>[]>;
    /**
     * Specifies how frequently updated findings are exported.
     */
    findingPublishingFrequency?: pulumi.Input<string>;
    /**
     * Specifies tags added to a new detector resource. Each tag consists of a key and an optional value, both of which you define.
     *
     * Currently, support is available only for creating and deleting a tag. No support exists for updating the tags.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
