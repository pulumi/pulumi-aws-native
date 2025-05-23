// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.M2
{
    [EnumType]
    public readonly struct ApplicationEngineType : IEquatable<ApplicationEngineType>
    {
        private readonly string _value;

        private ApplicationEngineType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplicationEngineType Microfocus { get; } = new ApplicationEngineType("microfocus");
        public static ApplicationEngineType Bluage { get; } = new ApplicationEngineType("bluage");

        public static bool operator ==(ApplicationEngineType left, ApplicationEngineType right) => left.Equals(right);
        public static bool operator !=(ApplicationEngineType left, ApplicationEngineType right) => !left.Equals(right);

        public static explicit operator string(ApplicationEngineType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplicationEngineType other && Equals(other);
        public bool Equals(ApplicationEngineType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The target platform for the environment.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentEngineType : IEquatable<EnvironmentEngineType>
    {
        private readonly string _value;

        private EnvironmentEngineType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentEngineType Microfocus { get; } = new EnvironmentEngineType("microfocus");
        public static EnvironmentEngineType Bluage { get; } = new EnvironmentEngineType("bluage");

        public static bool operator ==(EnvironmentEngineType left, EnvironmentEngineType right) => left.Equals(right);
        public static bool operator !=(EnvironmentEngineType left, EnvironmentEngineType right) => !left.Equals(right);

        public static explicit operator string(EnvironmentEngineType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentEngineType other && Equals(other);
        public bool Equals(EnvironmentEngineType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnvironmentNetworkType : IEquatable<EnvironmentNetworkType>
    {
        private readonly string _value;

        private EnvironmentNetworkType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentNetworkType Ipv4 { get; } = new EnvironmentNetworkType("ipv4");
        public static EnvironmentNetworkType Dual { get; } = new EnvironmentNetworkType("dual");

        public static bool operator ==(EnvironmentNetworkType left, EnvironmentNetworkType right) => left.Equals(right);
        public static bool operator !=(EnvironmentNetworkType left, EnvironmentNetworkType right) => !left.Equals(right);

        public static explicit operator string(EnvironmentNetworkType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentNetworkType other && Equals(other);
        public bool Equals(EnvironmentNetworkType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
