// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFilledMapFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The aggregated field well of the filled map.
        /// </summary>
        [Input("filledMapAggregatedFieldWells")]
        public Input<Inputs.DashboardFilledMapAggregatedFieldWellsArgs>? FilledMapAggregatedFieldWells { get; set; }

        public DashboardFilledMapFieldWellsArgs()
        {
        }
        public static new DashboardFilledMapFieldWellsArgs Empty => new DashboardFilledMapFieldWellsArgs();
    }
}
