// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.
 */
export function getLocalGatewayRouteTable(args: GetLocalGatewayRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayRouteTableResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getLocalGatewayRouteTable", {
        "localGatewayRouteTableId": args.localGatewayRouteTableId,
    }, opts);
}

export interface GetLocalGatewayRouteTableArgs {
    /**
     * The ID of the local gateway route table.
     */
    localGatewayRouteTableId: string;
}

export interface GetLocalGatewayRouteTableResult {
    /**
     * The ARN of the local gateway route table.
     */
    readonly localGatewayRouteTableArn?: string;
    /**
     * The ID of the local gateway route table.
     */
    readonly localGatewayRouteTableId?: string;
    /**
     * The ARN of the outpost.
     */
    readonly outpostArn?: string;
    /**
     * The owner of the local gateway route table.
     */
    readonly ownerId?: string;
    /**
     * The state of the local gateway route table.
     */
    readonly state?: string;
    /**
     * The tags for the local gateway route table.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.
 */
export function getLocalGatewayRouteTableOutput(args: GetLocalGatewayRouteTableOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLocalGatewayRouteTableResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getLocalGatewayRouteTable", {
        "localGatewayRouteTableId": args.localGatewayRouteTableId,
    }, opts);
}

export interface GetLocalGatewayRouteTableOutputArgs {
    /**
     * The ID of the local gateway route table.
     */
    localGatewayRouteTableId: pulumi.Input<string>;
}
