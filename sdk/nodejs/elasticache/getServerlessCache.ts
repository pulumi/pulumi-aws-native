// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::ElastiCache::ServerlessCache resource creates an Amazon ElastiCache Serverless Cache.
 */
export function getServerlessCache(args: GetServerlessCacheArgs, opts?: pulumi.InvokeOptions): Promise<GetServerlessCacheResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:elasticache:getServerlessCache", {
        "serverlessCacheName": args.serverlessCacheName,
    }, opts);
}

export interface GetServerlessCacheArgs {
    /**
     * The name of the Serverless Cache. This value must be unique.
     */
    serverlessCacheName: string;
}

export interface GetServerlessCacheResult {
    /**
     * The ARN of the Serverless Cache.
     */
    readonly arn?: string;
    /**
     * The cache usage limit for the serverless cache.
     */
    readonly cacheUsageLimits?: outputs.elasticache.ServerlessCacheCacheUsageLimits;
    /**
     * The creation time of the Serverless Cache.
     */
    readonly createTime?: string;
    /**
     * The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.
     */
    readonly dailySnapshotTime?: string;
    /**
     * The description of the Serverless Cache.
     */
    readonly description?: string;
    /**
     * Represents the information required for client programs to connect to a cache node. This value is read-only.
     */
    readonly endpoint?: outputs.elasticache.ServerlessCacheEndpoint;
    /**
     * The engine name of the Serverless Cache.
     */
    readonly engine?: string;
    /**
     * The full engine version of the Serverless Cache.
     */
    readonly fullEngineVersion?: string;
    /**
     * The major engine version of the Serverless Cache.
     */
    readonly majorEngineVersion?: string;
    /**
     * Represents the information required for client programs to connect to a cache node. This value is read-only.
     */
    readonly readerEndpoint?: outputs.elasticache.ServerlessCacheEndpoint;
    /**
     * One or more Amazon VPC security groups associated with this Serverless Cache.
     */
    readonly securityGroupIds?: string[];
    /**
     * The snapshot retention limit of the Serverless Cache.
     */
    readonly snapshotRetentionLimit?: number;
    /**
     * The status of the Serverless Cache.
     */
    readonly status?: string;
    /**
     * An array of key-value pairs to apply to this Serverless Cache.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The ID of the user group.
     */
    readonly userGroupId?: string;
}
/**
 * The AWS::ElastiCache::ServerlessCache resource creates an Amazon ElastiCache Serverless Cache.
 */
export function getServerlessCacheOutput(args: GetServerlessCacheOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerlessCacheResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:elasticache:getServerlessCache", {
        "serverlessCacheName": args.serverlessCacheName,
    }, opts);
}

export interface GetServerlessCacheOutputArgs {
    /**
     * The name of the Serverless Cache. This value must be unique.
     */
    serverlessCacheName: pulumi.Input<string>;
}
