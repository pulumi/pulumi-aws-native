// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AutoScaling::LifecycleHook
 */
export function getLifecycleHook(args: GetLifecycleHookArgs, opts?: pulumi.InvokeOptions): Promise<GetLifecycleHookResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:autoscaling:getLifecycleHook", {
        "autoScalingGroupName": args.autoScalingGroupName,
        "lifecycleHookName": args.lifecycleHookName,
    }, opts);
}

export interface GetLifecycleHookArgs {
    /**
     * The name of the Auto Scaling group for the lifecycle hook.
     */
    autoScalingGroupName: string;
    /**
     * The name of the lifecycle hook.
     */
    lifecycleHookName: string;
}

export interface GetLifecycleHookResult {
    /**
     * The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).
     */
    readonly defaultResult?: string;
    /**
     * The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.
     */
    readonly heartbeatTimeout?: number;
    /**
     * The instance state to which you want to attach the lifecycle hook.
     */
    readonly lifecycleTransition?: string;
    /**
     * Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.
     */
    readonly notificationMetadata?: string;
    /**
     * The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.
     */
    readonly notificationTargetARN?: string;
    /**
     * The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.
     */
    readonly roleARN?: string;
}

export function getLifecycleHookOutput(args: GetLifecycleHookOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLifecycleHookResult> {
    return pulumi.output(args).apply(a => getLifecycleHook(a, opts))
}

export interface GetLifecycleHookOutputArgs {
    /**
     * The name of the Auto Scaling group for the lifecycle hook.
     */
    autoScalingGroupName: pulumi.Input<string>;
    /**
     * The name of the lifecycle hook.
     */
    lifecycleHookName: pulumi.Input<string>;
}