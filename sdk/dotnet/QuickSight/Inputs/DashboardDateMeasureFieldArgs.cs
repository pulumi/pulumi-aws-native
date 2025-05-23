// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDateMeasureFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The aggregation function of the measure field.
        /// </summary>
        [Input("aggregationFunction")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardDateAggregationFunction>? AggregationFunction { get; set; }

        /// <summary>
        /// The column that is used in the `DateMeasureField` .
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        /// <summary>
        /// The custom field ID.
        /// </summary>
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        /// <summary>
        /// The format configuration of the field.
        /// </summary>
        [Input("formatConfiguration")]
        public Input<Inputs.DashboardDateTimeFormatConfigurationArgs>? FormatConfiguration { get; set; }

        public DashboardDateMeasureFieldArgs()
        {
        }
        public static new DashboardDateMeasureFieldArgs Empty => new DashboardDateMeasureFieldArgs();
    }
}
