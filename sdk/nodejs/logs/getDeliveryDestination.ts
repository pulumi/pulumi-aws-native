// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Logs::DeliveryDestination
 */
export function getDeliveryDestination(args: GetDeliveryDestinationArgs, opts?: pulumi.InvokeOptions): Promise<GetDeliveryDestinationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:logs:getDeliveryDestination", {
        "name": args.name,
    }, opts);
}

export interface GetDeliveryDestinationArgs {
    /**
     * The unique name of the Delivery Destination.
     */
    name: string;
}

export interface GetDeliveryDestinationResult {
    /**
     * The value of the Arn property for this object.
     */
    readonly arn?: string;
    /**
     * The value of the DeliveryDestinationType property for this object.
     */
    readonly deliveryDestinationType?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.logs.DeliveryDestinationTag[];
}
/**
 * Resource Type definition for AWS::Logs::DeliveryDestination
 */
export function getDeliveryDestinationOutput(args: GetDeliveryDestinationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeliveryDestinationResult> {
    return pulumi.output(args).apply((a: any) => getDeliveryDestination(a, opts))
}

export interface GetDeliveryDestinationOutputArgs {
    /**
     * The unique name of the Delivery Destination.
     */
    name: pulumi.Input<string>;
}