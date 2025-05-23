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
    public sealed class DashboardSliderControlDisplayOptions
    {
        /// <summary>
        /// The configuration of info icon label options.
        /// </summary>
        public readonly Outputs.DashboardSheetControlInfoIconLabelOptions? InfoIconLabelOptions;
        /// <summary>
        /// The options to configure the title visibility, name, and font size.
        /// </summary>
        public readonly Outputs.DashboardLabelOptions? TitleOptions;

        [OutputConstructor]
        private DashboardSliderControlDisplayOptions(
            Outputs.DashboardSheetControlInfoIconLabelOptions? infoIconLabelOptions,

            Outputs.DashboardLabelOptions? titleOptions)
        {
            InfoIconLabelOptions = infoIconLabelOptions;
            TitleOptions = titleOptions;
        }
    }
}
