// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application
 */
export function getMonitor(args: GetMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:internetmonitor:getMonitor", {
        "monitorName": args.monitorName,
    }, opts);
}

export interface GetMonitorArgs {
    monitorName: string;
}

export interface GetMonitorResult {
    readonly createdAt?: string;
    readonly modifiedAt?: string;
    readonly monitorArn?: string;
    readonly processingStatus?: enums.internetmonitor.MonitorProcessingStatusCode;
    readonly processingStatusInfo?: string;
    readonly resources?: string[];
    readonly status?: enums.internetmonitor.MonitorConfigState;
    readonly tags?: outputs.internetmonitor.MonitorTag[];
}
/**
 * Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application
 */
export function getMonitorOutput(args: GetMonitorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitorResult> {
    return pulumi.output(args).apply((a: any) => getMonitor(a, opts))
}

export interface GetMonitorOutputArgs {
    monitorName: pulumi.Input<string>;
}