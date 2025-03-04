// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::MediaLive::Cluster Resource Type
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:medialive:getCluster", {
        "id": args.id,
    }, opts);
}

export interface GetClusterArgs {
    /**
     * The unique ID of the Cluster.
     */
    id: string;
}

export interface GetClusterResult {
    /**
     * The ARN of the Cluster.
     */
    readonly arn?: string;
    /**
     * The MediaLive Channels that are currently running on Nodes in this Cluster.
     */
    readonly channelIds?: string[];
    /**
     * The unique ID of the Cluster.
     */
    readonly id?: string;
    /**
     * The user-specified name of the Cluster to be created.
     */
    readonly name?: string;
    readonly networkSettings?: outputs.medialive.ClusterNetworkSettings;
    readonly state?: enums.medialive.ClusterState;
    /**
     * A collection of key-value pairs.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Definition of AWS::MediaLive::Cluster Resource Type
 */
export function getClusterOutput(args: GetClusterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClusterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:medialive:getCluster", {
        "id": args.id,
    }, opts);
}

export interface GetClusterOutputArgs {
    /**
     * The unique ID of the Cluster.
     */
    id: pulumi.Input<string>;
}
