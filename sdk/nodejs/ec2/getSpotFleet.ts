// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::SpotFleet
 */
export function getSpotFleet(args: GetSpotFleetArgs, opts?: pulumi.InvokeOptions): Promise<GetSpotFleetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getSpotFleet", {
        "id": args.id,
    }, opts);
}

export interface GetSpotFleetArgs {
    /**
     * The ID of the Spot Fleet.
     */
    id: string;
}

export interface GetSpotFleetResult {
    /**
     * The ID of the Spot Fleet.
     */
    readonly id?: string;
    /**
     * Describes the configuration of a Spot Fleet request.
     */
    readonly spotFleetRequestConfigData?: outputs.ec2.SpotFleetRequestConfigData;
}
/**
 * Resource Type definition for AWS::EC2::SpotFleet
 */
export function getSpotFleetOutput(args: GetSpotFleetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSpotFleetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:ec2:getSpotFleet", {
        "id": args.id,
    }, opts);
}

export interface GetSpotFleetOutputArgs {
    /**
     * The ID of the Spot Fleet.
     */
    id: pulumi.Input<string>;
}
