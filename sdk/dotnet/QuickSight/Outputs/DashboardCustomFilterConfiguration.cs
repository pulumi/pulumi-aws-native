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
    public sealed class DashboardCustomFilterConfiguration
    {
        public readonly string? CategoryValue;
        public readonly Pulumi.AwsNative.QuickSight.DashboardCategoryFilterMatchOperator MatchOperator;
        public readonly Pulumi.AwsNative.QuickSight.DashboardFilterNullOption NullOption;
        public readonly string? ParameterName;
        public readonly Pulumi.AwsNative.QuickSight.DashboardCategoryFilterSelectAllOptions? SelectAllOptions;

        [OutputConstructor]
        private DashboardCustomFilterConfiguration(
            string? categoryValue,

            Pulumi.AwsNative.QuickSight.DashboardCategoryFilterMatchOperator matchOperator,

            Pulumi.AwsNative.QuickSight.DashboardFilterNullOption nullOption,

            string? parameterName,

            Pulumi.AwsNative.QuickSight.DashboardCategoryFilterSelectAllOptions? selectAllOptions)
        {
            CategoryValue = categoryValue;
            MatchOperator = matchOperator;
            NullOption = nullOption;
            ParameterName = parameterName;
            SelectAllOptions = selectAllOptions;
        }
    }
}