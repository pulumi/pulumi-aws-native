// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT::ThingType
 */
export function getThingType(args: GetThingTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetThingTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getThingType", {
        "thingTypeName": args.thingTypeName,
    }, opts);
}

export interface GetThingTypeArgs {
    /**
     * The name of the thing type.
     */
    thingTypeName: string;
}

export interface GetThingTypeResult {
    /**
     * The thing type arn.
     */
    readonly arn?: string;
    /**
     * Deprecates a thing type. You can not associate new things with deprecated thing type.
     *
     * Requires permission to access the [DeprecateThingType](https://docs.aws.amazon.com//service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions) action.
     */
    readonly deprecateThingType?: boolean;
    /**
     * The thing type id.
     */
    readonly id?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The thing type properties for the thing type to create. It contains information about the new thing type including a description, a list of searchable thing attribute names, and a list of propagating attributes. After a thing type is created, you can only update `Mqtt5Configuration` .
     */
    readonly thingTypeProperties?: outputs.iot.ThingTypePropertiesProperties;
}
/**
 * Resource Type definition for AWS::IoT::ThingType
 */
export function getThingTypeOutput(args: GetThingTypeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetThingTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iot:getThingType", {
        "thingTypeName": args.thingTypeName,
    }, opts);
}

export interface GetThingTypeOutputArgs {
    /**
     * The name of the thing type.
     */
    thingTypeName: pulumi.Input<string>;
}
