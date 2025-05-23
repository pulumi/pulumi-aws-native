// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when JSON is the record format on the streaming source.
    /// </summary>
    [OutputType]
    public sealed class ApplicationJsonMappingParameters
    {
        /// <summary>
        /// The path to the top-level parent that contains the records.
        /// </summary>
        public readonly string RecordRowPath;

        [OutputConstructor]
        private ApplicationJsonMappingParameters(string recordRowPath)
        {
            RecordRowPath = recordRowPath;
        }
    }
}
