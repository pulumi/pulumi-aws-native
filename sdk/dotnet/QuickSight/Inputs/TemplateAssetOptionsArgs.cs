// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAssetOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the timezone for the analysis.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// Determines the week start day for an analysis.
        /// </summary>
        [Input("weekStart")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateDayOfTheWeek>? WeekStart { get; set; }

        public TemplateAssetOptionsArgs()
        {
        }
        public static new TemplateAssetOptionsArgs Empty => new TemplateAssetOptionsArgs();
    }
}
