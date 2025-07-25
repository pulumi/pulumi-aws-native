// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Billing
{
    [EnumType]
    public readonly struct BillingViewDimensionKey : IEquatable<BillingViewDimensionKey>
    {
        private readonly string _value;

        private BillingViewDimensionKey(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BillingViewDimensionKey LinkedAccount { get; } = new BillingViewDimensionKey("LINKED_ACCOUNT");

        public static bool operator ==(BillingViewDimensionKey left, BillingViewDimensionKey right) => left.Equals(right);
        public static bool operator !=(BillingViewDimensionKey left, BillingViewDimensionKey right) => !left.Equals(right);

        public static explicit operator string(BillingViewDimensionKey value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BillingViewDimensionKey other && Equals(other);
        public bool Equals(BillingViewDimensionKey other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct BillingViewType : IEquatable<BillingViewType>
    {
        private readonly string _value;

        private BillingViewType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BillingViewType Primary { get; } = new BillingViewType("PRIMARY");
        public static BillingViewType BillingGroup { get; } = new BillingViewType("BILLING_GROUP");
        public static BillingViewType Custom { get; } = new BillingViewType("CUSTOM");

        public static bool operator ==(BillingViewType left, BillingViewType right) => left.Equals(right);
        public static bool operator !=(BillingViewType left, BillingViewType right) => !left.Equals(right);

        public static explicit operator string(BillingViewType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BillingViewType other && Equals(other);
        public bool Equals(BillingViewType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
