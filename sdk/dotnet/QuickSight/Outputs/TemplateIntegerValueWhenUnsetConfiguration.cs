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
    public sealed class TemplateIntegerValueWhenUnsetConfiguration
    {
        public readonly double? CustomValue;
        public readonly Pulumi.AwsNative.QuickSight.TemplateValueWhenUnsetOption? ValueWhenUnsetOption;

        [OutputConstructor]
        private TemplateIntegerValueWhenUnsetConfiguration(
            double? customValue,

            Pulumi.AwsNative.QuickSight.TemplateValueWhenUnsetOption? valueWhenUnsetOption)
        {
            CustomValue = customValue;
            ValueWhenUnsetOption = valueWhenUnsetOption;
        }
    }
}