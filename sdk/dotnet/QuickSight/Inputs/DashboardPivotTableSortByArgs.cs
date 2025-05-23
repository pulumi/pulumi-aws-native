// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableSortByArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The column sort (field id, direction) for the pivot table sort by options.
        /// </summary>
        [Input("column")]
        public Input<Inputs.DashboardColumnSortArgs>? Column { get; set; }

        /// <summary>
        /// The data path sort (data path value, direction) for the pivot table sort by options.
        /// </summary>
        [Input("dataPath")]
        public Input<Inputs.DashboardDataPathSortArgs>? DataPath { get; set; }

        /// <summary>
        /// The field sort (field id, direction) for the pivot table sort by options.
        /// </summary>
        [Input("field")]
        public Input<Inputs.DashboardFieldSortArgs>? Field { get; set; }

        public DashboardPivotTableSortByArgs()
        {
        }
        public static new DashboardPivotTableSortByArgs Empty => new DashboardPivotTableSortByArgs();
    }
}
