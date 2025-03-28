// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource type definition for the AWS::EC2::SecurityGroupVpcAssociation resource
 */
export function getSecurityGroupVpcAssociation(args: GetSecurityGroupVpcAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupVpcAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getSecurityGroupVpcAssociation", {
        "groupId": args.groupId,
        "vpcId": args.vpcId,
    }, opts);
}

export interface GetSecurityGroupVpcAssociationArgs {
    /**
     * The group ID of the specified security group.
     */
    groupId: string;
    /**
     * The ID of the VPC in the security group vpc association.
     */
    vpcId: string;
}

export interface GetSecurityGroupVpcAssociationResult {
    /**
     * The state of the security group vpc association.
     */
    readonly state?: enums.ec2.SecurityGroupVpcAssociationState;
    /**
     * The reason for the state of the security group vpc association.
     */
    readonly stateReason?: string;
    /**
     * The owner of the VPC in the security group vpc association.
     */
    readonly vpcOwnerId?: string;
}
/**
 * Resource type definition for the AWS::EC2::SecurityGroupVpcAssociation resource
 */
export function getSecurityGroupVpcAssociationOutput(args: GetSecurityGroupVpcAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSecurityGroupVpcAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getSecurityGroupVpcAssociation", {
        "groupId": args.groupId,
        "vpcId": args.vpcId,
    }, opts);
}

export interface GetSecurityGroupVpcAssociationOutputArgs {
    /**
     * The group ID of the specified security group.
     */
    groupId: pulumi.Input<string>;
    /**
     * The ID of the VPC in the security group vpc association.
     */
    vpcId: pulumi.Input<string>;
}
