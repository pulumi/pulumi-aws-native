// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class GraphQLApiAdditionalAuthenticationProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        [Input("lambdaAuthorizerConfig")]
        public Input<Inputs.GraphQLApiLambdaAuthorizerConfigArgs>? LambdaAuthorizerConfig { get; set; }

        [Input("openIDConnectConfig")]
        public Input<Inputs.GraphQLApiOpenIDConnectConfigArgs>? OpenIDConnectConfig { get; set; }

        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiCognitoUserPoolConfigArgs>? UserPoolConfig { get; set; }

        public GraphQLApiAdditionalAuthenticationProviderArgs()
        {
        }
        public static new GraphQLApiAdditionalAuthenticationProviderArgs Empty => new GraphQLApiAdditionalAuthenticationProviderArgs();
    }
}