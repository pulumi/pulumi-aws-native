// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.S3Outposts
{
    /// <summary>
    /// If `Enabled` , the rule is currently being applied. If `Disabled` , the rule is not currently being applied.
    /// </summary>
    [EnumType]
    public readonly struct BucketRuleStatus : IEquatable<BucketRuleStatus>
    {
        private readonly string _value;

        private BucketRuleStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BucketRuleStatus Enabled { get; } = new BucketRuleStatus("Enabled");
        public static BucketRuleStatus Disabled { get; } = new BucketRuleStatus("Disabled");

        public static bool operator ==(BucketRuleStatus left, BucketRuleStatus right) => left.Equals(right);
        public static bool operator !=(BucketRuleStatus left, BucketRuleStatus right) => !left.Equals(right);

        public static explicit operator string(BucketRuleStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BucketRuleStatus other && Equals(other);
        public bool Equals(BucketRuleStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
    /// </summary>
    [EnumType]
    public readonly struct EndpointAccessType : IEquatable<EndpointAccessType>
    {
        private readonly string _value;

        private EndpointAccessType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EndpointAccessType CustomerOwnedIp { get; } = new EndpointAccessType("CustomerOwnedIp");
        public static EndpointAccessType Private { get; } = new EndpointAccessType("Private");

        public static bool operator ==(EndpointAccessType left, EndpointAccessType right) => left.Equals(right);
        public static bool operator !=(EndpointAccessType left, EndpointAccessType right) => !left.Equals(right);

        public static explicit operator string(EndpointAccessType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EndpointAccessType other && Equals(other);
        public bool Equals(EndpointAccessType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the endpoint.
    /// </summary>
    [EnumType]
    public readonly struct EndpointStatus : IEquatable<EndpointStatus>
    {
        private readonly string _value;

        private EndpointStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EndpointStatus Available { get; } = new EndpointStatus("Available");
        public static EndpointStatus Pending { get; } = new EndpointStatus("Pending");
        public static EndpointStatus Deleting { get; } = new EndpointStatus("Deleting");
        public static EndpointStatus CreateFailed { get; } = new EndpointStatus("Create_Failed");
        public static EndpointStatus DeleteFailed { get; } = new EndpointStatus("Delete_Failed");

        public static bool operator ==(EndpointStatus left, EndpointStatus right) => left.Equals(right);
        public static bool operator !=(EndpointStatus left, EndpointStatus right) => !left.Equals(right);

        public static explicit operator string(EndpointStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EndpointStatus other && Equals(other);
        public bool Equals(EndpointStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
