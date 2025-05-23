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
    public sealed class AnalysisNumericalDimensionField
    {
        /// <summary>
        /// The column that is used in the `NumericalDimensionField` .
        /// </summary>
        public readonly Outputs.AnalysisColumnIdentifier Column;
        /// <summary>
        /// The custom field ID.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The format configuration of the field.
        /// </summary>
        public readonly Outputs.AnalysisNumberFormatConfiguration? FormatConfiguration;
        /// <summary>
        /// The custom hierarchy ID.
        /// </summary>
        public readonly string? HierarchyId;

        [OutputConstructor]
        private AnalysisNumericalDimensionField(
            Outputs.AnalysisColumnIdentifier column,

            string fieldId,

            Outputs.AnalysisNumberFormatConfiguration? formatConfiguration,

            string? hierarchyId)
        {
            Column = column;
            FieldId = fieldId;
            FormatConfiguration = formatConfiguration;
            HierarchyId = hierarchyId;
        }
    }
}
