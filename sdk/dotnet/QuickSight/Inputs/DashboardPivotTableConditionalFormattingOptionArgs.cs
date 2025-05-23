// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableConditionalFormattingOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cell conditional formatting option for a pivot table.
        /// </summary>
        [Input("cell")]
        public Input<Inputs.DashboardPivotTableCellConditionalFormattingArgs>? Cell { get; set; }

        public DashboardPivotTableConditionalFormattingOptionArgs()
        {
        }
        public static new DashboardPivotTableConditionalFormattingOptionArgs Empty => new DashboardPivotTableConditionalFormattingOptionArgs();
    }
}
