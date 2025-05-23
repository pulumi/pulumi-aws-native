// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Outputs
{

    [OutputType]
    public sealed class IdMappingWorkflowInputSource
    {
        /// <summary>
        /// An Glue table ARN for the input source table, MatchingWorkflow arn or IdNamespace ARN
        /// </summary>
        public readonly string InputSourceArn;
        /// <summary>
        /// The ARN (Amazon Resource Name) that AWS Entity Resolution generated for the `SchemaMapping` .
        /// </summary>
        public readonly string? SchemaArn;
        /// <summary>
        /// The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
        /// 
        /// The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
        /// 
        /// The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
        /// </summary>
        public readonly Pulumi.AwsNative.EntityResolution.IdMappingWorkflowInputSourceType? Type;

        [OutputConstructor]
        private IdMappingWorkflowInputSource(
            string inputSourceArn,

            string? schemaArn,

            Pulumi.AwsNative.EntityResolution.IdMappingWorkflowInputSourceType? type)
        {
            InputSourceArn = inputSourceArn;
            SchemaArn = schemaArn;
            Type = type;
        }
    }
}
