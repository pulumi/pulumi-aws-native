// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Comprehend
{
    public static class GetDocumentClassifier
    {
        /// <summary>
        /// Document Classifier enables training document classifier models.
        /// </summary>
        public static Task<GetDocumentClassifierResult> InvokeAsync(GetDocumentClassifierArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDocumentClassifierResult>("aws-native:comprehend:getDocumentClassifier", args ?? new GetDocumentClassifierArgs(), options.WithDefaults());

        /// <summary>
        /// Document Classifier enables training document classifier models.
        /// </summary>
        public static Output<GetDocumentClassifierResult> Invoke(GetDocumentClassifierInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDocumentClassifierResult>("aws-native:comprehend:getDocumentClassifier", args ?? new GetDocumentClassifierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDocumentClassifierArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetDocumentClassifierArgs()
        {
        }
        public static new GetDocumentClassifierArgs Empty => new GetDocumentClassifierArgs();
    }

    public sealed class GetDocumentClassifierInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetDocumentClassifierInvokeArgs()
        {
        }
        public static new GetDocumentClassifierInvokeArgs Empty => new GetDocumentClassifierInvokeArgs();
    }


    [OutputType]
    public sealed class GetDocumentClassifierResult
    {
        public readonly string? Arn;
        public readonly string? ModelPolicy;
        public readonly ImmutableArray<Outputs.DocumentClassifierTag> Tags;

        [OutputConstructor]
        private GetDocumentClassifierResult(
            string? arn,

            string? modelPolicy,

            ImmutableArray<Outputs.DocumentClassifierTag> tags)
        {
            Arn = arn;
            ModelPolicy = modelPolicy;
            Tags = tags;
        }
    }
}