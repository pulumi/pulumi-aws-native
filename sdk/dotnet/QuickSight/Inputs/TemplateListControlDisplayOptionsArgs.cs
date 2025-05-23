// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateListControlDisplayOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration of info icon label options.
        /// </summary>
        [Input("infoIconLabelOptions")]
        public Input<Inputs.TemplateSheetControlInfoIconLabelOptionsArgs>? InfoIconLabelOptions { get; set; }

        /// <summary>
        /// The configuration of the search options in a list control.
        /// </summary>
        [Input("searchOptions")]
        public Input<Inputs.TemplateListControlSearchOptionsArgs>? SearchOptions { get; set; }

        /// <summary>
        /// The configuration of the `Select all` options in a list control.
        /// </summary>
        [Input("selectAllOptions")]
        public Input<Inputs.TemplateListControlSelectAllOptionsArgs>? SelectAllOptions { get; set; }

        /// <summary>
        /// The options to configure the title visibility, name, and font size.
        /// </summary>
        [Input("titleOptions")]
        public Input<Inputs.TemplateLabelOptionsArgs>? TitleOptions { get; set; }

        public TemplateListControlDisplayOptionsArgs()
        {
        }
        public static new TemplateListControlDisplayOptionsArgs Empty => new TemplateListControlDisplayOptionsArgs();
    }
}
