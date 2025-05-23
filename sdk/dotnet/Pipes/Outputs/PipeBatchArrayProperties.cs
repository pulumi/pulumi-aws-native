// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeBatchArrayProperties
    {
        /// <summary>
        /// The size of the array, if this is an array batch job.
        /// </summary>
        public readonly int? Size;

        [OutputConstructor]
        private PipeBatchArrayProperties(int? size)
        {
            Size = size;
        }
    }
}
