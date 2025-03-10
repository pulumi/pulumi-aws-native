// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create and manage wireless gateways, including LoRa gateways.
 */
export function getWirelessGateway(args: GetWirelessGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetWirelessGatewayResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotwireless:getWirelessGateway", {
        "id": args.id,
    }, opts);
}

export interface GetWirelessGatewayArgs {
    /**
     * Id for Wireless Gateway. Returned upon successful create.
     */
    id: string;
}

export interface GetWirelessGatewayResult {
    /**
     * Arn for Wireless Gateway. Returned upon successful create.
     */
    readonly arn?: string;
    /**
     * Description of Wireless Gateway.
     */
    readonly description?: string;
    /**
     * Id for Wireless Gateway. Returned upon successful create.
     */
    readonly id?: string;
    /**
     * The date and time when the most recent uplink was received.
     */
    readonly lastUplinkReceivedAt?: string;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
     */
    readonly loRaWan?: outputs.iotwireless.WirelessGatewayLoRaWanGateway;
    /**
     * Name of Wireless Gateway.
     */
    readonly name?: string;
    /**
     * A list of key-value pairs that contain metadata for the gateway.
     */
    readonly tags?: outputs.Tag[];
    /**
     * Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
     */
    readonly thingArn?: string;
    /**
     * Thing Name. If there is a Thing created, this can be returned with a Get call.
     */
    readonly thingName?: string;
}
/**
 * Create and manage wireless gateways, including LoRa gateways.
 */
export function getWirelessGatewayOutput(args: GetWirelessGatewayOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWirelessGatewayResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iotwireless:getWirelessGateway", {
        "id": args.id,
    }, opts);
}

export interface GetWirelessGatewayOutputArgs {
    /**
     * Id for Wireless Gateway. Returned upon successful create.
     */
    id: pulumi.Input<string>;
}
