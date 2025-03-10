// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::SageMaker::DeviceFleet
 */
export function getDeviceFleet(args: GetDeviceFleetArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceFleetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getDeviceFleet", {
        "deviceFleetName": args.deviceFleetName,
    }, opts);
}

export interface GetDeviceFleetArgs {
    /**
     * The name of the edge device fleet
     */
    deviceFleetName: string;
}

export interface GetDeviceFleetResult {
    /**
     * Description for the edge device fleet
     */
    readonly description?: string;
    /**
     * S3 bucket and an ecryption key id (if available) to store outputs for the fleet
     */
    readonly outputConfig?: outputs.sagemaker.DeviceFleetEdgeOutputConfig;
    /**
     * Role associated with the device fleet
     */
    readonly roleArn?: string;
    /**
     * Associate tags with the resource
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::SageMaker::DeviceFleet
 */
export function getDeviceFleetOutput(args: GetDeviceFleetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDeviceFleetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sagemaker:getDeviceFleet", {
        "deviceFleetName": args.deviceFleetName,
    }, opts);
}

export interface GetDeviceFleetOutputArgs {
    /**
     * The name of the edge device fleet
     */
    deviceFleetName: pulumi.Input<string>;
}
