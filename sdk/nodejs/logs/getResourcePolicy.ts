// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The resource schema for AWSLogs ResourcePolicy
 */
export function getResourcePolicy(args: GetResourcePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:logs:getResourcePolicy", {
        "policyName": args.policyName,
    }, opts);
}

export interface GetResourcePolicyArgs {
    /**
     * A name for resource policy
     */
    policyName: string;
}

export interface GetResourcePolicyResult {
    /**
     * The policy document
     */
    readonly policyDocument?: string;
}
/**
 * The resource schema for AWSLogs ResourcePolicy
 */
export function getResourcePolicyOutput(args: GetResourcePolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourcePolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:logs:getResourcePolicy", {
        "policyName": args.policyName,
    }, opts);
}

export interface GetResourcePolicyOutputArgs {
    /**
     * A name for resource policy
     */
    policyName: pulumi.Input<string>;
}
