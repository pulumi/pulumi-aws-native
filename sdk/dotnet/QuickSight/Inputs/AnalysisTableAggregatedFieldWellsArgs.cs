// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupBy")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _groupBy;
        public InputList<Inputs.AnalysisDimensionFieldArgs> GroupBy
        {
            get => _groupBy ?? (_groupBy = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _groupBy = value;
        }

        [Input("values")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _values;
        public InputList<Inputs.AnalysisMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _values = value;
        }

        public AnalysisTableAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisTableAggregatedFieldWellsArgs Empty => new AnalysisTableAggregatedFieldWellsArgs();
    }
}