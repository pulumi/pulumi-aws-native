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
    /// The configuration used by AWS Lambda to access a self-managed event source.
    /// </summary>
    [OutputType]
    public sealed class EventSourceMappingSelfManagedEventSource
    {
        /// <summary>
        /// The endpoints for a self-managed event source.
        /// </summary>
        public readonly Outputs.EventSourceMappingEndpoints? Endpoints;

        [OutputConstructor]
        private EventSourceMappingSelfManagedEventSource(Outputs.EventSourceMappingEndpoints? endpoints)
        {
            Endpoints = endpoints;
        }
    }
}