// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    /// <summary>
    /// Resource Type definition for AWS::ApiGateway::Method
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:Method")]
    public partial class Method : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the method requires clients to submit a valid API key.
        /// </summary>
        [Output("apiKeyRequired")]
        public Output<bool?> ApiKeyRequired { get; private set; } = null!;

        /// <summary>
        /// A list of authorization scopes configured on the method.
        /// </summary>
        [Output("authorizationScopes")]
        public Output<ImmutableArray<string>> AuthorizationScopes { get; private set; } = null!;

        /// <summary>
        /// The method's authorization type.
        /// </summary>
        [Output("authorizationType")]
        public Output<Pulumi.AwsNative.ApiGateway.MethodAuthorizationType?> AuthorizationType { get; private set; } = null!;

        /// <summary>
        /// The identifier of the authorizer to use on this method.
        /// </summary>
        [Output("authorizerId")]
        public Output<string?> AuthorizerId { get; private set; } = null!;

        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        [Output("httpMethod")]
        public Output<string> HttpMethod { get; private set; } = null!;

        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        [Output("integration")]
        public Output<Outputs.MethodIntegration?> Integration { get; private set; } = null!;

        /// <summary>
        /// The responses that can be sent to the client who calls the method.
        /// </summary>
        [Output("methodResponses")]
        public Output<ImmutableArray<Outputs.MethodResponse>> MethodResponses { get; private set; } = null!;

        /// <summary>
        /// A friendly operation name for the method.
        /// </summary>
        [Output("operationName")]
        public Output<string?> OperationName { get; private set; } = null!;

        /// <summary>
        /// The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
        /// </summary>
        [Output("requestModels")]
        public Output<object?> RequestModels { get; private set; } = null!;

        /// <summary>
        /// The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
        /// </summary>
        [Output("requestParameters")]
        public Output<object?> RequestParameters { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated request validator.
        /// </summary>
        [Output("requestValidatorId")]
        public Output<string?> RequestValidatorId { get; private set; } = null!;

        /// <summary>
        /// The ID of an API Gateway resource.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the RestApi resource in which API Gateway creates the method.
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;


        /// <summary>
        /// Create a Method resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Method(string name, MethodArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:Method", name, args ?? new MethodArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Method(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:Method", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Method resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Method Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Method(name, id, options);
        }
    }

    public sealed class MethodArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the method requires clients to submit a valid API key.
        /// </summary>
        [Input("apiKeyRequired")]
        public Input<bool>? ApiKeyRequired { get; set; }

        [Input("authorizationScopes")]
        private InputList<string>? _authorizationScopes;

        /// <summary>
        /// A list of authorization scopes configured on the method.
        /// </summary>
        public InputList<string> AuthorizationScopes
        {
            get => _authorizationScopes ?? (_authorizationScopes = new InputList<string>());
            set => _authorizationScopes = value;
        }

        /// <summary>
        /// The method's authorization type.
        /// </summary>
        [Input("authorizationType")]
        public Input<Pulumi.AwsNative.ApiGateway.MethodAuthorizationType>? AuthorizationType { get; set; }

        /// <summary>
        /// The identifier of the authorizer to use on this method.
        /// </summary>
        [Input("authorizerId")]
        public Input<string>? AuthorizerId { get; set; }

        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        [Input("httpMethod", required: true)]
        public Input<string> HttpMethod { get; set; } = null!;

        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        [Input("integration")]
        public Input<Inputs.MethodIntegrationArgs>? Integration { get; set; }

        [Input("methodResponses")]
        private InputList<Inputs.MethodResponseArgs>? _methodResponses;

        /// <summary>
        /// The responses that can be sent to the client who calls the method.
        /// </summary>
        public InputList<Inputs.MethodResponseArgs> MethodResponses
        {
            get => _methodResponses ?? (_methodResponses = new InputList<Inputs.MethodResponseArgs>());
            set => _methodResponses = value;
        }

        /// <summary>
        /// A friendly operation name for the method.
        /// </summary>
        [Input("operationName")]
        public Input<string>? OperationName { get; set; }

        /// <summary>
        /// The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
        /// </summary>
        [Input("requestModels")]
        public Input<object>? RequestModels { get; set; }

        /// <summary>
        /// The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
        /// </summary>
        [Input("requestParameters")]
        public Input<object>? RequestParameters { get; set; }

        /// <summary>
        /// The ID of the associated request validator.
        /// </summary>
        [Input("requestValidatorId")]
        public Input<string>? RequestValidatorId { get; set; }

        /// <summary>
        /// The ID of an API Gateway resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The ID of the RestApi resource in which API Gateway creates the method.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public MethodArgs()
        {
        }
        public static new MethodArgs Empty => new MethodArgs();
    }
}