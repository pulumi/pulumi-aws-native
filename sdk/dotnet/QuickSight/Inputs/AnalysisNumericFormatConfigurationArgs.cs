// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisNumericFormatConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The options that determine the currency display format configuration.
        /// </summary>
        [Input("currencyDisplayFormatConfiguration")]
        public Input<Inputs.AnalysisCurrencyDisplayFormatConfigurationArgs>? CurrencyDisplayFormatConfiguration { get; set; }

        /// <summary>
        /// The options that determine the number display format configuration.
        /// </summary>
        [Input("numberDisplayFormatConfiguration")]
        public Input<Inputs.AnalysisNumberDisplayFormatConfigurationArgs>? NumberDisplayFormatConfiguration { get; set; }

        /// <summary>
        /// The options that determine the percentage display format configuration.
        /// </summary>
        [Input("percentageDisplayFormatConfiguration")]
        public Input<Inputs.AnalysisPercentageDisplayFormatConfigurationArgs>? PercentageDisplayFormatConfiguration { get; set; }

        public AnalysisNumericFormatConfigurationArgs()
        {
        }
        public static new AnalysisNumericFormatConfigurationArgs Empty => new AnalysisNumericFormatConfigurationArgs();
    }
}
