// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Outputs
{

    [OutputType]
    public sealed class FilterDateFilter
    {
        public readonly int? EndInclusive;
        public readonly int? StartInclusive;

        [OutputConstructor]
        private FilterDateFilter(
            int? endInclusive,

            int? startInclusive)
        {
            EndInclusive = endInclusive;
            StartInclusive = startInclusive;
        }
    }
}
