// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::ApiCache
 *
 * @deprecated ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ApiCache extends pulumi.CustomResource {
    /**
     * Get an existing ApiCache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiCache {
        pulumi.log.warn("ApiCache is deprecated: ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ApiCache(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appsync:ApiCache';

    /**
     * Returns true if the given object is an instance of ApiCache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiCache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiCache.__pulumiType;
    }

    public readonly apiCachingBehavior!: pulumi.Output<string>;
    public readonly apiId!: pulumi.Output<string>;
    public readonly atRestEncryptionEnabled!: pulumi.Output<boolean | undefined>;
    public readonly transitEncryptionEnabled!: pulumi.Output<boolean | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ApiCache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ApiCacheArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ApiCache is deprecated: ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiCachingBehavior === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiCachingBehavior'");
            }
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["apiCachingBehavior"] = args ? args.apiCachingBehavior : undefined;
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["atRestEncryptionEnabled"] = args ? args.atRestEncryptionEnabled : undefined;
            resourceInputs["transitEncryptionEnabled"] = args ? args.transitEncryptionEnabled : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        } else {
            resourceInputs["apiCachingBehavior"] = undefined /*out*/;
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["atRestEncryptionEnabled"] = undefined /*out*/;
            resourceInputs["transitEncryptionEnabled"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiCache.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiCache resource.
 */
export interface ApiCacheArgs {
    apiCachingBehavior: pulumi.Input<string>;
    apiId: pulumi.Input<string>;
    atRestEncryptionEnabled?: pulumi.Input<boolean>;
    transitEncryptionEnabled?: pulumi.Input<boolean>;
    ttl: pulumi.Input<number>;
    type: pulumi.Input<string>;
}