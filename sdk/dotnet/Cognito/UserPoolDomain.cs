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
    /// Resource Type definition for AWS::Cognito::UserPoolDomain
    /// </summary>
    [AwsNativeResourceType("aws-native:cognito:UserPoolDomain")]
    public partial class UserPoolDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon CloudFront endpoint that you use as the target of the alias that you set up with your Domain Name Service (DNS) provider.
        /// </summary>
        [Output("cloudFrontDistribution")]
        public Output<string> CloudFrontDistribution { get; private set; } = null!;

        /// <summary>
        /// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application. Use this object to specify an SSL certificate that is managed by ACM.
        /// 
        /// When you create a custom domain, the passkey RP ID defaults to the custom domain. If you had a prefix domain active, this will cause passkey integration for your prefix domain to stop working due to a mismatch in RP ID. To keep the prefix domain passkey integration working, you can explicitly set RP ID to the prefix domain.
        /// </summary>
        [Output("customDomainConfig")]
        public Output<Outputs.UserPoolDomainCustomDomainConfigType?> CustomDomainConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the domain that you want to update. For custom domains, this is the fully-qualified domain name, for example `auth.example.com` . For prefix domains, this is the prefix alone, such as `myprefix` .
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// A version number that indicates the state of managed login for your domain. Version `1` is hosted UI (classic). Version `2` is the newer managed login with the branding editor. For more information, see [Managed login](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html) .
        /// </summary>
        [Output("managedLoginVersion")]
        public Output<int?> ManagedLoginVersion { get; private set; } = null!;

        /// <summary>
        /// The ID of the user pool that is associated with the domain you're updating.
        /// </summary>
        [Output("userPoolId")]
        public Output<string> UserPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a UserPoolDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPoolDomain(string name, UserPoolDomainArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPoolDomain", name, args ?? new UserPoolDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPoolDomain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPoolDomain", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "domain",
                    "userPoolId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserPoolDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPoolDomain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserPoolDomain(name, id, options);
        }
    }

    public sealed class UserPoolDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application. Use this object to specify an SSL certificate that is managed by ACM.
        /// 
        /// When you create a custom domain, the passkey RP ID defaults to the custom domain. If you had a prefix domain active, this will cause passkey integration for your prefix domain to stop working due to a mismatch in RP ID. To keep the prefix domain passkey integration working, you can explicitly set RP ID to the prefix domain.
        /// </summary>
        [Input("customDomainConfig")]
        public Input<Inputs.UserPoolDomainCustomDomainConfigTypeArgs>? CustomDomainConfig { get; set; }

        /// <summary>
        /// The name of the domain that you want to update. For custom domains, this is the fully-qualified domain name, for example `auth.example.com` . For prefix domains, this is the prefix alone, such as `myprefix` .
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// A version number that indicates the state of managed login for your domain. Version `1` is hosted UI (classic). Version `2` is the newer managed login with the branding editor. For more information, see [Managed login](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html) .
        /// </summary>
        [Input("managedLoginVersion")]
        public Input<int>? ManagedLoginVersion { get; set; }

        /// <summary>
        /// The ID of the user pool that is associated with the domain you're updating.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public UserPoolDomainArgs()
        {
        }
        public static new UserPoolDomainArgs Empty => new UserPoolDomainArgs();
    }
}
