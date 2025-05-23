// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The infrastructure configuration for deploying the model to a real-time inference endpoint.
    /// </summary>
    [OutputType]
    public sealed class InferenceExperimentRealTimeInferenceConfig
    {
        /// <summary>
        /// The number of instances of the type specified by InstanceType.
        /// </summary>
        public readonly int InstanceCount;
        /// <summary>
        /// The instance type the model is deployed to.
        /// </summary>
        public readonly string InstanceType;

        [OutputConstructor]
        private InferenceExperimentRealTimeInferenceConfig(
            int instanceCount,

            string instanceType)
        {
            InstanceCount = instanceCount;
            InstanceType = instanceType;
        }
    }
}
