// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDefaultFilterListControlOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display options of a control.
        /// </summary>
        [Input("displayOptions")]
        public Input<Inputs.DashboardListControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        /// <summary>
        /// A list of selectable values that are used in a control.
        /// </summary>
        [Input("selectableValues")]
        public Input<Inputs.DashboardFilterSelectableValuesArgs>? SelectableValues { get; set; }

        /// <summary>
        /// The type of the `DefaultFilterListControlOptions` . Choose one of the following options:
        /// 
        /// - `MULTI_SELECT` : The user can select multiple entries from the list.
        /// - `SINGLE_SELECT` : The user can select a single entry from the list.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardSheetControlListType>? Type { get; set; }

        public DashboardDefaultFilterListControlOptionsArgs()
        {
        }
        public static new DashboardDefaultFilterListControlOptionsArgs Empty => new DashboardDefaultFilterListControlOptionsArgs();
    }
}