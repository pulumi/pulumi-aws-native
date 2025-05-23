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
    public sealed class AnalysisPivotTableFieldOptions
    {
        /// <summary>
        /// The collapse state options for the pivot table field options.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisPivotTableFieldCollapseStateOption> CollapseStateOptions;
        /// <summary>
        /// The data path options for the pivot table field options.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisPivotTableDataPathOption> DataPathOptions;
        /// <summary>
        /// The selected field options for the pivot table field options.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisPivotTableFieldOption> SelectedFieldOptions;

        [OutputConstructor]
        private AnalysisPivotTableFieldOptions(
            ImmutableArray<Outputs.AnalysisPivotTableFieldCollapseStateOption> collapseStateOptions,

            ImmutableArray<Outputs.AnalysisPivotTableDataPathOption> dataPathOptions,

            ImmutableArray<Outputs.AnalysisPivotTableFieldOption> selectedFieldOptions)
        {
            CollapseStateOptions = collapseStateOptions;
            DataPathOptions = dataPathOptions;
            SelectedFieldOptions = selectedFieldOptions;
        }
    }
}
