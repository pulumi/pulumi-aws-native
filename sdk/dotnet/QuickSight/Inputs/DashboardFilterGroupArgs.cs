// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFilterGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("crossDataset", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.DashboardCrossDatasetTypes> CrossDataset { get; set; } = null!;

        [Input("filterGroupId", required: true)]
        public Input<string> FilterGroupId { get; set; } = null!;

        [Input("filters", required: true)]
        private InputList<Inputs.DashboardFilterArgs>? _filters;
        public InputList<Inputs.DashboardFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.DashboardFilterArgs>());
            set => _filters = value;
        }

        [Input("scopeConfiguration", required: true)]
        public Input<Inputs.DashboardFilterScopeConfigurationArgs> ScopeConfiguration { get; set; } = null!;

        [Input("status")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardWidgetStatus>? Status { get; set; }

        public DashboardFilterGroupArgs()
        {
        }
        public static new DashboardFilterGroupArgs Empty => new DashboardFilterGroupArgs();
    }
}