// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobOutput
    {
        public readonly Pulumi.AwsNative.DataBrew.JobOutputCompressionFormat? CompressionFormat;
        public readonly Pulumi.AwsNative.DataBrew.JobOutputFormat? Format;
        public readonly Outputs.JobOutputFormatOptions? FormatOptions;
        public readonly Outputs.JobS3Location Location;
        public readonly int? MaxOutputFiles;
        public readonly bool? Overwrite;
        public readonly ImmutableArray<string> PartitionColumns;

        [OutputConstructor]
        private JobOutput(
            Pulumi.AwsNative.DataBrew.JobOutputCompressionFormat? compressionFormat,

            Pulumi.AwsNative.DataBrew.JobOutputFormat? format,

            Outputs.JobOutputFormatOptions? formatOptions,

            Outputs.JobS3Location location,

            int? maxOutputFiles,

            bool? overwrite,

            ImmutableArray<string> partitionColumns)
        {
            CompressionFormat = compressionFormat;
            Format = format;
            FormatOptions = formatOptions;
            Location = location;
            MaxOutputFiles = maxOutputFiles;
            Overwrite = overwrite;
            PartitionColumns = partitionColumns;
        }
    }
}