// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig
{
    /// <summary>
    /// Resource Type definition for AWS::AppConfig::Environment
    /// </summary>
    [AwsNativeResourceType("aws-native:appconfig:Environment")]
    public partial class Environment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
        /// </summary>
        [Output("deletionProtectionCheck")]
        public Output<Pulumi.AwsNative.AppConfig.EnvironmentDeletionProtectionCheck?> DeletionProtectionCheck { get; private set; } = null!;

        /// <summary>
        /// A description of the environment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The environment ID.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// Amazon CloudWatch alarms to monitor during the deployment process.
        /// </summary>
        [Output("monitors")]
        public Output<ImmutableArray<Outputs.EnvironmentMonitor>> Monitors { get; private set; } = null!;

        /// <summary>
        /// A name for the environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:Environment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "applicationId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, options);
        }
    }

    public sealed class EnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
        /// </summary>
        [Input("deletionProtectionCheck")]
        public Input<Pulumi.AwsNative.AppConfig.EnvironmentDeletionProtectionCheck>? DeletionProtectionCheck { get; set; }

        /// <summary>
        /// A description of the environment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("monitors")]
        private InputList<Inputs.EnvironmentMonitorArgs>? _monitors;

        /// <summary>
        /// Amazon CloudWatch alarms to monitor during the deployment process.
        /// </summary>
        public InputList<Inputs.EnvironmentMonitorArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.EnvironmentMonitorArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// A name for the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public EnvironmentArgs()
        {
        }
        public static new EnvironmentArgs Empty => new EnvironmentArgs();
    }
}
