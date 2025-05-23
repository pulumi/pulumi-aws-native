// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePercentageDisplayFormatConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The option that determines the decimal places configuration.
        /// </summary>
        [Input("decimalPlacesConfiguration")]
        public Input<Inputs.TemplateDecimalPlacesConfigurationArgs>? DecimalPlacesConfiguration { get; set; }

        /// <summary>
        /// The options that determine the negative value configuration.
        /// </summary>
        [Input("negativeValueConfiguration")]
        public Input<Inputs.TemplateNegativeValueConfigurationArgs>? NegativeValueConfiguration { get; set; }

        /// <summary>
        /// The options that determine the null value format configuration.
        /// </summary>
        [Input("nullValueFormatConfiguration")]
        public Input<Inputs.TemplateNullValueFormatConfigurationArgs>? NullValueFormatConfiguration { get; set; }

        /// <summary>
        /// Determines the prefix value of the percentage format.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The options that determine the numeric separator configuration.
        /// </summary>
        [Input("separatorConfiguration")]
        public Input<Inputs.TemplateNumericSeparatorConfigurationArgs>? SeparatorConfiguration { get; set; }

        /// <summary>
        /// Determines the suffix value of the percentage format.
        /// </summary>
        [Input("suffix")]
        public Input<string>? Suffix { get; set; }

        public TemplatePercentageDisplayFormatConfigurationArgs()
        {
        }
        public static new TemplatePercentageDisplayFormatConfigurationArgs Empty => new TemplatePercentageDisplayFormatConfigurationArgs();
    }
}
