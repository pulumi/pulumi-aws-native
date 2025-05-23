// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class GraphQlApiLambdaAuthorizerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds a response should be cached for.
        /// </summary>
        [Input("authorizerResultTtlInSeconds")]
        public Input<int>? AuthorizerResultTtlInSeconds { get; set; }

        /// <summary>
        /// The ARN of the Lambda function to be called for authorization.
        /// </summary>
        [Input("authorizerUri")]
        public Input<string>? AuthorizerUri { get; set; }

        /// <summary>
        /// A regular expression for validation of tokens before the Lambda function is called.
        /// </summary>
        [Input("identityValidationExpression")]
        public Input<string>? IdentityValidationExpression { get; set; }

        public GraphQlApiLambdaAuthorizerConfigArgs()
        {
        }
        public static new GraphQlApiLambdaAuthorizerConfigArgs Empty => new GraphQlApiLambdaAuthorizerConfigArgs();
    }
}
