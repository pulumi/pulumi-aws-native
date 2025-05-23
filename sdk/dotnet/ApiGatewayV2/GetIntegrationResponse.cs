// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    public static class GetIntegrationResponse
    {
        /// <summary>
        /// The ``AWS::ApiGatewayV2::IntegrationResponse`` resource updates an integration response for an WebSocket API. For more information, see [Set up WebSocket API Integration Responses in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-responses.html) in the *API Gateway Developer Guide*.
        /// </summary>
        public static Task<GetIntegrationResponseResult> InvokeAsync(GetIntegrationResponseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationResponseResult>("aws-native:apigatewayv2:getIntegrationResponse", args ?? new GetIntegrationResponseArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGatewayV2::IntegrationResponse`` resource updates an integration response for an WebSocket API. For more information, see [Set up WebSocket API Integration Responses in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-responses.html) in the *API Gateway Developer Guide*.
        /// </summary>
        public static Output<GetIntegrationResponseResult> Invoke(GetIntegrationResponseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationResponseResult>("aws-native:apigatewayv2:getIntegrationResponse", args ?? new GetIntegrationResponseInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGatewayV2::IntegrationResponse`` resource updates an integration response for an WebSocket API. For more information, see [Set up WebSocket API Integration Responses in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-responses.html) in the *API Gateway Developer Guide*.
        /// </summary>
        public static Output<GetIntegrationResponseResult> Invoke(GetIntegrationResponseInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationResponseResult>("aws-native:apigatewayv2:getIntegrationResponse", args ?? new GetIntegrationResponseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIntegrationResponseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// The integration ID.
        /// </summary>
        [Input("integrationId", required: true)]
        public string IntegrationId { get; set; } = null!;

        /// <summary>
        /// The integration response ID.
        /// </summary>
        [Input("integrationResponseId", required: true)]
        public string IntegrationResponseId { get; set; } = null!;

        public GetIntegrationResponseArgs()
        {
        }
        public static new GetIntegrationResponseArgs Empty => new GetIntegrationResponseArgs();
    }

    public sealed class GetIntegrationResponseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The integration ID.
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// The integration response ID.
        /// </summary>
        [Input("integrationResponseId", required: true)]
        public Input<string> IntegrationResponseId { get; set; } = null!;

        public GetIntegrationResponseInvokeArgs()
        {
        }
        public static new GetIntegrationResponseInvokeArgs Empty => new GetIntegrationResponseInvokeArgs();
    }


    [OutputType]
    public sealed class GetIntegrationResponseResult
    {
        /// <summary>
        /// Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are ``CONVERT_TO_BINARY`` and ``CONVERT_TO_TEXT``, with the following behaviors:
        ///   ``CONVERT_TO_BINARY``: Converts a response payload from a Base64-encoded string to the corresponding binary blob.
        ///   ``CONVERT_TO_TEXT``: Converts a response payload from a binary blob to a Base64-encoded string.
        ///  If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.
        /// </summary>
        public readonly string? ContentHandlingStrategy;
        /// <summary>
        /// The integration response ID.
        /// </summary>
        public readonly string? IntegrationResponseId;
        /// <summary>
        /// The integration response key.
        /// </summary>
        public readonly string? IntegrationResponseKey;
        /// <summary>
        /// A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ``method.response.header.{name}``, where name is a valid and unique header name. The mapped non-static value must match the pattern of ``integration.response.header.{name}`` or ``integration.response.body.{JSON-expression}``, where ``{name}`` is a valid and unique response header name and ``{JSON-expression}`` is a valid JSON expression without the ``$`` prefix.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::IntegrationResponse` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? ResponseParameters;
        /// <summary>
        /// The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::IntegrationResponse` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? ResponseTemplates;
        /// <summary>
        /// The template selection expression for the integration response. Supported only for WebSocket APIs.
        /// </summary>
        public readonly string? TemplateSelectionExpression;

        [OutputConstructor]
        private GetIntegrationResponseResult(
            string? contentHandlingStrategy,

            string? integrationResponseId,

            string? integrationResponseKey,

            object? responseParameters,

            object? responseTemplates,

            string? templateSelectionExpression)
        {
            ContentHandlingStrategy = contentHandlingStrategy;
            IntegrationResponseId = integrationResponseId;
            IntegrationResponseKey = integrationResponseKey;
            ResponseParameters = responseParameters;
            ResponseTemplates = responseTemplates;
            TemplateSelectionExpression = templateSelectionExpression;
        }
    }
}
