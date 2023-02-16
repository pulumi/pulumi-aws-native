// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Outputs
{

    [OutputType]
    public sealed class FilterNumberFilter
    {
        public readonly double? LowerInclusive;
        public readonly double? UpperInclusive;

        [OutputConstructor]
        private FilterNumberFilter(
            double? lowerInclusive,

            double? upperInclusive)
        {
            LowerInclusive = lowerInclusive;
            UpperInclusive = upperInclusive;
        }
    }
}