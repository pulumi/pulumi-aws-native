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
    /// Additional Inference Specification specifies details about inference jobs that can be run with models based on this model package.AdditionalInferenceSpecifications can be added to existing model packages using AdditionalInferenceSpecificationsToAdd.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageAdditionalInferenceSpecificationDefinition
    {
        /// <summary>
        /// The Amazon ECR registry path of the Docker image that contains the inference code.
        /// </summary>
        public readonly ImmutableArray<Outputs.ModelPackageContainerDefinition> Containers;
        /// <summary>
        /// A description of the additional Inference specification.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A unique name to identify the additional inference specification. The name must be unique within the list of your additional inference specifications for a particular model package.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The supported MIME types for the input data.
        /// </summary>
        public readonly ImmutableArray<string> SupportedContentTypes;
        /// <summary>
        /// A list of the instance types that are used to generate inferences in real-time
        /// </summary>
        public readonly ImmutableArray<string> SupportedRealtimeInferenceInstanceTypes;
        /// <summary>
        /// The supported MIME types for the output data.
        /// </summary>
        public readonly ImmutableArray<string> SupportedResponseMimeTypes;
        /// <summary>
        /// A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.
        /// </summary>
        public readonly ImmutableArray<string> SupportedTransformInstanceTypes;

        [OutputConstructor]
        private ModelPackageAdditionalInferenceSpecificationDefinition(
            ImmutableArray<Outputs.ModelPackageContainerDefinition> containers,

            string? description,

            string name,

            ImmutableArray<string> supportedContentTypes,

            ImmutableArray<string> supportedRealtimeInferenceInstanceTypes,

            ImmutableArray<string> supportedResponseMimeTypes,

            ImmutableArray<string> supportedTransformInstanceTypes)
        {
            Containers = containers;
            Description = description;
            Name = name;
            SupportedContentTypes = supportedContentTypes;
            SupportedRealtimeInferenceInstanceTypes = supportedRealtimeInferenceInstanceTypes;
            SupportedResponseMimeTypes = supportedResponseMimeTypes;
            SupportedTransformInstanceTypes = supportedTransformInstanceTypes;
        }
    }
}
