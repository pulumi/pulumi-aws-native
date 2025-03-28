// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.
 */
export function getGameSessionQueue(args: GetGameSessionQueueArgs, opts?: pulumi.InvokeOptions): Promise<GetGameSessionQueueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:gamelift:getGameSessionQueue", {
        "name": args.name,
    }, opts);
}

export interface GetGameSessionQueueArgs {
    /**
     * A descriptive label that is associated with game session queue. Queue names must be unique within each Region.
     */
    name: string;
}

export interface GetGameSessionQueueResult {
    /**
     * The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
     */
    readonly arn?: string;
    /**
     * Information that is added to all events that are related to this game session queue.
     */
    readonly customEventData?: string;
    /**
     * A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.
     */
    readonly destinations?: outputs.gamelift.GameSessionQueueDestination[];
    /**
     * A list of locations where a queue is allowed to place new game sessions.
     */
    readonly filterConfiguration?: outputs.gamelift.GameSessionQueueFilterConfiguration;
    /**
     * An SNS topic ARN that is set up to receive game session placement notifications.
     */
    readonly notificationTarget?: string;
    /**
     * A set of policies that act as a sliding cap on player latency.
     */
    readonly playerLatencyPolicies?: outputs.gamelift.GameSessionQueuePlayerLatencyPolicy[];
    /**
     * Custom settings to use when prioritizing destinations and locations for game session placements.
     */
    readonly priorityConfiguration?: outputs.gamelift.GameSessionQueuePriorityConfiguration;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The maximum time, in seconds, that a new game session placement request remains in the queue.
     */
    readonly timeoutInSeconds?: number;
}
/**
 * The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.
 */
export function getGameSessionQueueOutput(args: GetGameSessionQueueOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGameSessionQueueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:gamelift:getGameSessionQueue", {
        "name": args.name,
    }, opts);
}

export interface GetGameSessionQueueOutputArgs {
    /**
     * A descriptive label that is associated with game session queue. Queue names must be unique within each Region.
     */
    name: pulumi.Input<string>;
}
