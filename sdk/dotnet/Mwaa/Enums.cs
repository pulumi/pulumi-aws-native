// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Mwaa
{
    /// <summary>
    /// Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentEndpointManagement : IEquatable<EnvironmentEndpointManagement>
    {
        private readonly string _value;

        private EnvironmentEndpointManagement(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentEndpointManagement Customer { get; } = new EnvironmentEndpointManagement("CUSTOMER");
        public static EnvironmentEndpointManagement Service { get; } = new EnvironmentEndpointManagement("SERVICE");

        public static bool operator ==(EnvironmentEndpointManagement left, EnvironmentEndpointManagement right) => left.Equals(right);
        public static bool operator !=(EnvironmentEndpointManagement left, EnvironmentEndpointManagement right) => !left.Equals(right);

        public static explicit operator string(EnvironmentEndpointManagement value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentEndpointManagement other && Equals(other);
        public bool Equals(EnvironmentEndpointManagement other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnvironmentLoggingLevel : IEquatable<EnvironmentLoggingLevel>
    {
        private readonly string _value;

        private EnvironmentLoggingLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentLoggingLevel Critical { get; } = new EnvironmentLoggingLevel("CRITICAL");
        public static EnvironmentLoggingLevel Error { get; } = new EnvironmentLoggingLevel("ERROR");
        public static EnvironmentLoggingLevel Warning { get; } = new EnvironmentLoggingLevel("WARNING");
        public static EnvironmentLoggingLevel Info { get; } = new EnvironmentLoggingLevel("INFO");
        public static EnvironmentLoggingLevel Debug { get; } = new EnvironmentLoggingLevel("DEBUG");

        public static bool operator ==(EnvironmentLoggingLevel left, EnvironmentLoggingLevel right) => left.Equals(right);
        public static bool operator !=(EnvironmentLoggingLevel left, EnvironmentLoggingLevel right) => !left.Equals(right);

        public static explicit operator string(EnvironmentLoggingLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentLoggingLevel other && Equals(other);
        public bool Equals(EnvironmentLoggingLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Choice for mode of webserver access including over public internet or via private VPC endpoint.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentWebserverAccessMode : IEquatable<EnvironmentWebserverAccessMode>
    {
        private readonly string _value;

        private EnvironmentWebserverAccessMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentWebserverAccessMode PrivateOnly { get; } = new EnvironmentWebserverAccessMode("PRIVATE_ONLY");
        public static EnvironmentWebserverAccessMode PublicOnly { get; } = new EnvironmentWebserverAccessMode("PUBLIC_ONLY");

        public static bool operator ==(EnvironmentWebserverAccessMode left, EnvironmentWebserverAccessMode right) => left.Equals(right);
        public static bool operator !=(EnvironmentWebserverAccessMode left, EnvironmentWebserverAccessMode right) => !left.Equals(right);

        public static explicit operator string(EnvironmentWebserverAccessMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentWebserverAccessMode other && Equals(other);
        public bool Equals(EnvironmentWebserverAccessMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
