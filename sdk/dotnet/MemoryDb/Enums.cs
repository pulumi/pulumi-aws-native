// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.MemoryDb
{
    [EnumType]
    public readonly struct ClusterDataTieringStatus : IEquatable<ClusterDataTieringStatus>
    {
        private readonly string _value;

        private ClusterDataTieringStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterDataTieringStatus True { get; } = new ClusterDataTieringStatus("true");
        public static ClusterDataTieringStatus False { get; } = new ClusterDataTieringStatus("false");

        public static bool operator ==(ClusterDataTieringStatus left, ClusterDataTieringStatus right) => left.Equals(right);
        public static bool operator !=(ClusterDataTieringStatus left, ClusterDataTieringStatus right) => !left.Equals(right);

        public static explicit operator string(ClusterDataTieringStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterDataTieringStatus other && Equals(other);
        public bool Equals(ClusterDataTieringStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ClusterSupportedIpDiscoveryTypes : IEquatable<ClusterSupportedIpDiscoveryTypes>
    {
        private readonly string _value;

        private ClusterSupportedIpDiscoveryTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterSupportedIpDiscoveryTypes Ipv4 { get; } = new ClusterSupportedIpDiscoveryTypes("ipv4");
        public static ClusterSupportedIpDiscoveryTypes Ipv6 { get; } = new ClusterSupportedIpDiscoveryTypes("ipv6");

        public static bool operator ==(ClusterSupportedIpDiscoveryTypes left, ClusterSupportedIpDiscoveryTypes right) => left.Equals(right);
        public static bool operator !=(ClusterSupportedIpDiscoveryTypes left, ClusterSupportedIpDiscoveryTypes right) => !left.Equals(right);

        public static explicit operator string(ClusterSupportedIpDiscoveryTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterSupportedIpDiscoveryTypes other && Equals(other);
        public bool Equals(ClusterSupportedIpDiscoveryTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ClusterSupportedNetworkTypes : IEquatable<ClusterSupportedNetworkTypes>
    {
        private readonly string _value;

        private ClusterSupportedNetworkTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterSupportedNetworkTypes Ipv4 { get; } = new ClusterSupportedNetworkTypes("ipv4");
        public static ClusterSupportedNetworkTypes Ipv6 { get; } = new ClusterSupportedNetworkTypes("ipv6");
        public static ClusterSupportedNetworkTypes DualStack { get; } = new ClusterSupportedNetworkTypes("dual_stack");

        public static bool operator ==(ClusterSupportedNetworkTypes left, ClusterSupportedNetworkTypes right) => left.Equals(right);
        public static bool operator !=(ClusterSupportedNetworkTypes left, ClusterSupportedNetworkTypes right) => !left.Equals(right);

        public static explicit operator string(ClusterSupportedNetworkTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterSupportedNetworkTypes other && Equals(other);
        public bool Equals(ClusterSupportedNetworkTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
    /// </summary>
    [EnumType]
    public readonly struct MultiRegionClusterUpdateStrategy : IEquatable<MultiRegionClusterUpdateStrategy>
    {
        private readonly string _value;

        private MultiRegionClusterUpdateStrategy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MultiRegionClusterUpdateStrategy Coordinated { get; } = new MultiRegionClusterUpdateStrategy("COORDINATED");
        public static MultiRegionClusterUpdateStrategy Uncoordinated { get; } = new MultiRegionClusterUpdateStrategy("UNCOORDINATED");

        public static bool operator ==(MultiRegionClusterUpdateStrategy left, MultiRegionClusterUpdateStrategy right) => left.Equals(right);
        public static bool operator !=(MultiRegionClusterUpdateStrategy left, MultiRegionClusterUpdateStrategy right) => !left.Equals(right);

        public static explicit operator string(MultiRegionClusterUpdateStrategy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MultiRegionClusterUpdateStrategy other && Equals(other);
        public bool Equals(MultiRegionClusterUpdateStrategy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of authentication strategy for this user.
    /// </summary>
    [EnumType]
    public readonly struct UserAuthenticationModePropertiesType : IEquatable<UserAuthenticationModePropertiesType>
    {
        private readonly string _value;

        private UserAuthenticationModePropertiesType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UserAuthenticationModePropertiesType Password { get; } = new UserAuthenticationModePropertiesType("password");
        public static UserAuthenticationModePropertiesType Iam { get; } = new UserAuthenticationModePropertiesType("iam");

        public static bool operator ==(UserAuthenticationModePropertiesType left, UserAuthenticationModePropertiesType right) => left.Equals(right);
        public static bool operator !=(UserAuthenticationModePropertiesType left, UserAuthenticationModePropertiesType right) => !left.Equals(right);

        public static explicit operator string(UserAuthenticationModePropertiesType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UserAuthenticationModePropertiesType other && Equals(other);
        public bool Equals(UserAuthenticationModePropertiesType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
