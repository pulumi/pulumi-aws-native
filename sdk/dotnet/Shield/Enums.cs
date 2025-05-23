// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Shield
{
    /// <summary>
    /// If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
    /// If `DISABLED`, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.
    /// </summary>
    [EnumType]
    public readonly struct ProactiveEngagementStatus : IEquatable<ProactiveEngagementStatus>
    {
        private readonly string _value;

        private ProactiveEngagementStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProactiveEngagementStatus Enabled { get; } = new ProactiveEngagementStatus("ENABLED");
        public static ProactiveEngagementStatus Disabled { get; } = new ProactiveEngagementStatus("DISABLED");

        public static bool operator ==(ProactiveEngagementStatus left, ProactiveEngagementStatus right) => left.Equals(right);
        public static bool operator !=(ProactiveEngagementStatus left, ProactiveEngagementStatus right) => !left.Equals(right);

        public static explicit operator string(ProactiveEngagementStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProactiveEngagementStatus other && Equals(other);
        public bool Equals(ProactiveEngagementStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
    /// </summary>
    [EnumType]
    public readonly struct ProtectionApplicationLayerAutomaticResponseConfigurationStatus : IEquatable<ProtectionApplicationLayerAutomaticResponseConfigurationStatus>
    {
        private readonly string _value;

        private ProtectionApplicationLayerAutomaticResponseConfigurationStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtectionApplicationLayerAutomaticResponseConfigurationStatus Enabled { get; } = new ProtectionApplicationLayerAutomaticResponseConfigurationStatus("ENABLED");
        public static ProtectionApplicationLayerAutomaticResponseConfigurationStatus Disabled { get; } = new ProtectionApplicationLayerAutomaticResponseConfigurationStatus("DISABLED");

        public static bool operator ==(ProtectionApplicationLayerAutomaticResponseConfigurationStatus left, ProtectionApplicationLayerAutomaticResponseConfigurationStatus right) => left.Equals(right);
        public static bool operator !=(ProtectionApplicationLayerAutomaticResponseConfigurationStatus left, ProtectionApplicationLayerAutomaticResponseConfigurationStatus right) => !left.Equals(right);

        public static explicit operator string(ProtectionApplicationLayerAutomaticResponseConfigurationStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtectionApplicationLayerAutomaticResponseConfigurationStatus other && Equals(other);
        public bool Equals(ProtectionApplicationLayerAutomaticResponseConfigurationStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.
    /// * Sum - Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.
    /// * Mean - Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.
    /// * Max - Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront and origin resources for CloudFront distributions.
    /// </summary>
    [EnumType]
    public readonly struct ProtectionGroupAggregation : IEquatable<ProtectionGroupAggregation>
    {
        private readonly string _value;

        private ProtectionGroupAggregation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtectionGroupAggregation Sum { get; } = new ProtectionGroupAggregation("SUM");
        public static ProtectionGroupAggregation Mean { get; } = new ProtectionGroupAggregation("MEAN");
        public static ProtectionGroupAggregation Max { get; } = new ProtectionGroupAggregation("MAX");

        public static bool operator ==(ProtectionGroupAggregation left, ProtectionGroupAggregation right) => left.Equals(right);
        public static bool operator !=(ProtectionGroupAggregation left, ProtectionGroupAggregation right) => !left.Equals(right);

        public static explicit operator string(ProtectionGroupAggregation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtectionGroupAggregation other && Equals(other);
        public bool Equals(ProtectionGroupAggregation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The criteria to use to choose the protected resources for inclusion in the group. You can include all resources that have protections, provide a list of resource Amazon Resource Names (ARNs), or include all resources of a specified resource type.
    /// </summary>
    [EnumType]
    public readonly struct ProtectionGroupPattern : IEquatable<ProtectionGroupPattern>
    {
        private readonly string _value;

        private ProtectionGroupPattern(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtectionGroupPattern All { get; } = new ProtectionGroupPattern("ALL");
        public static ProtectionGroupPattern Arbitrary { get; } = new ProtectionGroupPattern("ARBITRARY");
        public static ProtectionGroupPattern ByResourceType { get; } = new ProtectionGroupPattern("BY_RESOURCE_TYPE");

        public static bool operator ==(ProtectionGroupPattern left, ProtectionGroupPattern right) => left.Equals(right);
        public static bool operator !=(ProtectionGroupPattern left, ProtectionGroupPattern right) => !left.Equals(right);

        public static explicit operator string(ProtectionGroupPattern value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtectionGroupPattern other && Equals(other);
        public bool Equals(ProtectionGroupPattern other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The resource type to include in the protection group. All protected resources of this type are included in the protection group. Newly protected resources of this type are automatically added to the group. You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting.
    /// </summary>
    [EnumType]
    public readonly struct ProtectionGroupResourceType : IEquatable<ProtectionGroupResourceType>
    {
        private readonly string _value;

        private ProtectionGroupResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtectionGroupResourceType CloudfrontDistribution { get; } = new ProtectionGroupResourceType("CLOUDFRONT_DISTRIBUTION");
        public static ProtectionGroupResourceType Route53HostedZone { get; } = new ProtectionGroupResourceType("ROUTE_53_HOSTED_ZONE");
        public static ProtectionGroupResourceType ElasticIpAllocation { get; } = new ProtectionGroupResourceType("ELASTIC_IP_ALLOCATION");
        public static ProtectionGroupResourceType ClassicLoadBalancer { get; } = new ProtectionGroupResourceType("CLASSIC_LOAD_BALANCER");
        public static ProtectionGroupResourceType ApplicationLoadBalancer { get; } = new ProtectionGroupResourceType("APPLICATION_LOAD_BALANCER");
        public static ProtectionGroupResourceType GlobalAccelerator { get; } = new ProtectionGroupResourceType("GLOBAL_ACCELERATOR");

        public static bool operator ==(ProtectionGroupResourceType left, ProtectionGroupResourceType right) => left.Equals(right);
        public static bool operator !=(ProtectionGroupResourceType left, ProtectionGroupResourceType right) => !left.Equals(right);

        public static explicit operator string(ProtectionGroupResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtectionGroupResourceType other && Equals(other);
        public bool Equals(ProtectionGroupResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
