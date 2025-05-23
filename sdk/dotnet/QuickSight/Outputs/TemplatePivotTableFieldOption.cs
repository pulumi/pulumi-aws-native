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
    public sealed class TemplatePivotTableFieldOption
    {
        /// <summary>
        /// The custom label of the pivot table field.
        /// </summary>
        public readonly string? CustomLabel;
        /// <summary>
        /// The field ID of the pivot table field.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The visibility of the pivot table field.
        /// </summary>
        public readonly object? Visibility;

        [OutputConstructor]
        private TemplatePivotTableFieldOption(
            string? customLabel,

            string fieldId,

            object? visibility)
        {
            CustomLabel = customLabel;
            FieldId = fieldId;
            Visibility = visibility;
        }
    }
}
