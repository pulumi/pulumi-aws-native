// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IVS::PlaybackRestrictionPolicy.
 */
export class PlaybackRestrictionPolicy extends pulumi.CustomResource {
    /**
     * Get an existing PlaybackRestrictionPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PlaybackRestrictionPolicy {
        return new PlaybackRestrictionPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ivs:PlaybackRestrictionPolicy';

    /**
     * Returns true if the given object is an instance of PlaybackRestrictionPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PlaybackRestrictionPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PlaybackRestrictionPolicy.__pulumiType;
    }

    /**
     * A list of country codes that control geoblocking restriction. Allowed values are the officially assigned ISO 3166-1 alpha-2 codes. Default: All countries (an empty array).
     */
    public readonly allowedCountries!: pulumi.Output<string[]>;
    /**
     * A list of origin sites that control CORS restriction. Allowed values are the same as valid values of the Origin header defined at https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
     */
    public readonly allowedOrigins!: pulumi.Output<string[]>;
    /**
     * Playback-restriction-policy identifier.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether channel playback is constrained by origin site.
     */
    public readonly enableStrictOriginEnforcement!: pulumi.Output<boolean | undefined>;
    /**
     * Playback-restriction-policy name.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a PlaybackRestrictionPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlaybackRestrictionPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.allowedCountries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedCountries'");
            }
            if ((!args || args.allowedOrigins === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedOrigins'");
            }
            resourceInputs["allowedCountries"] = args ? args.allowedCountries : undefined;
            resourceInputs["allowedOrigins"] = args ? args.allowedOrigins : undefined;
            resourceInputs["enableStrictOriginEnforcement"] = args ? args.enableStrictOriginEnforcement : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["allowedCountries"] = undefined /*out*/;
            resourceInputs["allowedOrigins"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["enableStrictOriginEnforcement"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PlaybackRestrictionPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PlaybackRestrictionPolicy resource.
 */
export interface PlaybackRestrictionPolicyArgs {
    /**
     * A list of country codes that control geoblocking restriction. Allowed values are the officially assigned ISO 3166-1 alpha-2 codes. Default: All countries (an empty array).
     */
    allowedCountries: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of origin sites that control CORS restriction. Allowed values are the same as valid values of the Origin header defined at https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
     */
    allowedOrigins: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether channel playback is constrained by origin site.
     */
    enableStrictOriginEnforcement?: pulumi.Input<boolean>;
    /**
     * Playback-restriction-policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}