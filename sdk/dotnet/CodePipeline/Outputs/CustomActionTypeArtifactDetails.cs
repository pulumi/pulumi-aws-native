// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Outputs
{

    /// <summary>
    /// Returns information about the details of an artifact.
    /// </summary>
    [OutputType]
    public sealed class CustomActionTypeArtifactDetails
    {
        /// <summary>
        /// The maximum number of artifacts allowed for the action type.
        /// </summary>
        public readonly int MaximumCount;
        /// <summary>
        /// The minimum number of artifacts allowed for the action type.
        /// </summary>
        public readonly int MinimumCount;

        [OutputConstructor]
        private CustomActionTypeArtifactDetails(
            int maximumCount,

            int minimumCount)
        {
            MaximumCount = maximumCount;
            MinimumCount = minimumCount;
        }
    }
}