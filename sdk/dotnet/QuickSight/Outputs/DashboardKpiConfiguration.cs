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
    public sealed class DashboardKpiConfiguration
    {
        public readonly Outputs.DashboardKpiFieldWells? FieldWells;
        public readonly Outputs.DashboardKpiOptions? KpiOptions;
        public readonly Outputs.DashboardKpiSortConfiguration? SortConfiguration;

        [OutputConstructor]
        private DashboardKpiConfiguration(
            Outputs.DashboardKpiFieldWells? fieldWells,

            Outputs.DashboardKpiOptions? kpiOptions,

            Outputs.DashboardKpiSortConfiguration? sortConfiguration)
        {
            FieldWells = fieldWells;
            KpiOptions = kpiOptions;
            SortConfiguration = sortConfiguration;
        }
    }
}