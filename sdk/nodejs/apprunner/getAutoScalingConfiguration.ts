// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.
 */
export function getAutoScalingConfiguration(args: GetAutoScalingConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetAutoScalingConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apprunner:getAutoScalingConfiguration", {
        "autoScalingConfigurationArn": args.autoScalingConfigurationArn,
    }, opts);
}

export interface GetAutoScalingConfigurationArgs {
    /**
     * The Amazon Resource Name (ARN) of this auto scaling configuration.
     */
    autoScalingConfigurationArn: string;
}

export interface GetAutoScalingConfigurationResult {
    /**
     * The Amazon Resource Name (ARN) of this auto scaling configuration.
     */
    readonly autoScalingConfigurationArn?: string;
    /**
     * The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName.
     */
    readonly autoScalingConfigurationRevision?: number;
    /**
     * It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
     */
    readonly latest?: boolean;
}
/**
 * Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.
 */
export function getAutoScalingConfigurationOutput(args: GetAutoScalingConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAutoScalingConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getAutoScalingConfiguration(a, opts))
}

export interface GetAutoScalingConfigurationOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of this auto scaling configuration.
     */
    autoScalingConfigurationArn: pulumi.Input<string>;
}