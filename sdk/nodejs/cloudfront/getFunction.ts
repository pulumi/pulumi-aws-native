// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFront::Function
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudfront:getFunction", {
        "functionARN": args.functionARN,
    }, opts);
}

export interface GetFunctionArgs {
    functionARN: string;
}

export interface GetFunctionResult {
    readonly functionARN?: string;
    readonly functionCode?: string;
    readonly functionConfig?: outputs.cloudfront.FunctionConfig;
    readonly functionMetadata?: outputs.cloudfront.FunctionMetadata;
    readonly name?: string;
    readonly stage?: string;
}

export function getFunctionOutput(args: GetFunctionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionResult> {
    return pulumi.output(args).apply(a => getFunction(a, opts))
}

export interface GetFunctionOutputArgs {
    functionARN: pulumi.Input<string>;
}