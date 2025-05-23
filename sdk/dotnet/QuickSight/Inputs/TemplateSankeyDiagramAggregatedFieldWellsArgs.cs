// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSankeyDiagramAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("destination")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _destination;

        /// <summary>
        /// The destination field wells of a sankey diagram.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Destination
        {
            get => _destination ?? (_destination = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _destination = value;
        }

        [Input("source")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _source;

        /// <summary>
        /// The source field wells of a sankey diagram.
        /// </summary>
        public InputList<Inputs.TemplateDimensionFieldArgs> Source
        {
            get => _source ?? (_source = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _source = value;
        }

        [Input("weight")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _weight;

        /// <summary>
        /// The weight field wells of a sankey diagram.
        /// </summary>
        public InputList<Inputs.TemplateMeasureFieldArgs> Weight
        {
            get => _weight ?? (_weight = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _weight = value;
        }

        public TemplateSankeyDiagramAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateSankeyDiagramAggregatedFieldWellsArgs Empty => new TemplateSankeyDiagramAggregatedFieldWellsArgs();
    }
}
