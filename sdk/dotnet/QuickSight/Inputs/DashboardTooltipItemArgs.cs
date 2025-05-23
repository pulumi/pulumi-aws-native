// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTooltipItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tooltip item for the columns that are not part of a field well.
        /// </summary>
        [Input("columnTooltipItem")]
        public Input<Inputs.DashboardColumnTooltipItemArgs>? ColumnTooltipItem { get; set; }

        /// <summary>
        /// The tooltip item for the fields.
        /// </summary>
        [Input("fieldTooltipItem")]
        public Input<Inputs.DashboardFieldTooltipItemArgs>? FieldTooltipItem { get; set; }

        public DashboardTooltipItemArgs()
        {
        }
        public static new DashboardTooltipItemArgs Empty => new DashboardTooltipItemArgs();
    }
}
