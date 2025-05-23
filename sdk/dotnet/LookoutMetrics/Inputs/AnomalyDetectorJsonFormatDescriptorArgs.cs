// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorJsonFormatDescriptorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The character set in which the source JSON file is written.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// The level of compression of the source CSV file.
        /// </summary>
        [Input("fileCompression")]
        public Input<Pulumi.AwsNative.LookoutMetrics.AnomalyDetectorJsonFormatDescriptorFileCompression>? FileCompression { get; set; }

        public AnomalyDetectorJsonFormatDescriptorArgs()
        {
        }
        public static new AnomalyDetectorJsonFormatDescriptorArgs Empty => new AnomalyDetectorJsonFormatDescriptorArgs();
    }
}
