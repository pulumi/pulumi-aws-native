// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetDocumentationVersion
    {
        /// <summary>
        /// The ``AWS::ApiGateway::DocumentationVersion`` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.
        /// </summary>
        public static Task<GetDocumentationVersionResult> InvokeAsync(GetDocumentationVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDocumentationVersionResult>("aws-native:apigateway:getDocumentationVersion", args ?? new GetDocumentationVersionArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGateway::DocumentationVersion`` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.
        /// </summary>
        public static Output<GetDocumentationVersionResult> Invoke(GetDocumentationVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDocumentationVersionResult>("aws-native:apigateway:getDocumentationVersion", args ?? new GetDocumentationVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGateway::DocumentationVersion`` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.
        /// </summary>
        public static Output<GetDocumentationVersionResult> Invoke(GetDocumentationVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDocumentationVersionResult>("aws-native:apigateway:getDocumentationVersion", args ?? new GetDocumentationVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDocumentationVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The version identifier of the to-be-updated documentation version.
        /// </summary>
        [Input("documentationVersion", required: true)]
        public string DocumentationVersionValue { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public string RestApiId { get; set; } = null!;

        public GetDocumentationVersionArgs()
        {
        }
        public static new GetDocumentationVersionArgs Empty => new GetDocumentationVersionArgs();
    }

    public sealed class GetDocumentationVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The version identifier of the to-be-updated documentation version.
        /// </summary>
        [Input("documentationVersion", required: true)]
        public Input<string> DocumentationVersion { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public GetDocumentationVersionInvokeArgs()
        {
        }
        public static new GetDocumentationVersionInvokeArgs Empty => new GetDocumentationVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDocumentationVersionResult
    {
        /// <summary>
        /// A description about the new documentation snapshot.
        /// </summary>
        public readonly string? Description;

        [OutputConstructor]
        private GetDocumentationVersionResult(string? description)
        {
            Description = description;
        }
    }
}
