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
    public sealed class DashboardPluginVisualSortConfiguration
    {
        /// <summary>
        /// The table query sorting options for the plugin visual.
        /// </summary>
        public readonly Outputs.DashboardPluginVisualTableQuerySort? PluginVisualTableQuerySort;

        [OutputConstructor]
        private DashboardPluginVisualSortConfiguration(Outputs.DashboardPluginVisualTableQuerySort? pluginVisualTableQuerySort)
        {
            PluginVisualTableQuerySort = pluginVisualTableQuerySort;
        }
    }
}
