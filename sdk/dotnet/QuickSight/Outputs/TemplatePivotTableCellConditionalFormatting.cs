// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplatePivotTableCellConditionalFormatting
    {
        /// <summary>
        /// The field ID of the cell for conditional formatting.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The scope of the cell for conditional formatting.
        /// </summary>
        public readonly Outputs.TemplatePivotTableConditionalFormattingScope? Scope;
        /// <summary>
        /// A list of cell scopes for conditional formatting.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplatePivotTableConditionalFormattingScope> Scopes;
        /// <summary>
        /// The text format of the cell for conditional formatting.
        /// </summary>
        public readonly Outputs.TemplateTextConditionalFormat? TextFormat;

        [OutputConstructor]
        private TemplatePivotTableCellConditionalFormatting(
            string fieldId,

            Outputs.TemplatePivotTableConditionalFormattingScope? scope,

            ImmutableArray<Outputs.TemplatePivotTableConditionalFormattingScope> scopes,

            Outputs.TemplateTextConditionalFormat? textFormat)
        {
            FieldId = fieldId;
            Scope = scope;
            Scopes = scopes;
            TextFormat = textFormat;
        }
    }
}
