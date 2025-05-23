// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableConditionalFormattingArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditionalFormattingOptions")]
        private InputList<Inputs.DashboardPivotTableConditionalFormattingOptionArgs>? _conditionalFormattingOptions;

        /// <summary>
        /// Conditional formatting options for a `PivotTableVisual` .
        /// </summary>
        public InputList<Inputs.DashboardPivotTableConditionalFormattingOptionArgs> ConditionalFormattingOptions
        {
            get => _conditionalFormattingOptions ?? (_conditionalFormattingOptions = new InputList<Inputs.DashboardPivotTableConditionalFormattingOptionArgs>());
            set => _conditionalFormattingOptions = value;
        }

        public DashboardPivotTableConditionalFormattingArgs()
        {
        }
        public static new DashboardPivotTableConditionalFormattingArgs Empty => new DashboardPivotTableConditionalFormattingArgs();
    }
}
