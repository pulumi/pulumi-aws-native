// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetDataQualityJobDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
        /// </summary>
        public static Task<GetDataQualityJobDefinitionResult> InvokeAsync(GetDataQualityJobDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataQualityJobDefinitionResult>("aws-native:sagemaker:getDataQualityJobDefinition", args ?? new GetDataQualityJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
        /// </summary>
        public static Output<GetDataQualityJobDefinitionResult> Invoke(GetDataQualityJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataQualityJobDefinitionResult>("aws-native:sagemaker:getDataQualityJobDefinition", args ?? new GetDataQualityJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataQualityJobDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public string JobDefinitionArn { get; set; } = null!;

        public GetDataQualityJobDefinitionArgs()
        {
        }
        public static new GetDataQualityJobDefinitionArgs Empty => new GetDataQualityJobDefinitionArgs();
    }

    public sealed class GetDataQualityJobDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public Input<string> JobDefinitionArn { get; set; } = null!;

        public GetDataQualityJobDefinitionInvokeArgs()
        {
        }
        public static new GetDataQualityJobDefinitionInvokeArgs Empty => new GetDataQualityJobDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataQualityJobDefinitionResult
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
        private GetDataQualityJobDefinitionResult(
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