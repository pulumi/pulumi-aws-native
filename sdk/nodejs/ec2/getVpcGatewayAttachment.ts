// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::VPCGatewayAttachment
 */
export function getVpcGatewayAttachment(args: GetVpcGatewayAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcGatewayAttachmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getVpcGatewayAttachment", {
        "attachmentType": args.attachmentType,
        "vpcId": args.vpcId,
    }, opts);
}

export interface GetVpcGatewayAttachmentArgs {
    /**
     * Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment 
     */
    attachmentType: string;
    /**
     * The ID of the VPC.
     */
    vpcId: string;
}

export interface GetVpcGatewayAttachmentResult {
    /**
     * Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment 
     */
    readonly attachmentType?: string;
    /**
     * The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
     */
    readonly internetGatewayId?: string;
    /**
     * The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
     */
    readonly vpnGatewayId?: string;
}
/**
 * Resource Type definition for AWS::EC2::VPCGatewayAttachment
 */
export function getVpcGatewayAttachmentOutput(args: GetVpcGatewayAttachmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcGatewayAttachmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getVpcGatewayAttachment", {
        "attachmentType": args.attachmentType,
        "vpcId": args.vpcId,
    }, opts);
}

export interface GetVpcGatewayAttachmentOutputArgs {
    /**
     * Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment 
     */
    attachmentType: pulumi.Input<string>;
    /**
     * The ID of the VPC.
     */
    vpcId: pulumi.Input<string>;
}
