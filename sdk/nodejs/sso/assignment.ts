// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for SSO assignmet
 */
export class Assignment extends pulumi.CustomResource {
    /**
     * Get an existing Assignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Assignment {
        return new Assignment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sso:Assignment';

    /**
     * Returns true if the given object is an instance of Assignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Assignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Assignment.__pulumiType;
    }

    /**
     * The sso instance that the permission set is owned.
     */
    public readonly instanceArn!: pulumi.Output<string>;
    /**
     * The permission set that the assignemt will be assigned
     */
    public readonly permissionSetArn!: pulumi.Output<string>;
    /**
     * The assignee's identifier, user id/group id
     */
    public readonly principalId!: pulumi.Output<string>;
    /**
     * The assignee's type, user/group
     */
    public readonly principalType!: pulumi.Output<enums.sso.AssignmentPrincipalType>;
    /**
     * The account id to be provisioned.
     */
    public readonly targetId!: pulumi.Output<string>;
    /**
     * The type of resource to be provsioned to, only aws account now
     */
    public readonly targetType!: pulumi.Output<enums.sso.AssignmentTargetType>;

    /**
     * Create a Assignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceArn'");
            }
            if ((!args || args.permissionSetArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissionSetArn'");
            }
            if ((!args || args.principalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalId'");
            }
            if ((!args || args.principalType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalType'");
            }
            if ((!args || args.targetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetId'");
            }
            if ((!args || args.targetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetType'");
            }
            resourceInputs["instanceArn"] = args ? args.instanceArn : undefined;
            resourceInputs["permissionSetArn"] = args ? args.permissionSetArn : undefined;
            resourceInputs["principalId"] = args ? args.principalId : undefined;
            resourceInputs["principalType"] = args ? args.principalType : undefined;
            resourceInputs["targetId"] = args ? args.targetId : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
        } else {
            resourceInputs["instanceArn"] = undefined /*out*/;
            resourceInputs["permissionSetArn"] = undefined /*out*/;
            resourceInputs["principalId"] = undefined /*out*/;
            resourceInputs["principalType"] = undefined /*out*/;
            resourceInputs["targetId"] = undefined /*out*/;
            resourceInputs["targetType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Assignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Assignment resource.
 */
export interface AssignmentArgs {
    /**
     * The sso instance that the permission set is owned.
     */
    instanceArn: pulumi.Input<string>;
    /**
     * The permission set that the assignemt will be assigned
     */
    permissionSetArn: pulumi.Input<string>;
    /**
     * The assignee's identifier, user id/group id
     */
    principalId: pulumi.Input<string>;
    /**
     * The assignee's type, user/group
     */
    principalType: pulumi.Input<enums.sso.AssignmentPrincipalType>;
    /**
     * The account id to be provisioned.
     */
    targetId: pulumi.Input<string>;
    /**
     * The type of resource to be provsioned to, only aws account now
     */
    targetType: pulumi.Input<enums.sso.AssignmentTargetType>;
}