// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive
{
    /// <summary>
    /// Definition of AWS::MediaLive::SdiSource Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:medialive:SdiSource")]
    public partial class SdiSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique arn of the SdiSource.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the SdiSource.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The list of inputs currently using this SDI source.
        /// </summary>
        [Output("inputs")]
        public Output<ImmutableArray<string>> Inputs { get; private set; } = null!;

        [Output("mode")]
        public Output<Pulumi.AwsNative.MediaLive.SdiSourceMode?> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the SdiSource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("state")]
        public Output<Pulumi.AwsNative.MediaLive.SdiSourceState> State { get; private set; } = null!;

        /// <summary>
        /// A collection of key-value pairs.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.AwsNative.MediaLive.SdiSourceType> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SdiSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SdiSource(string name, SdiSourceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:medialive:SdiSource", name, args ?? new SdiSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SdiSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:medialive:SdiSource", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SdiSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SdiSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SdiSource(name, id, options);
        }
    }

    public sealed class SdiSourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode")]
        public Input<Pulumi.AwsNative.MediaLive.SdiSourceMode>? Mode { get; set; }

        /// <summary>
        /// The name of the SdiSource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A collection of key-value pairs.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.MediaLive.SdiSourceType> Type { get; set; } = null!;

        public SdiSourceArgs()
        {
        }
        public static new SdiSourceArgs Empty => new SdiSourceArgs();
    }
}
