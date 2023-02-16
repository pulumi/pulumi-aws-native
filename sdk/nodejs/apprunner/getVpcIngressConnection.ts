// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::AppRunner::VpcIngressConnection resource is an App Runner resource that specifies an App Runner VpcIngressConnection.
 */
export function getVpcIngressConnection(args: GetVpcIngressConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcIngressConnectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apprunner:getVpcIngressConnection", {
        "vpcIngressConnectionArn": args.vpcIngressConnectionArn,
    }, opts);
}

export interface GetVpcIngressConnectionArgs {
    /**
     * The Amazon Resource Name (ARN) of the VpcIngressConnection.
     */
    vpcIngressConnectionArn: string;
}

export interface GetVpcIngressConnectionResult {
    /**
     * The Domain name associated with the VPC Ingress Connection.
     */
    readonly domainName?: string;
    readonly ingressVpcConfiguration?: outputs.apprunner.VpcIngressConnectionIngressVpcConfiguration;
    /**
     * The current status of the VpcIngressConnection.
     */
    readonly status?: enums.apprunner.VpcIngressConnectionStatus;
    /**
     * The Amazon Resource Name (ARN) of the VpcIngressConnection.
     */
    readonly vpcIngressConnectionArn?: string;
}
/**
 * The AWS::AppRunner::VpcIngressConnection resource is an App Runner resource that specifies an App Runner VpcIngressConnection.
 */
export function getVpcIngressConnectionOutput(args: GetVpcIngressConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcIngressConnectionResult> {
    return pulumi.output(args).apply((a: any) => getVpcIngressConnection(a, opts))
}

export interface GetVpcIngressConnectionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the VpcIngressConnection.
     */
    vpcIngressConnectionArn: pulumi.Input<string>;
}