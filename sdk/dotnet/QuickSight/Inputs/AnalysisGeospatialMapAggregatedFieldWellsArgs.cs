// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisGeospatialMapAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("colors")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _colors;

        /// <summary>
        /// The color field wells of a geospatial map.
        /// </summary>
        public InputList<Inputs.AnalysisDimensionFieldArgs> Colors
        {
            get => _colors ?? (_colors = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _colors = value;
        }

        [Input("geospatial")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _geospatial;

        /// <summary>
        /// The geospatial field wells of a geospatial map. Values are grouped by geospatial fields.
        /// </summary>
        public InputList<Inputs.AnalysisDimensionFieldArgs> Geospatial
        {
            get => _geospatial ?? (_geospatial = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _geospatial = value;
        }

        [Input("values")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _values;

        /// <summary>
        /// The size field wells of a geospatial map. Values are aggregated based on geospatial fields.
        /// </summary>
        public InputList<Inputs.AnalysisMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _values = value;
        }

        public AnalysisGeospatialMapAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisGeospatialMapAggregatedFieldWellsArgs Empty => new AnalysisGeospatialMapAggregatedFieldWellsArgs();
    }
}
