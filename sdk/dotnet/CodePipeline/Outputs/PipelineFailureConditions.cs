// *** WARNING: this file was generated by pulumi. ***
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
    public sealed class PipelineFailureConditions
    {
        /// <summary>
        /// The specified result for when the failure conditions are met, such as rolling back the stage
        /// </summary>
        public readonly Pulumi.AwsNative.CodePipeline.PipelineFailureConditionsResult? Result;

        [OutputConstructor]
        private PipelineFailureConditions(Pulumi.AwsNative.CodePipeline.PipelineFailureConditionsResult? result)
        {
            Result = result;
        }
    }
}