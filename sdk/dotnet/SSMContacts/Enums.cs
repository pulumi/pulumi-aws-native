// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.SSMContacts
{
    /// <summary>
    /// Device type, which specify notification channel. Currently supported values: “SMS”, “VOICE”, “EMAIL”, “CHATBOT.
    /// </summary>
    [EnumType]
    public readonly struct ContactChannelChannelType : IEquatable<ContactChannelChannelType>
    {
        private readonly string _value;

        private ContactChannelChannelType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContactChannelChannelType Sms { get; } = new ContactChannelChannelType("SMS");
        public static ContactChannelChannelType Voice { get; } = new ContactChannelChannelType("VOICE");
        public static ContactChannelChannelType Email { get; } = new ContactChannelChannelType("EMAIL");

        public static bool operator ==(ContactChannelChannelType left, ContactChannelChannelType right) => left.Equals(right);
        public static bool operator !=(ContactChannelChannelType left, ContactChannelChannelType right) => !left.Equals(right);

        public static explicit operator string(ContactChannelChannelType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactChannelChannelType other && Equals(other);
        public bool Equals(ContactChannelChannelType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“.
    /// </summary>
    [EnumType]
    public readonly struct ContactType : IEquatable<ContactType>
    {
        private readonly string _value;

        private ContactType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContactType Personal { get; } = new ContactType("PERSONAL");
        public static ContactType Custom { get; } = new ContactType("CUSTOM");
        public static ContactType Service { get; } = new ContactType("SERVICE");
        public static ContactType Escalation { get; } = new ContactType("ESCALATION");
        public static ContactType OncallSchedule { get; } = new ContactType("ONCALL_SCHEDULE");

        public static bool operator ==(ContactType left, ContactType right) => left.Equals(right);
        public static bool operator !=(ContactType left, ContactType right) => !left.Equals(right);

        public static explicit operator string(ContactType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactType other && Equals(other);
        public bool Equals(ContactType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The day of the week when weekly recurring on-call shift rotations begin. 
    /// </summary>
    [EnumType]
    public readonly struct RotationDayOfWeek : IEquatable<RotationDayOfWeek>
    {
        private readonly string _value;

        private RotationDayOfWeek(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RotationDayOfWeek Mon { get; } = new RotationDayOfWeek("MON");
        public static RotationDayOfWeek Tue { get; } = new RotationDayOfWeek("TUE");
        public static RotationDayOfWeek Wed { get; } = new RotationDayOfWeek("WED");
        public static RotationDayOfWeek Thu { get; } = new RotationDayOfWeek("THU");
        public static RotationDayOfWeek Fri { get; } = new RotationDayOfWeek("FRI");
        public static RotationDayOfWeek Sat { get; } = new RotationDayOfWeek("SAT");
        public static RotationDayOfWeek Sun { get; } = new RotationDayOfWeek("SUN");

        public static bool operator ==(RotationDayOfWeek left, RotationDayOfWeek right) => left.Equals(right);
        public static bool operator !=(RotationDayOfWeek left, RotationDayOfWeek right) => !left.Equals(right);

        public static explicit operator string(RotationDayOfWeek value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RotationDayOfWeek other && Equals(other);
        public bool Equals(RotationDayOfWeek other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}