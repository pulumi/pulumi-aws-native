// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchServerless.Outputs
{

    /// <summary>
    /// Index Mappings
    /// </summary>
    [OutputType]
    public sealed class MappingsProperties
    {
        /// <summary>
        /// Defines the fields within the mapping, including their types and configurations
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.IndexPropertyMapping>? Properties;

        [OutputConstructor]
        private MappingsProperties(ImmutableDictionary<string, Outputs.IndexPropertyMapping>? properties)
        {
            Properties = properties;
        }
    }
}
