// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.CodeStarConnections
{
    /// <summary>
    /// The name of the external provider where your third-party code repository is configured.
    /// </summary>
    [EnumType]
    public readonly struct RepositoryLinkProviderType : IEquatable<RepositoryLinkProviderType>
    {
        private readonly string _value;

        private RepositoryLinkProviderType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RepositoryLinkProviderType GitHub { get; } = new RepositoryLinkProviderType("GitHub");
        public static RepositoryLinkProviderType Bitbucket { get; } = new RepositoryLinkProviderType("Bitbucket");
        public static RepositoryLinkProviderType GitHubEnterprise { get; } = new RepositoryLinkProviderType("GitHubEnterprise");
        public static RepositoryLinkProviderType GitLab { get; } = new RepositoryLinkProviderType("GitLab");
        public static RepositoryLinkProviderType GitLabSelfManaged { get; } = new RepositoryLinkProviderType("GitLabSelfManaged");

        public static bool operator ==(RepositoryLinkProviderType left, RepositoryLinkProviderType right) => left.Equals(right);
        public static bool operator !=(RepositoryLinkProviderType left, RepositoryLinkProviderType right) => !left.Equals(right);

        public static explicit operator string(RepositoryLinkProviderType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RepositoryLinkProviderType other && Equals(other);
        public bool Equals(RepositoryLinkProviderType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the external provider where your third-party code repository is configured.
    /// </summary>
    [EnumType]
    public readonly struct SyncConfigurationProviderType : IEquatable<SyncConfigurationProviderType>
    {
        private readonly string _value;

        private SyncConfigurationProviderType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SyncConfigurationProviderType GitHub { get; } = new SyncConfigurationProviderType("GitHub");
        public static SyncConfigurationProviderType Bitbucket { get; } = new SyncConfigurationProviderType("Bitbucket");
        public static SyncConfigurationProviderType GitHubEnterprise { get; } = new SyncConfigurationProviderType("GitHubEnterprise");
        public static SyncConfigurationProviderType GitLab { get; } = new SyncConfigurationProviderType("GitLab");
        public static SyncConfigurationProviderType GitLabSelfManaged { get; } = new SyncConfigurationProviderType("GitLabSelfManaged");

        public static bool operator ==(SyncConfigurationProviderType left, SyncConfigurationProviderType right) => left.Equals(right);
        public static bool operator !=(SyncConfigurationProviderType left, SyncConfigurationProviderType right) => !left.Equals(right);

        public static explicit operator string(SyncConfigurationProviderType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SyncConfigurationProviderType other && Equals(other);
        public bool Equals(SyncConfigurationProviderType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Whether to enable or disable publishing of deployment status to source providers.
    /// </summary>
    [EnumType]
    public readonly struct SyncConfigurationPublishDeploymentStatus : IEquatable<SyncConfigurationPublishDeploymentStatus>
    {
        private readonly string _value;

        private SyncConfigurationPublishDeploymentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SyncConfigurationPublishDeploymentStatus Enabled { get; } = new SyncConfigurationPublishDeploymentStatus("ENABLED");
        public static SyncConfigurationPublishDeploymentStatus Disabled { get; } = new SyncConfigurationPublishDeploymentStatus("DISABLED");

        public static bool operator ==(SyncConfigurationPublishDeploymentStatus left, SyncConfigurationPublishDeploymentStatus right) => left.Equals(right);
        public static bool operator !=(SyncConfigurationPublishDeploymentStatus left, SyncConfigurationPublishDeploymentStatus right) => !left.Equals(right);

        public static explicit operator string(SyncConfigurationPublishDeploymentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SyncConfigurationPublishDeploymentStatus other && Equals(other);
        public bool Equals(SyncConfigurationPublishDeploymentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// When to trigger Git sync to begin the stack update.
    /// </summary>
    [EnumType]
    public readonly struct SyncConfigurationTriggerResourceUpdateOn : IEquatable<SyncConfigurationTriggerResourceUpdateOn>
    {
        private readonly string _value;

        private SyncConfigurationTriggerResourceUpdateOn(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SyncConfigurationTriggerResourceUpdateOn AnyChange { get; } = new SyncConfigurationTriggerResourceUpdateOn("ANY_CHANGE");
        public static SyncConfigurationTriggerResourceUpdateOn FileChange { get; } = new SyncConfigurationTriggerResourceUpdateOn("FILE_CHANGE");

        public static bool operator ==(SyncConfigurationTriggerResourceUpdateOn left, SyncConfigurationTriggerResourceUpdateOn right) => left.Equals(right);
        public static bool operator !=(SyncConfigurationTriggerResourceUpdateOn left, SyncConfigurationTriggerResourceUpdateOn right) => !left.Equals(right);

        public static explicit operator string(SyncConfigurationTriggerResourceUpdateOn value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SyncConfigurationTriggerResourceUpdateOn other && Equals(other);
        public bool Equals(SyncConfigurationTriggerResourceUpdateOn other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}