// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.AppConfig
{
    /// <summary>
    /// On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
    /// </summary>
    [EnumType]
    public readonly struct ConfigurationProfileDeletionProtectionCheck : IEquatable<ConfigurationProfileDeletionProtectionCheck>
    {
        private readonly string _value;

        private ConfigurationProfileDeletionProtectionCheck(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConfigurationProfileDeletionProtectionCheck AccountDefault { get; } = new ConfigurationProfileDeletionProtectionCheck("ACCOUNT_DEFAULT");
        public static ConfigurationProfileDeletionProtectionCheck Apply { get; } = new ConfigurationProfileDeletionProtectionCheck("APPLY");
        public static ConfigurationProfileDeletionProtectionCheck Bypass { get; } = new ConfigurationProfileDeletionProtectionCheck("BYPASS");

        public static bool operator ==(ConfigurationProfileDeletionProtectionCheck left, ConfigurationProfileDeletionProtectionCheck right) => left.Equals(right);
        public static bool operator !=(ConfigurationProfileDeletionProtectionCheck left, ConfigurationProfileDeletionProtectionCheck right) => !left.Equals(right);

        public static explicit operator string(ConfigurationProfileDeletionProtectionCheck value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConfigurationProfileDeletionProtectionCheck other && Equals(other);
        public bool Equals(ConfigurationProfileDeletionProtectionCheck other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The algorithm used to define how percentage grows over time. AWS AppConfig supports the following growth types:
    /// 
    /// Linear: For this type, AWS AppConfig processes the deployment by dividing the total number of targets by the value specified for Step percentage. For example, a linear deployment that uses a Step percentage of 10 deploys the configuration to 10 percent of the hosts. After those deployments are complete, the system deploys the configuration to the next 10 percent. This continues until 100% of the targets have successfully received the configuration.
    /// 
    /// Exponential: For this type, AWS AppConfig processes the deployment exponentially using the following formula: G*(2^N). In this formula, G is the growth factor specified by the user and N is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:
    /// 
    /// 2*(2^0)
    /// 
    /// 2*(2^1)
    /// 
    /// 2*(2^2)
    /// 
    /// Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.
    /// </summary>
    [EnumType]
    public readonly struct DeploymentStrategyGrowthType : IEquatable<DeploymentStrategyGrowthType>
    {
        private readonly string _value;

        private DeploymentStrategyGrowthType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentStrategyGrowthType Exponential { get; } = new DeploymentStrategyGrowthType("EXPONENTIAL");
        public static DeploymentStrategyGrowthType Linear { get; } = new DeploymentStrategyGrowthType("LINEAR");

        public static bool operator ==(DeploymentStrategyGrowthType left, DeploymentStrategyGrowthType right) => left.Equals(right);
        public static bool operator !=(DeploymentStrategyGrowthType left, DeploymentStrategyGrowthType right) => !left.Equals(right);

        public static explicit operator string(DeploymentStrategyGrowthType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentStrategyGrowthType other && Equals(other);
        public bool Equals(DeploymentStrategyGrowthType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Save the deployment strategy to a Systems Manager (SSM) document.
    /// </summary>
    [EnumType]
    public readonly struct DeploymentStrategyReplicateTo : IEquatable<DeploymentStrategyReplicateTo>
    {
        private readonly string _value;

        private DeploymentStrategyReplicateTo(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentStrategyReplicateTo None { get; } = new DeploymentStrategyReplicateTo("NONE");
        public static DeploymentStrategyReplicateTo SsmDocument { get; } = new DeploymentStrategyReplicateTo("SSM_DOCUMENT");

        public static bool operator ==(DeploymentStrategyReplicateTo left, DeploymentStrategyReplicateTo right) => left.Equals(right);
        public static bool operator !=(DeploymentStrategyReplicateTo left, DeploymentStrategyReplicateTo right) => !left.Equals(right);

        public static explicit operator string(DeploymentStrategyReplicateTo value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentStrategyReplicateTo other && Equals(other);
        public bool Equals(DeploymentStrategyReplicateTo other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentDeletionProtectionCheck : IEquatable<EnvironmentDeletionProtectionCheck>
    {
        private readonly string _value;

        private EnvironmentDeletionProtectionCheck(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentDeletionProtectionCheck AccountDefault { get; } = new EnvironmentDeletionProtectionCheck("ACCOUNT_DEFAULT");
        public static EnvironmentDeletionProtectionCheck Apply { get; } = new EnvironmentDeletionProtectionCheck("APPLY");
        public static EnvironmentDeletionProtectionCheck Bypass { get; } = new EnvironmentDeletionProtectionCheck("BYPASS");

        public static bool operator ==(EnvironmentDeletionProtectionCheck left, EnvironmentDeletionProtectionCheck right) => left.Equals(right);
        public static bool operator !=(EnvironmentDeletionProtectionCheck left, EnvironmentDeletionProtectionCheck right) => !left.Equals(right);

        public static explicit operator string(EnvironmentDeletionProtectionCheck value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentDeletionProtectionCheck other && Equals(other);
        public bool Equals(EnvironmentDeletionProtectionCheck other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
