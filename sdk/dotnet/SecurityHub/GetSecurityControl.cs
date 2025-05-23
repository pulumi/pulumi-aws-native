// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetSecurityControl
    {
        /// <summary>
        /// A security control in Security Hub describes a security best practice related to a specific resource.
        /// </summary>
        public static Task<GetSecurityControlResult> InvokeAsync(GetSecurityControlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityControlResult>("aws-native:securityhub:getSecurityControl", args ?? new GetSecurityControlArgs(), options.WithDefaults());

        /// <summary>
        /// A security control in Security Hub describes a security best practice related to a specific resource.
        /// </summary>
        public static Output<GetSecurityControlResult> Invoke(GetSecurityControlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityControlResult>("aws-native:securityhub:getSecurityControl", args ?? new GetSecurityControlInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A security control in Security Hub describes a security best practice related to a specific resource.
        /// </summary>
        public static Output<GetSecurityControlResult> Invoke(GetSecurityControlInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityControlResult>("aws-native:securityhub:getSecurityControl", args ?? new GetSecurityControlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityControlArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
        /// </summary>
        [Input("securityControlId", required: true)]
        public string SecurityControlId { get; set; } = null!;

        public GetSecurityControlArgs()
        {
        }
        public static new GetSecurityControlArgs Empty => new GetSecurityControlArgs();
    }

    public sealed class GetSecurityControlInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
        /// </summary>
        [Input("securityControlId", required: true)]
        public Input<string> SecurityControlId { get; set; } = null!;

        public GetSecurityControlInvokeArgs()
        {
        }
        public static new GetSecurityControlInvokeArgs Empty => new GetSecurityControlInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityControlResult
    {
        /// <summary>
        /// The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
        /// </summary>
        public readonly string? LastUpdateReason;
        /// <summary>
        /// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.SecurityControlParameterConfiguration>? Parameters;
        /// <summary>
        /// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.
        /// </summary>
        public readonly string? SecurityControlArn;

        [OutputConstructor]
        private GetSecurityControlResult(
            string? lastUpdateReason,

            ImmutableDictionary<string, Outputs.SecurityControlParameterConfiguration>? parameters,

            string? securityControlArn)
        {
            LastUpdateReason = lastUpdateReason;
            Parameters = parameters;
            SecurityControlArn = securityControlArn;
        }
    }
}
