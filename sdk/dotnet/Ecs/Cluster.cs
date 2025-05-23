// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs
{
    /// <summary>
    /// The ``AWS::ECS::Cluster`` resource creates an Amazon Elastic Container Service (Amazon ECS) cluster.
    /// </summary>
    [AwsNativeResourceType("aws-native:ecs:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon ECS cluster, such as `arn:aws:ecs:us-east-2:123456789012:cluster/MyECSCluster` .
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions.
        ///  If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation.
        ///  To use a FARGATElong capacity provider, specify either the ``FARGATE`` or ``FARGATE_SPOT`` capacity providers. The FARGATElong capacity providers are available to all accounts and only need to be associated with a cluster to be used.
        ///  The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.
        /// </summary>
        [Output("capacityProviders")]
        public Output<ImmutableArray<string>> CapacityProviders { get; private set; } = null!;

        /// <summary>
        /// A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name.
        /// </summary>
        [Output("clusterName")]
        public Output<string?> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights with enhanced observability or CloudWatch Container Insights for a cluster.
        ///  Container Insights with enhanced observability provides all the Container Insights metrics, plus additional task and container metrics. This version supports enhanced observability for Amazon ECS clusters using the Amazon EC2 and Fargate launch types. After you configure Container Insights with enhanced observability on Amazon ECS, Container Insights auto-collects detailed infrastructure telemetry from the cluster level down to the container level in your environment and displays these critical performance data in curated dashboards removing the heavy lifting in observability set-up. 
        ///  For more information, see [Monitor Amazon ECS containers using Container Insights with enhanced observability](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html) in the *Amazon Elastic Container Service Developer Guide*.
        /// </summary>
        [Output("clusterSettings")]
        public Output<ImmutableArray<Outputs.ClusterSettings>> ClusterSettings { get; private set; } = null!;

        /// <summary>
        /// The execute command and managed storage configuration for the cluster.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.ClusterConfiguration?> Configuration { get; private set; } = null!;

        /// <summary>
        /// The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
        /// </summary>
        [Output("defaultCapacityProviderStrategy")]
        public Output<ImmutableArray<Outputs.ClusterCapacityProviderStrategyItem>> DefaultCapacityProviderStrategy { get; private set; } = null!;

        /// <summary>
        /// Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter.
        ///  Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.
        /// </summary>
        [Output("serviceConnectDefaults")]
        public Output<Outputs.ClusterServiceConnectDefaults?> ServiceConnectDefaults { get; private set; } = null!;

        /// <summary>
        /// The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both.
        ///  The following basic restrictions apply to tags:
        ///   +  Maximum number of tags per resource - 50
        ///   +  For each resource, each tag key must be unique, and each tag key can have only one value.
        ///   +  Maximum key length - 128 Unicode characters in UTF-8
        ///   +  Maximum value length - 256 Unicode characters in UTF-8
        ///   +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
        ///   +  Tag keys and values are case-sensitive.
        ///   +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ecs:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecs:Cluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "clusterName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("capacityProviders")]
        private InputList<string>? _capacityProviders;

        /// <summary>
        /// The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions.
        ///  If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation.
        ///  To use a FARGATElong capacity provider, specify either the ``FARGATE`` or ``FARGATE_SPOT`` capacity providers. The FARGATElong capacity providers are available to all accounts and only need to be associated with a cluster to be used.
        ///  The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.
        /// </summary>
        public InputList<string> CapacityProviders
        {
            get => _capacityProviders ?? (_capacityProviders = new InputList<string>());
            set => _capacityProviders = value;
        }

        /// <summary>
        /// A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        [Input("clusterSettings")]
        private InputList<Inputs.ClusterSettingsArgs>? _clusterSettings;

        /// <summary>
        /// The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights with enhanced observability or CloudWatch Container Insights for a cluster.
        ///  Container Insights with enhanced observability provides all the Container Insights metrics, plus additional task and container metrics. This version supports enhanced observability for Amazon ECS clusters using the Amazon EC2 and Fargate launch types. After you configure Container Insights with enhanced observability on Amazon ECS, Container Insights auto-collects detailed infrastructure telemetry from the cluster level down to the container level in your environment and displays these critical performance data in curated dashboards removing the heavy lifting in observability set-up. 
        ///  For more information, see [Monitor Amazon ECS containers using Container Insights with enhanced observability](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html) in the *Amazon Elastic Container Service Developer Guide*.
        /// </summary>
        public InputList<Inputs.ClusterSettingsArgs> ClusterSettings
        {
            get => _clusterSettings ?? (_clusterSettings = new InputList<Inputs.ClusterSettingsArgs>());
            set => _clusterSettings = value;
        }

        /// <summary>
        /// The execute command and managed storage configuration for the cluster.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ClusterConfigurationArgs>? Configuration { get; set; }

        [Input("defaultCapacityProviderStrategy")]
        private InputList<Inputs.ClusterCapacityProviderStrategyItemArgs>? _defaultCapacityProviderStrategy;

        /// <summary>
        /// The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
        /// </summary>
        public InputList<Inputs.ClusterCapacityProviderStrategyItemArgs> DefaultCapacityProviderStrategy
        {
            get => _defaultCapacityProviderStrategy ?? (_defaultCapacityProviderStrategy = new InputList<Inputs.ClusterCapacityProviderStrategyItemArgs>());
            set => _defaultCapacityProviderStrategy = value;
        }

        /// <summary>
        /// Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter.
        ///  Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.
        /// </summary>
        [Input("serviceConnectDefaults")]
        public Input<Inputs.ClusterServiceConnectDefaultsArgs>? ServiceConnectDefaults { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both.
        ///  The following basic restrictions apply to tags:
        ///   +  Maximum number of tags per resource - 50
        ///   +  For each resource, each tag key must be unique, and each tag key can have only one value.
        ///   +  Maximum key length - 128 Unicode characters in UTF-8
        ///   +  Maximum value length - 256 Unicode characters in UTF-8
        ///   +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
        ///   +  Tag keys and values are case-sensitive.
        ///   +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }
}
