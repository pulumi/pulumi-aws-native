// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeTargetTimestreamParameters
    {
        /// <summary>
        /// Map source data to dimensions in the target Timestream for LiveAnalytics table.
        /// 
        /// For more information, see [Amazon Timestream for LiveAnalytics concepts](https://docs.aws.amazon.com/timestream/latest/developerguide/concepts.html)
        /// </summary>
        public readonly ImmutableArray<Outputs.PipeDimensionMapping> DimensionMappings;
        /// <summary>
        /// The granularity of the time units used. Default is `MILLISECONDS` .
        /// 
        /// Required if `TimeFieldType` is specified as `EPOCH` .
        /// </summary>
        public readonly Pulumi.AwsNative.Pipes.PipeEpochTimeUnit? EpochTimeUnit;
        /// <summary>
        /// Maps multiple measures from the source event to the same record in the specified Timestream for LiveAnalytics table.
        /// </summary>
        public readonly ImmutableArray<Outputs.PipeMultiMeasureMapping> MultiMeasureMappings;
        /// <summary>
        /// Mappings of single source data fields to individual records in the specified Timestream for LiveAnalytics table.
        /// </summary>
        public readonly ImmutableArray<Outputs.PipeSingleMeasureMapping> SingleMeasureMappings;
        /// <summary>
        /// The type of time value used.
        /// 
        /// The default is `EPOCH` .
        /// </summary>
        public readonly Pulumi.AwsNative.Pipes.PipeTimeFieldType? TimeFieldType;
        /// <summary>
        /// Dynamic path to the source data field that represents the time value for your data.
        /// </summary>
        public readonly string TimeValue;
        /// <summary>
        /// How to format the timestamps. For example, `yyyy-MM-dd'T'HH:mm:ss'Z'` .
        /// 
        /// Required if `TimeFieldType` is specified as `TIMESTAMP_FORMAT` .
        /// </summary>
        public readonly string? TimestampFormat;
        /// <summary>
        /// 64 bit version value or source data field that represents the version value for your data.
        /// 
        /// Write requests with a higher version number will update the existing measure values of the record and version. In cases where the measure value is the same, the version will still be updated.
        /// 
        /// Default value is 1.
        /// 
        /// Timestream for LiveAnalytics does not support updating partial measure values in a record.
        /// 
        /// Write requests for duplicate data with a higher version number will update the existing measure value and version. In cases where the measure value is the same, `Version` will still be updated. Default value is `1` .
        /// 
        /// &gt; `Version` must be `1` or greater, or you will receive a `ValidationException` error.
        /// </summary>
        public readonly string VersionValue;

        [OutputConstructor]
        private PipeTargetTimestreamParameters(
            ImmutableArray<Outputs.PipeDimensionMapping> dimensionMappings,

            Pulumi.AwsNative.Pipes.PipeEpochTimeUnit? epochTimeUnit,

            ImmutableArray<Outputs.PipeMultiMeasureMapping> multiMeasureMappings,

            ImmutableArray<Outputs.PipeSingleMeasureMapping> singleMeasureMappings,

            Pulumi.AwsNative.Pipes.PipeTimeFieldType? timeFieldType,

            string timeValue,

            string? timestampFormat,

            string versionValue)
        {
            DimensionMappings = dimensionMappings;
            EpochTimeUnit = epochTimeUnit;
            MultiMeasureMappings = multiMeasureMappings;
            SingleMeasureMappings = singleMeasureMappings;
            TimeFieldType = timeFieldType;
            TimeValue = timeValue;
            TimestampFormat = timestampFormat;
            VersionValue = versionValue;
        }
    }
}
