// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class GraphQLApiLambdaAuthorizerConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizerResultTtlInSeconds")]
        public Input<double>? AuthorizerResultTtlInSeconds { get; set; }

        [Input("authorizerUri")]
        public Input<string>? AuthorizerUri { get; set; }

        [Input("identityValidationExpression")]
        public Input<string>? IdentityValidationExpression { get; set; }

        public GraphQLApiLambdaAuthorizerConfigArgs()
        {
        }
        public static new GraphQLApiLambdaAuthorizerConfigArgs Empty => new GraphQLApiLambdaAuthorizerConfigArgs();
    }
}