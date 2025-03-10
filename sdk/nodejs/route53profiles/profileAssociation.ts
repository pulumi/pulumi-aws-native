// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Route53Profiles::ProfileAssociation
 */
export class ProfileAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ProfileAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProfileAssociation {
        return new ProfileAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53profiles:ProfileAssociation';

    /**
     * Returns true if the given object is an instance of ProfileAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileAssociation.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the profile association.
     */
    public readonly arn!: pulumi.Output<string | undefined>;
    /**
     * Primary Identifier for  Profile Association
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The name of an association between a  Profile and a VPC.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the  profile that you associated with the resource that is specified by ResourceId.
     */
    public readonly profileId!: pulumi.Output<string>;
    /**
     * The resource that you associated the  profile with.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a ProfileAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.profileId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profileId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["profileId"] = args ? args.profileId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["profileId"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "profileId", "resourceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ProfileAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProfileAssociation resource.
 */
export interface ProfileAssociationArgs {
    /**
     * The Amazon Resource Name (ARN) of the profile association.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of an association between a  Profile and a VPC.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the  profile that you associated with the resource that is specified by ResourceId.
     */
    profileId: pulumi.Input<string>;
    /**
     * The resource that you associated the  profile with.
     */
    resourceId: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
