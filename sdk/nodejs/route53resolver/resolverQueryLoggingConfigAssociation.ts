// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.
 */
export class ResolverQueryLoggingConfigAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ResolverQueryLoggingConfigAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResolverQueryLoggingConfigAssociation {
        return new ResolverQueryLoggingConfigAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53resolver:ResolverQueryLoggingConfigAssociation';

    /**
     * Returns true if the given object is an instance of ResolverQueryLoggingConfigAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverQueryLoggingConfigAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverQueryLoggingConfigAssociation.__pulumiType;
    }

    /**
     * Id
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Rfc3339TimeString
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * ResolverQueryLogConfigAssociationError
     */
    public /*out*/ readonly error!: pulumi.Output<enums.route53resolver.ResolverQueryLoggingConfigAssociationError>;
    /**
     * ResolverQueryLogConfigAssociationErrorMessage
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * ResolverQueryLogConfigId
     */
    public readonly resolverQueryLogConfigId!: pulumi.Output<string | undefined>;
    /**
     * ResourceId
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * ResolverQueryLogConfigAssociationStatus
     */
    public /*out*/ readonly status!: pulumi.Output<enums.route53resolver.ResolverQueryLoggingConfigAssociationStatus>;

    /**
     * Create a ResolverQueryLoggingConfigAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResolverQueryLoggingConfigAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["resolverQueryLogConfigId"] = args ? args.resolverQueryLogConfigId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["errorMessage"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["errorMessage"] = undefined /*out*/;
            resourceInputs["resolverQueryLogConfigId"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["resolverQueryLogConfigId", "resourceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ResolverQueryLoggingConfigAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResolverQueryLoggingConfigAssociation resource.
 */
export interface ResolverQueryLoggingConfigAssociationArgs {
    /**
     * ResolverQueryLogConfigId
     */
    resolverQueryLogConfigId?: pulumi.Input<string>;
    /**
     * ResourceId
     */
    resourceId?: pulumi.Input<string>;
}
