// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.IoTTwinMaker
{
    /// <summary>
    /// The underlying type of the data type.
    /// </summary>
    [EnumType]
    public readonly struct ComponentTypeDataTypeType : IEquatable<ComponentTypeDataTypeType>
    {
        private readonly string _value;

        private ComponentTypeDataTypeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentTypeDataTypeType Relationship { get; } = new ComponentTypeDataTypeType("RELATIONSHIP");
        public static ComponentTypeDataTypeType String { get; } = new ComponentTypeDataTypeType("STRING");
        public static ComponentTypeDataTypeType Long { get; } = new ComponentTypeDataTypeType("LONG");
        public static ComponentTypeDataTypeType Boolean { get; } = new ComponentTypeDataTypeType("BOOLEAN");
        public static ComponentTypeDataTypeType Integer { get; } = new ComponentTypeDataTypeType("INTEGER");
        public static ComponentTypeDataTypeType Double { get; } = new ComponentTypeDataTypeType("DOUBLE");
        public static ComponentTypeDataTypeType List { get; } = new ComponentTypeDataTypeType("LIST");
        public static ComponentTypeDataTypeType Map { get; } = new ComponentTypeDataTypeType("MAP");

        public static bool operator ==(ComponentTypeDataTypeType left, ComponentTypeDataTypeType right) => left.Equals(right);
        public static bool operator !=(ComponentTypeDataTypeType left, ComponentTypeDataTypeType right) => !left.Equals(right);

        public static explicit operator string(ComponentTypeDataTypeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentTypeDataTypeType other && Equals(other);
        public bool Equals(ComponentTypeDataTypeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The scope of the function.
    /// </summary>
    [EnumType]
    public readonly struct ComponentTypeFunctionScope : IEquatable<ComponentTypeFunctionScope>
    {
        private readonly string _value;

        private ComponentTypeFunctionScope(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentTypeFunctionScope Entity { get; } = new ComponentTypeFunctionScope("ENTITY");
        public static ComponentTypeFunctionScope Workspace { get; } = new ComponentTypeFunctionScope("WORKSPACE");

        public static bool operator ==(ComponentTypeFunctionScope left, ComponentTypeFunctionScope right) => left.Equals(right);
        public static bool operator !=(ComponentTypeFunctionScope left, ComponentTypeFunctionScope right) => !left.Equals(right);

        public static explicit operator string(ComponentTypeFunctionScope value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentTypeFunctionScope other && Equals(other);
        public bool Equals(ComponentTypeFunctionScope other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of property group.
    /// </summary>
    [EnumType]
    public readonly struct ComponentTypePropertyGroupGroupType : IEquatable<ComponentTypePropertyGroupGroupType>
    {
        private readonly string _value;

        private ComponentTypePropertyGroupGroupType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentTypePropertyGroupGroupType Tabular { get; } = new ComponentTypePropertyGroupGroupType("TABULAR");

        public static bool operator ==(ComponentTypePropertyGroupGroupType left, ComponentTypePropertyGroupGroupType right) => left.Equals(right);
        public static bool operator !=(ComponentTypePropertyGroupGroupType left, ComponentTypePropertyGroupGroupType right) => !left.Equals(right);

        public static explicit operator string(ComponentTypePropertyGroupGroupType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentTypePropertyGroupGroupType other && Equals(other);
        public bool Equals(ComponentTypePropertyGroupGroupType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComponentTypeStatusError1PropertiesCode : IEquatable<ComponentTypeStatusError1PropertiesCode>
    {
        private readonly string _value;

        private ComponentTypeStatusError1PropertiesCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentTypeStatusError1PropertiesCode ValidationError { get; } = new ComponentTypeStatusError1PropertiesCode("VALIDATION_ERROR");
        public static ComponentTypeStatusError1PropertiesCode InternalFailure { get; } = new ComponentTypeStatusError1PropertiesCode("INTERNAL_FAILURE");

        public static bool operator ==(ComponentTypeStatusError1PropertiesCode left, ComponentTypeStatusError1PropertiesCode right) => left.Equals(right);
        public static bool operator !=(ComponentTypeStatusError1PropertiesCode left, ComponentTypeStatusError1PropertiesCode right) => !left.Equals(right);

        public static explicit operator string(ComponentTypeStatusError1PropertiesCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentTypeStatusError1PropertiesCode other && Equals(other);
        public bool Equals(ComponentTypeStatusError1PropertiesCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComponentTypeStatusErrorPropertiesCode : IEquatable<ComponentTypeStatusErrorPropertiesCode>
    {
        private readonly string _value;

        private ComponentTypeStatusErrorPropertiesCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentTypeStatusErrorPropertiesCode ValidationError { get; } = new ComponentTypeStatusErrorPropertiesCode("VALIDATION_ERROR");
        public static ComponentTypeStatusErrorPropertiesCode InternalFailure { get; } = new ComponentTypeStatusErrorPropertiesCode("INTERNAL_FAILURE");

        public static bool operator ==(ComponentTypeStatusErrorPropertiesCode left, ComponentTypeStatusErrorPropertiesCode right) => left.Equals(right);
        public static bool operator !=(ComponentTypeStatusErrorPropertiesCode left, ComponentTypeStatusErrorPropertiesCode right) => !left.Equals(right);

        public static explicit operator string(ComponentTypeStatusErrorPropertiesCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentTypeStatusErrorPropertiesCode other && Equals(other);
        public bool Equals(ComponentTypeStatusErrorPropertiesCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The component type status state.
    /// </summary>
    [EnumType]
    public readonly struct ComponentTypeStatusState : IEquatable<ComponentTypeStatusState>
    {
        private readonly string _value;

        private ComponentTypeStatusState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentTypeStatusState Creating { get; } = new ComponentTypeStatusState("CREATING");
        public static ComponentTypeStatusState Updating { get; } = new ComponentTypeStatusState("UPDATING");
        public static ComponentTypeStatusState Deleting { get; } = new ComponentTypeStatusState("DELETING");
        public static ComponentTypeStatusState Active { get; } = new ComponentTypeStatusState("ACTIVE");
        public static ComponentTypeStatusState Error { get; } = new ComponentTypeStatusState("ERROR");

        public static bool operator ==(ComponentTypeStatusState left, ComponentTypeStatusState right) => left.Equals(right);
        public static bool operator !=(ComponentTypeStatusState left, ComponentTypeStatusState right) => !left.Equals(right);

        public static explicit operator string(ComponentTypeStatusState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentTypeStatusState other && Equals(other);
        public bool Equals(ComponentTypeStatusState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The underlying type of the data type.
    /// </summary>
    [EnumType]
    public readonly struct EntityDataTypeType : IEquatable<EntityDataTypeType>
    {
        private readonly string _value;

        private EntityDataTypeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EntityDataTypeType Relationship { get; } = new EntityDataTypeType("RELATIONSHIP");
        public static EntityDataTypeType String { get; } = new EntityDataTypeType("STRING");
        public static EntityDataTypeType Long { get; } = new EntityDataTypeType("LONG");
        public static EntityDataTypeType Boolean { get; } = new EntityDataTypeType("BOOLEAN");
        public static EntityDataTypeType Integer { get; } = new EntityDataTypeType("INTEGER");
        public static EntityDataTypeType Double { get; } = new EntityDataTypeType("DOUBLE");
        public static EntityDataTypeType List { get; } = new EntityDataTypeType("LIST");
        public static EntityDataTypeType Map { get; } = new EntityDataTypeType("MAP");

        public static bool operator ==(EntityDataTypeType left, EntityDataTypeType right) => left.Equals(right);
        public static bool operator !=(EntityDataTypeType left, EntityDataTypeType right) => !left.Equals(right);

        public static explicit operator string(EntityDataTypeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityDataTypeType other && Equals(other);
        public bool Equals(EntityDataTypeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of property group.
    /// </summary>
    [EnumType]
    public readonly struct EntityPropertyGroupGroupType : IEquatable<EntityPropertyGroupGroupType>
    {
        private readonly string _value;

        private EntityPropertyGroupGroupType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EntityPropertyGroupGroupType Tabular { get; } = new EntityPropertyGroupGroupType("TABULAR");

        public static bool operator ==(EntityPropertyGroupGroupType left, EntityPropertyGroupGroupType right) => left.Equals(right);
        public static bool operator !=(EntityPropertyGroupGroupType left, EntityPropertyGroupGroupType right) => !left.Equals(right);

        public static explicit operator string(EntityPropertyGroupGroupType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityPropertyGroupGroupType other && Equals(other);
        public bool Equals(EntityPropertyGroupGroupType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EntityStatusError1PropertiesCode : IEquatable<EntityStatusError1PropertiesCode>
    {
        private readonly string _value;

        private EntityStatusError1PropertiesCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EntityStatusError1PropertiesCode ValidationError { get; } = new EntityStatusError1PropertiesCode("VALIDATION_ERROR");
        public static EntityStatusError1PropertiesCode InternalFailure { get; } = new EntityStatusError1PropertiesCode("INTERNAL_FAILURE");

        public static bool operator ==(EntityStatusError1PropertiesCode left, EntityStatusError1PropertiesCode right) => left.Equals(right);
        public static bool operator !=(EntityStatusError1PropertiesCode left, EntityStatusError1PropertiesCode right) => !left.Equals(right);

        public static explicit operator string(EntityStatusError1PropertiesCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityStatusError1PropertiesCode other && Equals(other);
        public bool Equals(EntityStatusError1PropertiesCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EntityStatusErrorPropertiesCode : IEquatable<EntityStatusErrorPropertiesCode>
    {
        private readonly string _value;

        private EntityStatusErrorPropertiesCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EntityStatusErrorPropertiesCode ValidationError { get; } = new EntityStatusErrorPropertiesCode("VALIDATION_ERROR");
        public static EntityStatusErrorPropertiesCode InternalFailure { get; } = new EntityStatusErrorPropertiesCode("INTERNAL_FAILURE");

        public static bool operator ==(EntityStatusErrorPropertiesCode left, EntityStatusErrorPropertiesCode right) => left.Equals(right);
        public static bool operator !=(EntityStatusErrorPropertiesCode left, EntityStatusErrorPropertiesCode right) => !left.Equals(right);

        public static explicit operator string(EntityStatusErrorPropertiesCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityStatusErrorPropertiesCode other && Equals(other);
        public bool Equals(EntityStatusErrorPropertiesCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EntityStatusState : IEquatable<EntityStatusState>
    {
        private readonly string _value;

        private EntityStatusState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EntityStatusState Creating { get; } = new EntityStatusState("CREATING");
        public static EntityStatusState Updating { get; } = new EntityStatusState("UPDATING");
        public static EntityStatusState Deleting { get; } = new EntityStatusState("DELETING");
        public static EntityStatusState Active { get; } = new EntityStatusState("ACTIVE");
        public static EntityStatusState Error { get; } = new EntityStatusState("ERROR");

        public static bool operator ==(EntityStatusState left, EntityStatusState right) => left.Equals(right);
        public static bool operator !=(EntityStatusState left, EntityStatusState right) => !left.Equals(right);

        public static explicit operator string(EntityStatusState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityStatusState other && Equals(other);
        public bool Equals(EntityStatusState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
