// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// Specifies how much memory is available to the container.
    /// </summary>
    [OutputType]
    public sealed class ContainerGroupDefinitionMemoryLimits
    {
        /// <summary>
        /// The hard limit of memory to reserve for the container.
        /// </summary>
        public readonly int? HardLimit;
        /// <summary>
        /// The amount of memory that is reserved for the container.
        /// </summary>
        public readonly int? SoftLimit;

        [OutputConstructor]
        private ContainerGroupDefinitionMemoryLimits(
            int? hardLimit,

            int? softLimit)
        {
            HardLimit = hardLimit;
            SoftLimit = softLimit;
        }
    }
}