// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeEnrichmentParameters
    {
        public readonly Outputs.PipeEnrichmentHttpParameters? HttpParameters;
        public readonly string? InputTemplate;

        [OutputConstructor]
        private PipeEnrichmentParameters(
            Outputs.PipeEnrichmentHttpParameters? httpParameters,

            string? inputTemplate)
        {
            HttpParameters = httpParameters;
            InputTemplate = inputTemplate;
        }
    }
}