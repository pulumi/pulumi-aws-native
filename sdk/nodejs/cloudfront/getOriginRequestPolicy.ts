// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An origin request policy.
 *  When it's attached to a cache behavior, the origin request policy determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
 *   +  The request body and the URL path (without the domain name) from the viewer request.
 *   +  The headers that CloudFront automatically includes in every origin request, including ``Host``, ``User-Agent``, and ``X-Amz-Cf-Id``.
 *   +  All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
 *
 *  CloudFront sends a request when it can't find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use ``CachePolicy``.
 */
export function getOriginRequestPolicy(args: GetOriginRequestPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetOriginRequestPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudfront:getOriginRequestPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetOriginRequestPolicyArgs {
    /**
     * The unique identifier for the origin request policy. For example: `befd7079-9bbc-4ebf-8ade-498a3694176c` .
     */
    id: string;
}

export interface GetOriginRequestPolicyResult {
    /**
     * The unique identifier for the origin request policy. For example: `befd7079-9bbc-4ebf-8ade-498a3694176c` .
     */
    readonly id?: string;
    /**
     * The date and time when the origin request policy was last modified.
     */
    readonly lastModifiedTime?: string;
    /**
     * The origin request policy configuration.
     */
    readonly originRequestPolicyConfig?: outputs.cloudfront.OriginRequestPolicyConfig;
}
/**
 * An origin request policy.
 *  When it's attached to a cache behavior, the origin request policy determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
 *   +  The request body and the URL path (without the domain name) from the viewer request.
 *   +  The headers that CloudFront automatically includes in every origin request, including ``Host``, ``User-Agent``, and ``X-Amz-Cf-Id``.
 *   +  All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
 *
 *  CloudFront sends a request when it can't find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use ``CachePolicy``.
 */
export function getOriginRequestPolicyOutput(args: GetOriginRequestPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOriginRequestPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cloudfront:getOriginRequestPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetOriginRequestPolicyOutputArgs {
    /**
     * The unique identifier for the origin request policy. For example: `befd7079-9bbc-4ebf-8ade-498a3694176c` .
     */
    id: pulumi.Input<string>;
}
