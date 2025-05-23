// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateScatterPlotUnaggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("category")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _category;

        /// <summary>
        /// The category field well of a scatter plot.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("label")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _label;

        /// <summary>
        /// The label field well of a scatter plot.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Label
        {
            get => _label ?? (_label = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _label = value;
        }

        [Input("size")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _size;

        /// <summary>
        /// The size field well of a scatter plot.
        /// </summary>
        public InputList<Inputs.TemplateMeasureFieldArgs> Size
        {
            get => _size ?? (_size = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _size = value;
        }

        [Input("xAxis")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _xAxis;

        /// <summary>
        /// The x-axis field well of a scatter plot.
        /// 
        /// The x-axis is a dimension field and cannot be aggregated.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> XAxis
        {
            get => _xAxis ?? (_xAxis = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _xAxis = value;
        }

        [Input("yAxis")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _yAxis;

        /// <summary>
        /// The y-axis field well of a scatter plot.
        /// 
        /// The y-axis is a dimension field and cannot be aggregated.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> YAxis
        {
            get => _yAxis ?? (_yAxis = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _yAxis = value;
        }

        public TemplateScatterPlotUnaggregatedFieldWellsArgs()
        {
        }
        public static new TemplateScatterPlotUnaggregatedFieldWellsArgs Empty => new TemplateScatterPlotUnaggregatedFieldWellsArgs();
    }
}
