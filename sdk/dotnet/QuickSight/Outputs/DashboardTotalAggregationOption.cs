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
    public sealed class DashboardTotalAggregationOption
    {
        /// <summary>
        /// The field id that's associated with the total aggregation option.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The total aggregation function that you want to set for a specified field id.
        /// </summary>
        public readonly Outputs.DashboardTotalAggregationFunction TotalAggregationFunction;

        [OutputConstructor]
        private DashboardTotalAggregationOption(
            string fieldId,

            Outputs.DashboardTotalAggregationFunction totalAggregationFunction)
        {
            FieldId = fieldId;
            TotalAggregationFunction = totalAggregationFunction;
        }
    }
}
