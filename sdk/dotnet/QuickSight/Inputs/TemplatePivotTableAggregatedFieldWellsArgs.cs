// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePivotTableAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _columns;

        /// <summary>
        /// The columns field well for a pivot table. Values are grouped by columns fields.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _columns = value;
        }

        [Input("rows")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _rows;

        /// <summary>
        /// The rows field well for a pivot table. Values are grouped by rows fields.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Rows
        {
            get => _rows ?? (_rows = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _rows = value;
        }

        [Input("values")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _values;

        /// <summary>
        /// The values field well for a pivot table. Values are aggregated based on rows and columns fields.
        /// </summary>
        public InputList<Inputs.TemplateMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _values = value;
        }

        public TemplatePivotTableAggregatedFieldWellsArgs()
        {
        }
        public static new TemplatePivotTableAggregatedFieldWellsArgs Empty => new TemplatePivotTableAggregatedFieldWellsArgs();
    }
}
