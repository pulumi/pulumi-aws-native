// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// For a SQL-based Kinesis Data Analytics application, describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
    /// </summary>
    public sealed class ApplicationInputSchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("recordColumns", required: true)]
        private InputList<Inputs.ApplicationRecordColumnArgs>? _recordColumns;

        /// <summary>
        /// A list of `RecordColumn` objects.
        /// </summary>
        public InputList<Inputs.ApplicationRecordColumnArgs> RecordColumns
        {
            get => _recordColumns ?? (_recordColumns = new InputList<Inputs.ApplicationRecordColumnArgs>());
            set => _recordColumns = value;
        }

        /// <summary>
        /// Specifies the encoding of the records in the streaming source. For example, UTF-8.
        /// </summary>
        [Input("recordEncoding")]
        public Input<Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationInputSchemaRecordEncoding>? RecordEncoding { get; set; }

        /// <summary>
        /// Specifies the format of the records on the streaming source.
        /// </summary>
        [Input("recordFormat", required: true)]
        public Input<Inputs.ApplicationRecordFormatArgs> RecordFormat { get; set; } = null!;

        public ApplicationInputSchemaArgs()
        {
        }
        public static new ApplicationInputSchemaArgs Empty => new ApplicationInputSchemaArgs();
    }
}