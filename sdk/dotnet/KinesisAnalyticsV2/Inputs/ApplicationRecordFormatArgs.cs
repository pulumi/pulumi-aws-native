// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// For a SQL-based Kinesis Data Analytics application, describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
    /// </summary>
    public sealed class ApplicationRecordFormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When you configure application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
        /// </summary>
        [Input("mappingParameters")]
        public Input<Inputs.ApplicationMappingParametersArgs>? MappingParameters { get; set; }

        /// <summary>
        /// The type of record format.
        /// </summary>
        [Input("recordFormatType", required: true)]
        public Input<Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationRecordFormatRecordFormatType> RecordFormatType { get; set; } = null!;

        public ApplicationRecordFormatArgs()
        {
        }
        public static new ApplicationRecordFormatArgs Empty => new ApplicationRecordFormatArgs();
    }
}
