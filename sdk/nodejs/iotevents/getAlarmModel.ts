// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.
 */
export function getAlarmModel(args: GetAlarmModelArgs, opts?: pulumi.InvokeOptions): Promise<GetAlarmModelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotevents:getAlarmModel", {
        "alarmModelName": args.alarmModelName,
    }, opts);
}

export interface GetAlarmModelArgs {
    /**
     * The name of the alarm model.
     */
    alarmModelName: string;
}

export interface GetAlarmModelResult {
    /**
     * Contains the configuration information of alarm state changes.
     */
    readonly alarmCapabilities?: outputs.iotevents.AlarmModelAlarmCapabilities;
    /**
     * Contains information about one or more alarm actions.
     */
    readonly alarmEventActions?: outputs.iotevents.AlarmModelAlarmEventActions;
    /**
     * The description of the alarm model.
     */
    readonly alarmModelDescription?: string;
    /**
     * Defines when your alarm is invoked.
     */
    readonly alarmRule?: outputs.iotevents.AlarmModelAlarmRule;
    /**
     * The ARN of the IAM role that allows the alarm to perform actions and access AWS resources. For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.
     */
    readonly roleArn?: string;
    /**
     * A non-negative integer that reflects the severity level of the alarm.
     */
    readonly severity?: number;
    /**
     * A list of key-value pairs that contain metadata for the alarm model. The tags help you manage the alarm model. For more information, see [Tagging your resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *Developer Guide*.
     *  You can create up to 50 tags for one alarm model.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.
 */
export function getAlarmModelOutput(args: GetAlarmModelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAlarmModelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iotevents:getAlarmModel", {
        "alarmModelName": args.alarmModelName,
    }, opts);
}

export interface GetAlarmModelOutputArgs {
    /**
     * The name of the alarm model.
     */
    alarmModelName: pulumi.Input<string>;
}
