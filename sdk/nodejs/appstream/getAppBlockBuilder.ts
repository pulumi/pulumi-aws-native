// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::AppBlockBuilder.
 */
export function getAppBlockBuilder(args: GetAppBlockBuilderArgs, opts?: pulumi.InvokeOptions): Promise<GetAppBlockBuilderResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appstream:getAppBlockBuilder", {
        "name": args.name,
    }, opts);
}

export interface GetAppBlockBuilderArgs {
    name: string;
}

export interface GetAppBlockBuilderResult {
    readonly accessEndpoints?: outputs.appstream.AppBlockBuilderAccessEndpoint[];
    readonly arn?: string;
    readonly createdTime?: string;
    readonly description?: string;
    readonly displayName?: string;
    readonly enableDefaultInternetAccess?: boolean;
    readonly iamRoleArn?: string;
    readonly instanceType?: string;
    readonly platform?: string;
    readonly vpcConfig?: outputs.appstream.AppBlockBuilderVpcConfig;
}
/**
 * Resource Type definition for AWS::AppStream::AppBlockBuilder.
 */
export function getAppBlockBuilderOutput(args: GetAppBlockBuilderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppBlockBuilderResult> {
    return pulumi.output(args).apply((a: any) => getAppBlockBuilder(a, opts))
}

export interface GetAppBlockBuilderOutputArgs {
    name: pulumi.Input<string>;
}