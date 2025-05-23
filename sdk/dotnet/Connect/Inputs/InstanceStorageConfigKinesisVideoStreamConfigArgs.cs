// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    public sealed class InstanceStorageConfigKinesisVideoStreamConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encryption configuration.
        /// </summary>
        [Input("encryptionConfig", required: true)]
        public Input<Inputs.InstanceStorageConfigEncryptionConfigArgs> EncryptionConfig { get; set; } = null!;

        /// <summary>
        /// The prefix of the video stream.
        /// </summary>
        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        /// <summary>
        /// The number of hours data is retained in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream.
        /// 
        /// The default value is 0, indicating that the stream does not persist data.
        /// </summary>
        [Input("retentionPeriodHours", required: true)]
        public Input<double> RetentionPeriodHours { get; set; } = null!;

        public InstanceStorageConfigKinesisVideoStreamConfigArgs()
        {
        }
        public static new InstanceStorageConfigKinesisVideoStreamConfigArgs Empty => new InstanceStorageConfigKinesisVideoStreamConfigArgs();
    }
}
