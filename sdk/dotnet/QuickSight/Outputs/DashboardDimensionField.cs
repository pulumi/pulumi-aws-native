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
    public sealed class DashboardDimensionField
    {
        public readonly Outputs.DashboardCategoricalDimensionField? CategoricalDimensionField;
        public readonly Outputs.DashboardDateDimensionField? DateDimensionField;
        public readonly Outputs.DashboardNumericalDimensionField? NumericalDimensionField;

        [OutputConstructor]
        private DashboardDimensionField(
            Outputs.DashboardCategoricalDimensionField? categoricalDimensionField,

            Outputs.DashboardDateDimensionField? dateDimensionField,

            Outputs.DashboardNumericalDimensionField? numericalDimensionField)
        {
            CategoricalDimensionField = categoricalDimensionField;
            DateDimensionField = dateDimensionField;
            NumericalDimensionField = numericalDimensionField;
        }
    }
}