// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Inputs
{

    /// <summary>
    /// The Git repository branches specified as filter criteria to start the pipeline.
    /// </summary>
    public sealed class PipelineGitBranchFilterCriteriaArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludes")]
        private InputList<string>? _excludes;

        /// <summary>
        /// The list of patterns of Git branches that, when a commit is pushed, are to be excluded from starting the pipeline.
        /// </summary>
        public InputList<string> Excludes
        {
            get => _excludes ?? (_excludes = new InputList<string>());
            set => _excludes = value;
        }

        [Input("includes")]
        private InputList<string>? _includes;

        /// <summary>
        /// The list of patterns of Git branches that, when a commit is pushed, are to be included as criteria that starts the pipeline.
        /// </summary>
        public InputList<string> Includes
        {
            get => _includes ?? (_includes = new InputList<string>());
            set => _includes = value;
        }

        public PipelineGitBranchFilterCriteriaArgs()
        {
        }
        public static new PipelineGitBranchFilterCriteriaArgs Empty => new PipelineGitBranchFilterCriteriaArgs();
    }
}
