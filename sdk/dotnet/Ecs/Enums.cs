// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Ecs
{
    [EnumType]
    public readonly struct CapacityProviderAutoScalingGroupProviderManagedDraining : IEquatable<CapacityProviderAutoScalingGroupProviderManagedDraining>
    {
        private readonly string _value;

        private CapacityProviderAutoScalingGroupProviderManagedDraining(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CapacityProviderAutoScalingGroupProviderManagedDraining Disabled { get; } = new CapacityProviderAutoScalingGroupProviderManagedDraining("DISABLED");
        public static CapacityProviderAutoScalingGroupProviderManagedDraining Enabled { get; } = new CapacityProviderAutoScalingGroupProviderManagedDraining("ENABLED");

        public static bool operator ==(CapacityProviderAutoScalingGroupProviderManagedDraining left, CapacityProviderAutoScalingGroupProviderManagedDraining right) => left.Equals(right);
        public static bool operator !=(CapacityProviderAutoScalingGroupProviderManagedDraining left, CapacityProviderAutoScalingGroupProviderManagedDraining right) => !left.Equals(right);

        public static explicit operator string(CapacityProviderAutoScalingGroupProviderManagedDraining value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CapacityProviderAutoScalingGroupProviderManagedDraining other && Equals(other);
        public bool Equals(CapacityProviderAutoScalingGroupProviderManagedDraining other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CapacityProviderAutoScalingGroupProviderManagedTerminationProtection : IEquatable<CapacityProviderAutoScalingGroupProviderManagedTerminationProtection>
    {
        private readonly string _value;

        private CapacityProviderAutoScalingGroupProviderManagedTerminationProtection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CapacityProviderAutoScalingGroupProviderManagedTerminationProtection Disabled { get; } = new CapacityProviderAutoScalingGroupProviderManagedTerminationProtection("DISABLED");
        public static CapacityProviderAutoScalingGroupProviderManagedTerminationProtection Enabled { get; } = new CapacityProviderAutoScalingGroupProviderManagedTerminationProtection("ENABLED");

        public static bool operator ==(CapacityProviderAutoScalingGroupProviderManagedTerminationProtection left, CapacityProviderAutoScalingGroupProviderManagedTerminationProtection right) => left.Equals(right);
        public static bool operator !=(CapacityProviderAutoScalingGroupProviderManagedTerminationProtection left, CapacityProviderAutoScalingGroupProviderManagedTerminationProtection right) => !left.Equals(right);

        public static explicit operator string(CapacityProviderAutoScalingGroupProviderManagedTerminationProtection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CapacityProviderAutoScalingGroupProviderManagedTerminationProtection other && Equals(other);
        public bool Equals(CapacityProviderAutoScalingGroupProviderManagedTerminationProtection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CapacityProviderManagedScalingStatus : IEquatable<CapacityProviderManagedScalingStatus>
    {
        private readonly string _value;

        private CapacityProviderManagedScalingStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CapacityProviderManagedScalingStatus Disabled { get; } = new CapacityProviderManagedScalingStatus("DISABLED");
        public static CapacityProviderManagedScalingStatus Enabled { get; } = new CapacityProviderManagedScalingStatus("ENABLED");

        public static bool operator ==(CapacityProviderManagedScalingStatus left, CapacityProviderManagedScalingStatus right) => left.Equals(right);
        public static bool operator !=(CapacityProviderManagedScalingStatus left, CapacityProviderManagedScalingStatus right) => !left.Equals(right);

        public static explicit operator string(CapacityProviderManagedScalingStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CapacityProviderManagedScalingStatus other && Equals(other);
        public bool Equals(CapacityProviderManagedScalingStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.
    /// </summary>
    [EnumType]
    public readonly struct ClusterCapacityProviderAssociationsCapacityProvider : IEquatable<ClusterCapacityProviderAssociationsCapacityProvider>
    {
        private readonly string _value;

        private ClusterCapacityProviderAssociationsCapacityProvider(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterCapacityProviderAssociationsCapacityProvider Fargate { get; } = new ClusterCapacityProviderAssociationsCapacityProvider("FARGATE");
        public static ClusterCapacityProviderAssociationsCapacityProvider FargateSpot { get; } = new ClusterCapacityProviderAssociationsCapacityProvider("FARGATE_SPOT");

        public static bool operator ==(ClusterCapacityProviderAssociationsCapacityProvider left, ClusterCapacityProviderAssociationsCapacityProvider right) => left.Equals(right);
        public static bool operator !=(ClusterCapacityProviderAssociationsCapacityProvider left, ClusterCapacityProviderAssociationsCapacityProvider right) => !left.Equals(right);

        public static explicit operator string(ClusterCapacityProviderAssociationsCapacityProvider value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterCapacityProviderAssociationsCapacityProvider other && Equals(other);
        public bool Equals(ClusterCapacityProviderAssociationsCapacityProvider other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.
    /// </summary>
    [EnumType]
    public readonly struct ClusterCapacityProviderAssociationsCapacityProvider0 : IEquatable<ClusterCapacityProviderAssociationsCapacityProvider0>
    {
        private readonly string _value;

        private ClusterCapacityProviderAssociationsCapacityProvider0(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterCapacityProviderAssociationsCapacityProvider0 Fargate { get; } = new ClusterCapacityProviderAssociationsCapacityProvider0("FARGATE");
        public static ClusterCapacityProviderAssociationsCapacityProvider0 FargateSpot { get; } = new ClusterCapacityProviderAssociationsCapacityProvider0("FARGATE_SPOT");

        public static bool operator ==(ClusterCapacityProviderAssociationsCapacityProvider0 left, ClusterCapacityProviderAssociationsCapacityProvider0 right) => left.Equals(right);
        public static bool operator !=(ClusterCapacityProviderAssociationsCapacityProvider0 left, ClusterCapacityProviderAssociationsCapacityProvider0 right) => !left.Equals(right);

        public static explicit operator string(ClusterCapacityProviderAssociationsCapacityProvider0 value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterCapacityProviderAssociationsCapacityProvider0 other && Equals(other);
        public bool Equals(ClusterCapacityProviderAssociationsCapacityProvider0 other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceAwsVpcConfigurationAssignPublicIp : IEquatable<ServiceAwsVpcConfigurationAssignPublicIp>
    {
        private readonly string _value;

        private ServiceAwsVpcConfigurationAssignPublicIp(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceAwsVpcConfigurationAssignPublicIp Disabled { get; } = new ServiceAwsVpcConfigurationAssignPublicIp("DISABLED");
        public static ServiceAwsVpcConfigurationAssignPublicIp Enabled { get; } = new ServiceAwsVpcConfigurationAssignPublicIp("ENABLED");

        public static bool operator ==(ServiceAwsVpcConfigurationAssignPublicIp left, ServiceAwsVpcConfigurationAssignPublicIp right) => left.Equals(right);
        public static bool operator !=(ServiceAwsVpcConfigurationAssignPublicIp left, ServiceAwsVpcConfigurationAssignPublicIp right) => !left.Equals(right);

        public static explicit operator string(ServiceAwsVpcConfigurationAssignPublicIp value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceAwsVpcConfigurationAssignPublicIp other && Equals(other);
        public bool Equals(ServiceAwsVpcConfigurationAssignPublicIp other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceDeploymentControllerType : IEquatable<ServiceDeploymentControllerType>
    {
        private readonly string _value;

        private ServiceDeploymentControllerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceDeploymentControllerType CodeDeploy { get; } = new ServiceDeploymentControllerType("CODE_DEPLOY");
        public static ServiceDeploymentControllerType Ecs { get; } = new ServiceDeploymentControllerType("ECS");
        public static ServiceDeploymentControllerType External { get; } = new ServiceDeploymentControllerType("EXTERNAL");

        public static bool operator ==(ServiceDeploymentControllerType left, ServiceDeploymentControllerType right) => left.Equals(right);
        public static bool operator !=(ServiceDeploymentControllerType left, ServiceDeploymentControllerType right) => !left.Equals(right);

        public static explicit operator string(ServiceDeploymentControllerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceDeploymentControllerType other && Equals(other);
        public bool Equals(ServiceDeploymentControllerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceEbsTagSpecificationPropagateTags : IEquatable<ServiceEbsTagSpecificationPropagateTags>
    {
        private readonly string _value;

        private ServiceEbsTagSpecificationPropagateTags(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceEbsTagSpecificationPropagateTags Service { get; } = new ServiceEbsTagSpecificationPropagateTags("SERVICE");
        public static ServiceEbsTagSpecificationPropagateTags TaskDefinition { get; } = new ServiceEbsTagSpecificationPropagateTags("TASK_DEFINITION");

        public static bool operator ==(ServiceEbsTagSpecificationPropagateTags left, ServiceEbsTagSpecificationPropagateTags right) => left.Equals(right);
        public static bool operator !=(ServiceEbsTagSpecificationPropagateTags left, ServiceEbsTagSpecificationPropagateTags right) => !left.Equals(right);

        public static explicit operator string(ServiceEbsTagSpecificationPropagateTags value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceEbsTagSpecificationPropagateTags other && Equals(other);
        public bool Equals(ServiceEbsTagSpecificationPropagateTags other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceLaunchType : IEquatable<ServiceLaunchType>
    {
        private readonly string _value;

        private ServiceLaunchType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceLaunchType Ec2 { get; } = new ServiceLaunchType("EC2");
        public static ServiceLaunchType Fargate { get; } = new ServiceLaunchType("FARGATE");
        public static ServiceLaunchType External { get; } = new ServiceLaunchType("EXTERNAL");

        public static bool operator ==(ServiceLaunchType left, ServiceLaunchType right) => left.Equals(right);
        public static bool operator !=(ServiceLaunchType left, ServiceLaunchType right) => !left.Equals(right);

        public static explicit operator string(ServiceLaunchType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLaunchType other && Equals(other);
        public bool Equals(ServiceLaunchType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServicePlacementConstraintType : IEquatable<ServicePlacementConstraintType>
    {
        private readonly string _value;

        private ServicePlacementConstraintType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServicePlacementConstraintType DistinctInstance { get; } = new ServicePlacementConstraintType("distinctInstance");
        public static ServicePlacementConstraintType MemberOf { get; } = new ServicePlacementConstraintType("memberOf");

        public static bool operator ==(ServicePlacementConstraintType left, ServicePlacementConstraintType right) => left.Equals(right);
        public static bool operator !=(ServicePlacementConstraintType left, ServicePlacementConstraintType right) => !left.Equals(right);

        public static explicit operator string(ServicePlacementConstraintType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServicePlacementConstraintType other && Equals(other);
        public bool Equals(ServicePlacementConstraintType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServicePlacementStrategyType : IEquatable<ServicePlacementStrategyType>
    {
        private readonly string _value;

        private ServicePlacementStrategyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServicePlacementStrategyType Binpack { get; } = new ServicePlacementStrategyType("binpack");
        public static ServicePlacementStrategyType Random { get; } = new ServicePlacementStrategyType("random");
        public static ServicePlacementStrategyType Spread { get; } = new ServicePlacementStrategyType("spread");

        public static bool operator ==(ServicePlacementStrategyType left, ServicePlacementStrategyType right) => left.Equals(right);
        public static bool operator !=(ServicePlacementStrategyType left, ServicePlacementStrategyType right) => !left.Equals(right);

        public static explicit operator string(ServicePlacementStrategyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServicePlacementStrategyType other && Equals(other);
        public bool Equals(ServicePlacementStrategyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServicePropagateTags : IEquatable<ServicePropagateTags>
    {
        private readonly string _value;

        private ServicePropagateTags(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServicePropagateTags Service { get; } = new ServicePropagateTags("SERVICE");
        public static ServicePropagateTags TaskDefinition { get; } = new ServicePropagateTags("TASK_DEFINITION");

        public static bool operator ==(ServicePropagateTags left, ServicePropagateTags right) => left.Equals(right);
        public static bool operator !=(ServicePropagateTags left, ServicePropagateTags right) => !left.Equals(right);

        public static explicit operator string(ServicePropagateTags value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServicePropagateTags other && Equals(other);
        public bool Equals(ServicePropagateTags other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceSchedulingStrategy : IEquatable<ServiceSchedulingStrategy>
    {
        private readonly string _value;

        private ServiceSchedulingStrategy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceSchedulingStrategy Daemon { get; } = new ServiceSchedulingStrategy("DAEMON");
        public static ServiceSchedulingStrategy Replica { get; } = new ServiceSchedulingStrategy("REPLICA");

        public static bool operator ==(ServiceSchedulingStrategy left, ServiceSchedulingStrategy right) => left.Equals(right);
        public static bool operator !=(ServiceSchedulingStrategy left, ServiceSchedulingStrategy right) => !left.Equals(right);

        public static explicit operator string(ServiceSchedulingStrategy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceSchedulingStrategy other && Equals(other);
        public bool Equals(ServiceSchedulingStrategy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TaskDefinitionAuthorizationConfigIam : IEquatable<TaskDefinitionAuthorizationConfigIam>
    {
        private readonly string _value;

        private TaskDefinitionAuthorizationConfigIam(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskDefinitionAuthorizationConfigIam Enabled { get; } = new TaskDefinitionAuthorizationConfigIam("ENABLED");
        public static TaskDefinitionAuthorizationConfigIam Disabled { get; } = new TaskDefinitionAuthorizationConfigIam("DISABLED");

        public static bool operator ==(TaskDefinitionAuthorizationConfigIam left, TaskDefinitionAuthorizationConfigIam right) => left.Equals(right);
        public static bool operator !=(TaskDefinitionAuthorizationConfigIam left, TaskDefinitionAuthorizationConfigIam right) => !left.Equals(right);

        public static explicit operator string(TaskDefinitionAuthorizationConfigIam value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskDefinitionAuthorizationConfigIam other && Equals(other);
        public bool Equals(TaskDefinitionAuthorizationConfigIam other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TaskDefinitionEfsVolumeConfigurationTransitEncryption : IEquatable<TaskDefinitionEfsVolumeConfigurationTransitEncryption>
    {
        private readonly string _value;

        private TaskDefinitionEfsVolumeConfigurationTransitEncryption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskDefinitionEfsVolumeConfigurationTransitEncryption Enabled { get; } = new TaskDefinitionEfsVolumeConfigurationTransitEncryption("ENABLED");
        public static TaskDefinitionEfsVolumeConfigurationTransitEncryption Disabled { get; } = new TaskDefinitionEfsVolumeConfigurationTransitEncryption("DISABLED");

        public static bool operator ==(TaskDefinitionEfsVolumeConfigurationTransitEncryption left, TaskDefinitionEfsVolumeConfigurationTransitEncryption right) => left.Equals(right);
        public static bool operator !=(TaskDefinitionEfsVolumeConfigurationTransitEncryption left, TaskDefinitionEfsVolumeConfigurationTransitEncryption right) => !left.Equals(right);

        public static explicit operator string(TaskDefinitionEfsVolumeConfigurationTransitEncryption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskDefinitionEfsVolumeConfigurationTransitEncryption other && Equals(other);
        public bool Equals(TaskDefinitionEfsVolumeConfigurationTransitEncryption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TaskDefinitionPortMappingAppProtocol : IEquatable<TaskDefinitionPortMappingAppProtocol>
    {
        private readonly string _value;

        private TaskDefinitionPortMappingAppProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskDefinitionPortMappingAppProtocol Http { get; } = new TaskDefinitionPortMappingAppProtocol("http");
        public static TaskDefinitionPortMappingAppProtocol Http2 { get; } = new TaskDefinitionPortMappingAppProtocol("http2");
        public static TaskDefinitionPortMappingAppProtocol Grpc { get; } = new TaskDefinitionPortMappingAppProtocol("grpc");

        public static bool operator ==(TaskDefinitionPortMappingAppProtocol left, TaskDefinitionPortMappingAppProtocol right) => left.Equals(right);
        public static bool operator !=(TaskDefinitionPortMappingAppProtocol left, TaskDefinitionPortMappingAppProtocol right) => !left.Equals(right);

        public static explicit operator string(TaskDefinitionPortMappingAppProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskDefinitionPortMappingAppProtocol other && Equals(other);
        public bool Equals(TaskDefinitionPortMappingAppProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.
    /// </summary>
    [EnumType]
    public readonly struct TaskSetAwsVpcConfigurationAssignPublicIp : IEquatable<TaskSetAwsVpcConfigurationAssignPublicIp>
    {
        private readonly string _value;

        private TaskSetAwsVpcConfigurationAssignPublicIp(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskSetAwsVpcConfigurationAssignPublicIp Disabled { get; } = new TaskSetAwsVpcConfigurationAssignPublicIp("DISABLED");
        public static TaskSetAwsVpcConfigurationAssignPublicIp Enabled { get; } = new TaskSetAwsVpcConfigurationAssignPublicIp("ENABLED");

        public static bool operator ==(TaskSetAwsVpcConfigurationAssignPublicIp left, TaskSetAwsVpcConfigurationAssignPublicIp right) => left.Equals(right);
        public static bool operator !=(TaskSetAwsVpcConfigurationAssignPublicIp left, TaskSetAwsVpcConfigurationAssignPublicIp right) => !left.Equals(right);

        public static explicit operator string(TaskSetAwsVpcConfigurationAssignPublicIp value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskSetAwsVpcConfigurationAssignPublicIp other && Equals(other);
        public bool Equals(TaskSetAwsVpcConfigurationAssignPublicIp other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. 
    /// </summary>
    [EnumType]
    public readonly struct TaskSetLaunchType : IEquatable<TaskSetLaunchType>
    {
        private readonly string _value;

        private TaskSetLaunchType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskSetLaunchType Ec2 { get; } = new TaskSetLaunchType("EC2");
        public static TaskSetLaunchType Fargate { get; } = new TaskSetLaunchType("FARGATE");

        public static bool operator ==(TaskSetLaunchType left, TaskSetLaunchType right) => left.Equals(right);
        public static bool operator !=(TaskSetLaunchType left, TaskSetLaunchType right) => !left.Equals(right);

        public static explicit operator string(TaskSetLaunchType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskSetLaunchType other && Equals(other);
        public bool Equals(TaskSetLaunchType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The unit of measure for the scale value.
    /// </summary>
    [EnumType]
    public readonly struct TaskSetScaleUnit : IEquatable<TaskSetScaleUnit>
    {
        private readonly string _value;

        private TaskSetScaleUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskSetScaleUnit Percent { get; } = new TaskSetScaleUnit("PERCENT");

        public static bool operator ==(TaskSetScaleUnit left, TaskSetScaleUnit right) => left.Equals(right);
        public static bool operator !=(TaskSetScaleUnit left, TaskSetScaleUnit right) => !left.Equals(right);

        public static explicit operator string(TaskSetScaleUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskSetScaleUnit other && Equals(other);
        public bool Equals(TaskSetScaleUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}