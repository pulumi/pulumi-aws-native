// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSeriesItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data field series item configuration of a line chart.
        /// </summary>
        [Input("dataFieldSeriesItem")]
        public Input<Inputs.TemplateDataFieldSeriesItemArgs>? DataFieldSeriesItem { get; set; }

        /// <summary>
        /// The field series item configuration of a line chart.
        /// </summary>
        [Input("fieldSeriesItem")]
        public Input<Inputs.TemplateFieldSeriesItemArgs>? FieldSeriesItem { get; set; }

        public TemplateSeriesItemArgs()
        {
        }
        public static new TemplateSeriesItemArgs Empty => new TemplateSeriesItemArgs();
    }
}
