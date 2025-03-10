// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies a virtual private gateway. A virtual private gateway is the endpoint on the VPC side of your VPN connection. You can create a virtual private gateway before creating the VPC itself.
 *  For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.
 */
export function getVpnGateway(args: GetVpnGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnGatewayResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getVpnGateway", {
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

export interface GetVpnGatewayArgs {
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId: string;
}

export interface GetVpnGatewayResult {
    /**
     * Any tags assigned to the virtual private gateway.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The ID of the VPN gateway.
     */
    readonly vpnGatewayId?: string;
}
/**
 * Specifies a virtual private gateway. A virtual private gateway is the endpoint on the VPC side of your VPN connection. You can create a virtual private gateway before creating the VPC itself.
 *  For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.
 */
export function getVpnGatewayOutput(args: GetVpnGatewayOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpnGatewayResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getVpnGateway", {
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

export interface GetVpnGatewayOutputArgs {
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId: pulumi.Input<string>;
}
