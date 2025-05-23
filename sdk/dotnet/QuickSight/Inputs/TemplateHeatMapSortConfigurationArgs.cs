// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateHeatMapSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The limit on the number of columns that are displayed in a heat map.
        /// </summary>
        [Input("heatMapColumnItemsLimitConfiguration")]
        public Input<Inputs.TemplateItemsLimitConfigurationArgs>? HeatMapColumnItemsLimitConfiguration { get; set; }

        [Input("heatMapColumnSort")]
        private InputList<Inputs.TemplateFieldSortOptionsArgs>? _heatMapColumnSort;

        /// <summary>
        /// The column sort configuration for heat map for columns that aren't a part of a field well.
        /// </summary>
        public InputList<Inputs.TemplateFieldSortOptionsArgs> HeatMapColumnSort
        {
            get => _heatMapColumnSort ?? (_heatMapColumnSort = new InputList<Inputs.TemplateFieldSortOptionsArgs>());
            set => _heatMapColumnSort = value;
        }

        /// <summary>
        /// The limit on the number of rows that are displayed in a heat map.
        /// </summary>
        [Input("heatMapRowItemsLimitConfiguration")]
        public Input<Inputs.TemplateItemsLimitConfigurationArgs>? HeatMapRowItemsLimitConfiguration { get; set; }

        [Input("heatMapRowSort")]
        private InputList<Inputs.TemplateFieldSortOptionsArgs>? _heatMapRowSort;

        /// <summary>
        /// The field sort configuration of the rows fields.
        /// </summary>
        public InputList<Inputs.TemplateFieldSortOptionsArgs> HeatMapRowSort
        {
            get => _heatMapRowSort ?? (_heatMapRowSort = new InputList<Inputs.TemplateFieldSortOptionsArgs>());
            set => _heatMapRowSort = value;
        }

        public TemplateHeatMapSortConfigurationArgs()
        {
        }
        public static new TemplateHeatMapSortConfigurationArgs Empty => new TemplateHeatMapSortConfigurationArgs();
    }
}
