// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.SystemsManagerSap
{
    /// <summary>
    /// This string is the type of the component.
    /// 
    /// Accepted value is `WD` .
    /// </summary>
    [EnumType]
    public readonly struct ApplicationComponentInfoComponentType : IEquatable<ApplicationComponentInfoComponentType>
    {
        private readonly string _value;

        private ApplicationComponentInfoComponentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplicationComponentInfoComponentType Hana { get; } = new ApplicationComponentInfoComponentType("HANA");
        public static ApplicationComponentInfoComponentType HanaNode { get; } = new ApplicationComponentInfoComponentType("HANA_NODE");
        public static ApplicationComponentInfoComponentType Abap { get; } = new ApplicationComponentInfoComponentType("ABAP");
        public static ApplicationComponentInfoComponentType Ascs { get; } = new ApplicationComponentInfoComponentType("ASCS");
        public static ApplicationComponentInfoComponentType Dialog { get; } = new ApplicationComponentInfoComponentType("DIALOG");
        public static ApplicationComponentInfoComponentType Webdisp { get; } = new ApplicationComponentInfoComponentType("WEBDISP");
        public static ApplicationComponentInfoComponentType Wd { get; } = new ApplicationComponentInfoComponentType("WD");
        public static ApplicationComponentInfoComponentType Ers { get; } = new ApplicationComponentInfoComponentType("ERS");

        public static bool operator ==(ApplicationComponentInfoComponentType left, ApplicationComponentInfoComponentType right) => left.Equals(right);
        public static bool operator !=(ApplicationComponentInfoComponentType left, ApplicationComponentInfoComponentType right) => !left.Equals(right);

        public static explicit operator string(ApplicationComponentInfoComponentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplicationComponentInfoComponentType other && Equals(other);
        public bool Equals(ApplicationComponentInfoComponentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the application credentials.
    /// </summary>
    [EnumType]
    public readonly struct ApplicationCredentialCredentialType : IEquatable<ApplicationCredentialCredentialType>
    {
        private readonly string _value;

        private ApplicationCredentialCredentialType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplicationCredentialCredentialType Admin { get; } = new ApplicationCredentialCredentialType("ADMIN");

        public static bool operator ==(ApplicationCredentialCredentialType left, ApplicationCredentialCredentialType right) => left.Equals(right);
        public static bool operator !=(ApplicationCredentialCredentialType left, ApplicationCredentialCredentialType right) => !left.Equals(right);

        public static explicit operator string(ApplicationCredentialCredentialType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplicationCredentialCredentialType other && Equals(other);
        public bool Equals(ApplicationCredentialCredentialType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the application.
    /// </summary>
    [EnumType]
    public readonly struct ApplicationType : IEquatable<ApplicationType>
    {
        private readonly string _value;

        private ApplicationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplicationType Hana { get; } = new ApplicationType("HANA");
        public static ApplicationType SapAbap { get; } = new ApplicationType("SAP_ABAP");

        public static bool operator ==(ApplicationType left, ApplicationType right) => left.Equals(right);
        public static bool operator !=(ApplicationType left, ApplicationType right) => !left.Equals(right);

        public static explicit operator string(ApplicationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplicationType other && Equals(other);
        public bool Equals(ApplicationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
