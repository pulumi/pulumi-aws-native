// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElastiCache::User
 */
export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:elasticache:getUser", {
        "userId": args.userId,
    }, opts);
}

export interface GetUserArgs {
    /**
     * The ID of the user.
     */
    userId: string;
}

export interface GetUserResult {
    /**
     * The Amazon Resource Name (ARN) of the user account.
     */
    readonly arn?: string;
    /**
     * The target cache engine for the user.
     */
    readonly engine?: enums.elasticache.UserEngine;
    /**
     * Indicates the user status. Can be "active", "modifying" or "deleting".
     */
    readonly status?: string;
    /**
     * An array of key-value pairs to apply to this user.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for AWS::ElastiCache::User
 */
export function getUserOutput(args: GetUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:elasticache:getUser", {
        "userId": args.userId,
    }, opts);
}

export interface GetUserOutputArgs {
    /**
     * The ID of the user.
     */
    userId: pulumi.Input<string>;
}
