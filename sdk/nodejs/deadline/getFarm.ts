// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Deadline::Farm Resource Type
 */
export function getFarm(args: GetFarmArgs, opts?: pulumi.InvokeOptions): Promise<GetFarmResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:deadline:getFarm", {
        "arn": args.arn,
    }, opts);
}

export interface GetFarmArgs {
    arn: string;
}

export interface GetFarmResult {
    readonly arn?: string;
    readonly description?: string;
    readonly displayName?: string;
    readonly farmId?: string;
}
/**
 * Definition of AWS::Deadline::Farm Resource Type
 */
export function getFarmOutput(args: GetFarmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFarmResult> {
    return pulumi.output(args).apply((a: any) => getFarm(a, opts))
}

export interface GetFarmOutputArgs {
    arn: pulumi.Input<string>;
}