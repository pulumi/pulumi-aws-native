// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardExplicitHierarchyArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns", required: true)]
        private InputList<Inputs.DashboardColumnIdentifierArgs>? _columns;
        public InputList<Inputs.DashboardColumnIdentifierArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.DashboardColumnIdentifierArgs>());
            set => _columns = value;
        }

        [Input("drillDownFilters")]
        private InputList<Inputs.DashboardDrillDownFilterArgs>? _drillDownFilters;
        public InputList<Inputs.DashboardDrillDownFilterArgs> DrillDownFilters
        {
            get => _drillDownFilters ?? (_drillDownFilters = new InputList<Inputs.DashboardDrillDownFilterArgs>());
            set => _drillDownFilters = value;
        }

        [Input("hierarchyId", required: true)]
        public Input<string> HierarchyId { get; set; } = null!;

        public DashboardExplicitHierarchyArgs()
        {
        }
        public static new DashboardExplicitHierarchyArgs Empty => new DashboardExplicitHierarchyArgs();
    }
}