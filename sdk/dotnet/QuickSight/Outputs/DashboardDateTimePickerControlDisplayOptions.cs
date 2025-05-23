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
    public sealed class DashboardDateTimePickerControlDisplayOptions
    {
        /// <summary>
        /// The date icon visibility of the `DateTimePickerControlDisplayOptions` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? DateIconVisibility;
        /// <summary>
        /// Customize how dates are formatted in controls.
        /// </summary>
        public readonly string? DateTimeFormat;
        /// <summary>
        /// The helper text visibility of the `DateTimePickerControlDisplayOptions` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? HelperTextVisibility;
        /// <summary>
        /// The configuration of info icon label options.
        /// </summary>
        public readonly Outputs.DashboardSheetControlInfoIconLabelOptions? InfoIconLabelOptions;
        /// <summary>
        /// The options to configure the title visibility, name, and font size.
        /// </summary>
        public readonly Outputs.DashboardLabelOptions? TitleOptions;

        [OutputConstructor]
        private DashboardDateTimePickerControlDisplayOptions(
            Pulumi.AwsNative.QuickSight.DashboardVisibility? dateIconVisibility,

            string? dateTimeFormat,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? helperTextVisibility,

            Outputs.DashboardSheetControlInfoIconLabelOptions? infoIconLabelOptions,

            Outputs.DashboardLabelOptions? titleOptions)
        {
            DateIconVisibility = dateIconVisibility;
            DateTimeFormat = dateTimeFormat;
            HelperTextVisibility = helperTextVisibility;
            InfoIconLabelOptions = infoIconLabelOptions;
            TitleOptions = titleOptions;
        }
    }
}
