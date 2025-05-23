// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms
{
    public static class GetAnalysisTemplate
    {
        /// <summary>
        /// Represents a stored analysis within a collaboration
        /// </summary>
        public static Task<GetAnalysisTemplateResult> InvokeAsync(GetAnalysisTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAnalysisTemplateResult>("aws-native:cleanrooms:getAnalysisTemplate", args ?? new GetAnalysisTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a stored analysis within a collaboration
        /// </summary>
        public static Output<GetAnalysisTemplateResult> Invoke(GetAnalysisTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAnalysisTemplateResult>("aws-native:cleanrooms:getAnalysisTemplate", args ?? new GetAnalysisTemplateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a stored analysis within a collaboration
        /// </summary>
        public static Output<GetAnalysisTemplateResult> Invoke(GetAnalysisTemplateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAnalysisTemplateResult>("aws-native:cleanrooms:getAnalysisTemplate", args ?? new GetAnalysisTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAnalysisTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the identifier for the analysis template.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
        /// </summary>
        [Input("analysisTemplateIdentifier", required: true)]
        public string AnalysisTemplateIdentifier { get; set; } = null!;

        /// <summary>
        /// The identifier for a membership resource.
        /// </summary>
        [Input("membershipIdentifier", required: true)]
        public string MembershipIdentifier { get; set; } = null!;

        public GetAnalysisTemplateArgs()
        {
        }
        public static new GetAnalysisTemplateArgs Empty => new GetAnalysisTemplateArgs();
    }

    public sealed class GetAnalysisTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the identifier for the analysis template.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
        /// </summary>
        [Input("analysisTemplateIdentifier", required: true)]
        public Input<string> AnalysisTemplateIdentifier { get; set; } = null!;

        /// <summary>
        /// The identifier for a membership resource.
        /// </summary>
        [Input("membershipIdentifier", required: true)]
        public Input<string> MembershipIdentifier { get; set; } = null!;

        public GetAnalysisTemplateInvokeArgs()
        {
        }
        public static new GetAnalysisTemplateInvokeArgs Empty => new GetAnalysisTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetAnalysisTemplateResult
    {
        /// <summary>
        /// Returns the identifier for the analysis template.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
        /// </summary>
        public readonly string? AnalysisTemplateIdentifier;
        /// <summary>
        /// Returns the Amazon Resource Name (ARN) of the analysis template.
        /// 
        /// Example: `arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/analysistemplates/a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Returns the unique ARN for the analysis template’s associated collaboration.
        /// 
        /// Example: `arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
        /// </summary>
        public readonly string? CollaborationArn;
        /// <summary>
        /// Returns the unique ID for the associated collaboration of the analysis template.
        /// 
        /// Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
        /// </summary>
        public readonly string? CollaborationIdentifier;
        /// <summary>
        /// The description of the analysis template.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Returns the Amazon Resource Name (ARN) of the member who created the analysis template.
        /// 
        /// Example: `arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        /// </summary>
        public readonly string? MembershipArn;
        /// <summary>
        /// The source metadata for the analysis template.
        /// </summary>
        public readonly Outputs.AnalysisTemplateAnalysisSourceMetadataProperties? SourceMetadata;
        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetAnalysisTemplateResult(
            string? analysisTemplateIdentifier,

            string? arn,

            string? collaborationArn,

            string? collaborationIdentifier,

            string? description,

            string? membershipArn,

            Outputs.AnalysisTemplateAnalysisSourceMetadataProperties? sourceMetadata,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AnalysisTemplateIdentifier = analysisTemplateIdentifier;
            Arn = arn;
            CollaborationArn = collaborationArn;
            CollaborationIdentifier = collaborationIdentifier;
            Description = description;
            MembershipArn = membershipArn;
            SourceMetadata = sourceMetadata;
            Tags = tags;
        }
    }
}
