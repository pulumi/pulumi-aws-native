// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateHeaderFooterSectionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The layout configuration of the header or footer section.
        /// </summary>
        [Input("layout", required: true)]
        public Input<Inputs.TemplateSectionLayoutConfigurationArgs> Layout { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the header or footer section.
        /// </summary>
        [Input("sectionId", required: true)]
        public Input<string> SectionId { get; set; } = null!;

        /// <summary>
        /// The style options of a header or footer section.
        /// </summary>
        [Input("style")]
        public Input<Inputs.TemplateSectionStyleArgs>? Style { get; set; }

        public TemplateHeaderFooterSectionConfigurationArgs()
        {
        }
        public static new TemplateHeaderFooterSectionConfigurationArgs Empty => new TemplateHeaderFooterSectionConfigurationArgs();
    }
}
