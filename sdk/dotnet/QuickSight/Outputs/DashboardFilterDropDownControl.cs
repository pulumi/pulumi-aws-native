// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardFilterDropDownControl
    {
        public readonly Outputs.DashboardCascadingControlConfiguration? CascadingControlConfiguration;
        public readonly Outputs.DashboardDropDownControlDisplayOptions? DisplayOptions;
        public readonly string FilterControlId;
        public readonly Outputs.DashboardFilterSelectableValues? SelectableValues;
        public readonly string SourceFilterId;
        public readonly string Title;
        public readonly Pulumi.AwsNative.QuickSight.DashboardSheetControlListType? Type;

        [OutputConstructor]
        private DashboardFilterDropDownControl(
            Outputs.DashboardCascadingControlConfiguration? cascadingControlConfiguration,

            Outputs.DashboardDropDownControlDisplayOptions? displayOptions,

            string filterControlId,

            Outputs.DashboardFilterSelectableValues? selectableValues,

            string sourceFilterId,

            string title,

            Pulumi.AwsNative.QuickSight.DashboardSheetControlListType? type)
        {
            CascadingControlConfiguration = cascadingControlConfiguration;
            DisplayOptions = displayOptions;
            FilterControlId = filterControlId;
            SelectableValues = selectableValues;
            SourceFilterId = sourceFilterId;
            Title = title;
            Type = type;
        }
    }
}