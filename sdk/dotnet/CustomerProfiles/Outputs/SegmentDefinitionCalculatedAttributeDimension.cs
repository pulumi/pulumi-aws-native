// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    /// <summary>
    /// Specifies calculated attribute based criteria for a segment.
    /// </summary>
    [OutputType]
    public sealed class SegmentDefinitionCalculatedAttributeDimension
    {
        public readonly Outputs.SegmentDefinitionConditionOverrides? ConditionOverrides;
        public readonly Pulumi.AwsNative.CustomerProfiles.SegmentDefinitionAttributeDimensionType DimensionType;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private SegmentDefinitionCalculatedAttributeDimension(
            Outputs.SegmentDefinitionConditionOverrides? conditionOverrides,

            Pulumi.AwsNative.CustomerProfiles.SegmentDefinitionAttributeDimensionType dimensionType,

            ImmutableArray<string> values)
        {
            ConditionOverrides = conditionOverrides;
            DimensionType = dimensionType;
            Values = values;
        }
    }
}
