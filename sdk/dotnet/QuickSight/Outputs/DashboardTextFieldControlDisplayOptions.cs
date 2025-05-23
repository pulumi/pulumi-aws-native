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
    public sealed class DashboardTextFieldControlDisplayOptions
    {
        /// <summary>
        /// The configuration of info icon label options.
        /// </summary>
        public readonly Outputs.DashboardSheetControlInfoIconLabelOptions? InfoIconLabelOptions;
        /// <summary>
        /// The configuration of the placeholder options in a text field control.
        /// </summary>
        public readonly Outputs.DashboardTextControlPlaceholderOptions? PlaceholderOptions;
        /// <summary>
        /// The options to configure the title visibility, name, and font size.
        /// </summary>
        public readonly Outputs.DashboardLabelOptions? TitleOptions;

        [OutputConstructor]
        private DashboardTextFieldControlDisplayOptions(
            Outputs.DashboardSheetControlInfoIconLabelOptions? infoIconLabelOptions,

            Outputs.DashboardTextControlPlaceholderOptions? placeholderOptions,

            Outputs.DashboardLabelOptions? titleOptions)
        {
            InfoIconLabelOptions = infoIconLabelOptions;
            PlaceholderOptions = placeholderOptions;
            TitleOptions = titleOptions;
        }
    }
}
