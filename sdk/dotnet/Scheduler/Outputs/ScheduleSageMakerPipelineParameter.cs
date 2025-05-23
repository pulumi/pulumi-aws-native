// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler.Outputs
{

    /// <summary>
    /// Name/Value pair of a parameter to start execution of a SageMaker Model Building Pipeline.
    /// </summary>
    [OutputType]
    public sealed class ScheduleSageMakerPipelineParameter
    {
        /// <summary>
        /// Name of parameter to start execution of a SageMaker Model Building Pipeline.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Value of parameter to start execution of a SageMaker Model Building Pipeline.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ScheduleSageMakerPipelineParameter(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
