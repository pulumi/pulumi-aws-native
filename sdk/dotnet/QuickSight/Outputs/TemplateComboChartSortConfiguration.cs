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
    public sealed class TemplateComboChartSortConfiguration
    {
        /// <summary>
        /// The item limit configuration for the category field well of a combo chart.
        /// </summary>
        public readonly Outputs.TemplateItemsLimitConfiguration? CategoryItemsLimit;
        /// <summary>
        /// The sort configuration of the category field well in a combo chart.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> CategorySort;
        /// <summary>
        /// The item limit configuration of the color field well in a combo chart.
        /// </summary>
        public readonly Outputs.TemplateItemsLimitConfiguration? ColorItemsLimit;
        /// <summary>
        /// The sort configuration of the color field well in a combo chart.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> ColorSort;

        [OutputConstructor]
        private TemplateComboChartSortConfiguration(
            Outputs.TemplateItemsLimitConfiguration? categoryItemsLimit,

            ImmutableArray<Outputs.TemplateFieldSortOptions> categorySort,

            Outputs.TemplateItemsLimitConfiguration? colorItemsLimit,

            ImmutableArray<Outputs.TemplateFieldSortOptions> colorSort)
        {
            CategoryItemsLimit = categoryItemsLimit;
            CategorySort = categorySort;
            ColorItemsLimit = colorItemsLimit;
            ColorSort = colorSort;
        }
    }
}
