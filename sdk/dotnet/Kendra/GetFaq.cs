// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra
{
    public static class GetFaq
    {
        /// <summary>
        /// A Kendra FAQ resource
        /// </summary>
        public static Task<GetFaqResult> InvokeAsync(GetFaqArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFaqResult>("aws-native:kendra:getFaq", args ?? new GetFaqArgs(), options.WithDefaults());

        /// <summary>
        /// A Kendra FAQ resource
        /// </summary>
        public static Output<GetFaqResult> Invoke(GetFaqInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFaqResult>("aws-native:kendra:getFaq", args ?? new GetFaqInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A Kendra FAQ resource
        /// </summary>
        public static Output<GetFaqResult> Invoke(GetFaqInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFaqResult>("aws-native:kendra:getFaq", args ?? new GetFaqInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFaqArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the FAQ. For example:
        /// 
        /// `f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// Index ID
        /// </summary>
        [Input("indexId", required: true)]
        public string IndexId { get; set; } = null!;

        public GetFaqArgs()
        {
        }
        public static new GetFaqArgs Empty => new GetFaqArgs();
    }

    public sealed class GetFaqInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the FAQ. For example:
        /// 
        /// `f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Index ID
        /// </summary>
        [Input("indexId", required: true)]
        public Input<string> IndexId { get; set; } = null!;

        public GetFaqInvokeArgs()
        {
        }
        public static new GetFaqInvokeArgs Empty => new GetFaqInvokeArgs();
    }


    [OutputType]
    public sealed class GetFaqResult
    {
        /// <summary>
        /// `arn:aws:kendra:us-west-2:111122223333:index/335c3741-41df-46a6-b5d3-61f85b787884/faq/f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The identifier for the FAQ. For example:
        /// 
        /// `f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The code for a language. This shows a supported language for the FAQ document as part of the summary information for FAQs. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html) .
        /// </summary>
        public readonly string? LanguageCode;
        /// <summary>
        /// Tags for labeling the FAQ
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetFaqResult(
            string? arn,

            string? id,

            string? languageCode,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Id = id;
            LanguageCode = languageCode;
            Tags = tags;
        }
    }
}
