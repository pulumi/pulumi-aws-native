// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardNumericalMeasureField
    {
        /// <summary>
        /// The aggregation function of the measure field.
        /// </summary>
        public readonly Outputs.DashboardNumericalAggregationFunction? AggregationFunction;
        /// <summary>
        /// The column that is used in the `NumericalMeasureField` .
        /// </summary>
        public readonly Outputs.DashboardColumnIdentifier Column;
        /// <summary>
        /// The custom field ID.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The format configuration of the field.
        /// </summary>
        public readonly Outputs.DashboardNumberFormatConfiguration? FormatConfiguration;

        [OutputConstructor]
        private DashboardNumericalMeasureField(
            Outputs.DashboardNumericalAggregationFunction? aggregationFunction,

            Outputs.DashboardColumnIdentifier column,

            string fieldId,

            Outputs.DashboardNumberFormatConfiguration? formatConfiguration)
        {
            AggregationFunction = aggregationFunction;
            Column = column;
            FieldId = fieldId;
            FormatConfiguration = formatConfiguration;
        }
    }
}
