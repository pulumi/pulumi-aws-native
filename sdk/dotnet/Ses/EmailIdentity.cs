// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses
{
    /// <summary>
    /// Resource Type definition for AWS::SES::EmailIdentity
    /// </summary>
    [AwsNativeResourceType("aws-native:ses:EmailIdentity")]
    public partial class EmailIdentity : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Used to associate a configuration set with an email identity.
        /// </summary>
        [Output("configurationSetAttributes")]
        public Output<Outputs.EmailIdentityConfigurationSetAttributes?> ConfigurationSetAttributes { get; private set; } = null!;

        /// <summary>
        /// An object that contains information about the DKIM attributes for the identity.
        /// </summary>
        [Output("dkimAttributes")]
        public Output<Outputs.EmailIdentityDkimAttributes?> DkimAttributes { get; private set; } = null!;

        /// <summary>
        /// The host name for the first token that you have to add to the DNS configuration for your domain.
        /// </summary>
        [Output("dkimDnsTokenName1")]
        public Output<string> DkimDnsTokenName1 { get; private set; } = null!;

        /// <summary>
        /// The host name for the second token that you have to add to the DNS configuration for your domain.
        /// </summary>
        [Output("dkimDnsTokenName2")]
        public Output<string> DkimDnsTokenName2 { get; private set; } = null!;

        /// <summary>
        /// The host name for the third token that you have to add to the DNS configuration for your domain.
        /// </summary>
        [Output("dkimDnsTokenName3")]
        public Output<string> DkimDnsTokenName3 { get; private set; } = null!;

        /// <summary>
        /// The record value for the first token that you have to add to the DNS configuration for your domain.
        /// </summary>
        [Output("dkimDnsTokenValue1")]
        public Output<string> DkimDnsTokenValue1 { get; private set; } = null!;

        /// <summary>
        /// The record value for the second token that you have to add to the DNS configuration for your domain.
        /// </summary>
        [Output("dkimDnsTokenValue2")]
        public Output<string> DkimDnsTokenValue2 { get; private set; } = null!;

        /// <summary>
        /// The record value for the third token that you have to add to the DNS configuration for your domain.
        /// </summary>
        [Output("dkimDnsTokenValue3")]
        public Output<string> DkimDnsTokenValue3 { get; private set; } = null!;

        /// <summary>
        /// If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) .
        /// 
        /// You can only specify this object if the email identity is a domain, as opposed to an address.
        /// </summary>
        [Output("dkimSigningAttributes")]
        public Output<Outputs.EmailIdentityDkimSigningAttributes?> DkimSigningAttributes { get; private set; } = null!;

        /// <summary>
        /// The email address or domain to verify.
        /// </summary>
        [Output("emailIdentity")]
        public Output<string> EmailIdentityValue { get; private set; } = null!;

        /// <summary>
        /// Used to enable or disable feedback forwarding for an identity.
        /// </summary>
        [Output("feedbackAttributes")]
        public Output<Outputs.EmailIdentityFeedbackAttributes?> FeedbackAttributes { get; private set; } = null!;

        /// <summary>
        /// Used to enable or disable the custom Mail-From domain configuration for an email identity.
        /// </summary>
        [Output("mailFromAttributes")]
        public Output<Outputs.EmailIdentityMailFromAttributes?> MailFromAttributes { get; private set; } = null!;


        /// <summary>
        /// Create a EmailIdentity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EmailIdentity(string name, EmailIdentityArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ses:EmailIdentity", name, args ?? new EmailIdentityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EmailIdentity(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ses:EmailIdentity", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "emailIdentity",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EmailIdentity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EmailIdentity Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EmailIdentity(name, id, options);
        }
    }

    public sealed class EmailIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to associate a configuration set with an email identity.
        /// </summary>
        [Input("configurationSetAttributes")]
        public Input<Inputs.EmailIdentityConfigurationSetAttributesArgs>? ConfigurationSetAttributes { get; set; }

        /// <summary>
        /// An object that contains information about the DKIM attributes for the identity.
        /// </summary>
        [Input("dkimAttributes")]
        public Input<Inputs.EmailIdentityDkimAttributesArgs>? DkimAttributes { get; set; }

        /// <summary>
        /// If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for [Easy DKIM](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) .
        /// 
        /// You can only specify this object if the email identity is a domain, as opposed to an address.
        /// </summary>
        [Input("dkimSigningAttributes")]
        public Input<Inputs.EmailIdentityDkimSigningAttributesArgs>? DkimSigningAttributes { get; set; }

        /// <summary>
        /// The email address or domain to verify.
        /// </summary>
        [Input("emailIdentity", required: true)]
        public Input<string> EmailIdentityValue { get; set; } = null!;

        /// <summary>
        /// Used to enable or disable feedback forwarding for an identity.
        /// </summary>
        [Input("feedbackAttributes")]
        public Input<Inputs.EmailIdentityFeedbackAttributesArgs>? FeedbackAttributes { get; set; }

        /// <summary>
        /// Used to enable or disable the custom Mail-From domain configuration for an email identity.
        /// </summary>
        [Input("mailFromAttributes")]
        public Input<Inputs.EmailIdentityMailFromAttributesArgs>? MailFromAttributes { get; set; }

        public EmailIdentityArgs()
        {
        }
        public static new EmailIdentityArgs Empty => new EmailIdentityArgs();
    }
}
