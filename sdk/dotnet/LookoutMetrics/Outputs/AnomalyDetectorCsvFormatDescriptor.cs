// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorCsvFormatDescriptor
    {
        public readonly string? Charset;
        public readonly bool? ContainsHeader;
        public readonly string? Delimiter;
        public readonly Pulumi.AwsNative.LookoutMetrics.AnomalyDetectorCsvFormatDescriptorFileCompression? FileCompression;
        public readonly ImmutableArray<string> HeaderList;
        public readonly string? QuoteSymbol;

        [OutputConstructor]
        private AnomalyDetectorCsvFormatDescriptor(
            string? charset,

            bool? containsHeader,

            string? delimiter,

            Pulumi.AwsNative.LookoutMetrics.AnomalyDetectorCsvFormatDescriptorFileCompression? fileCompression,

            ImmutableArray<string> headerList,

            string? quoteSymbol)
        {
            Charset = charset;
            ContainsHeader = containsHeader;
            Delimiter = delimiter;
            FileCompression = fileCompression;
            HeaderList = headerList;
            QuoteSymbol = quoteSymbol;
        }
    }
}