// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
    /// </summary>
    [OutputType]
    public sealed class EventSourceMappingDestinationConfig
    {
        /// <summary>
        /// The destination configuration for failed invocations.
        /// </summary>
        public readonly Outputs.EventSourceMappingOnFailure? OnFailure;

        [OutputConstructor]
        private EventSourceMappingDestinationConfig(Outputs.EventSourceMappingOnFailure? onFailure)
        {
            OnFailure = onFailure;
        }
    }
}