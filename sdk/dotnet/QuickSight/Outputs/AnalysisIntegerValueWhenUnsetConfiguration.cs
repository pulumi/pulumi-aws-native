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
    public sealed class AnalysisIntegerValueWhenUnsetConfiguration
    {
        public readonly double? CustomValue;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisValueWhenUnsetOption? ValueWhenUnsetOption;

        [OutputConstructor]
        private AnalysisIntegerValueWhenUnsetConfiguration(
            double? customValue,

            Pulumi.AwsNative.QuickSight.AnalysisValueWhenUnsetOption? valueWhenUnsetOption)
        {
            CustomValue = customValue;
            ValueWhenUnsetOption = valueWhenUnsetOption;
        }
    }
}