// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for Identity Center (SSO) Instance
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sso:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The ID of the identity store associated with the created Identity Center (SSO) Instance
     */
    public /*out*/ readonly identityStoreId!: pulumi.Output<string>;
    /**
     * The SSO Instance ARN that is returned upon creation of the Identity Center (SSO) Instance
     */
    public /*out*/ readonly instanceArn!: pulumi.Output<string>;
    /**
     * The name you want to assign to this Identity Center (SSO) Instance
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The AWS accountId of the owner of the Identity Center (SSO) Instance
     */
    public /*out*/ readonly ownerAccountId!: pulumi.Output<string>;
    /**
     * The status of the Identity Center (SSO) Instance, create_in_progress/delete_in_progress/active
     */
    public /*out*/ readonly status!: pulumi.Output<enums.sso.InstanceStatus>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["identityStoreId"] = undefined /*out*/;
            resourceInputs["instanceArn"] = undefined /*out*/;
            resourceInputs["ownerAccountId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["identityStoreId"] = undefined /*out*/;
            resourceInputs["instanceArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["ownerAccountId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The name you want to assign to this Identity Center (SSO) Instance
     */
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}