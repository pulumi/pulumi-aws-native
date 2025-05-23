// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDataLabelTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The option that specifies individual data values for labels.
        /// </summary>
        [Input("dataPathLabelType")]
        public Input<Inputs.TemplateDataPathLabelTypeArgs>? DataPathLabelType { get; set; }

        /// <summary>
        /// Determines the label configuration for the entire field.
        /// </summary>
        [Input("fieldLabelType")]
        public Input<Inputs.TemplateFieldLabelTypeArgs>? FieldLabelType { get; set; }

        /// <summary>
        /// Determines the label configuration for the maximum value in a visual.
        /// </summary>
        [Input("maximumLabelType")]
        public Input<Inputs.TemplateMaximumLabelTypeArgs>? MaximumLabelType { get; set; }

        /// <summary>
        /// Determines the label configuration for the minimum value in a visual.
        /// </summary>
        [Input("minimumLabelType")]
        public Input<Inputs.TemplateMinimumLabelTypeArgs>? MinimumLabelType { get; set; }

        /// <summary>
        /// Determines the label configuration for range end value in a visual.
        /// </summary>
        [Input("rangeEndsLabelType")]
        public Input<Inputs.TemplateRangeEndsLabelTypeArgs>? RangeEndsLabelType { get; set; }

        public TemplateDataLabelTypeArgs()
        {
        }
        public static new TemplateDataLabelTypeArgs Empty => new TemplateDataLabelTypeArgs();
    }
}
