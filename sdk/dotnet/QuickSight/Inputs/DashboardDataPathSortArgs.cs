// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDataPathSortArgs : global::Pulumi.ResourceArgs
    {
        [Input("direction", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.DashboardSortDirection> Direction { get; set; } = null!;

        [Input("sortPaths", required: true)]
        private InputList<Inputs.DashboardDataPathValueArgs>? _sortPaths;
        public InputList<Inputs.DashboardDataPathValueArgs> SortPaths
        {
            get => _sortPaths ?? (_sortPaths = new InputList<Inputs.DashboardDataPathValueArgs>());
            set => _sortPaths = value;
        }

        public DashboardDataPathSortArgs()
        {
        }
        public static new DashboardDataPathSortArgs Empty => new DashboardDataPathSortArgs();
    }
}