// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateHistogramAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _values;

        /// <summary>
        /// The value field wells of a histogram. Values are aggregated by `COUNT` or `DISTINCT_COUNT` .
        /// </summary>
        public InputList<Inputs.TemplateMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _values = value;
        }

        public TemplateHistogramAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateHistogramAggregatedFieldWellsArgs Empty => new TemplateHistogramAggregatedFieldWellsArgs();
    }
}
