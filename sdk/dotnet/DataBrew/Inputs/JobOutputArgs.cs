// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    public sealed class JobOutputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compression algorithm used to compress the output text of the job.
        /// </summary>
        [Input("compressionFormat")]
        public Input<Pulumi.AwsNative.DataBrew.JobOutputCompressionFormat>? CompressionFormat { get; set; }

        /// <summary>
        /// The data format of the output of the job.
        /// </summary>
        [Input("format")]
        public Input<Pulumi.AwsNative.DataBrew.JobOutputFormat>? Format { get; set; }

        /// <summary>
        /// Represents options that define how DataBrew formats job output files.
        /// </summary>
        [Input("formatOptions")]
        public Input<Inputs.JobOutputFormatOptionsArgs>? FormatOptions { get; set; }

        /// <summary>
        /// The location in Amazon S3 where the job writes its output.
        /// </summary>
        [Input("location", required: true)]
        public Input<Inputs.JobS3LocationArgs> Location { get; set; } = null!;

        /// <summary>
        /// The maximum number of files to be generated by the job and written to the output folder.
        /// </summary>
        [Input("maxOutputFiles")]
        public Input<int>? MaxOutputFiles { get; set; }

        /// <summary>
        /// A value that, if true, means that any data in the location specified for output is overwritten with new output.
        /// </summary>
        [Input("overwrite")]
        public Input<bool>? Overwrite { get; set; }

        [Input("partitionColumns")]
        private InputList<string>? _partitionColumns;

        /// <summary>
        /// The names of one or more partition columns for the output of the job.
        /// </summary>
        public InputList<string> PartitionColumns
        {
            get => _partitionColumns ?? (_partitionColumns = new InputList<string>());
            set => _partitionColumns = value;
        }

        public JobOutputArgs()
        {
        }
        public static new JobOutputArgs Empty => new JobOutputArgs();
    }
}
