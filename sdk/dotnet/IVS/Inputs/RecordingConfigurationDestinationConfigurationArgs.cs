// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IVS.Inputs
{

    /// <summary>
    /// Recording Destination Configuration.
    /// </summary>
    public sealed class RecordingConfigurationDestinationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("s3", required: true)]
        public Input<Inputs.RecordingConfigurationS3DestinationConfigurationArgs> S3 { get; set; } = null!;

        public RecordingConfigurationDestinationConfigurationArgs()
        {
        }
        public static new RecordingConfigurationDestinationConfigurationArgs Empty => new RecordingConfigurationDestinationConfigurationArgs();
    }
}