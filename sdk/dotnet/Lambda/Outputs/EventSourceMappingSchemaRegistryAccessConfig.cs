// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    [OutputType]
    public sealed class EventSourceMappingSchemaRegistryAccessConfig
    {
        public readonly Pulumi.AwsNative.Lambda.EventSourceMappingSchemaRegistryAccessConfigType? Type;
        public readonly string? Uri;

        [OutputConstructor]
        private EventSourceMappingSchemaRegistryAccessConfig(
            Pulumi.AwsNative.Lambda.EventSourceMappingSchemaRegistryAccessConfigType? type,

            string? uri)
        {
            Type = type;
            Uri = uri;
        }
    }
}
