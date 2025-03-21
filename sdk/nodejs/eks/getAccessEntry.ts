// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An object representing an Amazon EKS AccessEntry.
 */
export function getAccessEntry(args: GetAccessEntryArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessEntryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:eks:getAccessEntry", {
        "clusterName": args.clusterName,
        "principalArn": args.principalArn,
    }, opts);
}

export interface GetAccessEntryArgs {
    /**
     * The cluster that the access entry is created for.
     */
    clusterName: string;
    /**
     * The principal ARN that the access entry is created for.
     */
    principalArn: string;
}

export interface GetAccessEntryResult {
    /**
     * The ARN of the access entry.
     */
    readonly accessEntryArn?: string;
    /**
     * An array of access policies that are associated with the access entry.
     */
    readonly accessPolicies?: outputs.eks.AccessEntryAccessPolicy[];
    /**
     * The Kubernetes groups that the access entry is associated with.
     */
    readonly kubernetesGroups?: string[];
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The Kubernetes user that the access entry is associated with.
     */
    readonly username?: string;
}
/**
 * An object representing an Amazon EKS AccessEntry.
 */
export function getAccessEntryOutput(args: GetAccessEntryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccessEntryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:eks:getAccessEntry", {
        "clusterName": args.clusterName,
        "principalArn": args.principalArn,
    }, opts);
}

export interface GetAccessEntryOutputArgs {
    /**
     * The cluster that the access entry is created for.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The principal ARN that the access entry is created for.
     */
    principalArn: pulumi.Input<string>;
}
