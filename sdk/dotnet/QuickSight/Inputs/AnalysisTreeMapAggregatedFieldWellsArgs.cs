// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTreeMapAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("colors")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _colors;
        public InputList<Inputs.AnalysisMeasureFieldArgs> Colors
        {
            get => _colors ?? (_colors = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _colors = value;
        }

        [Input("groups")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _groups;
        public InputList<Inputs.AnalysisDimensionFieldArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _groups = value;
        }

        [Input("sizes")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _sizes;
        public InputList<Inputs.AnalysisMeasureFieldArgs> Sizes
        {
            get => _sizes ?? (_sizes = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _sizes = value;
        }

        public AnalysisTreeMapAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisTreeMapAggregatedFieldWellsArgs Empty => new AnalysisTreeMapAggregatedFieldWellsArgs();
    }
}