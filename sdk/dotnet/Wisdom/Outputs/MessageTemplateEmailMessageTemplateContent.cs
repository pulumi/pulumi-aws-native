// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    /// <summary>
    /// The content of message template that applies to email channel subtype.
    /// </summary>
    [OutputType]
    public sealed class MessageTemplateEmailMessageTemplateContent
    {
        /// <summary>
        /// The body to use in email messages.
        /// </summary>
        public readonly Outputs.MessageTemplateEmailMessageTemplateContentBody Body;
        /// <summary>
        /// The email headers to include in email messages.
        /// </summary>
        public readonly ImmutableArray<Outputs.MessageTemplateEmailMessageTemplateHeader> Headers;
        /// <summary>
        /// The subject line, or title, to use in email messages.
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private MessageTemplateEmailMessageTemplateContent(
            Outputs.MessageTemplateEmailMessageTemplateContentBody body,

            ImmutableArray<Outputs.MessageTemplateEmailMessageTemplateHeader> headers,

            string subject)
        {
            Body = body;
            Headers = headers;
            Subject = subject;
        }
    }
}
