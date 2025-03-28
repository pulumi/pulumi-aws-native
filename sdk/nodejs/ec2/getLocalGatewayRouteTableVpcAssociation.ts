// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.
 */
export function getLocalGatewayRouteTableVpcAssociation(args: GetLocalGatewayRouteTableVpcAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayRouteTableVpcAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getLocalGatewayRouteTableVpcAssociation", {
        "localGatewayRouteTableVpcAssociationId": args.localGatewayRouteTableVpcAssociationId,
    }, opts);
}

export interface GetLocalGatewayRouteTableVpcAssociationArgs {
    /**
     * The ID of the association.
     */
    localGatewayRouteTableVpcAssociationId: string;
}

export interface GetLocalGatewayRouteTableVpcAssociationResult {
    /**
     * The ID of the local gateway.
     */
    readonly localGatewayId?: string;
    /**
     * The ID of the association.
     */
    readonly localGatewayRouteTableVpcAssociationId?: string;
    /**
     * The state of the association.
     */
    readonly state?: string;
    /**
     * The tags for the association.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.
 */
export function getLocalGatewayRouteTableVpcAssociationOutput(args: GetLocalGatewayRouteTableVpcAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLocalGatewayRouteTableVpcAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getLocalGatewayRouteTableVpcAssociation", {
        "localGatewayRouteTableVpcAssociationId": args.localGatewayRouteTableVpcAssociationId,
    }, opts);
}

export interface GetLocalGatewayRouteTableVpcAssociationOutputArgs {
    /**
     * The ID of the association.
     */
    localGatewayRouteTableVpcAssociationId: pulumi.Input<string>;
}
