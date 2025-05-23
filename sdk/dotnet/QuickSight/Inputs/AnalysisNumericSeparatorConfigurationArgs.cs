// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisNumericSeparatorConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the decimal separator.
        /// </summary>
        [Input("decimalSeparator")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisNumericSeparatorSymbol>? DecimalSeparator { get; set; }

        /// <summary>
        /// The options that determine the thousands separator configuration.
        /// </summary>
        [Input("thousandsSeparator")]
        public Input<Inputs.AnalysisThousandSeparatorOptionsArgs>? ThousandsSeparator { get; set; }

        public AnalysisNumericSeparatorConfigurationArgs()
        {
        }
        public static new AnalysisNumericSeparatorConfigurationArgs Empty => new AnalysisNumericSeparatorConfigurationArgs();
    }
}
