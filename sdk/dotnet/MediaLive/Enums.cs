// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.MediaLive
{
    [EnumType]
    public readonly struct MultiplexState : IEquatable<MultiplexState>
    {
        private readonly string _value;

        private MultiplexState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MultiplexState Creating { get; } = new MultiplexState("CREATING");
        public static MultiplexState CreateFailed { get; } = new MultiplexState("CREATE_FAILED");
        public static MultiplexState Idle { get; } = new MultiplexState("IDLE");
        public static MultiplexState Starting { get; } = new MultiplexState("STARTING");
        public static MultiplexState Running { get; } = new MultiplexState("RUNNING");
        public static MultiplexState Recovering { get; } = new MultiplexState("RECOVERING");
        public static MultiplexState Stopping { get; } = new MultiplexState("STOPPING");
        public static MultiplexState Deleting { get; } = new MultiplexState("DELETING");
        public static MultiplexState Deleted { get; } = new MultiplexState("DELETED");

        public static bool operator ==(MultiplexState left, MultiplexState right) => left.Equals(right);
        public static bool operator !=(MultiplexState left, MultiplexState right) => !left.Equals(right);

        public static explicit operator string(MultiplexState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MultiplexState other && Equals(other);
        public bool Equals(MultiplexState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates which pipeline is preferred by the multiplex for program ingest.
    /// If set to \"PIPELINE_0\" or \"PIPELINE_1\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,
    /// it will switch back once that ingest is healthy again. If set to \"CURRENTLY_ACTIVE\",
    /// it will not switch back to the other pipeline based on it recovering to a healthy state,
    /// it will only switch if the active pipeline becomes unhealthy.
    /// </summary>
    [EnumType]
    public readonly struct MultiplexprogramPreferredChannelPipeline : IEquatable<MultiplexprogramPreferredChannelPipeline>
    {
        private readonly string _value;

        private MultiplexprogramPreferredChannelPipeline(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MultiplexprogramPreferredChannelPipeline CurrentlyActive { get; } = new MultiplexprogramPreferredChannelPipeline("CURRENTLY_ACTIVE");
        public static MultiplexprogramPreferredChannelPipeline Pipeline0 { get; } = new MultiplexprogramPreferredChannelPipeline("PIPELINE_0");
        public static MultiplexprogramPreferredChannelPipeline Pipeline1 { get; } = new MultiplexprogramPreferredChannelPipeline("PIPELINE_1");

        public static bool operator ==(MultiplexprogramPreferredChannelPipeline left, MultiplexprogramPreferredChannelPipeline right) => left.Equals(right);
        public static bool operator !=(MultiplexprogramPreferredChannelPipeline left, MultiplexprogramPreferredChannelPipeline right) => !left.Equals(right);

        public static explicit operator string(MultiplexprogramPreferredChannelPipeline value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MultiplexprogramPreferredChannelPipeline other && Equals(other);
        public bool Equals(MultiplexprogramPreferredChannelPipeline other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}