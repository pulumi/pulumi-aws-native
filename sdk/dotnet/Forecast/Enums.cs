// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Forecast
{
    /// <summary>
    /// Data type of the field
    /// </summary>
    [EnumType]
    public readonly struct DatasetAttributesItemPropertiesAttributeType : IEquatable<DatasetAttributesItemPropertiesAttributeType>
    {
        private readonly string _value;

        private DatasetAttributesItemPropertiesAttributeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatasetAttributesItemPropertiesAttributeType String { get; } = new DatasetAttributesItemPropertiesAttributeType("string");
        public static DatasetAttributesItemPropertiesAttributeType Integer { get; } = new DatasetAttributesItemPropertiesAttributeType("integer");
        public static DatasetAttributesItemPropertiesAttributeType Float { get; } = new DatasetAttributesItemPropertiesAttributeType("float");
        public static DatasetAttributesItemPropertiesAttributeType Timestamp { get; } = new DatasetAttributesItemPropertiesAttributeType("timestamp");
        public static DatasetAttributesItemPropertiesAttributeType Geolocation { get; } = new DatasetAttributesItemPropertiesAttributeType("geolocation");

        public static bool operator ==(DatasetAttributesItemPropertiesAttributeType left, DatasetAttributesItemPropertiesAttributeType right) => left.Equals(right);
        public static bool operator !=(DatasetAttributesItemPropertiesAttributeType left, DatasetAttributesItemPropertiesAttributeType right) => !left.Equals(right);

        public static explicit operator string(DatasetAttributesItemPropertiesAttributeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatasetAttributesItemPropertiesAttributeType other && Equals(other);
        public bool Equals(DatasetAttributesItemPropertiesAttributeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The domain associated with the dataset
    /// </summary>
    [EnumType]
    public readonly struct DatasetDomain : IEquatable<DatasetDomain>
    {
        private readonly string _value;

        private DatasetDomain(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatasetDomain Retail { get; } = new DatasetDomain("RETAIL");
        public static DatasetDomain Custom { get; } = new DatasetDomain("CUSTOM");
        public static DatasetDomain InventoryPlanning { get; } = new DatasetDomain("INVENTORY_PLANNING");
        public static DatasetDomain Ec2Capacity { get; } = new DatasetDomain("EC2_CAPACITY");
        public static DatasetDomain WorkForce { get; } = new DatasetDomain("WORK_FORCE");
        public static DatasetDomain WebTraffic { get; } = new DatasetDomain("WEB_TRAFFIC");
        public static DatasetDomain Metrics { get; } = new DatasetDomain("METRICS");

        public static bool operator ==(DatasetDomain left, DatasetDomain right) => left.Equals(right);
        public static bool operator !=(DatasetDomain left, DatasetDomain right) => !left.Equals(right);

        public static explicit operator string(DatasetDomain value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatasetDomain other && Equals(other);
        public bool Equals(DatasetDomain other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.
    /// </summary>
    [EnumType]
    public readonly struct DatasetGroupDomain : IEquatable<DatasetGroupDomain>
    {
        private readonly string _value;

        private DatasetGroupDomain(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatasetGroupDomain Retail { get; } = new DatasetGroupDomain("RETAIL");
        public static DatasetGroupDomain Custom { get; } = new DatasetGroupDomain("CUSTOM");
        public static DatasetGroupDomain InventoryPlanning { get; } = new DatasetGroupDomain("INVENTORY_PLANNING");
        public static DatasetGroupDomain Ec2Capacity { get; } = new DatasetGroupDomain("EC2_CAPACITY");
        public static DatasetGroupDomain WorkForce { get; } = new DatasetGroupDomain("WORK_FORCE");
        public static DatasetGroupDomain WebTraffic { get; } = new DatasetGroupDomain("WEB_TRAFFIC");
        public static DatasetGroupDomain Metrics { get; } = new DatasetGroupDomain("METRICS");

        public static bool operator ==(DatasetGroupDomain left, DatasetGroupDomain right) => left.Equals(right);
        public static bool operator !=(DatasetGroupDomain left, DatasetGroupDomain right) => !left.Equals(right);

        public static explicit operator string(DatasetGroupDomain value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatasetGroupDomain other && Equals(other);
        public bool Equals(DatasetGroupDomain other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The dataset type
    /// </summary>
    [EnumType]
    public readonly struct DatasetType : IEquatable<DatasetType>
    {
        private readonly string _value;

        private DatasetType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatasetType TargetTimeSeries { get; } = new DatasetType("TARGET_TIME_SERIES");
        public static DatasetType RelatedTimeSeries { get; } = new DatasetType("RELATED_TIME_SERIES");
        public static DatasetType ItemMetadata { get; } = new DatasetType("ITEM_METADATA");

        public static bool operator ==(DatasetType left, DatasetType right) => left.Equals(right);
        public static bool operator !=(DatasetType left, DatasetType right) => !left.Equals(right);

        public static explicit operator string(DatasetType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatasetType other && Equals(other);
        public bool Equals(DatasetType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}