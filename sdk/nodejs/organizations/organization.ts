// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Organizations::Organization
 */
export class Organization extends pulumi.CustomResource {
    /**
     * Get an existing Organization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Organization {
        return new Organization(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:organizations:Organization';

    /**
     * Returns true if the given object is an instance of Organization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Organization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Organization.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of an organization.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The unique identifier (ID) of an organization.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
     */
    public readonly featureSet!: pulumi.Output<enums.organizations.OrganizationFeatureSet | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization.
     */
    public /*out*/ readonly managementAccountArn!: pulumi.Output<string>;
    /**
     * The email address that is associated with the AWS account that is designated as the management account for the organization.
     */
    public /*out*/ readonly managementAccountEmail!: pulumi.Output<string>;
    /**
     * The unique identifier (ID) of the management account of an organization.
     */
    public /*out*/ readonly managementAccountId!: pulumi.Output<string>;
    /**
     * The unique identifier (ID) for the root.
     */
    public /*out*/ readonly rootId!: pulumi.Output<string>;

    /**
     * Create a Organization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrganizationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["managementAccountArn"] = undefined /*out*/;
            resourceInputs["managementAccountEmail"] = undefined /*out*/;
            resourceInputs["managementAccountId"] = undefined /*out*/;
            resourceInputs["rootId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["featureSet"] = undefined /*out*/;
            resourceInputs["managementAccountArn"] = undefined /*out*/;
            resourceInputs["managementAccountEmail"] = undefined /*out*/;
            resourceInputs["managementAccountId"] = undefined /*out*/;
            resourceInputs["rootId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Organization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Organization resource.
 */
export interface OrganizationArgs {
    /**
     * Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
     */
    featureSet?: pulumi.Input<enums.organizations.OrganizationFeatureSet>;
}
