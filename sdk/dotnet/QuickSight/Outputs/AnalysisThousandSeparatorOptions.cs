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
    public sealed class AnalysisThousandSeparatorOptions
    {
        public readonly Pulumi.AwsNative.QuickSight.AnalysisNumericSeparatorSymbol? Symbol;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? Visibility;

        [OutputConstructor]
        private AnalysisThousandSeparatorOptions(
            Pulumi.AwsNative.QuickSight.AnalysisNumericSeparatorSymbol? symbol,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? visibility)
        {
            Symbol = symbol;
            Visibility = visibility;
        }
    }
}