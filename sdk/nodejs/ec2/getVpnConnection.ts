// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.
 *  To specify a VPN connection between a transit gateway and customer gateway, use the ``TransitGatewayId`` and ``CustomerGatewayId`` properties.
 *  To specify a VPN connection between a virtual private gateway and customer gateway, use the ``VpnGatewayId`` and ``CustomerGatewayId`` properties.
 *  For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.
 */
export function getVpnConnection(args: GetVpnConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getVpnConnection", {
        "vpnConnectionId": args.vpnConnectionId,
    }, opts);
}

export interface GetVpnConnectionArgs {
    /**
     * The ID of the VPN connection.
     */
    vpnConnectionId: string;
}

export interface GetVpnConnectionResult {
    /**
     * Any tags assigned to the VPN connection.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The ID of the VPN connection.
     */
    readonly vpnConnectionId?: string;
}
/**
 * Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.
 *  To specify a VPN connection between a transit gateway and customer gateway, use the ``TransitGatewayId`` and ``CustomerGatewayId`` properties.
 *  To specify a VPN connection between a virtual private gateway and customer gateway, use the ``VpnGatewayId`` and ``CustomerGatewayId`` properties.
 *  For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.
 */
export function getVpnConnectionOutput(args: GetVpnConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpnConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getVpnConnection", {
        "vpnConnectionId": args.vpnConnectionId,
    }, opts);
}

export interface GetVpnConnectionOutputArgs {
    /**
     * The ID of the VPN connection.
     */
    vpnConnectionId: pulumi.Input<string>;
}
