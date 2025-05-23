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
    /// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when JSON is the record format on the streaming source.
    /// </summary>
    public sealed class ApplicationJsonMappingParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path to the top-level parent that contains the records.
        /// </summary>
        [Input("recordRowPath", required: true)]
        public Input<string> RecordRowPath { get; set; } = null!;

        public ApplicationJsonMappingParametersArgs()
        {
        }
        public static new ApplicationJsonMappingParametersArgs Empty => new ApplicationJsonMappingParametersArgs();
    }
}
