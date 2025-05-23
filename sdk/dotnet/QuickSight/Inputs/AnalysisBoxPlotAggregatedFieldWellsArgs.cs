// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisBoxPlotAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupBy")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _groupBy;

        /// <summary>
        /// The group by field well of a box plot chart. Values are grouped based on group by fields.
        /// </summary>
        public InputList<Inputs.AnalysisDimensionFieldArgs> GroupBy
        {
            get => _groupBy ?? (_groupBy = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _groupBy = value;
        }

        [Input("values")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _values;

        /// <summary>
        /// The value field well of a box plot chart. Values are aggregated based on group by fields.
        /// </summary>
        public InputList<Inputs.AnalysisMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _values = value;
        }

        public AnalysisBoxPlotAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisBoxPlotAggregatedFieldWellsArgs Empty => new AnalysisBoxPlotAggregatedFieldWellsArgs();
    }
}
