// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    /// <summary>
    /// Resource Type definition for AWS::Cognito::IdentityPool
    /// </summary>
    [AwsNativeResourceType("aws-native:cognito:IdentityPool")]
    public partial class IdentityPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enables the Basic (Classic) authentication flow.
        /// </summary>
        [Output("allowClassicFlow")]
        public Output<bool?> AllowClassicFlow { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the identity pool supports unauthenticated logins.
        /// </summary>
        [Output("allowUnauthenticatedIdentities")]
        public Output<bool> AllowUnauthenticatedIdentities { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The events to configure.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPool` for more information about the expected schema for this property.
        /// </summary>
        [Output("cognitoEvents")]
        public Output<object?> CognitoEvents { get; private set; } = null!;

        /// <summary>
        /// The Amazon Cognito user pools and their client IDs.
        /// </summary>
        [Output("cognitoIdentityProviders")]
        public Output<ImmutableArray<Outputs.IdentityPoolCognitoIdentityProvider>> CognitoIdentityProviders { get; private set; } = null!;

        /// <summary>
        /// Configuration options for configuring Amazon Cognito streams.
        /// </summary>
        [Output("cognitoStreams")]
        public Output<Outputs.IdentityPoolCognitoStreams?> CognitoStreams { get; private set; } = null!;

        /// <summary>
        /// The "domain" Amazon Cognito uses when referencing your users. This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
        /// 
        /// *Minimum length* : 1
        /// 
        /// *Maximum length* : 100
        /// </summary>
        [Output("developerProviderName")]
        public Output<string?> DeveloperProviderName { get; private set; } = null!;

        /// <summary>
        /// The name of your Amazon Cognito identity pool.
        /// 
        /// *Minimum length* : 1
        /// 
        /// *Maximum length* : 128
        /// 
        /// *Pattern* : `[\w\s+=,.@-]+`
        /// </summary>
        [Output("identityPoolName")]
        public Output<string?> IdentityPoolName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("identityPoolTags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> IdentityPoolTags { get; private set; } = null!;

        /// <summary>
        /// The name of the Amazon Cognito identity pool, returned as a string.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the OpenID connect providers.
        /// </summary>
        [Output("openIdConnectProviderArns")]
        public Output<ImmutableArray<string>> OpenIdConnectProviderArns { get; private set; } = null!;

        /// <summary>
        /// The configuration options to be applied to the identity pool.
        /// </summary>
        [Output("pushSync")]
        public Output<Outputs.IdentityPoolPushSync?> PushSync { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
        /// </summary>
        [Output("samlProviderArns")]
        public Output<ImmutableArray<string>> SamlProviderArns { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs that map provider names to provider app IDs.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPool` for more information about the expected schema for this property.
        /// </summary>
        [Output("supportedLoginProviders")]
        public Output<object?> SupportedLoginProviders { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityPool(string name, IdentityPoolArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cognito:IdentityPool", name, args ?? new IdentityPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cognito:IdentityPool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IdentityPool(name, id, options);
        }
    }

    public sealed class IdentityPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables the Basic (Classic) authentication flow.
        /// </summary>
        [Input("allowClassicFlow")]
        public Input<bool>? AllowClassicFlow { get; set; }

        /// <summary>
        /// Specifies whether the identity pool supports unauthenticated logins.
        /// </summary>
        [Input("allowUnauthenticatedIdentities", required: true)]
        public Input<bool> AllowUnauthenticatedIdentities { get; set; } = null!;

        /// <summary>
        /// The events to configure.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPool` for more information about the expected schema for this property.
        /// </summary>
        [Input("cognitoEvents")]
        public Input<object>? CognitoEvents { get; set; }

        [Input("cognitoIdentityProviders")]
        private InputList<Inputs.IdentityPoolCognitoIdentityProviderArgs>? _cognitoIdentityProviders;

        /// <summary>
        /// The Amazon Cognito user pools and their client IDs.
        /// </summary>
        public InputList<Inputs.IdentityPoolCognitoIdentityProviderArgs> CognitoIdentityProviders
        {
            get => _cognitoIdentityProviders ?? (_cognitoIdentityProviders = new InputList<Inputs.IdentityPoolCognitoIdentityProviderArgs>());
            set => _cognitoIdentityProviders = value;
        }

        /// <summary>
        /// Configuration options for configuring Amazon Cognito streams.
        /// </summary>
        [Input("cognitoStreams")]
        public Input<Inputs.IdentityPoolCognitoStreamsArgs>? CognitoStreams { get; set; }

        /// <summary>
        /// The "domain" Amazon Cognito uses when referencing your users. This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
        /// 
        /// *Minimum length* : 1
        /// 
        /// *Maximum length* : 100
        /// </summary>
        [Input("developerProviderName")]
        public Input<string>? DeveloperProviderName { get; set; }

        /// <summary>
        /// The name of your Amazon Cognito identity pool.
        /// 
        /// *Minimum length* : 1
        /// 
        /// *Maximum length* : 128
        /// 
        /// *Pattern* : `[\w\s+=,.@-]+`
        /// </summary>
        [Input("identityPoolName")]
        public Input<string>? IdentityPoolName { get; set; }

        [Input("identityPoolTags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _identityPoolTags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> IdentityPoolTags
        {
            get => _identityPoolTags ?? (_identityPoolTags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _identityPoolTags = value;
        }

        [Input("openIdConnectProviderArns")]
        private InputList<string>? _openIdConnectProviderArns;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the OpenID connect providers.
        /// </summary>
        public InputList<string> OpenIdConnectProviderArns
        {
            get => _openIdConnectProviderArns ?? (_openIdConnectProviderArns = new InputList<string>());
            set => _openIdConnectProviderArns = value;
        }

        /// <summary>
        /// The configuration options to be applied to the identity pool.
        /// </summary>
        [Input("pushSync")]
        public Input<Inputs.IdentityPoolPushSyncArgs>? PushSync { get; set; }

        [Input("samlProviderArns")]
        private InputList<string>? _samlProviderArns;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
        /// </summary>
        public InputList<string> SamlProviderArns
        {
            get => _samlProviderArns ?? (_samlProviderArns = new InputList<string>());
            set => _samlProviderArns = value;
        }

        /// <summary>
        /// Key-value pairs that map provider names to provider app IDs.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPool` for more information about the expected schema for this property.
        /// </summary>
        [Input("supportedLoginProviders")]
        public Input<object>? SupportedLoginProviders { get; set; }

        public IdentityPoolArgs()
        {
        }
        public static new IdentityPoolArgs Empty => new IdentityPoolArgs();
    }
}
