// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPluginVisualItemsLimitConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines how many values are be fetched at once.
        /// </summary>
        [Input("itemsLimit")]
        public Input<double>? ItemsLimit { get; set; }

        public DashboardPluginVisualItemsLimitConfigurationArgs()
        {
        }
        public static new DashboardPluginVisualItemsLimitConfigurationArgs Empty => new DashboardPluginVisualItemsLimitConfigurationArgs();
    }
}
