// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for a Redshift-managed VPC endpoint.
 */
export function getEndpointAccess(args: GetEndpointAccessArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointAccessResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:redshift:getEndpointAccess", {
        "endpointName": args.endpointName,
    }, opts);
}

export interface GetEndpointAccessArgs {
    /**
     * The name of the endpoint.
     */
    endpointName: string;
}

export interface GetEndpointAccessResult {
    /**
     * The DNS address of the endpoint.
     */
    readonly address?: string;
    /**
     * The time (UTC) that the endpoint was created.
     */
    readonly endpointCreateTime?: string;
    /**
     * The status of the endpoint.
     */
    readonly endpointStatus?: string;
    /**
     * The port number on which the cluster accepts incoming connections.
     */
    readonly port?: number;
    /**
     * The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
     */
    readonly vpcEndpoint?: outputs.redshift.VpcEndpointProperties;
    /**
     * A list of vpc security group ids to apply to the created endpoint access.
     */
    readonly vpcSecurityGroupIds?: string[];
    /**
     * A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.
     */
    readonly vpcSecurityGroups?: outputs.redshift.EndpointAccessVpcSecurityGroup[];
}
/**
 * Resource schema for a Redshift-managed VPC endpoint.
 */
export function getEndpointAccessOutput(args: GetEndpointAccessOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEndpointAccessResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:redshift:getEndpointAccess", {
        "endpointName": args.endpointName,
    }, opts);
}

export interface GetEndpointAccessOutputArgs {
    /**
     * The name of the endpoint.
     */
    endpointName: pulumi.Input<string>;
}
