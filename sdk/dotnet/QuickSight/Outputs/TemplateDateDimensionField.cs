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
    public sealed class TemplateDateDimensionField
    {
        public readonly Outputs.TemplateColumnIdentifier Column;
        public readonly Pulumi.AwsNative.QuickSight.TemplateTimeGranularity? DateGranularity;
        public readonly string FieldId;
        public readonly Outputs.TemplateDateTimeFormatConfiguration? FormatConfiguration;
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