// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Organizations
{
    /// <summary>
    /// The method by which the account joined the organization.
    /// </summary>
    [EnumType]
    public readonly struct AccountJoinedMethod : IEquatable<AccountJoinedMethod>
    {
        private readonly string _value;

        private AccountJoinedMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccountJoinedMethod Invited { get; } = new AccountJoinedMethod("INVITED");
        public static AccountJoinedMethod Created { get; } = new AccountJoinedMethod("CREATED");

        public static bool operator ==(AccountJoinedMethod left, AccountJoinedMethod right) => left.Equals(right);
        public static bool operator !=(AccountJoinedMethod left, AccountJoinedMethod right) => !left.Equals(right);

        public static explicit operator string(AccountJoinedMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccountJoinedMethod other && Equals(other);
        public bool Equals(AccountJoinedMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the account in the organization.
    /// </summary>
    [EnumType]
    public readonly struct AccountStatus : IEquatable<AccountStatus>
    {
        private readonly string _value;

        private AccountStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccountStatus Active { get; } = new AccountStatus("ACTIVE");
        public static AccountStatus Suspended { get; } = new AccountStatus("SUSPENDED");
        public static AccountStatus PendingClosure { get; } = new AccountStatus("PENDING_CLOSURE");

        public static bool operator ==(AccountStatus left, AccountStatus right) => left.Equals(right);
        public static bool operator !=(AccountStatus left, AccountStatus right) => !left.Equals(right);

        public static explicit operator string(AccountStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccountStatus other && Equals(other);
        public bool Equals(AccountStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
    /// </summary>
    [EnumType]
    public readonly struct OrganizationFeatureSet : IEquatable<OrganizationFeatureSet>
    {
        private readonly string _value;

        private OrganizationFeatureSet(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OrganizationFeatureSet All { get; } = new OrganizationFeatureSet("ALL");
        public static OrganizationFeatureSet ConsolidatedBilling { get; } = new OrganizationFeatureSet("CONSOLIDATED_BILLING");

        public static bool operator ==(OrganizationFeatureSet left, OrganizationFeatureSet right) => left.Equals(right);
        public static bool operator !=(OrganizationFeatureSet left, OrganizationFeatureSet right) => !left.Equals(right);

        public static explicit operator string(OrganizationFeatureSet value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OrganizationFeatureSet other && Equals(other);
        public bool Equals(OrganizationFeatureSet other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY, CHATBOT_POLICY, RESOURCE_CONTROL_POLICY,DECLARATIVE_POLICY_EC2, SECURITYHUB_POLICY
    /// </summary>
    [EnumType]
    public readonly struct PolicyType : IEquatable<PolicyType>
    {
        private readonly string _value;

        private PolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyType ServiceControlPolicy { get; } = new PolicyType("SERVICE_CONTROL_POLICY");
        public static PolicyType AiservicesOptOutPolicy { get; } = new PolicyType("AISERVICES_OPT_OUT_POLICY");
        public static PolicyType BackupPolicy { get; } = new PolicyType("BACKUP_POLICY");
        public static PolicyType TagPolicy { get; } = new PolicyType("TAG_POLICY");
        public static PolicyType ChatbotPolicy { get; } = new PolicyType("CHATBOT_POLICY");
        public static PolicyType ResourceControlPolicy { get; } = new PolicyType("RESOURCE_CONTROL_POLICY");
        public static PolicyType DeclarativePolicyEc2 { get; } = new PolicyType("DECLARATIVE_POLICY_EC2");
        public static PolicyType SecurityhubPolicy { get; } = new PolicyType("SECURITYHUB_POLICY");

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
