// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    /// <summary>
    /// The ``AWS::ApiGatewayV2::RouteResponse`` resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigatewayv2:RouteResponse")]
    public partial class RouteResponse : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The model selection expression for the route response. Supported only for WebSocket APIs.
        /// </summary>
        [Output("modelSelectionExpression")]
        public Output<string?> ModelSelectionExpression { get; private set; } = null!;

        /// <summary>
        /// The response models for the route response.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
        /// </summary>
        [Output("responseModels")]
        public Output<object?> ResponseModels { get; private set; } = null!;

        /// <summary>
        /// The route response parameters.
        /// </summary>
        [Output("responseParameters")]
        public Output<ImmutableDictionary<string, Outputs.RouteResponseParameterConstraints>?> ResponseParameters { get; private set; } = null!;

        /// <summary>
        /// The route ID.
        /// </summary>
        [Output("routeId")]
        public Output<string> RouteId { get; private set; } = null!;

        /// <summary>
        /// The route response ID.
        /// </summary>
        [Output("routeResponseId")]
        public Output<string> RouteResponseId { get; private set; } = null!;

        /// <summary>
        /// The route response key.
        /// </summary>
        [Output("routeResponseKey")]
        public Output<string> RouteResponseKey { get; private set; } = null!;


        /// <summary>
        /// Create a RouteResponse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteResponse(string name, RouteResponseArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigatewayv2:RouteResponse", name, args ?? new RouteResponseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteResponse(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigatewayv2:RouteResponse", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiId",
                    "routeId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteResponse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteResponse Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RouteResponse(name, id, options);
        }
    }

    public sealed class RouteResponseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The model selection expression for the route response. Supported only for WebSocket APIs.
        /// </summary>
        [Input("modelSelectionExpression")]
        public Input<string>? ModelSelectionExpression { get; set; }

        /// <summary>
        /// The response models for the route response.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
        /// </summary>
        [Input("responseModels")]
        public Input<object>? ResponseModels { get; set; }

        [Input("responseParameters")]
        private InputMap<Inputs.RouteResponseParameterConstraintsArgs>? _responseParameters;

        /// <summary>
        /// The route response parameters.
        /// </summary>
        public InputMap<Inputs.RouteResponseParameterConstraintsArgs> ResponseParameters
        {
            get => _responseParameters ?? (_responseParameters = new InputMap<Inputs.RouteResponseParameterConstraintsArgs>());
            set => _responseParameters = value;
        }

        /// <summary>
        /// The route ID.
        /// </summary>
        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        /// <summary>
        /// The route response key.
        /// </summary>
        [Input("routeResponseKey", required: true)]
        public Input<string> RouteResponseKey { get; set; } = null!;

        public RouteResponseArgs()
        {
        }
        public static new RouteResponseArgs Empty => new RouteResponseArgs();
    }
}
