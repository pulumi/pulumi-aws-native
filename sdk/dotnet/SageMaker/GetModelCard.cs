// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetModelCard
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelCard.
        /// </summary>
        public static Task<GetModelCardResult> InvokeAsync(GetModelCardArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetModelCardResult>("aws-native:sagemaker:getModelCard", args ?? new GetModelCardArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelCard.
        /// </summary>
        public static Output<GetModelCardResult> Invoke(GetModelCardInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetModelCardResult>("aws-native:sagemaker:getModelCard", args ?? new GetModelCardInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelCard.
        /// </summary>
        public static Output<GetModelCardResult> Invoke(GetModelCardInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetModelCardResult>("aws-native:sagemaker:getModelCard", args ?? new GetModelCardInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelCardArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the model card.
        /// </summary>
        [Input("modelCardName", required: true)]
        public string ModelCardName { get; set; } = null!;

        public GetModelCardArgs()
        {
        }
        public static new GetModelCardArgs Empty => new GetModelCardArgs();
    }

    public sealed class GetModelCardInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the model card.
        /// </summary>
        [Input("modelCardName", required: true)]
        public Input<string> ModelCardName { get; set; } = null!;

        public GetModelCardInvokeArgs()
        {
        }
        public static new GetModelCardInvokeArgs Empty => new GetModelCardInvokeArgs();
    }


    [OutputType]
    public sealed class GetModelCardResult
    {
        /// <summary>
        /// The content of the model card. Content uses the [model card JSON schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
        /// </summary>
        public readonly Outputs.ModelCardContent? Content;
        /// <summary>
        /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
        /// </summary>
        public readonly Outputs.ModelCardUserContext? CreatedBy;
        /// <summary>
        /// The date and time the model card was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
        /// </summary>
        public readonly Outputs.ModelCardUserContext? LastModifiedBy;
        /// <summary>
        /// The date and time the model card was last modified.
        /// </summary>
        public readonly string? LastModifiedTime;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the successfully created model card.
        /// </summary>
        public readonly string? ModelCardArn;
        /// <summary>
        /// The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.ModelCardProcessingStatus? ModelCardProcessingStatus;
        /// <summary>
        /// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.ModelCardStatus? ModelCardStatus;
        /// <summary>
        /// A version of the model card.
        /// </summary>
        public readonly int? ModelCardVersion;
        /// <summary>
        /// Key-value pairs used to manage metadata for model cards.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetModelCardResult(
            Outputs.ModelCardContent? content,

            Outputs.ModelCardUserContext? createdBy,

            string? creationTime,

            Outputs.ModelCardUserContext? lastModifiedBy,

            string? lastModifiedTime,

            string? modelCardArn,

            Pulumi.AwsNative.SageMaker.ModelCardProcessingStatus? modelCardProcessingStatus,

            Pulumi.AwsNative.SageMaker.ModelCardStatus? modelCardStatus,

            int? modelCardVersion,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Content = content;
            CreatedBy = createdBy;
            CreationTime = creationTime;
            LastModifiedBy = lastModifiedBy;
            LastModifiedTime = lastModifiedTime;
            ModelCardArn = modelCardArn;
            ModelCardProcessingStatus = modelCardProcessingStatus;
            ModelCardStatus = modelCardStatus;
            ModelCardVersion = modelCardVersion;
            Tags = tags;
        }
    }
}
