// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IvsChat
{
    public static class GetRoom
    {
        /// <summary>
        /// Resource type definition for AWS::IVSChat::Room.
        /// </summary>
        public static Task<GetRoomResult> InvokeAsync(GetRoomArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRoomResult>("aws-native:ivschat:getRoom", args ?? new GetRoomArgs(), options.WithDefaults());

        /// <summary>
        /// Resource type definition for AWS::IVSChat::Room.
        /// </summary>
        public static Output<GetRoomResult> Invoke(GetRoomInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoomResult>("aws-native:ivschat:getRoom", args ?? new GetRoomInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource type definition for AWS::IVSChat::Room.
        /// </summary>
        public static Output<GetRoomResult> Invoke(GetRoomInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoomResult>("aws-native:ivschat:getRoom", args ?? new GetRoomInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoomArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Room ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRoomArgs()
        {
        }
        public static new GetRoomArgs Empty => new GetRoomArgs();
    }

    public sealed class GetRoomInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Room ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRoomInvokeArgs()
        {
        }
        public static new GetRoomInvokeArgs Empty => new GetRoomInvokeArgs();
    }


    [OutputType]
    public sealed class GetRoomResult
    {
        /// <summary>
        /// Room ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The system-generated ID of the room.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Array of logging configuration identifiers attached to the room.
        /// </summary>
        public readonly ImmutableArray<string> LoggingConfigurationIdentifiers;
        /// <summary>
        /// The maximum number of characters in a single message.
        /// </summary>
        public readonly int? MaximumMessageLength;
        /// <summary>
        /// The maximum number of messages per second that can be sent to the room.
        /// </summary>
        public readonly int? MaximumMessageRatePerSecond;
        /// <summary>
        /// Configuration information for optional review of messages.
        /// </summary>
        public readonly Outputs.RoomMessageReviewHandler? MessageReviewHandler;
        /// <summary>
        /// The name of the room. The value does not need to be unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetRoomResult(
            string? arn,

            string? id,

            ImmutableArray<string> loggingConfigurationIdentifiers,

            int? maximumMessageLength,

            int? maximumMessageRatePerSecond,

            Outputs.RoomMessageReviewHandler? messageReviewHandler,

            string? name,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Id = id;
            LoggingConfigurationIdentifiers = loggingConfigurationIdentifiers;
            MaximumMessageLength = maximumMessageLength;
            MaximumMessageRatePerSecond = maximumMessageRatePerSecond;
            MessageReviewHandler = messageReviewHandler;
            Name = name;
            Tags = tags;
        }
    }
}
