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
    public sealed class DashboardTreeMapAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Colors;
        public readonly ImmutableArray<Outputs.DashboardDimensionField> Groups;
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Sizes;

        [OutputConstructor]
        private DashboardTreeMapAggregatedFieldWells(
            ImmutableArray<Outputs.DashboardMeasureField> colors,

            ImmutableArray<Outputs.DashboardDimensionField> groups,

            ImmutableArray<Outputs.DashboardMeasureField> sizes)
        {
            Colors = colors;
            Groups = groups;
            Sizes = sizes;
        }
    }
}