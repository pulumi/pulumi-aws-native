// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder
{
    /// <summary>
    /// Resource schema for AWS::ImageBuilder::InfrastructureConfiguration
    /// </summary>
    [AwsNativeResourceType("aws-native:imagebuilder:InfrastructureConfiguration")]
    public partial class InfrastructureConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the infrastructure configuration.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the infrastructure configuration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The instance metadata option settings for the infrastructure configuration.
        /// </summary>
        [Output("instanceMetadataOptions")]
        public Output<Outputs.InfrastructureConfigurationInstanceMetadataOptions?> InstanceMetadataOptions { get; private set; } = null!;

        /// <summary>
        /// The instance profile of the infrastructure configuration.
        /// </summary>
        [Output("instanceProfileName")]
        public Output<string> InstanceProfileName { get; private set; } = null!;

        /// <summary>
        /// The instance types of the infrastructure configuration.
        /// </summary>
        [Output("instanceTypes")]
        public Output<ImmutableArray<string>> InstanceTypes { get; private set; } = null!;

        /// <summary>
        /// The EC2 key pair of the infrastructure configuration..
        /// </summary>
        [Output("keyPair")]
        public Output<string?> KeyPair { get; private set; } = null!;

        /// <summary>
        /// The logging configuration of the infrastructure configuration.
        /// </summary>
        [Output("logging")]
        public Output<Outputs.InfrastructureConfigurationLogging?> Logging { get; private set; } = null!;

        /// <summary>
        /// The name of the infrastructure configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The placement option settings for the infrastructure configuration.
        /// </summary>
        [Output("placement")]
        public Output<Outputs.InfrastructureConfigurationPlacement?> Placement { get; private set; } = null!;

        /// <summary>
        /// The tags attached to the resource created by Image Builder.
        /// </summary>
        [Output("resourceTags")]
        public Output<ImmutableDictionary<string, string>?> ResourceTags { get; private set; } = null!;

        /// <summary>
        /// The security group IDs of the infrastructure configuration.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.
        /// </summary>
        [Output("snsTopicArn")]
        public Output<string?> SnsTopicArn { get; private set; } = null!;

        /// <summary>
        /// The subnet ID of the infrastructure configuration.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the component.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The terminate instance on failure configuration of the infrastructure configuration.
        /// </summary>
        [Output("terminateInstanceOnFailure")]
        public Output<bool?> TerminateInstanceOnFailure { get; private set; } = null!;


        /// <summary>
        /// Create a InfrastructureConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InfrastructureConfiguration(string name, InfrastructureConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:imagebuilder:InfrastructureConfiguration", name, args ?? new InfrastructureConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InfrastructureConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:imagebuilder:InfrastructureConfiguration", name, null, MakeResourceOptions(options, id))
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
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InfrastructureConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InfrastructureConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InfrastructureConfiguration(name, id, options);
        }
    }

    public sealed class InfrastructureConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the infrastructure configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The instance metadata option settings for the infrastructure configuration.
        /// </summary>
        [Input("instanceMetadataOptions")]
        public Input<Inputs.InfrastructureConfigurationInstanceMetadataOptionsArgs>? InstanceMetadataOptions { get; set; }

        /// <summary>
        /// The instance profile of the infrastructure configuration.
        /// </summary>
        [Input("instanceProfileName", required: true)]
        public Input<string> InstanceProfileName { get; set; } = null!;

        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// The instance types of the infrastructure configuration.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        /// <summary>
        /// The EC2 key pair of the infrastructure configuration..
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        /// <summary>
        /// The logging configuration of the infrastructure configuration.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.InfrastructureConfigurationLoggingArgs>? Logging { get; set; }

        /// <summary>
        /// The name of the infrastructure configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The placement option settings for the infrastructure configuration.
        /// </summary>
        [Input("placement")]
        public Input<Inputs.InfrastructureConfigurationPlacementArgs>? Placement { get; set; }

        [Input("resourceTags")]
        private InputMap<string>? _resourceTags;

        /// <summary>
        /// The tags attached to the resource created by Image Builder.
        /// </summary>
        public InputMap<string> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputMap<string>());
            set => _resourceTags = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The security group IDs of the infrastructure configuration.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.
        /// </summary>
        [Input("snsTopicArn")]
        public Input<string>? SnsTopicArn { get; set; }

        /// <summary>
        /// The subnet ID of the infrastructure configuration.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags associated with the component.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The terminate instance on failure configuration of the infrastructure configuration.
        /// </summary>
        [Input("terminateInstanceOnFailure")]
        public Input<bool>? TerminateInstanceOnFailure { get; set; }

        public InfrastructureConfigurationArgs()
        {
        }
        public static new InfrastructureConfigurationArgs Empty => new InfrastructureConfigurationArgs();
    }
}
