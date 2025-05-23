// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateSectionBasedLayoutConfiguration
    {
        /// <summary>
        /// A list of body section configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateBodySectionConfiguration> BodySections;
        /// <summary>
        /// The options for the canvas of a section-based layout.
        /// </summary>
        public readonly Outputs.TemplateSectionBasedLayoutCanvasSizeOptions CanvasSizeOptions;
        /// <summary>
        /// A list of footer section configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateHeaderFooterSectionConfiguration> FooterSections;
        /// <summary>
        /// A list of header section configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateHeaderFooterSectionConfiguration> HeaderSections;

        [OutputConstructor]
        private TemplateSectionBasedLayoutConfiguration(
            ImmutableArray<Outputs.TemplateBodySectionConfiguration> bodySections,

            Outputs.TemplateSectionBasedLayoutCanvasSizeOptions canvasSizeOptions,

            ImmutableArray<Outputs.TemplateHeaderFooterSectionConfiguration> footerSections,

            ImmutableArray<Outputs.TemplateHeaderFooterSectionConfiguration> headerSections)
        {
            BodySections = bodySections;
            CanvasSizeOptions = canvasSizeOptions;
            FooterSections = footerSections;
            HeaderSections = headerSections;
        }
    }
}
