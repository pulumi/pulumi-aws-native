// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS Route53 Recovery Control Cluster resource schema
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53recoverycontrol:getCluster", {
        "clusterArn": args.clusterArn,
    }, opts);
}

export interface GetClusterArgs {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    clusterArn: string;
}

export interface GetClusterResult {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    readonly clusterArn?: string;
    /**
     * Endpoints for the cluster.
     */
    readonly clusterEndpoints?: outputs.route53recoverycontrol.ClusterEndpoint[];
    /**
     * Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
     */
    readonly status?: enums.route53recoverycontrol.ClusterStatus;
}

export function getClusterOutput(args: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply(a => getCluster(a, opts))
}

export interface GetClusterOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    clusterArn: pulumi.Input<string>;
}