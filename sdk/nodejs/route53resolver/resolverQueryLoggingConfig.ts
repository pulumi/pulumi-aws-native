// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.
 */
export class ResolverQueryLoggingConfig extends pulumi.CustomResource {
    /**
     * Get an existing ResolverQueryLoggingConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResolverQueryLoggingConfig {
        return new ResolverQueryLoggingConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53resolver:ResolverQueryLoggingConfig';

    /**
     * Returns true if the given object is an instance of ResolverQueryLoggingConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverQueryLoggingConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverQueryLoggingConfig.__pulumiType;
    }

    /**
     * Arn
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Count
     */
    public /*out*/ readonly associationCount!: pulumi.Output<number>;
    /**
     * Rfc3339TimeString
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The id of the creator request.
     */
    public /*out*/ readonly creatorRequestId!: pulumi.Output<string>;
    /**
     * destination arn
     */
    public readonly destinationArn!: pulumi.Output<string | undefined>;
    /**
     * ResolverQueryLogConfigName
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * AccountId
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
     */
    public /*out*/ readonly shareStatus!: pulumi.Output<enums.route53resolver.ResolverQueryLoggingConfigShareStatus>;
    /**
     * ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.route53resolver.ResolverQueryLoggingConfigStatus>;

    /**
     * Create a ResolverQueryLoggingConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResolverQueryLoggingConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["destinationArn"] = args ? args.destinationArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["associationCount"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["creatorRequestId"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["shareStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["associationCount"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["creatorRequestId"] = undefined /*out*/;
            resourceInputs["destinationArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["shareStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResolverQueryLoggingConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResolverQueryLoggingConfig resource.
 */
export interface ResolverQueryLoggingConfigArgs {
    /**
     * destination arn
     */
    destinationArn?: pulumi.Input<string>;
    /**
     * ResolverQueryLogConfigName
     */
    name?: pulumi.Input<string>;
}