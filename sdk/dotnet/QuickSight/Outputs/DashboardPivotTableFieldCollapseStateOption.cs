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
    public sealed class DashboardPivotTableFieldCollapseStateOption
    {
        /// <summary>
        /// The state of the field target of a pivot table. Choose one of the following options:
        /// 
        /// - `COLLAPSED`
        /// - `EXPANDED`
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardPivotTableFieldCollapseState? State;
        /// <summary>
        /// A tagged-union object that sets the collapse state.
        /// </summary>
        public readonly Outputs.DashboardPivotTableFieldCollapseStateTarget Target;

        [OutputConstructor]
        private DashboardPivotTableFieldCollapseStateOption(
            Pulumi.AwsNative.QuickSight.DashboardPivotTableFieldCollapseState? state,

            Outputs.DashboardPivotTableFieldCollapseStateTarget target)
        {
            State = state;
            Target = target;
        }
    }
}
