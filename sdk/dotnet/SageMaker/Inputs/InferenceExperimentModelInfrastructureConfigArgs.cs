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
    /// The configuration for the infrastructure that the model will be deployed to.
    /// </summary>
    public sealed class InferenceExperimentModelInfrastructureConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the inference experiment that you want to run.
        /// </summary>
        [Input("infrastructureType", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.InferenceExperimentModelInfrastructureConfigInfrastructureType> InfrastructureType { get; set; } = null!;

        /// <summary>
        /// The infrastructure configuration for deploying the model to real-time inference.
        /// </summary>
        [Input("realTimeInferenceConfig", required: true)]
        public Input<Inputs.InferenceExperimentRealTimeInferenceConfigArgs> RealTimeInferenceConfig { get; set; } = null!;

        public InferenceExperimentModelInfrastructureConfigArgs()
        {
        }
        public static new InferenceExperimentModelInfrastructureConfigArgs Empty => new InferenceExperimentModelInfrastructureConfigArgs();
    }
}
