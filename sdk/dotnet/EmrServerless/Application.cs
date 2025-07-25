// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EmrServerless
{
    /// <summary>
    /// Resource schema for AWS::EMRServerless::Application Type
    /// </summary>
    [AwsNativeResourceType("aws-native:emrserverless:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the EMR Serverless Application.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The CPU architecture of an application.
        /// </summary>
        [Output("architecture")]
        public Output<Pulumi.AwsNative.EmrServerless.ApplicationArchitecture?> Architecture { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the EMR Serverless Application.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration for Auto Start of Application.
        /// </summary>
        [Output("autoStartConfiguration")]
        public Output<Outputs.ApplicationAutoStartConfiguration?> AutoStartConfiguration { get; private set; } = null!;

        /// <summary>
        /// Configuration for Auto Stop of Application.
        /// </summary>
        [Output("autoStopConfiguration")]
        public Output<Outputs.ApplicationAutoStopConfiguration?> AutoStopConfiguration { get; private set; } = null!;

        /// <summary>
        /// The IAM IdentityCenter configuration for trusted-identity-propagation on this application. Supported with release labels emr-7.8.0 and above.
        /// </summary>
        [Output("identityCenterConfiguration")]
        public Output<Outputs.ApplicationIdentityCenterConfiguration?> IdentityCenterConfiguration { get; private set; } = null!;

        /// <summary>
        /// The image configuration applied to all worker types.
        /// </summary>
        [Output("imageConfiguration")]
        public Output<Outputs.ApplicationImageConfigurationInput?> ImageConfiguration { get; private set; } = null!;

        /// <summary>
        /// Initial capacity initialized when an Application is started.
        /// </summary>
        [Output("initialCapacity")]
        public Output<ImmutableArray<Outputs.ApplicationInitialCapacityConfigKeyValuePair>> InitialCapacity { get; private set; } = null!;

        /// <summary>
        /// The interactive configuration object that enables the interactive use cases for an application.
        /// </summary>
        [Output("interactiveConfiguration")]
        public Output<Outputs.ApplicationInteractiveConfiguration?> InteractiveConfiguration { get; private set; } = null!;

        /// <summary>
        /// Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.
        /// </summary>
        [Output("maximumCapacity")]
        public Output<Outputs.ApplicationMaximumAllowedResources?> MaximumCapacity { get; private set; } = null!;

        /// <summary>
        /// A configuration specification to be used when provisioning an application. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file.
        /// </summary>
        [Output("monitoringConfiguration")]
        public Output<Outputs.ApplicationMonitoringConfiguration?> MonitoringConfiguration { get; private set; } = null!;

        /// <summary>
        /// User friendly Application name.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Network Configuration for customer VPC connectivity.
        /// </summary>
        [Output("networkConfiguration")]
        public Output<Outputs.ApplicationNetworkConfiguration?> NetworkConfiguration { get; private set; } = null!;

        /// <summary>
        /// EMR release label.
        /// </summary>
        [Output("releaseLabel")]
        public Output<string> ReleaseLabel { get; private set; } = null!;

        /// <summary>
        /// The [Configuration](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_Configuration.html) specifications of an application. Each configuration consists of a classification and properties. You use this parameter when creating or updating an application. To see the runtimeConfiguration object of an application, run the [GetApplication](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_GetApplication.html) API operation.
        /// </summary>
        [Output("runtimeConfiguration")]
        public Output<ImmutableArray<Outputs.ApplicationConfigurationObject>> RuntimeConfiguration { get; private set; } = null!;

        /// <summary>
        /// The scheduler configuration for batch and streaming jobs running on this application. Supported with release labels emr-7.0.0 and above.
        /// </summary>
        [Output("schedulerConfiguration")]
        public Output<Outputs.ApplicationSchedulerConfiguration?> SchedulerConfiguration { get; private set; } = null!;

        /// <summary>
        /// Tag map with key and value
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the application
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.
        /// </summary>
        [Output("workerTypeSpecifications")]
        public Output<ImmutableDictionary<string, Outputs.ApplicationWorkerTypeSpecificationInput>?> WorkerTypeSpecifications { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:emrserverless:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:emrserverless:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU architecture of an application.
        /// </summary>
        [Input("architecture")]
        public Input<Pulumi.AwsNative.EmrServerless.ApplicationArchitecture>? Architecture { get; set; }

        /// <summary>
        /// Configuration for Auto Start of Application.
        /// </summary>
        [Input("autoStartConfiguration")]
        public Input<Inputs.ApplicationAutoStartConfigurationArgs>? AutoStartConfiguration { get; set; }

        /// <summary>
        /// Configuration for Auto Stop of Application.
        /// </summary>
        [Input("autoStopConfiguration")]
        public Input<Inputs.ApplicationAutoStopConfigurationArgs>? AutoStopConfiguration { get; set; }

        /// <summary>
        /// The IAM IdentityCenter configuration for trusted-identity-propagation on this application. Supported with release labels emr-7.8.0 and above.
        /// </summary>
        [Input("identityCenterConfiguration")]
        public Input<Inputs.ApplicationIdentityCenterConfigurationArgs>? IdentityCenterConfiguration { get; set; }

        /// <summary>
        /// The image configuration applied to all worker types.
        /// </summary>
        [Input("imageConfiguration")]
        public Input<Inputs.ApplicationImageConfigurationInputArgs>? ImageConfiguration { get; set; }

        [Input("initialCapacity")]
        private InputList<Inputs.ApplicationInitialCapacityConfigKeyValuePairArgs>? _initialCapacity;

        /// <summary>
        /// Initial capacity initialized when an Application is started.
        /// </summary>
        public InputList<Inputs.ApplicationInitialCapacityConfigKeyValuePairArgs> InitialCapacity
        {
            get => _initialCapacity ?? (_initialCapacity = new InputList<Inputs.ApplicationInitialCapacityConfigKeyValuePairArgs>());
            set => _initialCapacity = value;
        }

        /// <summary>
        /// The interactive configuration object that enables the interactive use cases for an application.
        /// </summary>
        [Input("interactiveConfiguration")]
        public Input<Inputs.ApplicationInteractiveConfigurationArgs>? InteractiveConfiguration { get; set; }

        /// <summary>
        /// Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.
        /// </summary>
        [Input("maximumCapacity")]
        public Input<Inputs.ApplicationMaximumAllowedResourcesArgs>? MaximumCapacity { get; set; }

        /// <summary>
        /// A configuration specification to be used when provisioning an application. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file.
        /// </summary>
        [Input("monitoringConfiguration")]
        public Input<Inputs.ApplicationMonitoringConfigurationArgs>? MonitoringConfiguration { get; set; }

        /// <summary>
        /// User friendly Application name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network Configuration for customer VPC connectivity.
        /// </summary>
        [Input("networkConfiguration")]
        public Input<Inputs.ApplicationNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        /// <summary>
        /// EMR release label.
        /// </summary>
        [Input("releaseLabel", required: true)]
        public Input<string> ReleaseLabel { get; set; } = null!;

        [Input("runtimeConfiguration")]
        private InputList<Inputs.ApplicationConfigurationObjectArgs>? _runtimeConfiguration;

        /// <summary>
        /// The [Configuration](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_Configuration.html) specifications of an application. Each configuration consists of a classification and properties. You use this parameter when creating or updating an application. To see the runtimeConfiguration object of an application, run the [GetApplication](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_GetApplication.html) API operation.
        /// </summary>
        public InputList<Inputs.ApplicationConfigurationObjectArgs> RuntimeConfiguration
        {
            get => _runtimeConfiguration ?? (_runtimeConfiguration = new InputList<Inputs.ApplicationConfigurationObjectArgs>());
            set => _runtimeConfiguration = value;
        }

        /// <summary>
        /// The scheduler configuration for batch and streaming jobs running on this application. Supported with release labels emr-7.0.0 and above.
        /// </summary>
        [Input("schedulerConfiguration")]
        public Input<Inputs.ApplicationSchedulerConfigurationArgs>? SchedulerConfiguration { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Tag map with key and value
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the application
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("workerTypeSpecifications")]
        private InputMap<Inputs.ApplicationWorkerTypeSpecificationInputArgs>? _workerTypeSpecifications;

        /// <summary>
        /// The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.
        /// </summary>
        public InputMap<Inputs.ApplicationWorkerTypeSpecificationInputArgs> WorkerTypeSpecifications
        {
            get => _workerTypeSpecifications ?? (_workerTypeSpecifications = new InputMap<Inputs.ApplicationWorkerTypeSpecificationInputArgs>());
            set => _workerTypeSpecifications = value;
        }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}
