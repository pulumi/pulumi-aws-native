// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorS3SourceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains information about a source file's formatting.
        /// </summary>
        [Input("fileFormatDescriptor", required: true)]
        public Input<Inputs.AnomalyDetectorFileFormatDescriptorArgs> FileFormatDescriptor { get; set; } = null!;

        [Input("historicalDataPathList")]
        private InputList<string>? _historicalDataPathList;

        /// <summary>
        /// A list of paths to the historical data files.
        /// </summary>
        public InputList<string> HistoricalDataPathList
        {
            get => _historicalDataPathList ?? (_historicalDataPathList = new InputList<string>());
            set => _historicalDataPathList = value;
        }

        /// <summary>
        /// The ARN of an IAM role that has read and write access permissions to the source S3 bucket.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("templatedPathList")]
        private InputList<string>? _templatedPathList;

        /// <summary>
        /// A list of templated paths to the source files.
        /// </summary>
        public InputList<string> TemplatedPathList
        {
            get => _templatedPathList ?? (_templatedPathList = new InputList<string>());
            set => _templatedPathList = value;
        }

        public AnomalyDetectorS3SourceConfigArgs()
        {
        }
        public static new AnomalyDetectorS3SourceConfigArgs Empty => new AnomalyDetectorS3SourceConfigArgs();
    }
}
