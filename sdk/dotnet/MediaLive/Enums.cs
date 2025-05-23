// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.MediaLive
{
    /// <summary>
    /// The current state of the ChannelPlacementGroupState
    /// </summary>
    [EnumType]
    public readonly struct ChannelPlacementGroupState : IEquatable<ChannelPlacementGroupState>
    {
        private readonly string _value;

        private ChannelPlacementGroupState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ChannelPlacementGroupState Unassigned { get; } = new ChannelPlacementGroupState("UNASSIGNED");
        public static ChannelPlacementGroupState Assigning { get; } = new ChannelPlacementGroupState("ASSIGNING");
        public static ChannelPlacementGroupState Assigned { get; } = new ChannelPlacementGroupState("ASSIGNED");
        public static ChannelPlacementGroupState Deleting { get; } = new ChannelPlacementGroupState("DELETING");
        public static ChannelPlacementGroupState Deleted { get; } = new ChannelPlacementGroupState("DELETED");
        public static ChannelPlacementGroupState Unassigning { get; } = new ChannelPlacementGroupState("UNASSIGNING");

        public static bool operator ==(ChannelPlacementGroupState left, ChannelPlacementGroupState right) => left.Equals(right);
        public static bool operator !=(ChannelPlacementGroupState left, ChannelPlacementGroupState right) => !left.Equals(right);

        public static explicit operator string(ChannelPlacementGroupState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ChannelPlacementGroupState other && Equals(other);
        public bool Equals(ChannelPlacementGroupState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The comparison operator used to compare the specified statistic and the threshold.
    /// </summary>
    [EnumType]
    public readonly struct CloudWatchAlarmTemplateComparisonOperator : IEquatable<CloudWatchAlarmTemplateComparisonOperator>
    {
        private readonly string _value;

        private CloudWatchAlarmTemplateComparisonOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudWatchAlarmTemplateComparisonOperator GreaterThanOrEqualToThreshold { get; } = new CloudWatchAlarmTemplateComparisonOperator("GreaterThanOrEqualToThreshold");
        public static CloudWatchAlarmTemplateComparisonOperator GreaterThanThreshold { get; } = new CloudWatchAlarmTemplateComparisonOperator("GreaterThanThreshold");
        public static CloudWatchAlarmTemplateComparisonOperator LessThanThreshold { get; } = new CloudWatchAlarmTemplateComparisonOperator("LessThanThreshold");
        public static CloudWatchAlarmTemplateComparisonOperator LessThanOrEqualToThreshold { get; } = new CloudWatchAlarmTemplateComparisonOperator("LessThanOrEqualToThreshold");

        public static bool operator ==(CloudWatchAlarmTemplateComparisonOperator left, CloudWatchAlarmTemplateComparisonOperator right) => left.Equals(right);
        public static bool operator !=(CloudWatchAlarmTemplateComparisonOperator left, CloudWatchAlarmTemplateComparisonOperator right) => !left.Equals(right);

        public static explicit operator string(CloudWatchAlarmTemplateComparisonOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudWatchAlarmTemplateComparisonOperator other && Equals(other);
        public bool Equals(CloudWatchAlarmTemplateComparisonOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The statistic to apply to the alarm's metric data.
    /// </summary>
    [EnumType]
    public readonly struct CloudWatchAlarmTemplateStatistic : IEquatable<CloudWatchAlarmTemplateStatistic>
    {
        private readonly string _value;

        private CloudWatchAlarmTemplateStatistic(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudWatchAlarmTemplateStatistic SampleCount { get; } = new CloudWatchAlarmTemplateStatistic("SampleCount");
        public static CloudWatchAlarmTemplateStatistic Average { get; } = new CloudWatchAlarmTemplateStatistic("Average");
        public static CloudWatchAlarmTemplateStatistic Sum { get; } = new CloudWatchAlarmTemplateStatistic("Sum");
        public static CloudWatchAlarmTemplateStatistic Minimum { get; } = new CloudWatchAlarmTemplateStatistic("Minimum");
        public static CloudWatchAlarmTemplateStatistic Maximum { get; } = new CloudWatchAlarmTemplateStatistic("Maximum");

        public static bool operator ==(CloudWatchAlarmTemplateStatistic left, CloudWatchAlarmTemplateStatistic right) => left.Equals(right);
        public static bool operator !=(CloudWatchAlarmTemplateStatistic left, CloudWatchAlarmTemplateStatistic right) => !left.Equals(right);

        public static explicit operator string(CloudWatchAlarmTemplateStatistic value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudWatchAlarmTemplateStatistic other && Equals(other);
        public bool Equals(CloudWatchAlarmTemplateStatistic other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The resource type this template should dynamically generate cloudwatch metric alarms for.
    /// </summary>
    [EnumType]
    public readonly struct CloudWatchAlarmTemplateTargetResourceType : IEquatable<CloudWatchAlarmTemplateTargetResourceType>
    {
        private readonly string _value;

        private CloudWatchAlarmTemplateTargetResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudWatchAlarmTemplateTargetResourceType CloudfrontDistribution { get; } = new CloudWatchAlarmTemplateTargetResourceType("CLOUDFRONT_DISTRIBUTION");
        public static CloudWatchAlarmTemplateTargetResourceType MedialiveMultiplex { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIALIVE_MULTIPLEX");
        public static CloudWatchAlarmTemplateTargetResourceType MedialiveChannel { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIALIVE_CHANNEL");
        public static CloudWatchAlarmTemplateTargetResourceType MedialiveInputDevice { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIALIVE_INPUT_DEVICE");
        public static CloudWatchAlarmTemplateTargetResourceType MediapackageChannel { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIAPACKAGE_CHANNEL");
        public static CloudWatchAlarmTemplateTargetResourceType MediapackageOriginEndpoint { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIAPACKAGE_ORIGIN_ENDPOINT");
        public static CloudWatchAlarmTemplateTargetResourceType MediaconnectFlow { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIACONNECT_FLOW");
        public static CloudWatchAlarmTemplateTargetResourceType MediatailorPlaybackConfiguration { get; } = new CloudWatchAlarmTemplateTargetResourceType("MEDIATAILOR_PLAYBACK_CONFIGURATION");
        public static CloudWatchAlarmTemplateTargetResourceType S3Bucket { get; } = new CloudWatchAlarmTemplateTargetResourceType("S3_BUCKET");

        public static bool operator ==(CloudWatchAlarmTemplateTargetResourceType left, CloudWatchAlarmTemplateTargetResourceType right) => left.Equals(right);
        public static bool operator !=(CloudWatchAlarmTemplateTargetResourceType left, CloudWatchAlarmTemplateTargetResourceType right) => !left.Equals(right);

        public static explicit operator string(CloudWatchAlarmTemplateTargetResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudWatchAlarmTemplateTargetResourceType other && Equals(other);
        public bool Equals(CloudWatchAlarmTemplateTargetResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies how missing data points are treated when evaluating the alarm's condition.
    /// </summary>
    [EnumType]
    public readonly struct CloudWatchAlarmTemplateTreatMissingData : IEquatable<CloudWatchAlarmTemplateTreatMissingData>
    {
        private readonly string _value;

        private CloudWatchAlarmTemplateTreatMissingData(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudWatchAlarmTemplateTreatMissingData NotBreaching { get; } = new CloudWatchAlarmTemplateTreatMissingData("notBreaching");
        public static CloudWatchAlarmTemplateTreatMissingData Breaching { get; } = new CloudWatchAlarmTemplateTreatMissingData("breaching");
        public static CloudWatchAlarmTemplateTreatMissingData Ignore { get; } = new CloudWatchAlarmTemplateTreatMissingData("ignore");
        public static CloudWatchAlarmTemplateTreatMissingData Missing { get; } = new CloudWatchAlarmTemplateTreatMissingData("missing");

        public static bool operator ==(CloudWatchAlarmTemplateTreatMissingData left, CloudWatchAlarmTemplateTreatMissingData right) => left.Equals(right);
        public static bool operator !=(CloudWatchAlarmTemplateTreatMissingData left, CloudWatchAlarmTemplateTreatMissingData right) => !left.Equals(right);

        public static explicit operator string(CloudWatchAlarmTemplateTreatMissingData value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudWatchAlarmTemplateTreatMissingData other && Equals(other);
        public bool Equals(CloudWatchAlarmTemplateTreatMissingData other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the Cluster.
    /// </summary>
    [EnumType]
    public readonly struct ClusterState : IEquatable<ClusterState>
    {
        private readonly string _value;

        private ClusterState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterState Creating { get; } = new ClusterState("CREATING");
        public static ClusterState CreateFailed { get; } = new ClusterState("CREATE_FAILED");
        public static ClusterState Active { get; } = new ClusterState("ACTIVE");
        public static ClusterState Deleting { get; } = new ClusterState("DELETING");
        public static ClusterState Deleted { get; } = new ClusterState("DELETED");

        public static bool operator ==(ClusterState left, ClusterState right) => left.Equals(right);
        public static bool operator !=(ClusterState left, ClusterState right) => !left.Equals(right);

        public static explicit operator string(ClusterState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterState other && Equals(other);
        public bool Equals(ClusterState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The hardware type for the cluster.
    /// </summary>
    [EnumType]
    public readonly struct ClusterType : IEquatable<ClusterType>
    {
        private readonly string _value;

        private ClusterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterType OnPremises { get; } = new ClusterType("ON_PREMISES");
        public static ClusterType OutpostsRack { get; } = new ClusterType("OUTPOSTS_RACK");
        public static ClusterType OutpostsServer { get; } = new ClusterType("OUTPOSTS_SERVER");
        public static ClusterType Ec2 { get; } = new ClusterType("EC2");

        public static bool operator ==(ClusterType left, ClusterType right) => left.Equals(right);
        public static bool operator !=(ClusterType left, ClusterType right) => !left.Equals(right);

        public static explicit operator string(ClusterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterType other && Equals(other);
        public bool Equals(ClusterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of event to match with the rule.
    /// </summary>
    [EnumType]
    public readonly struct EventBridgeRuleTemplateEventType : IEquatable<EventBridgeRuleTemplateEventType>
    {
        private readonly string _value;

        private EventBridgeRuleTemplateEventType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventBridgeRuleTemplateEventType MedialiveMultiplexAlert { get; } = new EventBridgeRuleTemplateEventType("MEDIALIVE_MULTIPLEX_ALERT");
        public static EventBridgeRuleTemplateEventType MedialiveMultiplexStateChange { get; } = new EventBridgeRuleTemplateEventType("MEDIALIVE_MULTIPLEX_STATE_CHANGE");
        public static EventBridgeRuleTemplateEventType MedialiveChannelAlert { get; } = new EventBridgeRuleTemplateEventType("MEDIALIVE_CHANNEL_ALERT");
        public static EventBridgeRuleTemplateEventType MedialiveChannelInputChange { get; } = new EventBridgeRuleTemplateEventType("MEDIALIVE_CHANNEL_INPUT_CHANGE");
        public static EventBridgeRuleTemplateEventType MedialiveChannelStateChange { get; } = new EventBridgeRuleTemplateEventType("MEDIALIVE_CHANNEL_STATE_CHANGE");
        public static EventBridgeRuleTemplateEventType MediapackageInputNotification { get; } = new EventBridgeRuleTemplateEventType("MEDIAPACKAGE_INPUT_NOTIFICATION");
        public static EventBridgeRuleTemplateEventType MediapackageKeyProviderNotification { get; } = new EventBridgeRuleTemplateEventType("MEDIAPACKAGE_KEY_PROVIDER_NOTIFICATION");
        public static EventBridgeRuleTemplateEventType MediapackageHarvestJobNotification { get; } = new EventBridgeRuleTemplateEventType("MEDIAPACKAGE_HARVEST_JOB_NOTIFICATION");
        public static EventBridgeRuleTemplateEventType SignalMapActiveAlarm { get; } = new EventBridgeRuleTemplateEventType("SIGNAL_MAP_ACTIVE_ALARM");
        public static EventBridgeRuleTemplateEventType MediaconnectAlert { get; } = new EventBridgeRuleTemplateEventType("MEDIACONNECT_ALERT");
        public static EventBridgeRuleTemplateEventType MediaconnectSourceHealth { get; } = new EventBridgeRuleTemplateEventType("MEDIACONNECT_SOURCE_HEALTH");
        public static EventBridgeRuleTemplateEventType MediaconnectOutputHealth { get; } = new EventBridgeRuleTemplateEventType("MEDIACONNECT_OUTPUT_HEALTH");
        public static EventBridgeRuleTemplateEventType MediaconnectFlowStatusChange { get; } = new EventBridgeRuleTemplateEventType("MEDIACONNECT_FLOW_STATUS_CHANGE");

        public static bool operator ==(EventBridgeRuleTemplateEventType left, EventBridgeRuleTemplateEventType right) => left.Equals(right);
        public static bool operator !=(EventBridgeRuleTemplateEventType left, EventBridgeRuleTemplateEventType right) => !left.Equals(right);

        public static explicit operator string(EventBridgeRuleTemplateEventType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventBridgeRuleTemplateEventType other && Equals(other);
        public bool Equals(EventBridgeRuleTemplateEventType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the multiplex.
    /// </summary>
    [EnumType]
    public readonly struct MultiplexState : IEquatable<MultiplexState>
    {
        private readonly string _value;

        private MultiplexState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MultiplexState Creating { get; } = new MultiplexState("CREATING");
        public static MultiplexState CreateFailed { get; } = new MultiplexState("CREATE_FAILED");
        public static MultiplexState Idle { get; } = new MultiplexState("IDLE");
        public static MultiplexState Starting { get; } = new MultiplexState("STARTING");
        public static MultiplexState Running { get; } = new MultiplexState("RUNNING");
        public static MultiplexState Recovering { get; } = new MultiplexState("RECOVERING");
        public static MultiplexState Stopping { get; } = new MultiplexState("STOPPING");
        public static MultiplexState Deleting { get; } = new MultiplexState("DELETING");
        public static MultiplexState Deleted { get; } = new MultiplexState("DELETED");

        public static bool operator ==(MultiplexState left, MultiplexState right) => left.Equals(right);
        public static bool operator !=(MultiplexState left, MultiplexState right) => !left.Equals(right);

        public static explicit operator string(MultiplexState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MultiplexState other && Equals(other);
        public bool Equals(MultiplexState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates which pipeline is preferred by the multiplex for program ingest.
    /// If set to \"PIPELINE_0\" or \"PIPELINE_1\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,
    /// it will switch back once that ingest is healthy again. If set to \"CURRENTLY_ACTIVE\",
    /// it will not switch back to the other pipeline based on it recovering to a healthy state,
    /// it will only switch if the active pipeline becomes unhealthy.
    /// </summary>
    [EnumType]
    public readonly struct MultiplexprogramPreferredChannelPipeline : IEquatable<MultiplexprogramPreferredChannelPipeline>
    {
        private readonly string _value;

        private MultiplexprogramPreferredChannelPipeline(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MultiplexprogramPreferredChannelPipeline CurrentlyActive { get; } = new MultiplexprogramPreferredChannelPipeline("CURRENTLY_ACTIVE");
        public static MultiplexprogramPreferredChannelPipeline Pipeline0 { get; } = new MultiplexprogramPreferredChannelPipeline("PIPELINE_0");
        public static MultiplexprogramPreferredChannelPipeline Pipeline1 { get; } = new MultiplexprogramPreferredChannelPipeline("PIPELINE_1");

        public static bool operator ==(MultiplexprogramPreferredChannelPipeline left, MultiplexprogramPreferredChannelPipeline right) => left.Equals(right);
        public static bool operator !=(MultiplexprogramPreferredChannelPipeline left, MultiplexprogramPreferredChannelPipeline right) => !left.Equals(right);

        public static explicit operator string(MultiplexprogramPreferredChannelPipeline value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MultiplexprogramPreferredChannelPipeline other && Equals(other);
        public bool Equals(MultiplexprogramPreferredChannelPipeline other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NetworkState : IEquatable<NetworkState>
    {
        private readonly string _value;

        private NetworkState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NetworkState Creating { get; } = new NetworkState("CREATING");
        public static NetworkState CreateFailed { get; } = new NetworkState("CREATE_FAILED");
        public static NetworkState Active { get; } = new NetworkState("ACTIVE");
        public static NetworkState Deleting { get; } = new NetworkState("DELETING");
        public static NetworkState Idle { get; } = new NetworkState("IDLE");
        public static NetworkState InUse { get; } = new NetworkState("IN_USE");
        public static NetworkState Updating { get; } = new NetworkState("UPDATING");
        public static NetworkState Deleted { get; } = new NetworkState("DELETED");
        public static NetworkState DeleteFailed { get; } = new NetworkState("DELETE_FAILED");

        public static bool operator ==(NetworkState left, NetworkState right) => left.Equals(right);
        public static bool operator !=(NetworkState left, NetworkState right) => !left.Equals(right);

        public static explicit operator string(NetworkState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkState other && Equals(other);
        public bool Equals(NetworkState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the SdiSource.
    /// </summary>
    [EnumType]
    public readonly struct SdiSourceMode : IEquatable<SdiSourceMode>
    {
        private readonly string _value;

        private SdiSourceMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SdiSourceMode Quadrant { get; } = new SdiSourceMode("QUADRANT");
        public static SdiSourceMode Interleave { get; } = new SdiSourceMode("INTERLEAVE");

        public static bool operator ==(SdiSourceMode left, SdiSourceMode right) => left.Equals(right);
        public static bool operator !=(SdiSourceMode left, SdiSourceMode right) => !left.Equals(right);

        public static explicit operator string(SdiSourceMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SdiSourceMode other && Equals(other);
        public bool Equals(SdiSourceMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The current state of the SdiSource.
    /// </summary>
    [EnumType]
    public readonly struct SdiSourceState : IEquatable<SdiSourceState>
    {
        private readonly string _value;

        private SdiSourceState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SdiSourceState Idle { get; } = new SdiSourceState("IDLE");
        public static SdiSourceState InUse { get; } = new SdiSourceState("IN_USE");
        public static SdiSourceState Deleted { get; } = new SdiSourceState("DELETED");

        public static bool operator ==(SdiSourceState left, SdiSourceState right) => left.Equals(right);
        public static bool operator !=(SdiSourceState left, SdiSourceState right) => !left.Equals(right);

        public static explicit operator string(SdiSourceState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SdiSourceState other && Equals(other);
        public bool Equals(SdiSourceState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The interface mode of the SdiSource.
    /// </summary>
    [EnumType]
    public readonly struct SdiSourceType : IEquatable<SdiSourceType>
    {
        private readonly string _value;

        private SdiSourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SdiSourceType Single { get; } = new SdiSourceType("SINGLE");
        public static SdiSourceType Quad { get; } = new SdiSourceType("QUAD");

        public static bool operator ==(SdiSourceType left, SdiSourceType right) => left.Equals(right);
        public static bool operator !=(SdiSourceType left, SdiSourceType right) => !left.Equals(right);

        public static explicit operator string(SdiSourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SdiSourceType other && Equals(other);
        public bool Equals(SdiSourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A signal map's monitor deployment status.
    /// </summary>
    [EnumType]
    public readonly struct SignalMapMonitorDeploymentStatus : IEquatable<SignalMapMonitorDeploymentStatus>
    {
        private readonly string _value;

        private SignalMapMonitorDeploymentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SignalMapMonitorDeploymentStatus NotDeployed { get; } = new SignalMapMonitorDeploymentStatus("NOT_DEPLOYED");
        public static SignalMapMonitorDeploymentStatus DryRunDeploymentComplete { get; } = new SignalMapMonitorDeploymentStatus("DRY_RUN_DEPLOYMENT_COMPLETE");
        public static SignalMapMonitorDeploymentStatus DryRunDeploymentFailed { get; } = new SignalMapMonitorDeploymentStatus("DRY_RUN_DEPLOYMENT_FAILED");
        public static SignalMapMonitorDeploymentStatus DryRunDeploymentInProgress { get; } = new SignalMapMonitorDeploymentStatus("DRY_RUN_DEPLOYMENT_IN_PROGRESS");
        public static SignalMapMonitorDeploymentStatus DeploymentComplete { get; } = new SignalMapMonitorDeploymentStatus("DEPLOYMENT_COMPLETE");
        public static SignalMapMonitorDeploymentStatus DeploymentFailed { get; } = new SignalMapMonitorDeploymentStatus("DEPLOYMENT_FAILED");
        public static SignalMapMonitorDeploymentStatus DeploymentInProgress { get; } = new SignalMapMonitorDeploymentStatus("DEPLOYMENT_IN_PROGRESS");
        public static SignalMapMonitorDeploymentStatus DeleteComplete { get; } = new SignalMapMonitorDeploymentStatus("DELETE_COMPLETE");
        public static SignalMapMonitorDeploymentStatus DeleteFailed { get; } = new SignalMapMonitorDeploymentStatus("DELETE_FAILED");
        public static SignalMapMonitorDeploymentStatus DeleteInProgress { get; } = new SignalMapMonitorDeploymentStatus("DELETE_IN_PROGRESS");

        public static bool operator ==(SignalMapMonitorDeploymentStatus left, SignalMapMonitorDeploymentStatus right) => left.Equals(right);
        public static bool operator !=(SignalMapMonitorDeploymentStatus left, SignalMapMonitorDeploymentStatus right) => !left.Equals(right);

        public static explicit operator string(SignalMapMonitorDeploymentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SignalMapMonitorDeploymentStatus other && Equals(other);
        public bool Equals(SignalMapMonitorDeploymentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A signal map's current status which is dependent on its lifecycle actions or associated jobs.
    /// </summary>
    [EnumType]
    public readonly struct SignalMapStatus : IEquatable<SignalMapStatus>
    {
        private readonly string _value;

        private SignalMapStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SignalMapStatus CreateInProgress { get; } = new SignalMapStatus("CREATE_IN_PROGRESS");
        public static SignalMapStatus CreateComplete { get; } = new SignalMapStatus("CREATE_COMPLETE");
        public static SignalMapStatus CreateFailed { get; } = new SignalMapStatus("CREATE_FAILED");
        public static SignalMapStatus UpdateInProgress { get; } = new SignalMapStatus("UPDATE_IN_PROGRESS");
        public static SignalMapStatus UpdateComplete { get; } = new SignalMapStatus("UPDATE_COMPLETE");
        public static SignalMapStatus UpdateReverted { get; } = new SignalMapStatus("UPDATE_REVERTED");
        public static SignalMapStatus UpdateFailed { get; } = new SignalMapStatus("UPDATE_FAILED");
        public static SignalMapStatus Ready { get; } = new SignalMapStatus("READY");
        public static SignalMapStatus NotReady { get; } = new SignalMapStatus("NOT_READY");

        public static bool operator ==(SignalMapStatus left, SignalMapStatus right) => left.Equals(right);
        public static bool operator !=(SignalMapStatus left, SignalMapStatus right) => !left.Equals(right);

        public static explicit operator string(SignalMapStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SignalMapStatus other && Equals(other);
        public bool Equals(SignalMapStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
