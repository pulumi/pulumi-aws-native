// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateCategoryFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The column that the filter is applied to.
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.TemplateColumnIdentifierArgs> Column { get; set; } = null!;

        /// <summary>
        /// The configuration for a `CategoryFilter` .
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.TemplateCategoryFilterConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.
        /// </summary>
        [Input("defaultFilterControlConfiguration")]
        public Input<Inputs.TemplateDefaultFilterControlConfigurationArgs>? DefaultFilterControlConfiguration { get; set; }

        /// <summary>
        /// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
        /// </summary>
        [Input("filterId", required: true)]
        public Input<string> FilterId { get; set; } = null!;

        public TemplateCategoryFilterArgs()
        {
        }
        public static new TemplateCategoryFilterArgs Empty => new TemplateCategoryFilterArgs();
    }
}
