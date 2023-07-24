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
    public sealed class DashboardFilterOperationSelectedFieldsConfiguration
    {
        public readonly ImmutableArray<Outputs.DashboardColumnIdentifier> SelectedColumns;
        public readonly Pulumi.AwsNative.QuickSight.DashboardSelectedFieldOptions? SelectedFieldOptions;
        public readonly ImmutableArray<string> SelectedFields;

        [OutputConstructor]
        private DashboardFilterOperationSelectedFieldsConfiguration(
            ImmutableArray<Outputs.DashboardColumnIdentifier> selectedColumns,

            Pulumi.AwsNative.QuickSight.DashboardSelectedFieldOptions? selectedFieldOptions,

            ImmutableArray<string> selectedFields)
        {
            SelectedColumns = selectedColumns;
            SelectedFieldOptions = selectedFieldOptions;
            SelectedFields = selectedFields;
        }
    }
}