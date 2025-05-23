// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs.Outputs
{

    /// <summary>
    /// Use this parameter to include the [grok](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-grok) processor in your transformer.
    /// </summary>
    [OutputType]
    public sealed class TransformerProcessorGrokProperties
    {
        public readonly string Match;
        public readonly string? Source;

        [OutputConstructor]
        private TransformerProcessorGrokProperties(
            string match,

            string? source)
        {
            Match = match;
            Source = source;
        }
    }
}
