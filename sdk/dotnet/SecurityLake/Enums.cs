// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.SecurityLake
{
    [EnumType]
    public readonly struct SubscriberAccessTypesItem : IEquatable<SubscriberAccessTypesItem>
    {
        private readonly string _value;

        private SubscriberAccessTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SubscriberAccessTypesItem Lakeformation { get; } = new SubscriberAccessTypesItem("LAKEFORMATION");
        public static SubscriberAccessTypesItem S3 { get; } = new SubscriberAccessTypesItem("S3");

        public static bool operator ==(SubscriberAccessTypesItem left, SubscriberAccessTypesItem right) => left.Equals(right);
        public static bool operator !=(SubscriberAccessTypesItem left, SubscriberAccessTypesItem right) => !left.Equals(right);

        public static explicit operator string(SubscriberAccessTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SubscriberAccessTypesItem other && Equals(other);
        public bool Equals(SubscriberAccessTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}