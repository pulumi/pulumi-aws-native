// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Inspector
{
    public static class GetAssessmentTemplate
    {
        /// <summary>
        /// Resource Type definition for AWS::Inspector::AssessmentTemplate
        /// </summary>
        public static Task<GetAssessmentTemplateResult> InvokeAsync(GetAssessmentTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssessmentTemplateResult>("aws-native:inspector:getAssessmentTemplate", args ?? new GetAssessmentTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Inspector::AssessmentTemplate
        /// </summary>
        public static Output<GetAssessmentTemplateResult> Invoke(GetAssessmentTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssessmentTemplateResult>("aws-native:inspector:getAssessmentTemplate", args ?? new GetAssessmentTemplateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Inspector::AssessmentTemplate
        /// </summary>
        public static Output<GetAssessmentTemplateResult> Invoke(GetAssessmentTemplateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssessmentTemplateResult>("aws-native:inspector:getAssessmentTemplate", args ?? new GetAssessmentTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssessmentTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that specifies the assessment template that is created.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetAssessmentTemplateArgs()
        {
        }
        public static new GetAssessmentTemplateArgs Empty => new GetAssessmentTemplateArgs();
    }

    public sealed class GetAssessmentTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that specifies the assessment template that is created.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetAssessmentTemplateInvokeArgs()
        {
        }
        public static new GetAssessmentTemplateInvokeArgs Empty => new GetAssessmentTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetAssessmentTemplateResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that specifies the assessment template that is created.
        /// </summary>
        public readonly string? Arn;

        [OutputConstructor]
        private GetAssessmentTemplateResult(string? arn)
        {
            Arn = arn;
        }
    }
}
