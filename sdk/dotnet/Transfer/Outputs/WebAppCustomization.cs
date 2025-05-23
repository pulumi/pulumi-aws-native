// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Outputs
{

    [OutputType]
    public sealed class WebAppCustomization
    {
        /// <summary>
        /// Specifies a favicon to display in the browser tab.
        /// </summary>
        public readonly string? FaviconFile;
        /// <summary>
        /// Specifies a logo to display on the web app.
        /// </summary>
        public readonly string? LogoFile;
        /// <summary>
        /// Specifies a title to display on the web app.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private WebAppCustomization(
            string? faviconFile,

            string? logoFile,

            string? title)
        {
            FaviconFile = faviconFile;
            LogoFile = logoFile;
            Title = title;
        }
    }
}
