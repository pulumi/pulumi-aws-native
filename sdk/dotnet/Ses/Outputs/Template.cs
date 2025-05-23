// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    /// <summary>
    /// The content of the email, composed of a subject line, an HTML part, and a text-only part
    /// </summary>
    [OutputType]
    public sealed class Template
    {
        /// <summary>
        /// The HTML body of the email.
        /// </summary>
        public readonly string? HtmlPart;
        /// <summary>
        /// The subject line of the email.
        /// </summary>
        public readonly string SubjectPart;
        /// <summary>
        /// The name of the template.
        /// </summary>
        public readonly string? TemplateName;
        /// <summary>
        /// The email body that is visible to recipients whose email clients do not display HTML content.
        /// </summary>
        public readonly string? TextPart;

        [OutputConstructor]
        private Template(
            string? htmlPart,

            string subjectPart,

            string? templateName,

            string? textPart)
        {
            HtmlPart = htmlPart;
            SubjectPart = subjectPart;
            TemplateName = templateName;
            TextPart = textPart;
        }
    }
}
