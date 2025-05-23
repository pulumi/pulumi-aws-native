// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class PipelineRemoveAttributes
    {
        /// <summary>
        /// A list of 1-50 attributes to remove from the message.
        /// </summary>
        public readonly ImmutableArray<string> Attributes;
        /// <summary>
        /// The name of the 'removeAttributes' activity.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The next activity in the pipeline.
        /// </summary>
        public readonly string? Next;

        [OutputConstructor]
        private PipelineRemoveAttributes(
            ImmutableArray<string> attributes,

            string name,

            string? next)
        {
            Attributes = attributes;
            Name = name;
            Next = next;
        }
    }
}
