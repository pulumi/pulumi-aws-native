// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.IoTFleetWise
{
    [EnumType]
    public readonly struct CampaignCompression : IEquatable<CampaignCompression>
    {
        private readonly string _value;

        private CampaignCompression(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignCompression Off { get; } = new CampaignCompression("OFF");
        public static CampaignCompression Snappy { get; } = new CampaignCompression("SNAPPY");

        public static bool operator ==(CampaignCompression left, CampaignCompression right) => left.Equals(right);
        public static bool operator !=(CampaignCompression left, CampaignCompression right) => !left.Equals(right);

        public static explicit operator string(CampaignCompression value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignCompression other && Equals(other);
        public bool Equals(CampaignCompression other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CampaignDiagnosticsMode : IEquatable<CampaignDiagnosticsMode>
    {
        private readonly string _value;

        private CampaignDiagnosticsMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignDiagnosticsMode Off { get; } = new CampaignDiagnosticsMode("OFF");
        public static CampaignDiagnosticsMode SendActiveDtcs { get; } = new CampaignDiagnosticsMode("SEND_ACTIVE_DTCS");

        public static bool operator ==(CampaignDiagnosticsMode left, CampaignDiagnosticsMode right) => left.Equals(right);
        public static bool operator !=(CampaignDiagnosticsMode left, CampaignDiagnosticsMode right) => !left.Equals(right);

        public static explicit operator string(CampaignDiagnosticsMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignDiagnosticsMode other && Equals(other);
        public bool Equals(CampaignDiagnosticsMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CampaignSpoolingMode : IEquatable<CampaignSpoolingMode>
    {
        private readonly string _value;

        private CampaignSpoolingMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignSpoolingMode Off { get; } = new CampaignSpoolingMode("OFF");
        public static CampaignSpoolingMode ToDisk { get; } = new CampaignSpoolingMode("TO_DISK");

        public static bool operator ==(CampaignSpoolingMode left, CampaignSpoolingMode right) => left.Equals(right);
        public static bool operator !=(CampaignSpoolingMode left, CampaignSpoolingMode right) => !left.Equals(right);

        public static explicit operator string(CampaignSpoolingMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignSpoolingMode other && Equals(other);
        public bool Equals(CampaignSpoolingMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CampaignStatus : IEquatable<CampaignStatus>
    {
        private readonly string _value;

        private CampaignStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignStatus Creating { get; } = new CampaignStatus("CREATING");
        public static CampaignStatus WaitingForApproval { get; } = new CampaignStatus("WAITING_FOR_APPROVAL");
        public static CampaignStatus Running { get; } = new CampaignStatus("RUNNING");
        public static CampaignStatus Suspended { get; } = new CampaignStatus("SUSPENDED");

        public static bool operator ==(CampaignStatus left, CampaignStatus right) => left.Equals(right);
        public static bool operator !=(CampaignStatus left, CampaignStatus right) => !left.Equals(right);

        public static explicit operator string(CampaignStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignStatus other && Equals(other);
        public bool Equals(CampaignStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CampaignUpdateCampaignAction : IEquatable<CampaignUpdateCampaignAction>
    {
        private readonly string _value;

        private CampaignUpdateCampaignAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CampaignUpdateCampaignAction Approve { get; } = new CampaignUpdateCampaignAction("APPROVE");
        public static CampaignUpdateCampaignAction Suspend { get; } = new CampaignUpdateCampaignAction("SUSPEND");
        public static CampaignUpdateCampaignAction Resume { get; } = new CampaignUpdateCampaignAction("RESUME");
        public static CampaignUpdateCampaignAction Update { get; } = new CampaignUpdateCampaignAction("UPDATE");

        public static bool operator ==(CampaignUpdateCampaignAction left, CampaignUpdateCampaignAction right) => left.Equals(right);
        public static bool operator !=(CampaignUpdateCampaignAction left, CampaignUpdateCampaignAction right) => !left.Equals(right);

        public static explicit operator string(CampaignUpdateCampaignAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CampaignUpdateCampaignAction other && Equals(other);
        public bool Equals(CampaignUpdateCampaignAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DecoderManifestCanNetworkInterfaceType : IEquatable<DecoderManifestCanNetworkInterfaceType>
    {
        private readonly string _value;

        private DecoderManifestCanNetworkInterfaceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DecoderManifestCanNetworkInterfaceType CanInterface { get; } = new DecoderManifestCanNetworkInterfaceType("CAN_INTERFACE");

        public static bool operator ==(DecoderManifestCanNetworkInterfaceType left, DecoderManifestCanNetworkInterfaceType right) => left.Equals(right);
        public static bool operator !=(DecoderManifestCanNetworkInterfaceType left, DecoderManifestCanNetworkInterfaceType right) => !left.Equals(right);

        public static explicit operator string(DecoderManifestCanNetworkInterfaceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DecoderManifestCanNetworkInterfaceType other && Equals(other);
        public bool Equals(DecoderManifestCanNetworkInterfaceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DecoderManifestCanSignalDecoderType : IEquatable<DecoderManifestCanSignalDecoderType>
    {
        private readonly string _value;

        private DecoderManifestCanSignalDecoderType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DecoderManifestCanSignalDecoderType CanSignal { get; } = new DecoderManifestCanSignalDecoderType("CAN_SIGNAL");

        public static bool operator ==(DecoderManifestCanSignalDecoderType left, DecoderManifestCanSignalDecoderType right) => left.Equals(right);
        public static bool operator !=(DecoderManifestCanSignalDecoderType left, DecoderManifestCanSignalDecoderType right) => !left.Equals(right);

        public static explicit operator string(DecoderManifestCanSignalDecoderType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DecoderManifestCanSignalDecoderType other && Equals(other);
        public bool Equals(DecoderManifestCanSignalDecoderType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DecoderManifestManifestStatus : IEquatable<DecoderManifestManifestStatus>
    {
        private readonly string _value;

        private DecoderManifestManifestStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DecoderManifestManifestStatus Active { get; } = new DecoderManifestManifestStatus("ACTIVE");
        public static DecoderManifestManifestStatus Draft { get; } = new DecoderManifestManifestStatus("DRAFT");

        public static bool operator ==(DecoderManifestManifestStatus left, DecoderManifestManifestStatus right) => left.Equals(right);
        public static bool operator !=(DecoderManifestManifestStatus left, DecoderManifestManifestStatus right) => !left.Equals(right);

        public static explicit operator string(DecoderManifestManifestStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DecoderManifestManifestStatus other && Equals(other);
        public bool Equals(DecoderManifestManifestStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DecoderManifestObdNetworkInterfaceType : IEquatable<DecoderManifestObdNetworkInterfaceType>
    {
        private readonly string _value;

        private DecoderManifestObdNetworkInterfaceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DecoderManifestObdNetworkInterfaceType ObdInterface { get; } = new DecoderManifestObdNetworkInterfaceType("OBD_INTERFACE");

        public static bool operator ==(DecoderManifestObdNetworkInterfaceType left, DecoderManifestObdNetworkInterfaceType right) => left.Equals(right);
        public static bool operator !=(DecoderManifestObdNetworkInterfaceType left, DecoderManifestObdNetworkInterfaceType right) => !left.Equals(right);

        public static explicit operator string(DecoderManifestObdNetworkInterfaceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DecoderManifestObdNetworkInterfaceType other && Equals(other);
        public bool Equals(DecoderManifestObdNetworkInterfaceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DecoderManifestObdSignalDecoderType : IEquatable<DecoderManifestObdSignalDecoderType>
    {
        private readonly string _value;

        private DecoderManifestObdSignalDecoderType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DecoderManifestObdSignalDecoderType ObdSignal { get; } = new DecoderManifestObdSignalDecoderType("OBD_SIGNAL");

        public static bool operator ==(DecoderManifestObdSignalDecoderType left, DecoderManifestObdSignalDecoderType right) => left.Equals(right);
        public static bool operator !=(DecoderManifestObdSignalDecoderType left, DecoderManifestObdSignalDecoderType right) => !left.Equals(right);

        public static explicit operator string(DecoderManifestObdSignalDecoderType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DecoderManifestObdSignalDecoderType other && Equals(other);
        public bool Equals(DecoderManifestObdSignalDecoderType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ModelManifestManifestStatus : IEquatable<ModelManifestManifestStatus>
    {
        private readonly string _value;

        private ModelManifestManifestStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ModelManifestManifestStatus Active { get; } = new ModelManifestManifestStatus("ACTIVE");
        public static ModelManifestManifestStatus Draft { get; } = new ModelManifestManifestStatus("DRAFT");

        public static bool operator ==(ModelManifestManifestStatus left, ModelManifestManifestStatus right) => left.Equals(right);
        public static bool operator !=(ModelManifestManifestStatus left, ModelManifestManifestStatus right) => !left.Equals(right);

        public static explicit operator string(ModelManifestManifestStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ModelManifestManifestStatus other && Equals(other);
        public bool Equals(ModelManifestManifestStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct VehicleAssociationBehavior : IEquatable<VehicleAssociationBehavior>
    {
        private readonly string _value;

        private VehicleAssociationBehavior(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VehicleAssociationBehavior CreateIotThing { get; } = new VehicleAssociationBehavior("CreateIotThing");
        public static VehicleAssociationBehavior ValidateIotThingExists { get; } = new VehicleAssociationBehavior("ValidateIotThingExists");

        public static bool operator ==(VehicleAssociationBehavior left, VehicleAssociationBehavior right) => left.Equals(right);
        public static bool operator !=(VehicleAssociationBehavior left, VehicleAssociationBehavior right) => !left.Equals(right);

        public static explicit operator string(VehicleAssociationBehavior value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VehicleAssociationBehavior other && Equals(other);
        public bool Equals(VehicleAssociationBehavior other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}