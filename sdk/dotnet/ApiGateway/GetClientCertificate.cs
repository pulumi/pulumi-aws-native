// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetClientCertificate
    {
        /// <summary>
        /// The ``AWS::ApiGateway::ClientCertificate`` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.
        /// </summary>
        public static Task<GetClientCertificateResult> InvokeAsync(GetClientCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientCertificateResult>("aws-native:apigateway:getClientCertificate", args ?? new GetClientCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGateway::ClientCertificate`` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.
        /// </summary>
        public static Output<GetClientCertificateResult> Invoke(GetClientCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientCertificateResult>("aws-native:apigateway:getClientCertificate", args ?? new GetClientCertificateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGateway::ClientCertificate`` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.
        /// </summary>
        public static Output<GetClientCertificateResult> Invoke(GetClientCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientCertificateResult>("aws-native:apigateway:getClientCertificate", args ?? new GetClientCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClientCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID for the client certificate. For example: `abc123` .
        /// </summary>
        [Input("clientCertificateId", required: true)]
        public string ClientCertificateId { get; set; } = null!;

        public GetClientCertificateArgs()
        {
        }
        public static new GetClientCertificateArgs Empty => new GetClientCertificateArgs();
    }

    public sealed class GetClientCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID for the client certificate. For example: `abc123` .
        /// </summary>
        [Input("clientCertificateId", required: true)]
        public Input<string> ClientCertificateId { get; set; } = null!;

        public GetClientCertificateInvokeArgs()
        {
        }
        public static new GetClientCertificateInvokeArgs Empty => new GetClientCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetClientCertificateResult
    {
        /// <summary>
        /// The ID for the client certificate. For example: `abc123` .
        /// </summary>
        public readonly string? ClientCertificateId;
        /// <summary>
        /// The description of the client certificate.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The collection of tags. Each tag element is associated with a given resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetClientCertificateResult(
            string? clientCertificateId,

            string? description,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            ClientCertificateId = clientCertificateId;
            Description = description;
            Tags = tags;
        }
    }
}
