// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53::HealthCheck.
 */
export function getHealthCheck(args: GetHealthCheckArgs, opts?: pulumi.InvokeOptions): Promise<GetHealthCheckResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53:getHealthCheck", {
        "healthCheckId": args.healthCheckId,
    }, opts);
}

export interface GetHealthCheckArgs {
    healthCheckId: string;
}

export interface GetHealthCheckResult {
    /**
     * A complex type that contains information about the health check.
     */
    readonly healthCheckConfig?: outputs.route53.HealthCheckConfigProperties;
    readonly healthCheckId?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly healthCheckTags?: outputs.route53.HealthCheckTag[];
}

export function getHealthCheckOutput(args: GetHealthCheckOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHealthCheckResult> {
    return pulumi.output(args).apply(a => getHealthCheck(a, opts))
}

export interface GetHealthCheckOutputArgs {
    healthCheckId: pulumi.Input<string>;
}