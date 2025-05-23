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
    public sealed class TemplatePluginVisualSortConfiguration
    {
        /// <summary>
        /// The table query sorting options for the plugin visual.
        /// </summary>
        public readonly Outputs.TemplatePluginVisualTableQuerySort? PluginVisualTableQuerySort;

        [OutputConstructor]
        private TemplatePluginVisualSortConfiguration(Outputs.TemplatePluginVisualTableQuerySort? pluginVisualTableQuerySort)
        {
            PluginVisualTableQuerySort = pluginVisualTableQuerySort;
        }
    }
}
