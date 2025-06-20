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
    public sealed class EventSourceMappingSchemaRegistryConfig
    {
        public readonly ImmutableArray<Outputs.EventSourceMappingSchemaRegistryAccessConfig> AccessConfigs;
        public readonly Pulumi.AwsNative.Lambda.EventSourceMappingSchemaRegistryConfigEventRecordFormat? EventRecordFormat;
        public readonly string? SchemaRegistryUri;
        public readonly ImmutableArray<Outputs.EventSourceMappingSchemaValidationConfig> SchemaValidationConfigs;

        [OutputConstructor]
        private EventSourceMappingSchemaRegistryConfig(
            ImmutableArray<Outputs.EventSourceMappingSchemaRegistryAccessConfig> accessConfigs,

            Pulumi.AwsNative.Lambda.EventSourceMappingSchemaRegistryConfigEventRecordFormat? eventRecordFormat,

            string? schemaRegistryUri,

            ImmutableArray<Outputs.EventSourceMappingSchemaValidationConfig> schemaValidationConfigs)
        {
            AccessConfigs = accessConfigs;
            EventRecordFormat = eventRecordFormat;
            SchemaRegistryUri = schemaRegistryUri;
            SchemaValidationConfigs = schemaValidationConfigs;
        }
    }
}
