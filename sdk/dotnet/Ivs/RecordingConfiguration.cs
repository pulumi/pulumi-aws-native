// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs
{
    /// <summary>
    /// Resource Type definition for AWS::IVS::RecordingConfiguration
    /// </summary>
    [AwsNativeResourceType("aws-native:ivs:RecordingConfiguration")]
    public partial class RecordingConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A destination configuration describes an S3 bucket where recorded video will be stored. See the DestinationConfiguration property type for more information.
        /// </summary>
        [Output("destinationConfiguration")]
        public Output<Outputs.RecordingConfigurationDestinationConfiguration> DestinationConfiguration { get; private set; } = null!;

        /// <summary>
        /// Recording Configuration Name.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Recording Reconnect Window Seconds. (0 means disabled)
        /// </summary>
        [Output("recordingReconnectWindowSeconds")]
        public Output<int?> RecordingReconnectWindowSeconds { get; private set; } = null!;

        /// <summary>
        /// A rendition configuration describes which renditions should be recorded for a stream. See the RenditionConfiguration property type for more information.
        /// </summary>
        [Output("renditionConfiguration")]
        public Output<Outputs.RecordingConfigurationRenditionConfiguration?> RenditionConfiguration { get; private set; } = null!;

        /// <summary>
        /// Recording Configuration State.
        /// </summary>
        [Output("state")]
        public Output<Pulumi.AwsNative.Ivs.RecordingConfigurationState> State { get; private set; } = null!;

        /// <summary>
        /// A list of key-value pairs that contain metadata for the asset model.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// A thumbnail configuration enables/disables the recording of thumbnails for a live session and controls the interval at which thumbnails are generated for the live session. See the ThumbnailConfiguration property type for more information.
        /// </summary>
        [Output("thumbnailConfiguration")]
        public Output<Outputs.RecordingConfigurationThumbnailConfiguration?> ThumbnailConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a RecordingConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecordingConfiguration(string name, RecordingConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ivs:RecordingConfiguration", name, args ?? new RecordingConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecordingConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ivs:RecordingConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "destinationConfiguration",
                    "name",
                    "recordingReconnectWindowSeconds",
                    "renditionConfiguration",
                    "thumbnailConfiguration",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RecordingConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecordingConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RecordingConfiguration(name, id, options);
        }
    }

    public sealed class RecordingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A destination configuration describes an S3 bucket where recorded video will be stored. See the DestinationConfiguration property type for more information.
        /// </summary>
        [Input("destinationConfiguration", required: true)]
        public Input<Inputs.RecordingConfigurationDestinationConfigurationArgs> DestinationConfiguration { get; set; } = null!;

        /// <summary>
        /// Recording Configuration Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Recording Reconnect Window Seconds. (0 means disabled)
        /// </summary>
        [Input("recordingReconnectWindowSeconds")]
        public Input<int>? RecordingReconnectWindowSeconds { get; set; }

        /// <summary>
        /// A rendition configuration describes which renditions should be recorded for a stream. See the RenditionConfiguration property type for more information.
        /// </summary>
        [Input("renditionConfiguration")]
        public Input<Inputs.RecordingConfigurationRenditionConfigurationArgs>? RenditionConfiguration { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A list of key-value pairs that contain metadata for the asset model.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// A thumbnail configuration enables/disables the recording of thumbnails for a live session and controls the interval at which thumbnails are generated for the live session. See the ThumbnailConfiguration property type for more information.
        /// </summary>
        [Input("thumbnailConfiguration")]
        public Input<Inputs.RecordingConfigurationThumbnailConfigurationArgs>? ThumbnailConfiguration { get; set; }

        public RecordingConfigurationArgs()
        {
        }
        public static new RecordingConfigurationArgs Empty => new RecordingConfigurationArgs();
    }
}
