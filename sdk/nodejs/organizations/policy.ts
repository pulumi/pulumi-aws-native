// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Policies in AWS Organizations enable you to manage different features of the AWS accounts in your organization.  You can use policies when all features are enabled in your organization.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:organizations:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * ARN of the Policy
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.
     */
    public /*out*/ readonly awsManaged!: pulumi.Output<boolean>;
    /**
     * The Policy text content
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Human readable description of the policy
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the Policy
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
     */
    public readonly tags!: pulumi.Output<outputs.organizations.PolicyTag[] | undefined>;
    /**
     * List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to
     */
    public readonly targetIds!: pulumi.Output<string[] | undefined>;
    /**
     * The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY
     */
    public readonly type!: pulumi.Output<enums.organizations.PolicyType>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetIds"] = args ? args.targetIds : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsManaged"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsManaged"] = undefined /*out*/;
            resourceInputs["content"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targetIds"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The Policy text content
     */
    content: pulumi.Input<string>;
    /**
     * Human readable description of the policy
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Policy
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.organizations.PolicyTagArgs>[]>;
    /**
     * List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to
     */
    targetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY
     */
    type: pulumi.Input<enums.organizations.PolicyType>;
}