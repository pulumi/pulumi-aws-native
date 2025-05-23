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
    public sealed class TemplateDateDimensionField
    {
        /// <summary>
        /// The column that is used in the `DateDimensionField` .
        /// </summary>
        public readonly Outputs.TemplateColumnIdentifier Column;
        /// <summary>
        /// The date granularity of the `DateDimensionField` . Choose one of the following options:
        /// 
        /// - `YEAR`
        /// - `QUARTER`
        /// - `MONTH`
        /// - `WEEK`
        /// - `DAY`
        /// - `HOUR`
        /// - `MINUTE`
        /// - `SECOND`
        /// - `MILLISECOND`
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateTimeGranularity? DateGranularity;
        /// <summary>
        /// The custom field ID.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The format configuration of the field.
        /// </summary>
        public readonly Outputs.TemplateDateTimeFormatConfiguration? FormatConfiguration;
        /// <summary>
        /// The custom hierarchy ID.
        /// </summary>
        public readonly string? HierarchyId;

        [OutputConstructor]
        private TemplateDateDimensionField(
            Outputs.TemplateColumnIdentifier column,

            Pulumi.AwsNative.QuickSight.TemplateTimeGranularity? dateGranularity,

            string fieldId,

            Outputs.TemplateDateTimeFormatConfiguration? formatConfiguration,

            string? hierarchyId)
        {
            Column = column;
            DateGranularity = dateGranularity;
            FieldId = fieldId;
            FormatConfiguration = formatConfiguration;
            HierarchyId = hierarchyId;
        }
    }
}
