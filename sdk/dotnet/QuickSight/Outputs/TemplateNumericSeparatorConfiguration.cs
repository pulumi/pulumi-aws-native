// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateNumericSeparatorConfiguration
    {
        public readonly Pulumi.AwsNative.QuickSight.TemplateNumericSeparatorSymbol? DecimalSeparator;
        public readonly Outputs.TemplateThousandSeparatorOptions? ThousandsSeparator;

        [OutputConstructor]
        private TemplateNumericSeparatorConfiguration(
            Pulumi.AwsNative.QuickSight.TemplateNumericSeparatorSymbol? decimalSeparator,

            Outputs.TemplateThousandSeparatorOptions? thousandsSeparator)
        {
            DecimalSeparator = decimalSeparator;
            ThousandsSeparator = thousandsSeparator;
        }
    }
}