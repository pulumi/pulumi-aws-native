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
    public sealed class AnalysisNumericFormatConfiguration
    {
        /// <summary>
        /// The options that determine the currency display format configuration.
        /// </summary>
        public readonly Outputs.AnalysisCurrencyDisplayFormatConfiguration? CurrencyDisplayFormatConfiguration;
        /// <summary>
        /// The options that determine the number display format configuration.
        /// </summary>
        public readonly Outputs.AnalysisNumberDisplayFormatConfiguration? NumberDisplayFormatConfiguration;
        /// <summary>
        /// The options that determine the percentage display format configuration.
        /// </summary>
        public readonly Outputs.AnalysisPercentageDisplayFormatConfiguration? PercentageDisplayFormatConfiguration;

        [OutputConstructor]
        private AnalysisNumericFormatConfiguration(
            Outputs.AnalysisCurrencyDisplayFormatConfiguration? currencyDisplayFormatConfiguration,

            Outputs.AnalysisNumberDisplayFormatConfiguration? numberDisplayFormatConfiguration,

            Outputs.AnalysisPercentageDisplayFormatConfiguration? percentageDisplayFormatConfiguration)
        {
            CurrencyDisplayFormatConfiguration = currencyDisplayFormatConfiguration;
            NumberDisplayFormatConfiguration = numberDisplayFormatConfiguration;
            PercentageDisplayFormatConfiguration = percentageDisplayFormatConfiguration;
        }
    }
}
