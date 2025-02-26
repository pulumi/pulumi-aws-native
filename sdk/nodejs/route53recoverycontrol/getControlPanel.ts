// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS Route53 Recovery Control Control Panel resource schema .
 */
export function getControlPanel(args: GetControlPanelArgs, opts?: pulumi.InvokeOptions): Promise<GetControlPanelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:route53recoverycontrol:getControlPanel", {
        "controlPanelArn": args.controlPanelArn,
    }, opts);
}

export interface GetControlPanelArgs {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    controlPanelArn: string;
}

export interface GetControlPanelResult {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    readonly controlPanelArn?: string;
    /**
     * A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.
     */
    readonly defaultControlPanel?: boolean;
    /**
     * The name of the control panel. You can use any non-white space character in the name.
     */
    readonly name?: string;
    /**
     * Count of associated routing controls
     */
    readonly routingControlCount?: number;
    /**
     * The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
     */
    readonly status?: enums.route53recoverycontrol.ControlPanelStatus;
}
/**
 * AWS Route53 Recovery Control Control Panel resource schema .
 */
export function getControlPanelOutput(args: GetControlPanelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetControlPanelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:route53recoverycontrol:getControlPanel", {
        "controlPanelArn": args.controlPanelArn,
    }, opts);
}

export interface GetControlPanelOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    controlPanelArn: pulumi.Input<string>;
}
