// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Outputs
{

    [OutputType]
    public sealed class FeatureVariationObject
    {
        public readonly bool? BooleanValue;
        public readonly double? DoubleValue;
        public readonly double? LongValue;
        public readonly string? StringValue;
        public readonly string? VariationName;

        [OutputConstructor]
        private FeatureVariationObject(
            bool? booleanValue,

            double? doubleValue,

            double? longValue,

            string? stringValue,

            string? variationName)
        {
            BooleanValue = booleanValue;
            DoubleValue = doubleValue;
            LongValue = longValue;
            StringValue = stringValue;
            VariationName = variationName;
        }
    }
}