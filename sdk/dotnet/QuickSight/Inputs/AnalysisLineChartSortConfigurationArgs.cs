// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisLineChartSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The limit on the number of categories that are displayed in a line chart.
        /// </summary>
        [Input("categoryItemsLimitConfiguration")]
        public Input<Inputs.AnalysisItemsLimitConfigurationArgs>? CategoryItemsLimitConfiguration { get; set; }

        [Input("categorySort")]
        private InputList<Inputs.AnalysisFieldSortOptionsArgs>? _categorySort;

        /// <summary>
        /// The sort configuration of the category fields.
        /// </summary>
        public InputList<Inputs.AnalysisFieldSortOptionsArgs> CategorySort
        {
            get => _categorySort ?? (_categorySort = new InputList<Inputs.AnalysisFieldSortOptionsArgs>());
            set => _categorySort = value;
        }

        /// <summary>
        /// The limit on the number of lines that are displayed in a line chart.
        /// </summary>
        [Input("colorItemsLimitConfiguration")]
        public Input<Inputs.AnalysisItemsLimitConfigurationArgs>? ColorItemsLimitConfiguration { get; set; }

        /// <summary>
        /// The limit on the number of small multiples panels that are displayed.
        /// </summary>
        [Input("smallMultiplesLimitConfiguration")]
        public Input<Inputs.AnalysisItemsLimitConfigurationArgs>? SmallMultiplesLimitConfiguration { get; set; }

        [Input("smallMultiplesSort")]
        private InputList<Inputs.AnalysisFieldSortOptionsArgs>? _smallMultiplesSort;

        /// <summary>
        /// The sort configuration of the small multiples field.
        /// </summary>
        public InputList<Inputs.AnalysisFieldSortOptionsArgs> SmallMultiplesSort
        {
            get => _smallMultiplesSort ?? (_smallMultiplesSort = new InputList<Inputs.AnalysisFieldSortOptionsArgs>());
            set => _smallMultiplesSort = value;
        }

        public AnalysisLineChartSortConfigurationArgs()
        {
        }
        public static new AnalysisLineChartSortConfigurationArgs Empty => new AnalysisLineChartSortConfigurationArgs();
    }
}
