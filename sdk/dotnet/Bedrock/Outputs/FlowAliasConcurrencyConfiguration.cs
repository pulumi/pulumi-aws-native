// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    [OutputType]
    public sealed class FlowAliasConcurrencyConfiguration
    {
        /// <summary>
        /// Number of nodes executed concurrently at a time
        /// </summary>
        public readonly double? MaxConcurrency;
        public readonly Pulumi.AwsNative.Bedrock.FlowAliasConcurrencyType Type;

        [OutputConstructor]
        private FlowAliasConcurrencyConfiguration(
            double? maxConcurrency,

            Pulumi.AwsNative.Bedrock.FlowAliasConcurrencyType type)
        {
            MaxConcurrency = maxConcurrency;
            Type = type;
        }
    }
}
