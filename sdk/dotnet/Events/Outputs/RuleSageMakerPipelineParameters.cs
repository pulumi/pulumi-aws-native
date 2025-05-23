// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class RuleSageMakerPipelineParameters
    {
        /// <summary>
        /// List of Parameter names and values for SageMaker AI Model Building Pipeline execution.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleSageMakerPipelineParameter> PipelineParameterList;

        [OutputConstructor]
        private RuleSageMakerPipelineParameters(ImmutableArray<Outputs.RuleSageMakerPipelineParameter> pipelineParameterList)
        {
            PipelineParameterList = pipelineParameterList;
        }
    }
}
