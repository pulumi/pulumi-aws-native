// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    /// <summary>
    /// Resource Type definition for AWS::Config::ConfigurationAggregator
    /// </summary>
    [AwsNativeResourceType("aws-native:configuration:ConfigurationAggregator")]
    public partial class ConfigurationAggregator : global::Pulumi.CustomResource
    {
        [Output("accountAggregationSources")]
        public Output<ImmutableArray<Outputs.ConfigurationAggregatorAccountAggregationSource>> AccountAggregationSources { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the aggregator.
        /// </summary>
        [Output("configurationAggregatorArn")]
        public Output<string> ConfigurationAggregatorArn { get; private set; } = null!;

        /// <summary>
        /// The name of the aggregator.
        /// </summary>
        [Output("configurationAggregatorName")]
        public Output<string?> ConfigurationAggregatorName { get; private set; } = null!;

        [Output("organizationAggregationSource")]
        public Output<Outputs.ConfigurationAggregatorOrganizationAggregationSource?> OrganizationAggregationSource { get; private set; } = null!;

        /// <summary>
        /// The tags for the configuration aggregator.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ConfigurationAggregatorTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationAggregator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationAggregator(string name, ConfigurationAggregatorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConfigurationAggregator", name, args ?? new ConfigurationAggregatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationAggregator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConfigurationAggregator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConfigurationAggregator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationAggregator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationAggregator(name, id, options);
        }
    }

    public sealed class ConfigurationAggregatorArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountAggregationSources")]
        private InputList<Inputs.ConfigurationAggregatorAccountAggregationSourceArgs>? _accountAggregationSources;
        public InputList<Inputs.ConfigurationAggregatorAccountAggregationSourceArgs> AccountAggregationSources
        {
            get => _accountAggregationSources ?? (_accountAggregationSources = new InputList<Inputs.ConfigurationAggregatorAccountAggregationSourceArgs>());
            set => _accountAggregationSources = value;
        }

        /// <summary>
        /// The name of the aggregator.
        /// </summary>
        [Input("configurationAggregatorName")]
        public Input<string>? ConfigurationAggregatorName { get; set; }

        [Input("organizationAggregationSource")]
        public Input<Inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs>? OrganizationAggregationSource { get; set; }

        [Input("tags")]
        private InputList<Inputs.ConfigurationAggregatorTagArgs>? _tags;

        /// <summary>
        /// The tags for the configuration aggregator.
        /// </summary>
        public InputList<Inputs.ConfigurationAggregatorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ConfigurationAggregatorTagArgs>());
            set => _tags = value;
        }

        public ConfigurationAggregatorArgs()
        {
        }
        public static new ConfigurationAggregatorArgs Empty => new ConfigurationAggregatorArgs();
    }
}