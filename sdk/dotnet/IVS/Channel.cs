// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IVS
{
    /// <summary>
    /// Resource Type definition for AWS::IVS::Channel
    /// </summary>
    [AwsNativeResourceType("aws-native:ivs:Channel")]
    public partial class Channel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Channel ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Whether the channel is authorized.
        /// </summary>
        [Output("authorized")]
        public Output<bool?> Authorized { get; private set; } = null!;

        /// <summary>
        /// Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.
        /// </summary>
        [Output("ingestEndpoint")]
        public Output<string> IngestEndpoint { get; private set; } = null!;

        /// <summary>
        /// Channel latency mode.
        /// </summary>
        [Output("latencyMode")]
        public Output<Pulumi.AwsNative.IVS.ChannelLatencyMode?> LatencyMode { get; private set; } = null!;

        /// <summary>
        /// Channel
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Channel Playback URL.
        /// </summary>
        [Output("playbackUrl")]
        public Output<string> PlaybackUrl { get; private set; } = null!;

        /// <summary>
        /// Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: "" (recording is disabled).
        /// </summary>
        [Output("recordingConfigurationArn")]
        public Output<string?> RecordingConfigurationArn { get; private set; } = null!;

        /// <summary>
        /// A list of key-value pairs that contain metadata for the asset model.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ChannelTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
        /// </summary>
        [Output("type")]
        public Output<Pulumi.AwsNative.IVS.ChannelType?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ivs:Channel", name, args ?? new ChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ivs:Channel", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Channel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Channel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Channel(name, id, options);
        }
    }

    public sealed class ChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the channel is authorized.
        /// </summary>
        [Input("authorized")]
        public Input<bool>? Authorized { get; set; }

        /// <summary>
        /// Channel latency mode.
        /// </summary>
        [Input("latencyMode")]
        public Input<Pulumi.AwsNative.IVS.ChannelLatencyMode>? LatencyMode { get; set; }

        /// <summary>
        /// Channel
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: "" (recording is disabled).
        /// </summary>
        [Input("recordingConfigurationArn")]
        public Input<string>? RecordingConfigurationArn { get; set; }

        [Input("tags")]
        private InputList<Inputs.ChannelTagArgs>? _tags;

        /// <summary>
        /// A list of key-value pairs that contain metadata for the asset model.
        /// </summary>
        public InputList<Inputs.ChannelTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ChannelTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.IVS.ChannelType>? Type { get; set; }

        public ChannelArgs()
        {
        }
        public static new ChannelArgs Empty => new ChannelArgs();
    }
}