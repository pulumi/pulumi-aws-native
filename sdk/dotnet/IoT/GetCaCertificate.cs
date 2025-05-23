// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetCaCertificate
    {
        /// <summary>
        /// Registers a CA Certificate in IoT.
        /// </summary>
        public static Task<GetCaCertificateResult> InvokeAsync(GetCaCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCaCertificateResult>("aws-native:iot:getCaCertificate", args ?? new GetCaCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Registers a CA Certificate in IoT.
        /// </summary>
        public static Output<GetCaCertificateResult> Invoke(GetCaCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCaCertificateResult>("aws-native:iot:getCaCertificate", args ?? new GetCaCertificateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Registers a CA Certificate in IoT.
        /// </summary>
        public static Output<GetCaCertificateResult> Invoke(GetCaCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCaCertificateResult>("aws-native:iot:getCaCertificate", args ?? new GetCaCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCaCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CA certificate ID.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCaCertificateArgs()
        {
        }
        public static new GetCaCertificateArgs Empty => new GetCaCertificateArgs();
    }

    public sealed class GetCaCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CA certificate ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCaCertificateInvokeArgs()
        {
        }
        public static new GetCaCertificateInvokeArgs Empty => new GetCaCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCaCertificateResult
    {
        /// <summary>
        /// Returns the Amazon Resource Name (ARN) for the CA certificate. For example:
        /// 
        /// `{ "Fn::GetAtt": ["MyCACertificate", "Arn"] }`
        /// 
        /// A value similar to the following is returned:
        /// 
        /// `arn:aws:iot:us-east-1:123456789012:cacert/a6be6b84559801927e35a8f901fae08b5971d78d1562e29504ff9663b276a5f5`
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Whether the CA certificate is configured for auto registration of device certificates. Valid values are "ENABLE" and "DISABLE".
        /// </summary>
        public readonly Pulumi.AwsNative.IoT.CaCertificateAutoRegistrationStatus? AutoRegistrationStatus;
        /// <summary>
        /// The CA certificate ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Information about the registration configuration.
        /// </summary>
        public readonly Outputs.CaCertificateRegistrationConfig? RegistrationConfig;
        /// <summary>
        /// The status of the CA certificate.
        /// 
        /// Valid values are "ACTIVE" and "INACTIVE".
        /// </summary>
        public readonly Pulumi.AwsNative.IoT.CaCertificateStatus? Status;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetCaCertificateResult(
            string? arn,

            Pulumi.AwsNative.IoT.CaCertificateAutoRegistrationStatus? autoRegistrationStatus,

            string? id,

            Outputs.CaCertificateRegistrationConfig? registrationConfig,

            Pulumi.AwsNative.IoT.CaCertificateStatus? status,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            AutoRegistrationStatus = autoRegistrationStatus;
            Id = id;
            RegistrationConfig = registrationConfig;
            Status = status;
            Tags = tags;
        }
    }
}
