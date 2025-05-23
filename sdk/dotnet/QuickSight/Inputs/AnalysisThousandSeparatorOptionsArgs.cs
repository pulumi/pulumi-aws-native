// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisThousandSeparatorOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the way numbers are styled to accommodate different readability standards. The `DEFAULT` value uses the standard international grouping system and groups numbers by the thousands. The `LAKHS` value uses the Indian numbering system and groups numbers by lakhs and crores.
        /// </summary>
        [Input("groupingStyle")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisDigitGroupingStyle>? GroupingStyle { get; set; }

        /// <summary>
        /// Determines the thousands separator symbol.
        /// </summary>
        [Input("symbol")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisNumericSeparatorSymbol>? Symbol { get; set; }

        /// <summary>
        /// Determines the visibility of the thousands separator.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        public AnalysisThousandSeparatorOptionsArgs()
        {
        }
        public static new AnalysisThousandSeparatorOptionsArgs Empty => new AnalysisThousandSeparatorOptionsArgs();
    }
}
