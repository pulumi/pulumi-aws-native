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
    /// The body to use in SMS messages.
    /// </summary>
    [OutputType]
    public sealed class MessageTemplateSmsMessageTemplateContentBody
    {
        /// <summary>
        /// The message body to use in SMS messages.
        /// </summary>
        public readonly Outputs.MessageTemplateBodyContentProvider? PlainText;

        [OutputConstructor]
        private MessageTemplateSmsMessageTemplateContentBody(Outputs.MessageTemplateBodyContentProvider? plainText)
        {
            PlainText = plainText;
        }
    }
}
