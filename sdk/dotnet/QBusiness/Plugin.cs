// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness
{
    /// <summary>
    /// Definition of AWS::QBusiness::Plugin Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:qbusiness:Plugin")]
    public partial class Plugin : global::Pulumi.CustomResource
    {
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        [Output("authConfiguration")]
        public Output<object> AuthConfiguration { get; private set; } = null!;

        [Output("buildStatus")]
        public Output<Pulumi.AwsNative.QBusiness.PluginBuildStatus> BuildStatus { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("customPluginConfiguration")]
        public Output<Outputs.PluginCustomPluginConfiguration?> CustomPluginConfiguration { get; private set; } = null!;

        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        [Output("pluginArn")]
        public Output<string> PluginArn { get; private set; } = null!;

        [Output("pluginId")]
        public Output<string> PluginId { get; private set; } = null!;

        [Output("serverUrl")]
        public Output<string?> ServerUrl { get; private set; } = null!;

        [Output("state")]
        public Output<Pulumi.AwsNative.QBusiness.PluginState?> State { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.AwsNative.QBusiness.PluginType> Type { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Plugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Plugin(string name, PluginArgs args, CustomResourceOptions? options = null)
            : base("aws-native:qbusiness:Plugin", name, args ?? new PluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Plugin(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:qbusiness:Plugin", name, null, MakeResourceOptions(options, id))
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
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Plugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Plugin Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Plugin(name, id, options);
        }
    }

    public sealed class PluginArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        [Input("authConfiguration", required: true)]
        public object AuthConfiguration { get; set; } = null!;

        [Input("customPluginConfiguration")]
        public Input<Inputs.PluginCustomPluginConfigurationArgs>? CustomPluginConfiguration { get; set; }

        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("serverUrl")]
        public Input<string>? ServerUrl { get; set; }

        [Input("state")]
        public Input<Pulumi.AwsNative.QBusiness.PluginState>? State { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.QBusiness.PluginType> Type { get; set; } = null!;

        public PluginArgs()
        {
        }
        public static new PluginArgs Empty => new PluginArgs();
    }
}