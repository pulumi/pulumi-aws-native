// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGatewayV2::VpcLink
 */
export function getVpcLink(args: GetVpcLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcLinkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigatewayv2:getVpcLink", {
        "vpcLinkId": args.vpcLinkId,
    }, opts);
}

export interface GetVpcLinkArgs {
    vpcLinkId: string;
}

export interface GetVpcLinkResult {
    readonly name?: string;
    /**
     * This resource type use map for Tags, suggest to use List of Tag
     */
    readonly tags?: any;
    readonly vpcLinkId?: string;
}

export function getVpcLinkOutput(args: GetVpcLinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcLinkResult> {
    return pulumi.output(args).apply(a => getVpcLink(a, opts))
}

export interface GetVpcLinkOutputArgs {
    vpcLinkId: pulumi.Input<string>;
}