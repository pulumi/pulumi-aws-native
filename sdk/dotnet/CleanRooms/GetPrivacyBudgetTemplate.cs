// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms
{
    public static class GetPrivacyBudgetTemplate
    {
        /// <summary>
        /// Represents a privacy budget within a collaboration
        /// </summary>
        public static Task<GetPrivacyBudgetTemplateResult> InvokeAsync(GetPrivacyBudgetTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivacyBudgetTemplateResult>("aws-native:cleanrooms:getPrivacyBudgetTemplate", args ?? new GetPrivacyBudgetTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a privacy budget within a collaboration
        /// </summary>
        public static Output<GetPrivacyBudgetTemplateResult> Invoke(GetPrivacyBudgetTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivacyBudgetTemplateResult>("aws-native:cleanrooms:getPrivacyBudgetTemplate", args ?? new GetPrivacyBudgetTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivacyBudgetTemplateArgs : global::Pulumi.InvokeArgs
    {
        [Input("membershipIdentifier", required: true)]
        public string MembershipIdentifier { get; set; } = null!;

        [Input("privacyBudgetTemplateIdentifier", required: true)]
        public string PrivacyBudgetTemplateIdentifier { get; set; } = null!;

        public GetPrivacyBudgetTemplateArgs()
        {
        }
        public static new GetPrivacyBudgetTemplateArgs Empty => new GetPrivacyBudgetTemplateArgs();
    }

    public sealed class GetPrivacyBudgetTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("membershipIdentifier", required: true)]
        public Input<string> MembershipIdentifier { get; set; } = null!;

        [Input("privacyBudgetTemplateIdentifier", required: true)]
        public Input<string> PrivacyBudgetTemplateIdentifier { get; set; } = null!;

        public GetPrivacyBudgetTemplateInvokeArgs()
        {
        }
        public static new GetPrivacyBudgetTemplateInvokeArgs Empty => new GetPrivacyBudgetTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivacyBudgetTemplateResult
    {
        public readonly string? Arn;
        public readonly string? CollaborationArn;
        public readonly string? CollaborationIdentifier;
        public readonly string? MembershipArn;
        public readonly Outputs.ParametersProperties? Parameters;
        public readonly string? PrivacyBudgetTemplateIdentifier;
        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetPrivacyBudgetTemplateResult(
            string? arn,

            string? collaborationArn,

            string? collaborationIdentifier,

            string? membershipArn,

            Outputs.ParametersProperties? parameters,

            string? privacyBudgetTemplateIdentifier,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            CollaborationArn = collaborationArn;
            CollaborationIdentifier = collaborationIdentifier;
            MembershipArn = membershipArn;
            Parameters = parameters;
            PrivacyBudgetTemplateIdentifier = privacyBudgetTemplateIdentifier;
            Tags = tags;
        }
    }
}