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
    public sealed class TemplateDefaultNewSheetConfiguration
    {
        /// <summary>
        /// The options that determine the default settings for interactive layout configuration.
        /// </summary>
        public readonly Outputs.TemplateDefaultInteractiveLayoutConfiguration? InteractiveLayoutConfiguration;
        /// <summary>
        /// The options that determine the default settings for a paginated layout configuration.
        /// </summary>
        public readonly Outputs.TemplateDefaultPaginatedLayoutConfiguration? PaginatedLayoutConfiguration;
        /// <summary>
        /// The option that determines the sheet content type.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateSheetContentType? SheetContentType;

        [OutputConstructor]
        private TemplateDefaultNewSheetConfiguration(
            Outputs.TemplateDefaultInteractiveLayoutConfiguration? interactiveLayoutConfiguration,

            Outputs.TemplateDefaultPaginatedLayoutConfiguration? paginatedLayoutConfiguration,

            Pulumi.AwsNative.QuickSight.TemplateSheetContentType? sheetContentType)
        {
            InteractiveLayoutConfiguration = interactiveLayoutConfiguration;
            PaginatedLayoutConfiguration = paginatedLayoutConfiguration;
            SheetContentType = sheetContentType;
        }
    }
}
