// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.ResourceExplorer2
{
    [EnumType]
    public readonly struct IndexState : IEquatable<IndexState>
    {
        private readonly string _value;

        private IndexState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IndexState Active { get; } = new IndexState("ACTIVE");
        public static IndexState Creating { get; } = new IndexState("CREATING");
        public static IndexState Deleting { get; } = new IndexState("DELETING");
        public static IndexState Deleted { get; } = new IndexState("DELETED");
        public static IndexState Updating { get; } = new IndexState("UPDATING");

        public static bool operator ==(IndexState left, IndexState right) => left.Equals(right);
        public static bool operator !=(IndexState left, IndexState right) => !left.Equals(right);

        public static explicit operator string(IndexState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IndexState other && Equals(other);
        public bool Equals(IndexState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct IndexType : IEquatable<IndexType>
    {
        private readonly string _value;

        private IndexType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IndexType Local { get; } = new IndexType("LOCAL");
        public static IndexType Aggregator { get; } = new IndexType("AGGREGATOR");

        public static bool operator ==(IndexType left, IndexType right) => left.Equals(right);
        public static bool operator !=(IndexType left, IndexType right) => !left.Equals(right);

        public static explicit operator string(IndexType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IndexType other && Equals(other);
        public bool Equals(IndexType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
