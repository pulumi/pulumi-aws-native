// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFunnelChartSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The limit on the number of categories displayed.
        /// </summary>
        [Input("categoryItemsLimit")]
        public Input<Inputs.TemplateItemsLimitConfigurationArgs>? CategoryItemsLimit { get; set; }

        [Input("categorySort")]
        private InputList<Inputs.TemplateFieldSortOptionsArgs>? _categorySort;

        /// <summary>
        /// The sort configuration of the category fields.
        /// </summary>
        public InputList<Inputs.TemplateFieldSortOptionsArgs> CategorySort
        {
            get => _categorySort ?? (_categorySort = new InputList<Inputs.TemplateFieldSortOptionsArgs>());
            set => _categorySort = value;
        }

        public TemplateFunnelChartSortConfigurationArgs()
        {
        }
        public static new TemplateFunnelChartSortConfigurationArgs Empty => new TemplateFunnelChartSortConfigurationArgs();
    }
}
