// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type Definition for AWS::S3Outposts::Endpoint
 */
export function getEndpoint(args: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:s3outposts:getEndpoint", {
        "arn": args.arn,
    }, opts);
}

export interface GetEndpointArgs {
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    arn: string;
}

export interface GetEndpointResult {
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    readonly arn?: string;
    /**
     * The VPC CIDR committed by this endpoint.
     */
    readonly cidrBlock?: string;
    /**
     * The time the endpoint was created.
     */
    readonly creationTime?: string;
    /**
     * The failure reason, if any, for a create or delete endpoint operation.
     */
    readonly failedReason?: outputs.s3outposts.EndpointFailedReason;
    /**
     * The ID of the endpoint.
     */
    readonly id?: string;
    /**
     * The network interfaces of the endpoint.
     */
    readonly networkInterfaces?: outputs.s3outposts.EndpointNetworkInterface[];
    /**
     * The status of the endpoint.
     */
    readonly status?: enums.s3outposts.EndpointStatus;
}
/**
 * Resource Type Definition for AWS::S3Outposts::Endpoint
 */
export function getEndpointOutput(args: GetEndpointOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEndpointResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:s3outposts:getEndpoint", {
        "arn": args.arn,
    }, opts);
}

export interface GetEndpointOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    arn: pulumi.Input<string>;
}
