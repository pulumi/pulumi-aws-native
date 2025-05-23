// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardGeospatialMapFieldWells
    {
        /// <summary>
        /// The aggregated field well for a geospatial map.
        /// </summary>
        public readonly Outputs.DashboardGeospatialMapAggregatedFieldWells? GeospatialMapAggregatedFieldWells;

        [OutputConstructor]
        private DashboardGeospatialMapFieldWells(Outputs.DashboardGeospatialMapAggregatedFieldWells? geospatialMapAggregatedFieldWells)
        {
            GeospatialMapAggregatedFieldWells = geospatialMapAggregatedFieldWells;
        }
    }
}
