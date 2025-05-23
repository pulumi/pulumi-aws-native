// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Outputs
{

    /// <summary>
    /// The Git tags specified as filter criteria for whether a Git tag repository event will start the pipeline.
    /// </summary>
    [OutputType]
    public sealed class PipelineGitTagFilterCriteria
    {
        /// <summary>
        /// The list of patterns of Git tags that, when pushed, are to be excluded from starting the pipeline.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// The list of patterns of Git tags that, when pushed, are to be included as criteria that starts the pipeline.
        /// </summary>
        public readonly ImmutableArray<string> Includes;

        [OutputConstructor]
        private PipelineGitTagFilterCriteria(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes)
        {
            Excludes = excludes;
            Includes = includes;
        }
    }
}
