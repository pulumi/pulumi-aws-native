// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPredefinedHierarchyArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns", required: true)]
        private InputList<Inputs.DashboardColumnIdentifierArgs>? _columns;

        /// <summary>
        /// The list of columns that define the predefined hierarchy.
        /// </summary>
        public InputList<Inputs.DashboardColumnIdentifierArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.DashboardColumnIdentifierArgs>());
            set => _columns = value;
        }

        [Input("drillDownFilters")]
        private InputList<Inputs.DashboardDrillDownFilterArgs>? _drillDownFilters;

        /// <summary>
        /// The option that determines the drill down filters for the predefined hierarchy.
        /// </summary>
        public InputList<Inputs.DashboardDrillDownFilterArgs> DrillDownFilters
        {
            get => _drillDownFilters ?? (_drillDownFilters = new InputList<Inputs.DashboardDrillDownFilterArgs>());
            set => _drillDownFilters = value;
        }

        /// <summary>
        /// The hierarchy ID of the predefined hierarchy.
        /// </summary>
        [Input("hierarchyId", required: true)]
        public Input<string> HierarchyId { get; set; } = null!;

        public DashboardPredefinedHierarchyArgs()
        {
        }
        public static new DashboardPredefinedHierarchyArgs Empty => new DashboardPredefinedHierarchyArgs();
    }
}
