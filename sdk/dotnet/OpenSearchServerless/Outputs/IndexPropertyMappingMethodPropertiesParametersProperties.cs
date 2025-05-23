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
    /// Additional parameters for the k-NN algorithm
    /// </summary>
    [OutputType]
    public sealed class IndexPropertyMappingMethodPropertiesParametersProperties
    {
        /// <summary>
        /// The size of the dynamic list used during k-NN graph creation
        /// </summary>
        public readonly int? EfConstruction;
        /// <summary>
        /// Number of neighbors to consider during k-NN search
        /// </summary>
        public readonly int? M;

        [OutputConstructor]
        private IndexPropertyMappingMethodPropertiesParametersProperties(
            int? efConstruction,

            int? m)
        {
            EfConstruction = efConstruction;
            M = m;
        }
    }
}
