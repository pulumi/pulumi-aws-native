// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    /// <summary>
    /// A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode. Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
    /// </summary>
    public sealed class ApiLambdaAuthorizerConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizerResultTtlInSeconds")]
        public Input<int>? AuthorizerResultTtlInSeconds { get; set; }

        [Input("authorizerUri", required: true)]
        public Input<string> AuthorizerUri { get; set; } = null!;

        [Input("identityValidationExpression")]
        public Input<string>? IdentityValidationExpression { get; set; }

        public ApiLambdaAuthorizerConfigArgs()
        {
        }
        public static new ApiLambdaAuthorizerConfigArgs Empty => new ApiLambdaAuthorizerConfigArgs();
    }
}
