// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The connection group for your distribution tenants. When you first create a distribution tenant and you don't specify a connection group, CloudFront will automatically create a default connection group for you. When you create a new distribution tenant and don't specify a connection group, the default one will be associated with your distribution tenant.
 */
export function getConnectionGroup(args: GetConnectionGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudfront:getConnectionGroup", {
        "id": args.id,
    }, opts);
}

export interface GetConnectionGroupArgs {
    /**
     * The ID of the connection group.
     */
    id: string;
}

export interface GetConnectionGroupResult {
    /**
     * The ID of the Anycast static IP list.
     */
    readonly anycastIpListId?: string;
    /**
     * The Amazon Resource Name (ARN) of the connection group.
     */
    readonly arn?: string;
    /**
     * The date and time when the connection group was created.
     */
    readonly createdTime?: string;
    /**
     * The current version of the connection group.
     */
    readonly eTag?: string;
    /**
     * Whether the connection group is enabled.
     */
    readonly enabled?: boolean;
    /**
     * The ID of the connection group.
     */
    readonly id?: string;
    /**
     * IPv6 is enabled for the connection group.
     */
    readonly ipv6Enabled?: boolean;
    /**
     * Whether the connection group is the default connection group for the distribution tenants.
     */
    readonly isDefault?: boolean;
    /**
     * The date and time when the connection group was updated.
     */
    readonly lastModifiedTime?: string;
    /**
     * The routing endpoint (also known as the DNS name) that is assigned to the connection group, such as d111111abcdef8.cloudfront.net.
     */
    readonly routingEndpoint?: string;
    /**
     * The status of the connection group.
     */
    readonly status?: string;
    /**
     * A complex type that contains zero or more ``Tag`` elements.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * The connection group for your distribution tenants. When you first create a distribution tenant and you don't specify a connection group, CloudFront will automatically create a default connection group for you. When you create a new distribution tenant and don't specify a connection group, the default one will be associated with your distribution tenant.
 */
export function getConnectionGroupOutput(args: GetConnectionGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConnectionGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cloudfront:getConnectionGroup", {
        "id": args.id,
    }, opts);
}

export interface GetConnectionGroupOutputArgs {
    /**
     * The ID of the connection group.
     */
    id: pulumi.Input<string>;
}
