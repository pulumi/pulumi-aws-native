// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Fms
{
    /// <summary>
    /// Firewall deployment mode.
    /// </summary>
    [EnumType]
    public readonly struct PolicyFirewallDeploymentModel : IEquatable<PolicyFirewallDeploymentModel>
    {
        private readonly string _value;

        private PolicyFirewallDeploymentModel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyFirewallDeploymentModel Distributed { get; } = new PolicyFirewallDeploymentModel("DISTRIBUTED");
        public static PolicyFirewallDeploymentModel Centralized { get; } = new PolicyFirewallDeploymentModel("CENTRALIZED");

        public static bool operator ==(PolicyFirewallDeploymentModel left, PolicyFirewallDeploymentModel right) => left.Equals(right);
        public static bool operator !=(PolicyFirewallDeploymentModel left, PolicyFirewallDeploymentModel right) => !left.Equals(right);

        public static explicit operator string(PolicyFirewallDeploymentModel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyFirewallDeploymentModel other && Equals(other);
        public bool Equals(PolicyFirewallDeploymentModel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Rule Action.
    /// </summary>
    [EnumType]
    public readonly struct PolicyNetworkAclEntryRuleAction : IEquatable<PolicyNetworkAclEntryRuleAction>
    {
        private readonly string _value;

        private PolicyNetworkAclEntryRuleAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyNetworkAclEntryRuleAction Allow { get; } = new PolicyNetworkAclEntryRuleAction("allow");
        public static PolicyNetworkAclEntryRuleAction Deny { get; } = new PolicyNetworkAclEntryRuleAction("deny");

        public static bool operator ==(PolicyNetworkAclEntryRuleAction left, PolicyNetworkAclEntryRuleAction right) => left.Equals(right);
        public static bool operator !=(PolicyNetworkAclEntryRuleAction left, PolicyNetworkAclEntryRuleAction right) => !left.Equals(right);

        public static explicit operator string(PolicyNetworkAclEntryRuleAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyNetworkAclEntryRuleAction other && Equals(other);
        public bool Equals(PolicyNetworkAclEntryRuleAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies whether to combine multiple resource tags with AND, so that a resource must have all tags to be included or excluded, or OR, so that a resource must have at least one tag.
    /// 
    /// Default: `AND`
    /// </summary>
    [EnumType]
    public readonly struct PolicyResourceTagLogicalOperator : IEquatable<PolicyResourceTagLogicalOperator>
    {
        private readonly string _value;

        private PolicyResourceTagLogicalOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyResourceTagLogicalOperator And { get; } = new PolicyResourceTagLogicalOperator("AND");
        public static PolicyResourceTagLogicalOperator Or { get; } = new PolicyResourceTagLogicalOperator("OR");

        public static bool operator ==(PolicyResourceTagLogicalOperator left, PolicyResourceTagLogicalOperator right) => left.Equals(right);
        public static bool operator !=(PolicyResourceTagLogicalOperator left, PolicyResourceTagLogicalOperator right) => !left.Equals(right);

        public static explicit operator string(PolicyResourceTagLogicalOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyResourceTagLogicalOperator other && Equals(other);
        public bool Equals(PolicyResourceTagLogicalOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Firewall policy type.
    /// </summary>
    [EnumType]
    public readonly struct PolicyType : IEquatable<PolicyType>
    {
        private readonly string _value;

        private PolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyType Waf { get; } = new PolicyType("WAF");
        public static PolicyType Wafv2 { get; } = new PolicyType("WAFV2");
        public static PolicyType ShieldAdvanced { get; } = new PolicyType("SHIELD_ADVANCED");
        public static PolicyType SecurityGroupsCommon { get; } = new PolicyType("SECURITY_GROUPS_COMMON");
        public static PolicyType SecurityGroupsContentAudit { get; } = new PolicyType("SECURITY_GROUPS_CONTENT_AUDIT");
        public static PolicyType SecurityGroupsUsageAudit { get; } = new PolicyType("SECURITY_GROUPS_USAGE_AUDIT");
        public static PolicyType NetworkFirewall { get; } = new PolicyType("NETWORK_FIREWALL");
        public static PolicyType ThirdPartyFirewall { get; } = new PolicyType("THIRD_PARTY_FIREWALL");
        public static PolicyType DnsFirewall { get; } = new PolicyType("DNS_FIREWALL");
        public static PolicyType ImportNetworkFirewall { get; } = new PolicyType("IMPORT_NETWORK_FIREWALL");
        public static PolicyType NetworkAclCommon { get; } = new PolicyType("NETWORK_ACL_COMMON");

        public static bool operator ==(PolicyType left, PolicyType right) => left.Equals(right);
        public static bool operator !=(PolicyType left, PolicyType right) => !left.Equals(right);

        public static explicit operator string(PolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyType other && Equals(other);
        public bool Equals(PolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
