// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    public static class GetGameSessionQueue
    {
        /// <summary>
        /// The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.
        /// </summary>
        public static Task<GetGameSessionQueueResult> InvokeAsync(GetGameSessionQueueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGameSessionQueueResult>("aws-native:gamelift:getGameSessionQueue", args ?? new GetGameSessionQueueArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.
        /// </summary>
        public static Output<GetGameSessionQueueResult> Invoke(GetGameSessionQueueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGameSessionQueueResult>("aws-native:gamelift:getGameSessionQueue", args ?? new GetGameSessionQueueInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::GameLift::GameSessionQueue resource creates an Amazon GameLift (GameLift) game session queue.
        /// </summary>
        public static Output<GetGameSessionQueueResult> Invoke(GetGameSessionQueueInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGameSessionQueueResult>("aws-native:gamelift:getGameSessionQueue", args ?? new GetGameSessionQueueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGameSessionQueueArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A descriptive label that is associated with game session queue. Queue names must be unique within each Region.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetGameSessionQueueArgs()
        {
        }
        public static new GetGameSessionQueueArgs Empty => new GetGameSessionQueueArgs();
    }

    public sealed class GetGameSessionQueueInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A descriptive label that is associated with game session queue. Queue names must be unique within each Region.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetGameSessionQueueInvokeArgs()
        {
        }
        public static new GetGameSessionQueueInvokeArgs Empty => new GetGameSessionQueueInvokeArgs();
    }


    [OutputType]
    public sealed class GetGameSessionQueueResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Information that is added to all events that are related to this game session queue.
        /// </summary>
        public readonly string? CustomEventData;
        /// <summary>
        /// A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.
        /// </summary>
        public readonly ImmutableArray<Outputs.GameSessionQueueDestination> Destinations;
        /// <summary>
        /// A list of locations where a queue is allowed to place new game sessions.
        /// </summary>
        public readonly Outputs.GameSessionQueueFilterConfiguration? FilterConfiguration;
        /// <summary>
        /// An SNS topic ARN that is set up to receive game session placement notifications.
        /// </summary>
        public readonly string? NotificationTarget;
        /// <summary>
        /// A set of policies that act as a sliding cap on player latency.
        /// </summary>
        public readonly ImmutableArray<Outputs.GameSessionQueuePlayerLatencyPolicy> PlayerLatencyPolicies;
        /// <summary>
        /// Custom settings to use when prioritizing destinations and locations for game session placements.
        /// </summary>
        public readonly Outputs.GameSessionQueuePriorityConfiguration? PriorityConfiguration;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The maximum time, in seconds, that a new game session placement request remains in the queue.
        /// </summary>
        public readonly int? TimeoutInSeconds;

        [OutputConstructor]
        private GetGameSessionQueueResult(
            string? arn,

            string? customEventData,

            ImmutableArray<Outputs.GameSessionQueueDestination> destinations,

            Outputs.GameSessionQueueFilterConfiguration? filterConfiguration,

            string? notificationTarget,

            ImmutableArray<Outputs.GameSessionQueuePlayerLatencyPolicy> playerLatencyPolicies,

            Outputs.GameSessionQueuePriorityConfiguration? priorityConfiguration,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            int? timeoutInSeconds)
        {
            Arn = arn;
            CustomEventData = customEventData;
            Destinations = destinations;
            FilterConfiguration = filterConfiguration;
            NotificationTarget = notificationTarget;
            PlayerLatencyPolicies = playerLatencyPolicies;
            PriorityConfiguration = priorityConfiguration;
            Tags = tags;
            TimeoutInSeconds = timeoutInSeconds;
        }
    }
}
