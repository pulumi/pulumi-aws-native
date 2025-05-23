// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class InAppTemplateInAppMessageContent
    {
        /// <summary>
        /// The background color for an in-app message banner, expressed as a hex color code (such as #000000 for black).
        /// </summary>
        public readonly string? BackgroundColor;
        /// <summary>
        /// An object that contains configuration information about the header or title text of the in-app message.
        /// </summary>
        public readonly Outputs.InAppTemplateBodyConfig? BodyConfig;
        /// <summary>
        /// An object that contains configuration information about the header or title text of the in-app message.
        /// </summary>
        public readonly Outputs.InAppTemplateHeaderConfig? HeaderConfig;
        /// <summary>
        /// The URL of the image that appears on an in-app message banner.
        /// </summary>
        public readonly string? ImageUrl;
        /// <summary>
        /// An object that contains configuration information about the primary button in an in-app message.
        /// </summary>
        public readonly Outputs.InAppTemplateButtonConfig? PrimaryBtn;
        /// <summary>
        /// An object that contains configuration information about the secondary button in an in-app message.
        /// </summary>
        public readonly Outputs.InAppTemplateButtonConfig? SecondaryBtn;

        [OutputConstructor]
        private InAppTemplateInAppMessageContent(
            string? backgroundColor,

            Outputs.InAppTemplateBodyConfig? bodyConfig,

            Outputs.InAppTemplateHeaderConfig? headerConfig,

            string? imageUrl,

            Outputs.InAppTemplateButtonConfig? primaryBtn,

            Outputs.InAppTemplateButtonConfig? secondaryBtn)
        {
            BackgroundColor = backgroundColor;
            BodyConfig = bodyConfig;
            HeaderConfig = headerConfig;
            ImageUrl = imageUrl;
            PrimaryBtn = primaryBtn;
            SecondaryBtn = secondaryBtn;
        }
    }
}
