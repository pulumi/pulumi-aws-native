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
    public sealed class TemplatePieChartSortConfiguration
    {
        /// <summary>
        /// The limit on the number of categories that are displayed in a pie chart.
        /// </summary>
        public readonly Outputs.TemplateItemsLimitConfiguration? CategoryItemsLimit;
        /// <summary>
        /// The sort configuration of the category fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> CategorySort;
        /// <summary>
        /// The limit on the number of small multiples panels that are displayed.
        /// </summary>
        public readonly Outputs.TemplateItemsLimitConfiguration? SmallMultiplesLimitConfiguration;
        /// <summary>
        /// The sort configuration of the small multiples field.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> SmallMultiplesSort;

        [OutputConstructor]
        private TemplatePieChartSortConfiguration(
            Outputs.TemplateItemsLimitConfiguration? categoryItemsLimit,

            ImmutableArray<Outputs.TemplateFieldSortOptions> categorySort,

            Outputs.TemplateItemsLimitConfiguration? smallMultiplesLimitConfiguration,

            ImmutableArray<Outputs.TemplateFieldSortOptions> smallMultiplesSort)
        {
            CategoryItemsLimit = categoryItemsLimit;
            CategorySort = categorySort;
            SmallMultiplesLimitConfiguration = smallMultiplesLimitConfiguration;
            SmallMultiplesSort = smallMultiplesSort;
        }
    }
}
