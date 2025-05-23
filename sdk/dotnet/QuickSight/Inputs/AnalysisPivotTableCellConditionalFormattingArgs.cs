// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPivotTableCellConditionalFormattingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field ID of the cell for conditional formatting.
        /// </summary>
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        /// <summary>
        /// The scope of the cell for conditional formatting.
        /// </summary>
        [Input("scope")]
        public Input<Inputs.AnalysisPivotTableConditionalFormattingScopeArgs>? Scope { get; set; }

        [Input("scopes")]
        private InputList<Inputs.AnalysisPivotTableConditionalFormattingScopeArgs>? _scopes;

        /// <summary>
        /// A list of cell scopes for conditional formatting.
        /// </summary>
        public InputList<Inputs.AnalysisPivotTableConditionalFormattingScopeArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.AnalysisPivotTableConditionalFormattingScopeArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// The text format of the cell for conditional formatting.
        /// </summary>
        [Input("textFormat")]
        public Input<Inputs.AnalysisTextConditionalFormatArgs>? TextFormat { get; set; }

        public AnalysisPivotTableCellConditionalFormattingArgs()
        {
        }
        public static new AnalysisPivotTableCellConditionalFormattingArgs Empty => new AnalysisPivotTableCellConditionalFormattingArgs();
    }
}
