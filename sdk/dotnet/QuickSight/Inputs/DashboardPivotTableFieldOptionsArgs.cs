// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableFieldOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataPathOptions")]
        private InputList<Inputs.DashboardPivotTableDataPathOptionArgs>? _dataPathOptions;
        public InputList<Inputs.DashboardPivotTableDataPathOptionArgs> DataPathOptions
        {
            get => _dataPathOptions ?? (_dataPathOptions = new InputList<Inputs.DashboardPivotTableDataPathOptionArgs>());
            set => _dataPathOptions = value;
        }

        [Input("selectedFieldOptions")]
        private InputList<Inputs.DashboardPivotTableFieldOptionArgs>? _selectedFieldOptions;
        public InputList<Inputs.DashboardPivotTableFieldOptionArgs> SelectedFieldOptions
        {
            get => _selectedFieldOptions ?? (_selectedFieldOptions = new InputList<Inputs.DashboardPivotTableFieldOptionArgs>());
            set => _selectedFieldOptions = value;
        }

        public DashboardPivotTableFieldOptionsArgs()
        {
        }
        public static new DashboardPivotTableFieldOptionsArgs Empty => new DashboardPivotTableFieldOptionsArgs();
    }
}