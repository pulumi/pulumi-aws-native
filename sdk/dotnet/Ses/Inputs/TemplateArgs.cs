// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    /// <summary>
    /// The content of the email, composed of a subject line, an HTML part, and a text-only part
    /// </summary>
    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTML body of the email.
        /// </summary>
        [Input("htmlPart")]
        public Input<string>? HtmlPart { get; set; }

        /// <summary>
        /// The subject line of the email.
        /// </summary>
        [Input("subjectPart", required: true)]
        public Input<string> SubjectPart { get; set; } = null!;

        /// <summary>
        /// The name of the template.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// The email body that is visible to recipients whose email clients do not display HTML content.
        /// </summary>
        [Input("textPart")]
        public Input<string>? TextPart { get; set; }

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }
}