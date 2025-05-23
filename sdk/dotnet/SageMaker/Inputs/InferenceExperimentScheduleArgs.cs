// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The duration for which you want the inference experiment to run.
    /// </summary>
    public sealed class InferenceExperimentScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp at which the inference experiment ended or will end.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The timestamp at which the inference experiment started or will start.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public InferenceExperimentScheduleArgs()
        {
        }
        public static new InferenceExperimentScheduleArgs Empty => new InferenceExperimentScheduleArgs();
    }
}
