// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.WorkspacesInstances
{
    /// <summary>
    /// Mode to use when disassociating the volume
    /// </summary>
    [EnumType]
    public readonly struct VolumeAssociationDisassociateMode : IEquatable<VolumeAssociationDisassociateMode>
    {
        private readonly string _value;

        private VolumeAssociationDisassociateMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VolumeAssociationDisassociateMode Force { get; } = new VolumeAssociationDisassociateMode("FORCE");
        public static VolumeAssociationDisassociateMode NoForce { get; } = new VolumeAssociationDisassociateMode("NO_FORCE");

        public static bool operator ==(VolumeAssociationDisassociateMode left, VolumeAssociationDisassociateMode right) => left.Equals(right);
        public static bool operator !=(VolumeAssociationDisassociateMode left, VolumeAssociationDisassociateMode right) => !left.Equals(right);

        public static explicit operator string(VolumeAssociationDisassociateMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeAssociationDisassociateMode other && Equals(other);
        public bool Equals(VolumeAssociationDisassociateMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct VolumeTagSpecificationResourceType : IEquatable<VolumeTagSpecificationResourceType>
    {
        private readonly string _value;

        private VolumeTagSpecificationResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VolumeTagSpecificationResourceType Instance { get; } = new VolumeTagSpecificationResourceType("instance");
        public static VolumeTagSpecificationResourceType Volume { get; } = new VolumeTagSpecificationResourceType("volume");
        public static VolumeTagSpecificationResourceType SpotInstancesRequest { get; } = new VolumeTagSpecificationResourceType("spot-instances-request");
        public static VolumeTagSpecificationResourceType NetworkInterface { get; } = new VolumeTagSpecificationResourceType("network-interface");

        public static bool operator ==(VolumeTagSpecificationResourceType left, VolumeTagSpecificationResourceType right) => left.Equals(right);
        public static bool operator !=(VolumeTagSpecificationResourceType left, VolumeTagSpecificationResourceType right) => !left.Equals(right);

        public static explicit operator string(VolumeTagSpecificationResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeTagSpecificationResourceType other && Equals(other);
        public bool Equals(VolumeTagSpecificationResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The volume type
    /// </summary>
    [EnumType]
    public readonly struct VolumeType : IEquatable<VolumeType>
    {
        private readonly string _value;

        private VolumeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VolumeType Standard { get; } = new VolumeType("standard");
        public static VolumeType Io1 { get; } = new VolumeType("io1");
        public static VolumeType Io2 { get; } = new VolumeType("io2");
        public static VolumeType Gp2 { get; } = new VolumeType("gp2");
        public static VolumeType Sc1 { get; } = new VolumeType("sc1");
        public static VolumeType St1 { get; } = new VolumeType("st1");
        public static VolumeType Gp3 { get; } = new VolumeType("gp3");

        public static bool operator ==(VolumeType left, VolumeType right) => left.Equals(right);
        public static bool operator !=(VolumeType left, VolumeType right) => !left.Equals(right);

        public static explicit operator string(VolumeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VolumeType other && Equals(other);
        public bool Equals(VolumeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceCreditSpecificationRequestCpuCredits : IEquatable<WorkspaceInstanceCreditSpecificationRequestCpuCredits>
    {
        private readonly string _value;

        private WorkspaceInstanceCreditSpecificationRequestCpuCredits(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceCreditSpecificationRequestCpuCredits Standard { get; } = new WorkspaceInstanceCreditSpecificationRequestCpuCredits("standard");
        public static WorkspaceInstanceCreditSpecificationRequestCpuCredits Unlimited { get; } = new WorkspaceInstanceCreditSpecificationRequestCpuCredits("unlimited");

        public static bool operator ==(WorkspaceInstanceCreditSpecificationRequestCpuCredits left, WorkspaceInstanceCreditSpecificationRequestCpuCredits right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceCreditSpecificationRequestCpuCredits left, WorkspaceInstanceCreditSpecificationRequestCpuCredits right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceCreditSpecificationRequestCpuCredits value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceCreditSpecificationRequestCpuCredits other && Equals(other);
        public bool Equals(WorkspaceInstanceCreditSpecificationRequestCpuCredits other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceEbsBlockDeviceVolumeType : IEquatable<WorkspaceInstanceEbsBlockDeviceVolumeType>
    {
        private readonly string _value;

        private WorkspaceInstanceEbsBlockDeviceVolumeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceEbsBlockDeviceVolumeType Standard { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("standard");
        public static WorkspaceInstanceEbsBlockDeviceVolumeType Io1 { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("io1");
        public static WorkspaceInstanceEbsBlockDeviceVolumeType Io2 { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("io2");
        public static WorkspaceInstanceEbsBlockDeviceVolumeType Gp2 { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("gp2");
        public static WorkspaceInstanceEbsBlockDeviceVolumeType Sc1 { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("sc1");
        public static WorkspaceInstanceEbsBlockDeviceVolumeType St1 { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("st1");
        public static WorkspaceInstanceEbsBlockDeviceVolumeType Gp3 { get; } = new WorkspaceInstanceEbsBlockDeviceVolumeType("gp3");

        public static bool operator ==(WorkspaceInstanceEbsBlockDeviceVolumeType left, WorkspaceInstanceEbsBlockDeviceVolumeType right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceEbsBlockDeviceVolumeType left, WorkspaceInstanceEbsBlockDeviceVolumeType right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceEbsBlockDeviceVolumeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceEbsBlockDeviceVolumeType other && Equals(other);
        public bool Equals(WorkspaceInstanceEbsBlockDeviceVolumeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery : IEquatable<WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery>
    {
        private readonly string _value;

        private WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery Disabled { get; } = new WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery("disabled");
        public static WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery Default { get; } = new WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery("default");

        public static bool operator ==(WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery left, WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery left, WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery other && Equals(other);
        public bool Equals(WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint : IEquatable<WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint>
    {
        private readonly string _value;

        private WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint Enabled { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint("enabled");
        public static WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint Disabled { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint("disabled");

        public static bool operator ==(WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint left, WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint left, WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint other && Equals(other);
        public bool Equals(WorkspaceInstanceInstanceMetadataOptionsRequestHttpEndpoint other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 : IEquatable<WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6>
    {
        private readonly string _value;

        private WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 Enabled { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6("enabled");
        public static WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 Disabled { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6("disabled");

        public static bool operator ==(WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 left, WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 left, WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 other && Equals(other);
        public bool Equals(WorkspaceInstanceInstanceMetadataOptionsRequestHttpProtocolIpv6 other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens : IEquatable<WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens>
    {
        private readonly string _value;

        private WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens Optional { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens("optional");
        public static WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens Required { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens("required");

        public static bool operator ==(WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens left, WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens left, WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens other && Equals(other);
        public bool Equals(WorkspaceInstanceInstanceMetadataOptionsRequestHttpTokens other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags : IEquatable<WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags>
    {
        private readonly string _value;

        private WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags Enabled { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags("enabled");
        public static WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags Disabled { get; } = new WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags("disabled");

        public static bool operator ==(WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags left, WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags left, WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags other && Equals(other);
        public bool Equals(WorkspaceInstanceInstanceMetadataOptionsRequestInstanceMetadataTags other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting : IEquatable<WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting>
    {
        private readonly string _value;

        private WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting Default { get; } = new WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting("default");
        public static WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting Vpc1 { get; } = new WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting("vpc-1");
        public static WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting Ebs1 { get; } = new WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting("ebs-1");

        public static bool operator ==(WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting left, WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting left, WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting other && Equals(other);
        public bool Equals(WorkspaceInstanceInstanceNetworkPerformanceOptionsRequestBandwidthWeighting other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstancePlacementTenancy : IEquatable<WorkspaceInstancePlacementTenancy>
    {
        private readonly string _value;

        private WorkspaceInstancePlacementTenancy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstancePlacementTenancy Default { get; } = new WorkspaceInstancePlacementTenancy("default");
        public static WorkspaceInstancePlacementTenancy Dedicated { get; } = new WorkspaceInstancePlacementTenancy("dedicated");
        public static WorkspaceInstancePlacementTenancy Host { get; } = new WorkspaceInstancePlacementTenancy("host");

        public static bool operator ==(WorkspaceInstancePlacementTenancy left, WorkspaceInstancePlacementTenancy right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstancePlacementTenancy left, WorkspaceInstancePlacementTenancy right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstancePlacementTenancy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstancePlacementTenancy other && Equals(other);
        public bool Equals(WorkspaceInstancePlacementTenancy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType : IEquatable<WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType>
    {
        private readonly string _value;

        private WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType IpName { get; } = new WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType("ip-name");
        public static WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType ResourceName { get; } = new WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType("resource-name");

        public static bool operator ==(WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType left, WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType left, WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType other && Equals(other);
        public bool Equals(WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the workspace instance
    /// </summary>
    [EnumType]
    public readonly struct WorkspaceInstanceProvisionState : IEquatable<WorkspaceInstanceProvisionState>
    {
        private readonly string _value;

        private WorkspaceInstanceProvisionState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceProvisionState Allocating { get; } = new WorkspaceInstanceProvisionState("ALLOCATING");
        public static WorkspaceInstanceProvisionState Allocated { get; } = new WorkspaceInstanceProvisionState("ALLOCATED");
        public static WorkspaceInstanceProvisionState Deallocating { get; } = new WorkspaceInstanceProvisionState("DEALLOCATING");
        public static WorkspaceInstanceProvisionState Deallocated { get; } = new WorkspaceInstanceProvisionState("DEALLOCATED");
        public static WorkspaceInstanceProvisionState ErrorAllocating { get; } = new WorkspaceInstanceProvisionState("ERROR_ALLOCATING");
        public static WorkspaceInstanceProvisionState ErrorDeallocating { get; } = new WorkspaceInstanceProvisionState("ERROR_DEALLOCATING");

        public static bool operator ==(WorkspaceInstanceProvisionState left, WorkspaceInstanceProvisionState right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceProvisionState left, WorkspaceInstanceProvisionState right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceProvisionState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceProvisionState other && Equals(other);
        public bool Equals(WorkspaceInstanceProvisionState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WorkspaceInstanceTagSpecificationResourceType : IEquatable<WorkspaceInstanceTagSpecificationResourceType>
    {
        private readonly string _value;

        private WorkspaceInstanceTagSpecificationResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkspaceInstanceTagSpecificationResourceType Instance { get; } = new WorkspaceInstanceTagSpecificationResourceType("instance");
        public static WorkspaceInstanceTagSpecificationResourceType Volume { get; } = new WorkspaceInstanceTagSpecificationResourceType("volume");
        public static WorkspaceInstanceTagSpecificationResourceType SpotInstancesRequest { get; } = new WorkspaceInstanceTagSpecificationResourceType("spot-instances-request");
        public static WorkspaceInstanceTagSpecificationResourceType NetworkInterface { get; } = new WorkspaceInstanceTagSpecificationResourceType("network-interface");

        public static bool operator ==(WorkspaceInstanceTagSpecificationResourceType left, WorkspaceInstanceTagSpecificationResourceType right) => left.Equals(right);
        public static bool operator !=(WorkspaceInstanceTagSpecificationResourceType left, WorkspaceInstanceTagSpecificationResourceType right) => !left.Equals(right);

        public static explicit operator string(WorkspaceInstanceTagSpecificationResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkspaceInstanceTagSpecificationResourceType other && Equals(other);
        public bool Equals(WorkspaceInstanceTagSpecificationResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
