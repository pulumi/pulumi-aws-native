// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The media stream that is associated with the source, and the parameters for that association.
    /// </summary>
    public sealed class FlowMediaStreamSourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The format that was used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.
        /// </summary>
        [Input("encodingName", required: true)]
        public Input<Pulumi.AwsNative.MediaConnect.FlowMediaStreamSourceConfigurationEncodingName> EncodingName { get; set; } = null!;

        [Input("inputConfigurations")]
        private InputList<Inputs.FlowInputConfigurationArgs>? _inputConfigurations;

        /// <summary>
        /// The media streams that you want to associate with the source.
        /// </summary>
        public InputList<Inputs.FlowInputConfigurationArgs> InputConfigurations
        {
            get => _inputConfigurations ?? (_inputConfigurations = new InputList<Inputs.FlowInputConfigurationArgs>());
            set => _inputConfigurations = value;
        }

        /// <summary>
        /// A name that helps you distinguish one media stream from another.
        /// </summary>
        [Input("mediaStreamName", required: true)]
        public Input<string> MediaStreamName { get; set; } = null!;

        public FlowMediaStreamSourceConfigurationArgs()
        {
        }
        public static new FlowMediaStreamSourceConfigurationArgs Empty => new FlowMediaStreamSourceConfigurationArgs();
    }
}