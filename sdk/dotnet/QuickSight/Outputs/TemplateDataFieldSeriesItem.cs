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
    public sealed class TemplateDataFieldSeriesItem
    {
        /// <summary>
        /// The axis that you are binding the field to.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateAxisBinding AxisBinding;
        /// <summary>
        /// The field ID of the field that you are setting the axis binding to.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The field value of the field that you are setting the axis binding to.
        /// </summary>
        public readonly string? FieldValue;
        /// <summary>
        /// The options that determine the presentation of line series associated to the field.
        /// </summary>
        public readonly Outputs.TemplateLineChartSeriesSettings? Settings;

        [OutputConstructor]
        private TemplateDataFieldSeriesItem(
            Pulumi.AwsNative.QuickSight.TemplateAxisBinding axisBinding,

            string fieldId,

            string? fieldValue,

            Outputs.TemplateLineChartSeriesSettings? settings)
        {
            AxisBinding = axisBinding;
            FieldId = fieldId;
            FieldValue = fieldValue;
            Settings = settings;
        }
    }
}
