// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::Dashboard
 */
export function getDashboard(args: GetDashboardArgs, opts?: pulumi.InvokeOptions): Promise<GetDashboardResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotsitewise:getDashboard", {
        "dashboardId": args.dashboardId,
    }, opts);
}

export interface GetDashboardArgs {
    /**
     * The ID of the dashboard.
     */
    dashboardId: string;
}

export interface GetDashboardResult {
    /**
     * The ARN of the dashboard.
     */
    readonly dashboardArn?: string;
    /**
     * The dashboard definition specified in a JSON literal.
     */
    readonly dashboardDefinition?: string;
    /**
     * A description for the dashboard.
     */
    readonly dashboardDescription?: string;
    /**
     * The ID of the dashboard.
     */
    readonly dashboardId?: string;
    /**
     * A friendly name for the dashboard.
     */
    readonly dashboardName?: string;
    /**
     * A list of key-value pairs that contain metadata for the dashboard.
     */
    readonly tags?: outputs.iotsitewise.DashboardTag[];
}

export function getDashboardOutput(args: GetDashboardOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDashboardResult> {
    return pulumi.output(args).apply(a => getDashboard(a, opts))
}

export interface GetDashboardOutputArgs {
    /**
     * The ID of the dashboard.
     */
    dashboardId: pulumi.Input<string>;
}