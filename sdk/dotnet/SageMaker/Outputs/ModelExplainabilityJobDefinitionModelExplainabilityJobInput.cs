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
    /// The inputs for a monitoring job.
    /// </summary>
    [OutputType]
    public sealed class ModelExplainabilityJobDefinitionModelExplainabilityJobInput
    {
        /// <summary>
        /// Input object for the batch transform job.
        /// </summary>
        public readonly Outputs.ModelExplainabilityJobDefinitionBatchTransformInput? BatchTransformInput;
        /// <summary>
        /// Input object for the endpoint
        /// </summary>
        public readonly Outputs.ModelExplainabilityJobDefinitionEndpointInput? EndpointInput;

        [OutputConstructor]
        private ModelExplainabilityJobDefinitionModelExplainabilityJobInput(
            Outputs.ModelExplainabilityJobDefinitionBatchTransformInput? batchTransformInput,

            Outputs.ModelExplainabilityJobDefinitionEndpointInput? endpointInput)
        {
            BatchTransformInput = batchTransformInput;
            EndpointInput = endpointInput;
        }
    }
}
