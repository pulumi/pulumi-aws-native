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
    public sealed class TemplateCategoryFilterConfiguration
    {
        /// <summary>
        /// A custom filter that filters based on a single value. This filter can be partially matched.
        /// </summary>
        public readonly Outputs.TemplateCustomFilterConfiguration? CustomFilterConfiguration;
        /// <summary>
        /// A list of custom filter values. In the Amazon QuickSight console, this filter type is called a custom filter list.
        /// </summary>
        public readonly Outputs.TemplateCustomFilterListConfiguration? CustomFilterListConfiguration;
        /// <summary>
        /// A list of filter configurations. In the Amazon QuickSight console, this filter type is called a filter list.
        /// </summary>
        public readonly Outputs.TemplateFilterListConfiguration? FilterListConfiguration;

        [OutputConstructor]
        private TemplateCategoryFilterConfiguration(
            Outputs.TemplateCustomFilterConfiguration? customFilterConfiguration,

            Outputs.TemplateCustomFilterListConfiguration? customFilterListConfiguration,

            Outputs.TemplateFilterListConfiguration? filterListConfiguration)
        {
            CustomFilterConfiguration = customFilterConfiguration;
            CustomFilterListConfiguration = customFilterListConfiguration;
            FilterListConfiguration = filterListConfiguration;
        }
    }
}
