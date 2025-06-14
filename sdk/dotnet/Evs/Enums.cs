// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Evs
{
    [EnumType]
    public readonly struct EnvironmentCheckResult : IEquatable<EnvironmentCheckResult>
    {
        private readonly string _value;

        private EnvironmentCheckResult(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentCheckResult Passed { get; } = new EnvironmentCheckResult("PASSED");
        public static EnvironmentCheckResult Failed { get; } = new EnvironmentCheckResult("FAILED");
        public static EnvironmentCheckResult Unknown { get; } = new EnvironmentCheckResult("UNKNOWN");

        public static bool operator ==(EnvironmentCheckResult left, EnvironmentCheckResult right) => left.Equals(right);
        public static bool operator !=(EnvironmentCheckResult left, EnvironmentCheckResult right) => !left.Equals(right);

        public static explicit operator string(EnvironmentCheckResult value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentCheckResult other && Equals(other);
        public bool Equals(EnvironmentCheckResult other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnvironmentCheckType : IEquatable<EnvironmentCheckType>
    {
        private readonly string _value;

        private EnvironmentCheckType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentCheckType KeyReuse { get; } = new EnvironmentCheckType("KEY_REUSE");
        public static EnvironmentCheckType KeyCoverage { get; } = new EnvironmentCheckType("KEY_COVERAGE");
        public static EnvironmentCheckType Reachability { get; } = new EnvironmentCheckType("REACHABILITY");
        public static EnvironmentCheckType VcfVersion { get; } = new EnvironmentCheckType("VCF_VERSION");
        public static EnvironmentCheckType HostCount { get; } = new EnvironmentCheckType("HOST_COUNT");

        public static bool operator ==(EnvironmentCheckType left, EnvironmentCheckType right) => left.Equals(right);
        public static bool operator !=(EnvironmentCheckType left, EnvironmentCheckType right) => !left.Equals(right);

        public static explicit operator string(EnvironmentCheckType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentCheckType other && Equals(other);
        public bool Equals(EnvironmentCheckType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnvironmentHostInfoForCreateInstanceType : IEquatable<EnvironmentHostInfoForCreateInstanceType>
    {
        private readonly string _value;

        private EnvironmentHostInfoForCreateInstanceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentHostInfoForCreateInstanceType I4iMetal { get; } = new EnvironmentHostInfoForCreateInstanceType("i4i.metal");

        public static bool operator ==(EnvironmentHostInfoForCreateInstanceType left, EnvironmentHostInfoForCreateInstanceType right) => left.Equals(right);
        public static bool operator !=(EnvironmentHostInfoForCreateInstanceType left, EnvironmentHostInfoForCreateInstanceType right) => !left.Equals(right);

        public static explicit operator string(EnvironmentHostInfoForCreateInstanceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentHostInfoForCreateInstanceType other && Equals(other);
        public bool Equals(EnvironmentHostInfoForCreateInstanceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnvironmentState : IEquatable<EnvironmentState>
    {
        private readonly string _value;

        private EnvironmentState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentState Creating { get; } = new EnvironmentState("CREATING");
        public static EnvironmentState Created { get; } = new EnvironmentState("CREATED");
        public static EnvironmentState Deleting { get; } = new EnvironmentState("DELETING");
        public static EnvironmentState Deleted { get; } = new EnvironmentState("DELETED");
        public static EnvironmentState CreateFailed { get; } = new EnvironmentState("CREATE_FAILED");

        public static bool operator ==(EnvironmentState left, EnvironmentState right) => left.Equals(right);
        public static bool operator !=(EnvironmentState left, EnvironmentState right) => !left.Equals(right);

        public static explicit operator string(EnvironmentState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentState other && Equals(other);
        public bool Equals(EnvironmentState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct EnvironmentVcfVersion : IEquatable<EnvironmentVcfVersion>
    {
        private readonly string _value;

        private EnvironmentVcfVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentVcfVersion Vcf521 { get; } = new EnvironmentVcfVersion("VCF-5.2.1");

        public static bool operator ==(EnvironmentVcfVersion left, EnvironmentVcfVersion right) => left.Equals(right);
        public static bool operator !=(EnvironmentVcfVersion left, EnvironmentVcfVersion right) => !left.Equals(right);

        public static explicit operator string(EnvironmentVcfVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentVcfVersion other && Equals(other);
        public bool Equals(EnvironmentVcfVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
