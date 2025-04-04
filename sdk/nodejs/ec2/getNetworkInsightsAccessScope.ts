// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::EC2::NetworkInsightsAccessScope
 */
export function getNetworkInsightsAccessScope(args: GetNetworkInsightsAccessScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInsightsAccessScopeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getNetworkInsightsAccessScope", {
        "networkInsightsAccessScopeId": args.networkInsightsAccessScopeId,
    }, opts);
}

export interface GetNetworkInsightsAccessScopeArgs {
    /**
     * The ID of the Network Access Scope.
     */
    networkInsightsAccessScopeId: string;
}

export interface GetNetworkInsightsAccessScopeResult {
    /**
     * The creation date.
     */
    readonly createdDate?: string;
    /**
     * The ARN of the Network Access Scope.
     */
    readonly networkInsightsAccessScopeArn?: string;
    /**
     * The ID of the Network Access Scope.
     */
    readonly networkInsightsAccessScopeId?: string;
    /**
     * The tags.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The last updated date.
     */
    readonly updatedDate?: string;
}
/**
 * Resource schema for AWS::EC2::NetworkInsightsAccessScope
 */
export function getNetworkInsightsAccessScopeOutput(args: GetNetworkInsightsAccessScopeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNetworkInsightsAccessScopeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getNetworkInsightsAccessScope", {
        "networkInsightsAccessScopeId": args.networkInsightsAccessScopeId,
    }, opts);
}

export interface GetNetworkInsightsAccessScopeOutputArgs {
    /**
     * The ID of the Network Access Scope.
     */
    networkInsightsAccessScopeId: pulumi.Input<string>;
}
