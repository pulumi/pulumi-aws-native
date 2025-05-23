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
    public sealed class TemplatePercentageDisplayFormatConfiguration
    {
        /// <summary>
        /// The option that determines the decimal places configuration.
        /// </summary>
        public readonly Outputs.TemplateDecimalPlacesConfiguration? DecimalPlacesConfiguration;
        /// <summary>
        /// The options that determine the negative value configuration.
        /// </summary>
        public readonly Outputs.TemplateNegativeValueConfiguration? NegativeValueConfiguration;
        /// <summary>
        /// The options that determine the null value format configuration.
        /// </summary>
        public readonly Outputs.TemplateNullValueFormatConfiguration? NullValueFormatConfiguration;
        /// <summary>
        /// Determines the prefix value of the percentage format.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// The options that determine the numeric separator configuration.
        /// </summary>
        public readonly Outputs.TemplateNumericSeparatorConfiguration? SeparatorConfiguration;
        /// <summary>
        /// Determines the suffix value of the percentage format.
        /// </summary>
        public readonly string? Suffix;

        [OutputConstructor]
        private TemplatePercentageDisplayFormatConfiguration(
            Outputs.TemplateDecimalPlacesConfiguration? decimalPlacesConfiguration,

            Outputs.TemplateNegativeValueConfiguration? negativeValueConfiguration,

            Outputs.TemplateNullValueFormatConfiguration? nullValueFormatConfiguration,

            string? prefix,

            Outputs.TemplateNumericSeparatorConfiguration? separatorConfiguration,

            string? suffix)
        {
            DecimalPlacesConfiguration = decimalPlacesConfiguration;
            NegativeValueConfiguration = negativeValueConfiguration;
            NullValueFormatConfiguration = nullValueFormatConfiguration;
            Prefix = prefix;
            SeparatorConfiguration = separatorConfiguration;
            Suffix = suffix;
        }
    }
}
