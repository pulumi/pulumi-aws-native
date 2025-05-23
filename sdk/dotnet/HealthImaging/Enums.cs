// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.HealthImaging
{
    /// <summary>
    /// A string to denote the Datastore's state.
    /// </summary>
    [EnumType]
    public readonly struct DatastoreStatus : IEquatable<DatastoreStatus>
    {
        private readonly string _value;

        private DatastoreStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatastoreStatus Creating { get; } = new DatastoreStatus("CREATING");
        public static DatastoreStatus CreateFailed { get; } = new DatastoreStatus("CREATE_FAILED");
        public static DatastoreStatus Active { get; } = new DatastoreStatus("ACTIVE");
        public static DatastoreStatus Deleting { get; } = new DatastoreStatus("DELETING");
        public static DatastoreStatus Deleted { get; } = new DatastoreStatus("DELETED");

        public static bool operator ==(DatastoreStatus left, DatastoreStatus right) => left.Equals(right);
        public static bool operator !=(DatastoreStatus left, DatastoreStatus right) => !left.Equals(right);

        public static explicit operator string(DatastoreStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatastoreStatus other && Equals(other);
        public bool Equals(DatastoreStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
