// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElastiCache::CacheCluster
 */
export function getCacheCluster(args: GetCacheClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetCacheClusterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticache:getCacheCluster", {
        "id": args.id,
    }, opts);
}

export interface GetCacheClusterArgs {
    id: string;
}

export interface GetCacheClusterResult {
    readonly aZMode?: string;
    readonly autoMinorVersionUpgrade?: boolean;
    readonly cacheNodeType?: string;
    readonly cacheParameterGroupName?: string;
    readonly cacheSecurityGroupNames?: string[];
    readonly configurationEndpointAddress?: string;
    readonly configurationEndpointPort?: string;
    readonly engineVersion?: string;
    readonly id?: string;
    readonly logDeliveryConfigurations?: outputs.elasticache.CacheClusterLogDeliveryConfigurationRequest[];
    readonly notificationTopicArn?: string;
    readonly numCacheNodes?: number;
    readonly preferredAvailabilityZone?: string;
    readonly preferredAvailabilityZones?: string[];
    readonly preferredMaintenanceWindow?: string;
    readonly redisEndpointAddress?: string;
    readonly redisEndpointPort?: string;
    readonly snapshotRetentionLimit?: number;
    readonly snapshotWindow?: string;
    readonly tags?: outputs.elasticache.CacheClusterTag[];
    readonly transitEncryptionEnabled?: boolean;
    readonly vpcSecurityGroupIds?: string[];
}

export function getCacheClusterOutput(args: GetCacheClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCacheClusterResult> {
    return pulumi.output(args).apply(a => getCacheCluster(a, opts))
}

export interface GetCacheClusterOutputArgs {
    id: pulumi.Input<string>;
}