// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTableAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupBy")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _groupBy;
        public InputList<Inputs.TemplateDimensionFieldArgs> GroupBy
        {
            get => _groupBy ?? (_groupBy = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _groupBy = value;
        }

        [Input("values")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _values;
        public InputList<Inputs.TemplateMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _values = value;
        }

        public TemplateTableAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateTableAggregatedFieldWellsArgs Empty => new TemplateTableAggregatedFieldWellsArgs();
    }
}