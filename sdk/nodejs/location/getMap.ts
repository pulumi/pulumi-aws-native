// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Location::Map Resource Type
 */
export function getMap(args: GetMapArgs, opts?: pulumi.InvokeOptions): Promise<GetMapResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:location:getMap", {
        "mapName": args.mapName,
    }, opts);
}

export interface GetMapArgs {
    /**
     * The name for the map resource.
     *
     * Requirements:
     *
     * - Must contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
     * - Must be a unique map resource name.
     * - No spaces allowed. For example, `ExampleMap` .
     */
    mapName: string;
}

export interface GetMapResult {
    /**
     * The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across all AWS .
     *
     * - Format example: `arn:aws:geo:region:account-id:maps/ExampleMap`
     */
    readonly arn?: string;
    /**
     * The timestamp for when the map resource was created in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
     */
    readonly createTime?: string;
    /**
     * An optional description for the map resource.
     */
    readonly description?: string;
    /**
     * Synonym for `Arn` . The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across all AWS .
     *
     * - Format example: `arn:aws:geo:region:account-id:maps/ExampleMap`
     */
    readonly mapArn?: string;
    /**
     * No longer used. If included, the only allowed value is `RequestBasedUsage` .
     *
     * *Allowed Values* : `RequestBasedUsage`
     */
    readonly pricingPlan?: enums.location.MapPricingPlan;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The timestamp for when the map resource was last updated in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
     */
    readonly updateTime?: string;
}
/**
 * Definition of AWS::Location::Map Resource Type
 */
export function getMapOutput(args: GetMapOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMapResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:location:getMap", {
        "mapName": args.mapName,
    }, opts);
}

export interface GetMapOutputArgs {
    /**
     * The name for the map resource.
     *
     * Requirements:
     *
     * - Must contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
     * - Must be a unique map resource name.
     * - No spaces allowed. For example, `ExampleMap` .
     */
    mapName: pulumi.Input<string>;
}
