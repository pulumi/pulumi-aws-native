// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetModelBiasJobDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition
        /// </summary>
        public static Task<GetModelBiasJobDefinitionResult> InvokeAsync(GetModelBiasJobDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetModelBiasJobDefinitionResult>("aws-native:sagemaker:getModelBiasJobDefinition", args ?? new GetModelBiasJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition
        /// </summary>
        public static Output<GetModelBiasJobDefinitionResult> Invoke(GetModelBiasJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetModelBiasJobDefinitionResult>("aws-native:sagemaker:getModelBiasJobDefinition", args ?? new GetModelBiasJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelBiasJobDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public string JobDefinitionArn { get; set; } = null!;

        public GetModelBiasJobDefinitionArgs()
        {
        }
        public static new GetModelBiasJobDefinitionArgs Empty => new GetModelBiasJobDefinitionArgs();
    }

    public sealed class GetModelBiasJobDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public Input<string> JobDefinitionArn { get; set; } = null!;

        public GetModelBiasJobDefinitionInvokeArgs()
        {
        }
        public static new GetModelBiasJobDefinitionInvokeArgs Empty => new GetModelBiasJobDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetModelBiasJobDefinitionResult
    {
        /// <summary>
        /// The time at which the job definition was created.
        /// </summary>
        public readonly string? CreationTime;
        public readonly string? EndpointName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        public readonly string? JobDefinitionArn;

        [OutputConstructor]
        private GetModelBiasJobDefinitionResult(
            string? creationTime,

            string? endpointName,

            string? jobDefinitionArn)
        {
            CreationTime = creationTime;
            EndpointName = endpointName;
            JobDefinitionArn = jobDefinitionArn;
        }
    }
}