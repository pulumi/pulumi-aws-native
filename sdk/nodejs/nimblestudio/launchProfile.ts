// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::NimbleStudio::LaunchProfile
 */
export class LaunchProfile extends pulumi.CustomResource {
    /**
     * Get an existing LaunchProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LaunchProfile {
        return new LaunchProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:nimblestudio:LaunchProfile';

    /**
     * Returns true if the given object is an instance of LaunchProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchProfile.__pulumiType;
    }

    /**
     * A human-readable description of the launch profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Unique identifiers for a collection of EC2 subnets.
     */
    public readonly ec2SubnetIds!: pulumi.Output<string[]>;
    /**
     * The unique identifier for the launch profile resource.
     */
    public /*out*/ readonly launchProfileId!: pulumi.Output<string>;
    /**
     * The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".
     */
    public readonly launchProfileProtocolVersions!: pulumi.Output<string[]>;
    /**
     * A friendly name for the launch profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A configuration for a streaming session.
     */
    public readonly streamConfiguration!: pulumi.Output<outputs.nimblestudio.LaunchProfileStreamConfiguration>;
    /**
     * Unique identifiers for a collection of studio components that can be used with this launch profile.
     */
    public readonly studioComponentIds!: pulumi.Output<string[]>;
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
     * Create a LaunchProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LaunchProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.ec2SubnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ec2SubnetIds'");
            }
            if ((!args || args.launchProfileProtocolVersions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'launchProfileProtocolVersions'");
            }
            if ((!args || args.streamConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamConfiguration'");
            }
            if ((!args || args.studioComponentIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'studioComponentIds'");
            }
            if ((!args || args.studioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'studioId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ec2SubnetIds"] = args ? args.ec2SubnetIds : undefined;
            resourceInputs["launchProfileProtocolVersions"] = args ? args.launchProfileProtocolVersions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["streamConfiguration"] = args ? args.streamConfiguration : undefined;
            resourceInputs["studioComponentIds"] = args ? args.studioComponentIds : undefined;
            resourceInputs["studioId"] = args ? args.studioId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["launchProfileId"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["ec2SubnetIds"] = undefined /*out*/;
            resourceInputs["launchProfileId"] = undefined /*out*/;
            resourceInputs["launchProfileProtocolVersions"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["streamConfiguration"] = undefined /*out*/;
            resourceInputs["studioComponentIds"] = undefined /*out*/;
            resourceInputs["studioId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["ec2SubnetIds[*]", "studioId", "tags.*"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(LaunchProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LaunchProfile resource.
 */
export interface LaunchProfileArgs {
    /**
     * A human-readable description of the launch profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique identifiers for a collection of EC2 subnets.
     */
    ec2SubnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".
     */
    launchProfileProtocolVersions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A friendly name for the launch profile.
     */
    name?: pulumi.Input<string>;
    /**
     * A configuration for a streaming session.
     */
    streamConfiguration: pulumi.Input<inputs.nimblestudio.LaunchProfileStreamConfigurationArgs>;
    /**
     * Unique identifiers for a collection of studio components that can be used with this launch profile.
     */
    studioComponentIds: pulumi.Input<pulumi.Input<string>[]>;
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
