// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A signing profile is a signing template that can be used to carry out a pre-defined signing job.
 */
export class SigningProfile extends pulumi.CustomResource {
    /**
     * Get an existing SigningProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SigningProfile {
        return new SigningProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:signer:SigningProfile';

    /**
     * Returns true if the given object is an instance of SigningProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SigningProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SigningProfile.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the specified signing profile.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the target signing platform.
     */
    public readonly platformId!: pulumi.Output<enums.signer.SigningProfilePlatformId>;
    /**
     * A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. 
     */
    public /*out*/ readonly profileName!: pulumi.Output<string>;
    /**
     * A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.
     */
    public /*out*/ readonly profileVersion!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the specified signing profile version.
     */
    public /*out*/ readonly profileVersionArn!: pulumi.Output<string>;
    /**
     * Signature validity period of the profile.
     */
    public readonly signatureValidityPeriod!: pulumi.Output<outputs.signer.SigningProfileSignatureValidityPeriod | undefined>;
    /**
     * A list of tags associated with the signing profile.
     */
    public readonly tags!: pulumi.Output<outputs.signer.SigningProfileTag[] | undefined>;

    /**
     * Create a SigningProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SigningProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.platformId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platformId'");
            }
            resourceInputs["platformId"] = args ? args.platformId : undefined;
            resourceInputs["signatureValidityPeriod"] = args ? args.signatureValidityPeriod : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["profileName"] = undefined /*out*/;
            resourceInputs["profileVersion"] = undefined /*out*/;
            resourceInputs["profileVersionArn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["platformId"] = undefined /*out*/;
            resourceInputs["profileName"] = undefined /*out*/;
            resourceInputs["profileVersion"] = undefined /*out*/;
            resourceInputs["profileVersionArn"] = undefined /*out*/;
            resourceInputs["signatureValidityPeriod"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SigningProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SigningProfile resource.
 */
export interface SigningProfileArgs {
    /**
     * The ID of the target signing platform.
     */
    platformId: pulumi.Input<enums.signer.SigningProfilePlatformId>;
    /**
     * Signature validity period of the profile.
     */
    signatureValidityPeriod?: pulumi.Input<inputs.signer.SigningProfileSignatureValidityPeriodArgs>;
    /**
     * A list of tags associated with the signing profile.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.signer.SigningProfileTagArgs>[]>;
}