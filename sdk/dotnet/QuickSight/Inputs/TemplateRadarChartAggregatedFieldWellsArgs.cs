// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateRadarChartAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("category")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _category;

        /// <summary>
        /// The aggregated field well categories of a radar chart.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("color")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _color;

        /// <summary>
        /// The color that are assigned to the aggregated field wells of a radar chart.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Color
        {
            get => _color ?? (_color = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _color = value;
        }

        [Input("values")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _values;

        /// <summary>
        /// The values that are assigned to the aggregated field wells of a radar chart.
        /// </summary>
        public InputList<Inputs.TemplateMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _values = value;
        }

        public TemplateRadarChartAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateRadarChartAggregatedFieldWellsArgs Empty => new TemplateRadarChartAggregatedFieldWellsArgs();
    }
}
