// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Outputs
{

    [OutputType]
    public sealed class CanaryVisualReference
    {
        /// <summary>
        /// Canary run id to be used as base reference for visual testing
        /// </summary>
        public readonly string BaseCanaryRunId;
        /// <summary>
        /// List of screenshots used as base reference for visual testing
        /// </summary>
        public readonly ImmutableArray<Outputs.CanaryBaseScreenshot> BaseScreenshots;

        [OutputConstructor]
        private CanaryVisualReference(
            string baseCanaryRunId,

            ImmutableArray<Outputs.CanaryBaseScreenshot> baseScreenshots)
        {
            BaseCanaryRunId = baseCanaryRunId;
            BaseScreenshots = baseScreenshots;
        }
    }
}
