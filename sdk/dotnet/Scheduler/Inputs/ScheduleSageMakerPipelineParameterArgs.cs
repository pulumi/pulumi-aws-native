// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler.Inputs
{

    /// <summary>
    /// Name/Value pair of a parameter to start execution of a SageMaker Model Building Pipeline.
    /// </summary>
    public sealed class ScheduleSageMakerPipelineParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of parameter to start execution of a SageMaker Model Building Pipeline.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of parameter to start execution of a SageMaker Model Building Pipeline.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ScheduleSageMakerPipelineParameterArgs()
        {
        }
        public static new ScheduleSageMakerPipelineParameterArgs Empty => new ScheduleSageMakerPipelineParameterArgs();
    }
}
