// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorFileFormatDescriptor
    {
        /// <summary>
        /// Contains information about how a source CSV data file should be analyzed.
        /// </summary>
        public readonly Outputs.AnomalyDetectorCsvFormatDescriptor? CsvFormatDescriptor;
        /// <summary>
        /// Contains information about how a source JSON data file should be analyzed.
        /// </summary>
        public readonly Outputs.AnomalyDetectorJsonFormatDescriptor? JsonFormatDescriptor;

        [OutputConstructor]
        private AnomalyDetectorFileFormatDescriptor(
            Outputs.AnomalyDetectorCsvFormatDescriptor? csvFormatDescriptor,

            Outputs.AnomalyDetectorJsonFormatDescriptor? jsonFormatDescriptor)
        {
            CsvFormatDescriptor = csvFormatDescriptor;
            JsonFormatDescriptor = jsonFormatDescriptor;
        }
    }
}
