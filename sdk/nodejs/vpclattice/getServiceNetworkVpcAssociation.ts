// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Association between a Service Network and VPC to allow the VPC to access Services being exposed within the Service Network
 */
export function getServiceNetworkVpcAssociation(args: GetServiceNetworkVpcAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceNetworkVpcAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:vpclattice:getServiceNetworkVpcAssociation", {
        "arn": args.arn,
    }, opts);
}

export interface GetServiceNetworkVpcAssociationArgs {
    arn: string;
}

export interface GetServiceNetworkVpcAssociationResult {
    readonly arn?: string;
    readonly createdAt?: string;
    readonly id?: string;
    readonly securityGroupIds?: string[];
    readonly serviceNetworkArn?: string;
    readonly serviceNetworkId?: string;
    readonly serviceNetworkName?: string;
    readonly status?: enums.vpclattice.ServiceNetworkVpcAssociationStatus;
    readonly tags?: outputs.vpclattice.ServiceNetworkVpcAssociationTag[];
    readonly vpcId?: string;
}
/**
 * Association between a Service Network and VPC to allow the VPC to access Services being exposed within the Service Network
 */
export function getServiceNetworkVpcAssociationOutput(args: GetServiceNetworkVpcAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceNetworkVpcAssociationResult> {
    return pulumi.output(args).apply((a: any) => getServiceNetworkVpcAssociation(a, opts))
}

export interface GetServiceNetworkVpcAssociationOutputArgs {
    arn: pulumi.Input<string>;
}