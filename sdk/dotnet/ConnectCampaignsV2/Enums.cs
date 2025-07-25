// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.ConnectCampaignsV2
{
    /// <summary>
    /// The communication limit time unit
    /// </summary>
    [EnumType]
    public readonly struct CampaignCommunicationLimitTimeUnit : IEquatable<CampaignCommunicationLimitTimeUnit>
    {
        private readonly string _value;

        private CampaignCommunicationLimitTimeUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignCommunicationLimitTimeUnit Day { get; } = new CampaignCommunicationLimitTimeUnit("DAY");

        public static bool operator ==(CampaignCommunicationLimitTimeUnit left, CampaignCommunicationLimitTimeUnit right) => left.Equals(right);
        public static bool operator !=(CampaignCommunicationLimitTimeUnit left, CampaignCommunicationLimitTimeUnit right) => !left.Equals(right);

        public static explicit operator string(CampaignCommunicationLimitTimeUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignCommunicationLimitTimeUnit other && Equals(other);
        public bool Equals(CampaignCommunicationLimitTimeUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Day of week
    /// </summary>
    [EnumType]
    public readonly struct CampaignDayOfWeek : IEquatable<CampaignDayOfWeek>
    {
        private readonly string _value;

        private CampaignDayOfWeek(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignDayOfWeek Monday { get; } = new CampaignDayOfWeek("MONDAY");
        public static CampaignDayOfWeek Tuesday { get; } = new CampaignDayOfWeek("TUESDAY");
        public static CampaignDayOfWeek Wednesday { get; } = new CampaignDayOfWeek("WEDNESDAY");
        public static CampaignDayOfWeek Thursday { get; } = new CampaignDayOfWeek("THURSDAY");
        public static CampaignDayOfWeek Friday { get; } = new CampaignDayOfWeek("FRIDAY");
        public static CampaignDayOfWeek Saturday { get; } = new CampaignDayOfWeek("SATURDAY");
        public static CampaignDayOfWeek Sunday { get; } = new CampaignDayOfWeek("SUNDAY");

        public static bool operator ==(CampaignDayOfWeek left, CampaignDayOfWeek right) => left.Equals(right);
        public static bool operator !=(CampaignDayOfWeek left, CampaignDayOfWeek right) => !left.Equals(right);

        public static explicit operator string(CampaignDayOfWeek value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignDayOfWeek other && Equals(other);
        public bool Equals(CampaignDayOfWeek other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Enumeration of Instance Limits handling in a Campaign
    /// </summary>
    [EnumType]
    public readonly struct CampaignInstanceLimitsHandling : IEquatable<CampaignInstanceLimitsHandling>
    {
        private readonly string _value;

        private CampaignInstanceLimitsHandling(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignInstanceLimitsHandling OptIn { get; } = new CampaignInstanceLimitsHandling("OPT_IN");
        public static CampaignInstanceLimitsHandling OptOut { get; } = new CampaignInstanceLimitsHandling("OPT_OUT");

        public static bool operator ==(CampaignInstanceLimitsHandling left, CampaignInstanceLimitsHandling right) => left.Equals(right);
        public static bool operator !=(CampaignInstanceLimitsHandling left, CampaignInstanceLimitsHandling right) => !left.Equals(right);

        public static explicit operator string(CampaignInstanceLimitsHandling value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignInstanceLimitsHandling other && Equals(other);
        public bool Equals(CampaignInstanceLimitsHandling other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Local TimeZone Detection method
    /// </summary>
    [EnumType]
    public readonly struct CampaignLocalTimeZoneDetectionType : IEquatable<CampaignLocalTimeZoneDetectionType>
    {
        private readonly string _value;

        private CampaignLocalTimeZoneDetectionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignLocalTimeZoneDetectionType ZipCode { get; } = new CampaignLocalTimeZoneDetectionType("ZIP_CODE");
        public static CampaignLocalTimeZoneDetectionType AreaCode { get; } = new CampaignLocalTimeZoneDetectionType("AREA_CODE");

        public static bool operator ==(CampaignLocalTimeZoneDetectionType left, CampaignLocalTimeZoneDetectionType right) => left.Equals(right);
        public static bool operator !=(CampaignLocalTimeZoneDetectionType left, CampaignLocalTimeZoneDetectionType right) => !left.Equals(right);

        public static explicit operator string(CampaignLocalTimeZoneDetectionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignLocalTimeZoneDetectionType other && Equals(other);
        public bool Equals(CampaignLocalTimeZoneDetectionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
