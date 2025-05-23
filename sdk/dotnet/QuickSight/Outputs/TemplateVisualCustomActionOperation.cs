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
    public sealed class TemplateVisualCustomActionOperation
    {
        /// <summary>
        /// The filter operation that filters data included in a visual or in an entire sheet.
        /// </summary>
        public readonly Outputs.TemplateCustomActionFilterOperation? FilterOperation;
        /// <summary>
        /// The navigation operation that navigates between different sheets in the same analysis.
        /// </summary>
        public readonly Outputs.TemplateCustomActionNavigationOperation? NavigationOperation;
        /// <summary>
        /// The set parameter operation that sets parameters in custom action.
        /// </summary>
        public readonly Outputs.TemplateCustomActionSetParametersOperation? SetParametersOperation;
        /// <summary>
        /// The URL operation that opens a link to another webpage.
        /// </summary>
        public readonly Outputs.TemplateCustomActionUrlOperation? UrlOperation;

        [OutputConstructor]
        private TemplateVisualCustomActionOperation(
            Outputs.TemplateCustomActionFilterOperation? filterOperation,

            Outputs.TemplateCustomActionNavigationOperation? navigationOperation,

            Outputs.TemplateCustomActionSetParametersOperation? setParametersOperation,

            Outputs.TemplateCustomActionUrlOperation? urlOperation)
        {
            FilterOperation = filterOperation;
            NavigationOperation = navigationOperation;
            SetParametersOperation = setParametersOperation;
            UrlOperation = urlOperation;
        }
    }
}
