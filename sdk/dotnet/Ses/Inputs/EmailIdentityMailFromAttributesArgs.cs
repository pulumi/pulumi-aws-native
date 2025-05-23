// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    /// <summary>
    /// Used to enable or disable the custom Mail-From domain configuration for an email identity.
    /// </summary>
    public sealed class EmailIdentityMailFromAttributesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.
        /// </summary>
        [Input("behaviorOnMxFailure")]
        public Input<string>? BehaviorOnMxFailure { get; set; }

        /// <summary>
        /// The custom MAIL FROM domain that you want the verified identity to use
        /// </summary>
        [Input("mailFromDomain")]
        public Input<string>? MailFromDomain { get; set; }

        public EmailIdentityMailFromAttributesArgs()
        {
        }
        public static new EmailIdentityMailFromAttributesArgs Empty => new EmailIdentityMailFromAttributesArgs();
    }
}
