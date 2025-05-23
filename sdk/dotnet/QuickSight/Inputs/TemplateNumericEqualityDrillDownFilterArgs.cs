// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateNumericEqualityDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The column that the filter is applied to.
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.TemplateColumnIdentifierArgs> Column { get; set; } = null!;

        /// <summary>
        /// The value of the double input numeric drill down filter.
        /// </summary>
        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public TemplateNumericEqualityDrillDownFilterArgs()
        {
        }
        public static new TemplateNumericEqualityDrillDownFilterArgs Empty => new TemplateNumericEqualityDrillDownFilterArgs();
    }
}
