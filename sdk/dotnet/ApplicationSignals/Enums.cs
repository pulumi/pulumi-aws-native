// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.ApplicationSignals
{
    /// <summary>
    /// Specifies the interval unit.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLevelObjectiveDurationUnit : IEquatable<ServiceLevelObjectiveDurationUnit>
    {
        private readonly string _value;

        private ServiceLevelObjectiveDurationUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLevelObjectiveDurationUnit Minute { get; } = new ServiceLevelObjectiveDurationUnit("MINUTE");
        public static ServiceLevelObjectiveDurationUnit Hour { get; } = new ServiceLevelObjectiveDurationUnit("HOUR");
        public static ServiceLevelObjectiveDurationUnit Day { get; } = new ServiceLevelObjectiveDurationUnit("DAY");
        public static ServiceLevelObjectiveDurationUnit Month { get; } = new ServiceLevelObjectiveDurationUnit("MONTH");

        public static bool operator ==(ServiceLevelObjectiveDurationUnit left, ServiceLevelObjectiveDurationUnit right) => left.Equals(right);
        public static bool operator !=(ServiceLevelObjectiveDurationUnit left, ServiceLevelObjectiveDurationUnit right) => !left.Equals(right);

        public static explicit operator string(ServiceLevelObjectiveDurationUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLevelObjectiveDurationUnit other && Equals(other);
        public bool Equals(ServiceLevelObjectiveDurationUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Displays whether this is a period-based SLO or a request-based SLO.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLevelObjectiveEvaluationType : IEquatable<ServiceLevelObjectiveEvaluationType>
    {
        private readonly string _value;

        private ServiceLevelObjectiveEvaluationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLevelObjectiveEvaluationType PeriodBased { get; } = new ServiceLevelObjectiveEvaluationType("PeriodBased");
        public static ServiceLevelObjectiveEvaluationType RequestBased { get; } = new ServiceLevelObjectiveEvaluationType("RequestBased");

        public static bool operator ==(ServiceLevelObjectiveEvaluationType left, ServiceLevelObjectiveEvaluationType right) => left.Equals(right);
        public static bool operator !=(ServiceLevelObjectiveEvaluationType left, ServiceLevelObjectiveEvaluationType right) => !left.Equals(right);

        public static explicit operator string(ServiceLevelObjectiveEvaluationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLevelObjectiveEvaluationType other && Equals(other);
        public bool Equals(ServiceLevelObjectiveEvaluationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The arithmetic operation used when comparing the specified metric to the threshold.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLevelObjectiveRequestBasedSliComparisonOperator : IEquatable<ServiceLevelObjectiveRequestBasedSliComparisonOperator>
    {
        private readonly string _value;

        private ServiceLevelObjectiveRequestBasedSliComparisonOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLevelObjectiveRequestBasedSliComparisonOperator GreaterThanOrEqualTo { get; } = new ServiceLevelObjectiveRequestBasedSliComparisonOperator("GreaterThanOrEqualTo");
        public static ServiceLevelObjectiveRequestBasedSliComparisonOperator LessThanOrEqualTo { get; } = new ServiceLevelObjectiveRequestBasedSliComparisonOperator("LessThanOrEqualTo");
        public static ServiceLevelObjectiveRequestBasedSliComparisonOperator LessThan { get; } = new ServiceLevelObjectiveRequestBasedSliComparisonOperator("LessThan");
        public static ServiceLevelObjectiveRequestBasedSliComparisonOperator GreaterThan { get; } = new ServiceLevelObjectiveRequestBasedSliComparisonOperator("GreaterThan");

        public static bool operator ==(ServiceLevelObjectiveRequestBasedSliComparisonOperator left, ServiceLevelObjectiveRequestBasedSliComparisonOperator right) => left.Equals(right);
        public static bool operator !=(ServiceLevelObjectiveRequestBasedSliComparisonOperator left, ServiceLevelObjectiveRequestBasedSliComparisonOperator right) => !left.Equals(right);

        public static explicit operator string(ServiceLevelObjectiveRequestBasedSliComparisonOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLevelObjectiveRequestBasedSliComparisonOperator other && Equals(other);
        public bool Equals(ServiceLevelObjectiveRequestBasedSliComparisonOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLevelObjectiveRequestBasedSliMetricMetricType : IEquatable<ServiceLevelObjectiveRequestBasedSliMetricMetricType>
    {
        private readonly string _value;

        private ServiceLevelObjectiveRequestBasedSliMetricMetricType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLevelObjectiveRequestBasedSliMetricMetricType Latency { get; } = new ServiceLevelObjectiveRequestBasedSliMetricMetricType("LATENCY");
        public static ServiceLevelObjectiveRequestBasedSliMetricMetricType Availability { get; } = new ServiceLevelObjectiveRequestBasedSliMetricMetricType("AVAILABILITY");

        public static bool operator ==(ServiceLevelObjectiveRequestBasedSliMetricMetricType left, ServiceLevelObjectiveRequestBasedSliMetricMetricType right) => left.Equals(right);
        public static bool operator !=(ServiceLevelObjectiveRequestBasedSliMetricMetricType left, ServiceLevelObjectiveRequestBasedSliMetricMetricType right) => !left.Equals(right);

        public static explicit operator string(ServiceLevelObjectiveRequestBasedSliMetricMetricType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLevelObjectiveRequestBasedSliMetricMetricType other && Equals(other);
        public bool Equals(ServiceLevelObjectiveRequestBasedSliMetricMetricType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The arithmetic operation used when comparing the specified metric to the threshold.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLevelObjectiveSliComparisonOperator : IEquatable<ServiceLevelObjectiveSliComparisonOperator>
    {
        private readonly string _value;

        private ServiceLevelObjectiveSliComparisonOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLevelObjectiveSliComparisonOperator GreaterThanOrEqualTo { get; } = new ServiceLevelObjectiveSliComparisonOperator("GreaterThanOrEqualTo");
        public static ServiceLevelObjectiveSliComparisonOperator LessThanOrEqualTo { get; } = new ServiceLevelObjectiveSliComparisonOperator("LessThanOrEqualTo");
        public static ServiceLevelObjectiveSliComparisonOperator LessThan { get; } = new ServiceLevelObjectiveSliComparisonOperator("LessThan");
        public static ServiceLevelObjectiveSliComparisonOperator GreaterThan { get; } = new ServiceLevelObjectiveSliComparisonOperator("GreaterThan");

        public static bool operator ==(ServiceLevelObjectiveSliComparisonOperator left, ServiceLevelObjectiveSliComparisonOperator right) => left.Equals(right);
        public static bool operator !=(ServiceLevelObjectiveSliComparisonOperator left, ServiceLevelObjectiveSliComparisonOperator right) => !left.Equals(right);

        public static explicit operator string(ServiceLevelObjectiveSliComparisonOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLevelObjectiveSliComparisonOperator other && Equals(other);
        public bool Equals(ServiceLevelObjectiveSliComparisonOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLevelObjectiveSliMetricMetricType : IEquatable<ServiceLevelObjectiveSliMetricMetricType>
    {
        private readonly string _value;

        private ServiceLevelObjectiveSliMetricMetricType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLevelObjectiveSliMetricMetricType Latency { get; } = new ServiceLevelObjectiveSliMetricMetricType("LATENCY");
        public static ServiceLevelObjectiveSliMetricMetricType Availability { get; } = new ServiceLevelObjectiveSliMetricMetricType("AVAILABILITY");

        public static bool operator ==(ServiceLevelObjectiveSliMetricMetricType left, ServiceLevelObjectiveSliMetricMetricType right) => left.Equals(right);
        public static bool operator !=(ServiceLevelObjectiveSliMetricMetricType left, ServiceLevelObjectiveSliMetricMetricType right) => !left.Equals(right);

        public static explicit operator string(ServiceLevelObjectiveSliMetricMetricType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLevelObjectiveSliMetricMetricType other && Equals(other);
        public bool Equals(ServiceLevelObjectiveSliMetricMetricType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
