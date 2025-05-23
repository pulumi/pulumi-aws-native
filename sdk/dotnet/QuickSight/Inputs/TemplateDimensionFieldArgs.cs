// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDimensionFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dimension type field with categorical type columns.
        /// </summary>
        [Input("categoricalDimensionField")]
        public Input<Inputs.TemplateCategoricalDimensionFieldArgs>? CategoricalDimensionField { get; set; }

        /// <summary>
        /// The dimension type field with date type columns.
        /// </summary>
        [Input("dateDimensionField")]
        public Input<Inputs.TemplateDateDimensionFieldArgs>? DateDimensionField { get; set; }

        /// <summary>
        /// The dimension type field with numerical type columns.
        /// </summary>
        [Input("numericalDimensionField")]
        public Input<Inputs.TemplateNumericalDimensionFieldArgs>? NumericalDimensionField { get; set; }

        public TemplateDimensionFieldArgs()
        {
        }
        public static new TemplateDimensionFieldArgs Empty => new TemplateDimensionFieldArgs();
    }
}
