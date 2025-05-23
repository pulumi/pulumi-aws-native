// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableFieldOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("order")]
        private InputList<string>? _order;

        /// <summary>
        /// The order of the field IDs that are configured as field options for a table visual.
        /// </summary>
        public InputList<string> Order
        {
            get => _order ?? (_order = new InputList<string>());
            set => _order = value;
        }

        /// <summary>
        /// The settings for the pinned columns of a table visual.
        /// </summary>
        [Input("pinnedFieldOptions")]
        public Input<Inputs.AnalysisTablePinnedFieldOptionsArgs>? PinnedFieldOptions { get; set; }

        [Input("selectedFieldOptions")]
        private InputList<Inputs.AnalysisTableFieldOptionArgs>? _selectedFieldOptions;

        /// <summary>
        /// The field options to be configured to a table.
        /// </summary>
        public InputList<Inputs.AnalysisTableFieldOptionArgs> SelectedFieldOptions
        {
            get => _selectedFieldOptions ?? (_selectedFieldOptions = new InputList<Inputs.AnalysisTableFieldOptionArgs>());
            set => _selectedFieldOptions = value;
        }

        [Input("transposedTableOptions")]
        private InputList<Inputs.AnalysisTransposedTableOptionArgs>? _transposedTableOptions;

        /// <summary>
        /// The `TableOptions` of a transposed table.
        /// </summary>
        public InputList<Inputs.AnalysisTransposedTableOptionArgs> TransposedTableOptions
        {
            get => _transposedTableOptions ?? (_transposedTableOptions = new InputList<Inputs.AnalysisTransposedTableOptionArgs>());
            set => _transposedTableOptions = value;
        }

        public AnalysisTableFieldOptionsArgs()
        {
        }
        public static new AnalysisTableFieldOptionsArgs Empty => new AnalysisTableFieldOptionsArgs();
    }
}
