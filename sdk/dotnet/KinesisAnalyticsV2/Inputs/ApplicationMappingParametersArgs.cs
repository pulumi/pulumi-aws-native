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
    /// When you configure a SQL-based Kinesis Data Analytics application's input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
    /// </summary>
    public sealed class ApplicationMappingParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provides additional mapping information when the record format uses delimiters (for example, CSV).
        /// </summary>
        [Input("cSVMappingParameters")]
        public Input<Inputs.ApplicationCSVMappingParametersArgs>? CSVMappingParameters { get; set; }

        /// <summary>
        /// Provides additional mapping information when JSON is the record format on the streaming source.
        /// </summary>
        [Input("jSONMappingParameters")]
        public Input<Inputs.ApplicationJSONMappingParametersArgs>? JSONMappingParameters { get; set; }

        public ApplicationMappingParametersArgs()
        {
        }
        public static new ApplicationMappingParametersArgs Empty => new ApplicationMappingParametersArgs();
    }
}