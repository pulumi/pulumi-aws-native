// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolInviteMessageTemplate
    {
        /// <summary>
        /// The message template for email messages. EmailMessage is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
        /// </summary>
        public readonly string? EmailMessage;
        /// <summary>
        /// The subject line for email messages. EmailSubject is allowed only if [EmailSendingAccount](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.
        /// </summary>
        public readonly string? EmailSubject;
        /// <summary>
        /// The message template for SMS messages.
        /// </summary>
        public readonly string? SmsMessage;

        [OutputConstructor]
        private UserPoolInviteMessageTemplate(
            string? emailMessage,

            string? emailSubject,

            string? smsMessage)
        {
            EmailMessage = emailMessage;
            EmailSubject = emailSubject;
            SmsMessage = smsMessage;
        }
    }
}
