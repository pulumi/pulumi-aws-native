// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    public static class GetApi
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::Api
        /// </summary>
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("aws-native:apigatewayv2:getApi", args ?? new GetApiArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::Api
        /// </summary>
        public static Output<GetApiResult> Invoke(GetApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiResult>("aws-native:apigatewayv2:getApi", args ?? new GetApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetApiArgs()
        {
        }
        public static new GetApiArgs Empty => new GetApiArgs();
    }

    public sealed class GetApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetApiInvokeArgs()
        {
        }
        public static new GetApiInvokeArgs Empty => new GetApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiResult
    {
        public readonly string? ApiEndpoint;
        public readonly string? ApiKeySelectionExpression;
        public readonly string? BasePath;
        public readonly object? Body;
        public readonly Outputs.ApiBodyS3Location? BodyS3Location;
        public readonly Outputs.ApiCors? CorsConfiguration;
        public readonly string? CredentialsArn;
        public readonly string? Description;
        public readonly bool? DisableExecuteApiEndpoint;
        public readonly bool? DisableSchemaValidation;
        public readonly bool? FailOnWarnings;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? RouteKey;
        public readonly string? RouteSelectionExpression;
        public readonly object? Tags;
        public readonly string? Target;
        public readonly string? Version;

        [OutputConstructor]
        private GetApiResult(
            string? apiEndpoint,

            string? apiKeySelectionExpression,

            string? basePath,

            object? body,

            Outputs.ApiBodyS3Location? bodyS3Location,

            Outputs.ApiCors? corsConfiguration,

            string? credentialsArn,

            string? description,

            bool? disableExecuteApiEndpoint,

            bool? disableSchemaValidation,

            bool? failOnWarnings,

            string? id,

            string? name,

            string? routeKey,

            string? routeSelectionExpression,

            object? tags,

            string? target,

            string? version)
        {
            ApiEndpoint = apiEndpoint;
            ApiKeySelectionExpression = apiKeySelectionExpression;
            BasePath = basePath;
            Body = body;
            BodyS3Location = bodyS3Location;
            CorsConfiguration = corsConfiguration;
            CredentialsArn = credentialsArn;
            Description = description;
            DisableExecuteApiEndpoint = disableExecuteApiEndpoint;
            DisableSchemaValidation = disableSchemaValidation;
            FailOnWarnings = failOnWarnings;
            Id = id;
            Name = name;
            RouteKey = routeKey;
            RouteSelectionExpression = routeSelectionExpression;
            Tags = tags;
            Target = target;
            Version = version;
        }
    }
}