// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The aggregated field well for the pivot table.
        /// </summary>
        [Input("pivotTableAggregatedFieldWells")]
        public Input<Inputs.DashboardPivotTableAggregatedFieldWellsArgs>? PivotTableAggregatedFieldWells { get; set; }

        public DashboardPivotTableFieldWellsArgs()
        {
        }
        public static new DashboardPivotTableFieldWellsArgs Empty => new DashboardPivotTableFieldWellsArgs();
    }
}
