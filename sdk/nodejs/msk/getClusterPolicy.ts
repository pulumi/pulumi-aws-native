// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::MSK::ClusterPolicy
 */
export function getClusterPolicy(args: GetClusterPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:msk:getClusterPolicy", {
        "clusterArn": args.clusterArn,
    }, opts);
}

export interface GetClusterPolicyArgs {
    /**
     * The arn of the cluster for the resource policy.
     */
    clusterArn: string;
}

export interface GetClusterPolicyResult {
    /**
     * The current version of the policy attached to the specified cluster
     */
    readonly currentVersion?: string;
    /**
     * A policy document containing permissions to add to the specified cluster.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MSK::ClusterPolicy` for more information about the expected schema for this property.
     */
    readonly policy?: any;
}
/**
 * Resource Type definition for AWS::MSK::ClusterPolicy
 */
export function getClusterPolicyOutput(args: GetClusterPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClusterPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:msk:getClusterPolicy", {
        "clusterArn": args.clusterArn,
    }, opts);
}

export interface GetClusterPolicyOutputArgs {
    /**
     * The arn of the cluster for the resource policy.
     */
    clusterArn: pulumi.Input<string>;
}
