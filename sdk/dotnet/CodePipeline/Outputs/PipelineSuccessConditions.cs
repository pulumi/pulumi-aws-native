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
    /// The configuration that specifies the result, such as rollback, to occur upon stage failure
    /// </summary>
    [OutputType]
    public sealed class PipelineSuccessConditions
    {
        /// <summary>
        /// The conditions that are success conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.PipelineCondition> Conditions;

        [OutputConstructor]
        private PipelineSuccessConditions(ImmutableArray<Outputs.PipelineCondition> conditions)
        {
            Conditions = conditions;
        }
    }
}
