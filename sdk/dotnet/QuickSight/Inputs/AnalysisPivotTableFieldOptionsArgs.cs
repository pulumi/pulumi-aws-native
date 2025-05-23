// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPivotTableFieldOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("collapseStateOptions")]
        private InputList<Inputs.AnalysisPivotTableFieldCollapseStateOptionArgs>? _collapseStateOptions;

        /// <summary>
        /// The collapse state options for the pivot table field options.
        /// </summary>
        public InputList<Inputs.AnalysisPivotTableFieldCollapseStateOptionArgs> CollapseStateOptions
        {
            get => _collapseStateOptions ?? (_collapseStateOptions = new InputList<Inputs.AnalysisPivotTableFieldCollapseStateOptionArgs>());
            set => _collapseStateOptions = value;
        }

        [Input("dataPathOptions")]
        private InputList<Inputs.AnalysisPivotTableDataPathOptionArgs>? _dataPathOptions;

        /// <summary>
        /// The data path options for the pivot table field options.
        /// </summary>
        public InputList<Inputs.AnalysisPivotTableDataPathOptionArgs> DataPathOptions
        {
            get => _dataPathOptions ?? (_dataPathOptions = new InputList<Inputs.AnalysisPivotTableDataPathOptionArgs>());
            set => _dataPathOptions = value;
        }

        [Input("selectedFieldOptions")]
        private InputList<Inputs.AnalysisPivotTableFieldOptionArgs>? _selectedFieldOptions;

        /// <summary>
        /// The selected field options for the pivot table field options.
        /// </summary>
        public InputList<Inputs.AnalysisPivotTableFieldOptionArgs> SelectedFieldOptions
        {
            get => _selectedFieldOptions ?? (_selectedFieldOptions = new InputList<Inputs.AnalysisPivotTableFieldOptionArgs>());
            set => _selectedFieldOptions = value;
        }

        public AnalysisPivotTableFieldOptionsArgs()
        {
        }
        public static new AnalysisPivotTableFieldOptionsArgs Empty => new AnalysisPivotTableFieldOptionsArgs();
    }
}
