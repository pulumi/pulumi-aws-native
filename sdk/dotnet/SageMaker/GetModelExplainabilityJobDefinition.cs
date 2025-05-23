// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetModelExplainabilityJobDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
        /// </summary>
        public static Task<GetModelExplainabilityJobDefinitionResult> InvokeAsync(GetModelExplainabilityJobDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetModelExplainabilityJobDefinitionResult>("aws-native:sagemaker:getModelExplainabilityJobDefinition", args ?? new GetModelExplainabilityJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
        /// </summary>
        public static Output<GetModelExplainabilityJobDefinitionResult> Invoke(GetModelExplainabilityJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetModelExplainabilityJobDefinitionResult>("aws-native:sagemaker:getModelExplainabilityJobDefinition", args ?? new GetModelExplainabilityJobDefinitionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
        /// </summary>
        public static Output<GetModelExplainabilityJobDefinitionResult> Invoke(GetModelExplainabilityJobDefinitionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetModelExplainabilityJobDefinitionResult>("aws-native:sagemaker:getModelExplainabilityJobDefinition", args ?? new GetModelExplainabilityJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelExplainabilityJobDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public string JobDefinitionArn { get; set; } = null!;

        public GetModelExplainabilityJobDefinitionArgs()
        {
        }
        public static new GetModelExplainabilityJobDefinitionArgs Empty => new GetModelExplainabilityJobDefinitionArgs();
    }

    public sealed class GetModelExplainabilityJobDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public Input<string> JobDefinitionArn { get; set; } = null!;

        public GetModelExplainabilityJobDefinitionInvokeArgs()
        {
        }
        public static new GetModelExplainabilityJobDefinitionInvokeArgs Empty => new GetModelExplainabilityJobDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetModelExplainabilityJobDefinitionResult
    {
        /// <summary>
        /// The time at which the job definition was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        public readonly string? JobDefinitionArn;

        [OutputConstructor]
        private GetModelExplainabilityJobDefinitionResult(
            string? creationTime,

            string? jobDefinitionArn)
        {
            CreationTime = creationTime;
            JobDefinitionArn = jobDefinitionArn;
        }
    }
}
