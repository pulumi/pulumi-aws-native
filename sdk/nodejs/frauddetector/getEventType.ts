// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A resource schema for an EventType in Amazon Fraud Detector.
 */
export function getEventType(args: GetEventTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetEventTypeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:frauddetector:getEventType", {
        "arn": args.arn,
    }, opts);
}

export interface GetEventTypeArgs {
    /**
     * The ARN of the event type.
     */
    arn: string;
}

export interface GetEventTypeResult {
    /**
     * The ARN of the event type.
     */
    readonly arn?: string;
    /**
     * The time when the event type was created.
     */
    readonly createdTime?: string;
    /**
     * The description of the event type.
     */
    readonly description?: string;
    readonly entityTypes?: outputs.frauddetector.EventTypeEntityType[];
    readonly eventVariables?: outputs.frauddetector.EventTypeEventVariable[];
    readonly labels?: outputs.frauddetector.EventTypeLabel[];
    /**
     * The time when the event type was last updated.
     */
    readonly lastUpdatedTime?: string;
    /**
     * Tags associated with this event type.
     */
    readonly tags?: outputs.frauddetector.EventTypeTag[];
}

export function getEventTypeOutput(args: GetEventTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEventTypeResult> {
    return pulumi.output(args).apply(a => getEventType(a, opts))
}

export interface GetEventTypeOutputArgs {
    /**
     * The ARN of the event type.
     */
    arn: pulumi.Input<string>;
}