// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// Resource Type definition for AWS::AppSync::GraphQLApi
    /// </summary>
    [Obsolete(@"GraphQlApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:appsync:GraphQlApi")]
    public partial class GraphQlApi : global::Pulumi.CustomResource
    {
        [Output("additionalAuthenticationProviders")]
        public Output<ImmutableArray<Outputs.GraphQlApiAdditionalAuthenticationProvider>> AdditionalAuthenticationProviders { get; private set; } = null!;

        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        [Output("apiType")]
        public Output<string?> ApiType { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("authenticationType")]
        public Output<string> AuthenticationType { get; private set; } = null!;

        [Output("graphQlDns")]
        public Output<string> GraphQlDns { get; private set; } = null!;

        [Output("graphQlEndpointArn")]
        public Output<string> GraphQlEndpointArn { get; private set; } = null!;

        [Output("graphQlUrl")]
        public Output<string> GraphQlUrl { get; private set; } = null!;

        [Output("introspectionConfig")]
        public Output<string?> IntrospectionConfig { get; private set; } = null!;

        [Output("lambdaAuthorizerConfig")]
        public Output<Outputs.GraphQlApiLambdaAuthorizerConfig?> LambdaAuthorizerConfig { get; private set; } = null!;

        [Output("logConfig")]
        public Output<Outputs.GraphQlApiLogConfig?> LogConfig { get; private set; } = null!;

        [Output("mergedApiExecutionRoleArn")]
        public Output<string?> MergedApiExecutionRoleArn { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("openIdConnectConfig")]
        public Output<Outputs.GraphQlApiOpenIdConnectConfig?> OpenIdConnectConfig { get; private set; } = null!;

        [Output("ownerContact")]
        public Output<string?> OwnerContact { get; private set; } = null!;

        [Output("queryDepthLimit")]
        public Output<int?> QueryDepthLimit { get; private set; } = null!;

        [Output("realtimeDns")]
        public Output<string> RealtimeDns { get; private set; } = null!;

        [Output("realtimeUrl")]
        public Output<string> RealtimeUrl { get; private set; } = null!;

        [Output("resolverCountLimit")]
        public Output<int?> ResolverCountLimit { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.GraphQlApiTag>> Tags { get; private set; } = null!;

        [Output("userPoolConfig")]
        public Output<Outputs.GraphQlApiUserPoolConfig?> UserPoolConfig { get; private set; } = null!;

        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;

        [Output("xrayEnabled")]
        public Output<bool?> XrayEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a GraphQlApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GraphQlApi(string name, GraphQlApiArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:GraphQlApi", name, args ?? new GraphQlApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GraphQlApi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:GraphQlApi", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GraphQlApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GraphQlApi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GraphQlApi(name, id, options);
        }
    }

    public sealed class GraphQlApiArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticationProviders")]
        private InputList<Inputs.GraphQlApiAdditionalAuthenticationProviderArgs>? _additionalAuthenticationProviders;
        public InputList<Inputs.GraphQlApiAdditionalAuthenticationProviderArgs> AdditionalAuthenticationProviders
        {
            get => _additionalAuthenticationProviders ?? (_additionalAuthenticationProviders = new InputList<Inputs.GraphQlApiAdditionalAuthenticationProviderArgs>());
            set => _additionalAuthenticationProviders = value;
        }

        [Input("apiType")]
        public Input<string>? ApiType { get; set; }

        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        [Input("introspectionConfig")]
        public Input<string>? IntrospectionConfig { get; set; }

        [Input("lambdaAuthorizerConfig")]
        public Input<Inputs.GraphQlApiLambdaAuthorizerConfigArgs>? LambdaAuthorizerConfig { get; set; }

        [Input("logConfig")]
        public Input<Inputs.GraphQlApiLogConfigArgs>? LogConfig { get; set; }

        [Input("mergedApiExecutionRoleArn")]
        public Input<string>? MergedApiExecutionRoleArn { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("openIdConnectConfig")]
        public Input<Inputs.GraphQlApiOpenIdConnectConfigArgs>? OpenIdConnectConfig { get; set; }

        [Input("ownerContact")]
        public Input<string>? OwnerContact { get; set; }

        [Input("queryDepthLimit")]
        public Input<int>? QueryDepthLimit { get; set; }

        [Input("resolverCountLimit")]
        public Input<int>? ResolverCountLimit { get; set; }

        [Input("tags")]
        private InputList<Inputs.GraphQlApiTagArgs>? _tags;
        public InputList<Inputs.GraphQlApiTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GraphQlApiTagArgs>());
            set => _tags = value;
        }

        [Input("userPoolConfig")]
        public Input<Inputs.GraphQlApiUserPoolConfigArgs>? UserPoolConfig { get; set; }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        [Input("xrayEnabled")]
        public Input<bool>? XrayEnabled { get; set; }

        public GraphQlApiArgs()
        {
        }
        public static new GraphQlApiArgs Empty => new GraphQlApiArgs();
    }
}