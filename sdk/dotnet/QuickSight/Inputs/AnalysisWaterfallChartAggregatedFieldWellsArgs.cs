// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisWaterfallChartAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("breakdowns")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _breakdowns;

        /// <summary>
        /// The breakdown field wells of a waterfall visual.
        /// </summary>
        public InputList<Inputs.AnalysisDimensionFieldArgs> Breakdowns
        {
            get => _breakdowns ?? (_breakdowns = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _breakdowns = value;
        }

        [Input("categories")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _categories;

        /// <summary>
        /// The category field wells of a waterfall visual.
        /// </summary>
        public InputList<Inputs.AnalysisDimensionFieldArgs> Categories
        {
            get => _categories ?? (_categories = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _categories = value;
        }

        [Input("values")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _values;

        /// <summary>
        /// The value field wells of a waterfall visual.
        /// </summary>
        public InputList<Inputs.AnalysisMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _values = value;
        }

        public AnalysisWaterfallChartAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisWaterfallChartAggregatedFieldWellsArgs Empty => new AnalysisWaterfallChartAggregatedFieldWellsArgs();
    }
}
