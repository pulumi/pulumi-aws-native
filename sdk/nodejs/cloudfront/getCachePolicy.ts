// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFront::CachePolicy
 */
export function getCachePolicy(args: GetCachePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetCachePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudfront:getCachePolicy", {
        "id": args.id,
    }, opts);
}

export interface GetCachePolicyArgs {
    /**
     * The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
     */
    id: string;
}

export interface GetCachePolicyResult {
    /**
     * The cache policy configuration.
     */
    readonly cachePolicyConfig?: outputs.cloudfront.CachePolicyConfig;
    /**
     * The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
     */
    readonly id?: string;
    /**
     * The date and time when the cache policy was last modified.
     */
    readonly lastModifiedTime?: string;
}
/**
 * Resource Type definition for AWS::CloudFront::CachePolicy
 */
export function getCachePolicyOutput(args: GetCachePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCachePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cloudfront:getCachePolicy", {
        "id": args.id,
    }, opts);
}

export interface GetCachePolicyOutputArgs {
    /**
     * The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
     */
    id: pulumi.Input<string>;
}
