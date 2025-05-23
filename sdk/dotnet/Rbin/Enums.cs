// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Rbin
{
    /// <summary>
    /// The resource type retained by the retention rule.
    /// </summary>
    [EnumType]
    public readonly struct RuleResourceType : IEquatable<RuleResourceType>
    {
        private readonly string _value;

        private RuleResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RuleResourceType EbsSnapshot { get; } = new RuleResourceType("EBS_SNAPSHOT");
        public static RuleResourceType Ec2Image { get; } = new RuleResourceType("EC2_IMAGE");

        public static bool operator ==(RuleResourceType left, RuleResourceType right) => left.Equals(right);
        public static bool operator !=(RuleResourceType left, RuleResourceType right) => !left.Equals(right);

        public static explicit operator string(RuleResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RuleResourceType other && Equals(other);
        public bool Equals(RuleResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The retention period unit of the rule
    /// </summary>
    [EnumType]
    public readonly struct RuleRetentionPeriodRetentionPeriodUnit : IEquatable<RuleRetentionPeriodRetentionPeriodUnit>
    {
        private readonly string _value;

        private RuleRetentionPeriodRetentionPeriodUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RuleRetentionPeriodRetentionPeriodUnit Days { get; } = new RuleRetentionPeriodRetentionPeriodUnit("DAYS");

        public static bool operator ==(RuleRetentionPeriodRetentionPeriodUnit left, RuleRetentionPeriodRetentionPeriodUnit right) => left.Equals(right);
        public static bool operator !=(RuleRetentionPeriodRetentionPeriodUnit left, RuleRetentionPeriodRetentionPeriodUnit right) => !left.Equals(right);

        public static explicit operator string(RuleRetentionPeriodRetentionPeriodUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RuleRetentionPeriodRetentionPeriodUnit other && Equals(other);
        public bool Equals(RuleRetentionPeriodRetentionPeriodUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The unit of time in which to measure the unlock delay. Currently, the unlock delay can be measure only in days.
    /// </summary>
    [EnumType]
    public readonly struct RuleUnlockDelayUnlockDelayUnit : IEquatable<RuleUnlockDelayUnlockDelayUnit>
    {
        private readonly string _value;

        private RuleUnlockDelayUnlockDelayUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RuleUnlockDelayUnlockDelayUnit Days { get; } = new RuleUnlockDelayUnlockDelayUnit("DAYS");

        public static bool operator ==(RuleUnlockDelayUnlockDelayUnit left, RuleUnlockDelayUnlockDelayUnit right) => left.Equals(right);
        public static bool operator !=(RuleUnlockDelayUnlockDelayUnit left, RuleUnlockDelayUnlockDelayUnit right) => !left.Equals(right);

        public static explicit operator string(RuleUnlockDelayUnlockDelayUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RuleUnlockDelayUnlockDelayUnit other && Equals(other);
        public bool Equals(RuleUnlockDelayUnlockDelayUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
