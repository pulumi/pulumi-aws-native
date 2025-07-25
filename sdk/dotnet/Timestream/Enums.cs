// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Timestream
{
    /// <summary>
    /// The compute instance of the InfluxDB instance.
    /// </summary>
    [EnumType]
    public readonly struct InfluxDbInstanceDbInstanceType : IEquatable<InfluxDbInstanceDbInstanceType>
    {
        private readonly string _value;

        private InfluxDbInstanceDbInstanceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfluxDbInstanceDbInstanceType DbInfluxMedium { get; } = new InfluxDbInstanceDbInstanceType("db.influx.medium");
        public static InfluxDbInstanceDbInstanceType DbInfluxLarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.large");
        public static InfluxDbInstanceDbInstanceType DbInfluxXlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.xlarge");
        public static InfluxDbInstanceDbInstanceType DbInflux2xlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.2xlarge");
        public static InfluxDbInstanceDbInstanceType DbInflux4xlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.4xlarge");
        public static InfluxDbInstanceDbInstanceType DbInflux8xlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.8xlarge");
        public static InfluxDbInstanceDbInstanceType DbInflux12xlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.12xlarge");
        public static InfluxDbInstanceDbInstanceType DbInflux16xlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.16xlarge");
        public static InfluxDbInstanceDbInstanceType DbInflux24xlarge { get; } = new InfluxDbInstanceDbInstanceType("db.influx.24xlarge");

        public static bool operator ==(InfluxDbInstanceDbInstanceType left, InfluxDbInstanceDbInstanceType right) => left.Equals(right);
        public static bool operator !=(InfluxDbInstanceDbInstanceType left, InfluxDbInstanceDbInstanceType right) => !left.Equals(right);

        public static explicit operator string(InfluxDbInstanceDbInstanceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfluxDbInstanceDbInstanceType other && Equals(other);
        public bool Equals(InfluxDbInstanceDbInstanceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The storage type of the InfluxDB instance.
    /// </summary>
    [EnumType]
    public readonly struct InfluxDbInstanceDbStorageType : IEquatable<InfluxDbInstanceDbStorageType>
    {
        private readonly string _value;

        private InfluxDbInstanceDbStorageType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfluxDbInstanceDbStorageType InfluxIoIncludedT1 { get; } = new InfluxDbInstanceDbStorageType("InfluxIOIncludedT1");
        public static InfluxDbInstanceDbStorageType InfluxIoIncludedT2 { get; } = new InfluxDbInstanceDbStorageType("InfluxIOIncludedT2");
        public static InfluxDbInstanceDbStorageType InfluxIoIncludedT3 { get; } = new InfluxDbInstanceDbStorageType("InfluxIOIncludedT3");

        public static bool operator ==(InfluxDbInstanceDbStorageType left, InfluxDbInstanceDbStorageType right) => left.Equals(right);
        public static bool operator !=(InfluxDbInstanceDbStorageType left, InfluxDbInstanceDbStorageType right) => !left.Equals(right);

        public static explicit operator string(InfluxDbInstanceDbStorageType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfluxDbInstanceDbStorageType other && Equals(other);
        public bool Equals(InfluxDbInstanceDbStorageType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Deployment type of the InfluxDB Instance.
    /// </summary>
    [EnumType]
    public readonly struct InfluxDbInstanceDeploymentType : IEquatable<InfluxDbInstanceDeploymentType>
    {
        private readonly string _value;

        private InfluxDbInstanceDeploymentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfluxDbInstanceDeploymentType SingleAz { get; } = new InfluxDbInstanceDeploymentType("SINGLE_AZ");
        public static InfluxDbInstanceDeploymentType WithMultiazStandby { get; } = new InfluxDbInstanceDeploymentType("WITH_MULTIAZ_STANDBY");

        public static bool operator ==(InfluxDbInstanceDeploymentType left, InfluxDbInstanceDeploymentType right) => left.Equals(right);
        public static bool operator !=(InfluxDbInstanceDeploymentType left, InfluxDbInstanceDeploymentType right) => !left.Equals(right);

        public static explicit operator string(InfluxDbInstanceDeploymentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfluxDbInstanceDeploymentType other && Equals(other);
        public bool Equals(InfluxDbInstanceDeploymentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Network type of the InfluxDB Instance.
    /// </summary>
    [EnumType]
    public readonly struct InfluxDbInstanceNetworkType : IEquatable<InfluxDbInstanceNetworkType>
    {
        private readonly string _value;

        private InfluxDbInstanceNetworkType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfluxDbInstanceNetworkType Ipv4 { get; } = new InfluxDbInstanceNetworkType("IPV4");
        public static InfluxDbInstanceNetworkType Dual { get; } = new InfluxDbInstanceNetworkType("DUAL");

        public static bool operator ==(InfluxDbInstanceNetworkType left, InfluxDbInstanceNetworkType right) => left.Equals(right);
        public static bool operator !=(InfluxDbInstanceNetworkType left, InfluxDbInstanceNetworkType right) => !left.Equals(right);

        public static explicit operator string(InfluxDbInstanceNetworkType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfluxDbInstanceNetworkType other && Equals(other);
        public bool Equals(InfluxDbInstanceNetworkType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Status of the InfluxDB Instance.
    /// </summary>
    [EnumType]
    public readonly struct InfluxDbInstanceStatus : IEquatable<InfluxDbInstanceStatus>
    {
        private readonly string _value;

        private InfluxDbInstanceStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InfluxDbInstanceStatus Creating { get; } = new InfluxDbInstanceStatus("CREATING");
        public static InfluxDbInstanceStatus Available { get; } = new InfluxDbInstanceStatus("AVAILABLE");
        public static InfluxDbInstanceStatus Deleting { get; } = new InfluxDbInstanceStatus("DELETING");
        public static InfluxDbInstanceStatus Modifying { get; } = new InfluxDbInstanceStatus("MODIFYING");
        public static InfluxDbInstanceStatus Updating { get; } = new InfluxDbInstanceStatus("UPDATING");
        public static InfluxDbInstanceStatus UpdatingDeploymentType { get; } = new InfluxDbInstanceStatus("UPDATING_DEPLOYMENT_TYPE");
        public static InfluxDbInstanceStatus UpdatingInstanceType { get; } = new InfluxDbInstanceStatus("UPDATING_INSTANCE_TYPE");
        public static InfluxDbInstanceStatus Deleted { get; } = new InfluxDbInstanceStatus("DELETED");
        public static InfluxDbInstanceStatus Failed { get; } = new InfluxDbInstanceStatus("FAILED");

        public static bool operator ==(InfluxDbInstanceStatus left, InfluxDbInstanceStatus right) => left.Equals(right);
        public static bool operator !=(InfluxDbInstanceStatus left, InfluxDbInstanceStatus right) => !left.Equals(right);

        public static explicit operator string(InfluxDbInstanceStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InfluxDbInstanceStatus other && Equals(other);
        public bool Equals(InfluxDbInstanceStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type for the dimension.
    /// </summary>
    [EnumType]
    public readonly struct ScheduledQueryDimensionValueType : IEquatable<ScheduledQueryDimensionValueType>
    {
        private readonly string _value;

        private ScheduledQueryDimensionValueType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduledQueryDimensionValueType Varchar { get; } = new ScheduledQueryDimensionValueType("VARCHAR");

        public static bool operator ==(ScheduledQueryDimensionValueType left, ScheduledQueryDimensionValueType right) => left.Equals(right);
        public static bool operator !=(ScheduledQueryDimensionValueType left, ScheduledQueryDimensionValueType right) => !left.Equals(right);

        public static explicit operator string(ScheduledQueryDimensionValueType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduledQueryDimensionValueType other && Equals(other);
        public bool Equals(ScheduledQueryDimensionValueType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Encryption at rest options for the error reports. If no encryption option is specified, Timestream will choose SSE_S3 as default.
    /// </summary>
    [EnumType]
    public readonly struct ScheduledQueryEncryptionOption : IEquatable<ScheduledQueryEncryptionOption>
    {
        private readonly string _value;

        private ScheduledQueryEncryptionOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduledQueryEncryptionOption SseS3 { get; } = new ScheduledQueryEncryptionOption("SSE_S3");
        public static ScheduledQueryEncryptionOption SseKms { get; } = new ScheduledQueryEncryptionOption("SSE_KMS");

        public static bool operator ==(ScheduledQueryEncryptionOption left, ScheduledQueryEncryptionOption right) => left.Equals(right);
        public static bool operator !=(ScheduledQueryEncryptionOption left, ScheduledQueryEncryptionOption right) => !left.Equals(right);

        public static explicit operator string(ScheduledQueryEncryptionOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduledQueryEncryptionOption other && Equals(other);
        public bool Equals(ScheduledQueryEncryptionOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the value that is to be read from SourceColumn. If the mapping is for MULTI, use MeasureValueType.MULTI.
    /// </summary>
    [EnumType]
    public readonly struct ScheduledQueryMixedMeasureMappingMeasureValueType : IEquatable<ScheduledQueryMixedMeasureMappingMeasureValueType>
    {
        private readonly string _value;

        private ScheduledQueryMixedMeasureMappingMeasureValueType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduledQueryMixedMeasureMappingMeasureValueType Bigint { get; } = new ScheduledQueryMixedMeasureMappingMeasureValueType("BIGINT");
        public static ScheduledQueryMixedMeasureMappingMeasureValueType Boolean { get; } = new ScheduledQueryMixedMeasureMappingMeasureValueType("BOOLEAN");
        public static ScheduledQueryMixedMeasureMappingMeasureValueType Double { get; } = new ScheduledQueryMixedMeasureMappingMeasureValueType("DOUBLE");
        public static ScheduledQueryMixedMeasureMappingMeasureValueType Varchar { get; } = new ScheduledQueryMixedMeasureMappingMeasureValueType("VARCHAR");
        public static ScheduledQueryMixedMeasureMappingMeasureValueType Multi { get; } = new ScheduledQueryMixedMeasureMappingMeasureValueType("MULTI");

        public static bool operator ==(ScheduledQueryMixedMeasureMappingMeasureValueType left, ScheduledQueryMixedMeasureMappingMeasureValueType right) => left.Equals(right);
        public static bool operator !=(ScheduledQueryMixedMeasureMappingMeasureValueType left, ScheduledQueryMixedMeasureMappingMeasureValueType right) => !left.Equals(right);

        public static explicit operator string(ScheduledQueryMixedMeasureMappingMeasureValueType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduledQueryMixedMeasureMappingMeasureValueType other && Equals(other);
        public bool Equals(ScheduledQueryMixedMeasureMappingMeasureValueType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Value type of the measure value column to be read from the query result.
    /// </summary>
    [EnumType]
    public readonly struct ScheduledQueryMultiMeasureAttributeMappingMeasureValueType : IEquatable<ScheduledQueryMultiMeasureAttributeMappingMeasureValueType>
    {
        private readonly string _value;

        private ScheduledQueryMultiMeasureAttributeMappingMeasureValueType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduledQueryMultiMeasureAttributeMappingMeasureValueType Bigint { get; } = new ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("BIGINT");
        public static ScheduledQueryMultiMeasureAttributeMappingMeasureValueType Boolean { get; } = new ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("BOOLEAN");
        public static ScheduledQueryMultiMeasureAttributeMappingMeasureValueType Double { get; } = new ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("DOUBLE");
        public static ScheduledQueryMultiMeasureAttributeMappingMeasureValueType Varchar { get; } = new ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("VARCHAR");
        public static ScheduledQueryMultiMeasureAttributeMappingMeasureValueType Timestamp { get; } = new ScheduledQueryMultiMeasureAttributeMappingMeasureValueType("TIMESTAMP");

        public static bool operator ==(ScheduledQueryMultiMeasureAttributeMappingMeasureValueType left, ScheduledQueryMultiMeasureAttributeMappingMeasureValueType right) => left.Equals(right);
        public static bool operator !=(ScheduledQueryMultiMeasureAttributeMappingMeasureValueType left, ScheduledQueryMultiMeasureAttributeMappingMeasureValueType right) => !left.Equals(right);

        public static explicit operator string(ScheduledQueryMultiMeasureAttributeMappingMeasureValueType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduledQueryMultiMeasureAttributeMappingMeasureValueType other && Equals(other);
        public bool Equals(ScheduledQueryMultiMeasureAttributeMappingMeasureValueType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The level of enforcement for the specification of a dimension key in ingested records. Options are REQUIRED (dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).
    /// </summary>
    [EnumType]
    public readonly struct TablePartitionKeyEnforcementLevel : IEquatable<TablePartitionKeyEnforcementLevel>
    {
        private readonly string _value;

        private TablePartitionKeyEnforcementLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TablePartitionKeyEnforcementLevel Required { get; } = new TablePartitionKeyEnforcementLevel("REQUIRED");
        public static TablePartitionKeyEnforcementLevel Optional { get; } = new TablePartitionKeyEnforcementLevel("OPTIONAL");

        public static bool operator ==(TablePartitionKeyEnforcementLevel left, TablePartitionKeyEnforcementLevel right) => left.Equals(right);
        public static bool operator !=(TablePartitionKeyEnforcementLevel left, TablePartitionKeyEnforcementLevel right) => !left.Equals(right);

        public static explicit operator string(TablePartitionKeyEnforcementLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TablePartitionKeyEnforcementLevel other && Equals(other);
        public bool Equals(TablePartitionKeyEnforcementLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the partition key. Options are DIMENSION (dimension key) and MEASURE (measure key).
    /// </summary>
    [EnumType]
    public readonly struct TablePartitionKeyType : IEquatable<TablePartitionKeyType>
    {
        private readonly string _value;

        private TablePartitionKeyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TablePartitionKeyType Dimension { get; } = new TablePartitionKeyType("DIMENSION");
        public static TablePartitionKeyType Measure { get; } = new TablePartitionKeyType("MEASURE");

        public static bool operator ==(TablePartitionKeyType left, TablePartitionKeyType right) => left.Equals(right);
        public static bool operator !=(TablePartitionKeyType left, TablePartitionKeyType right) => !left.Equals(right);

        public static explicit operator string(TablePartitionKeyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TablePartitionKeyType other && Equals(other);
        public bool Equals(TablePartitionKeyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
