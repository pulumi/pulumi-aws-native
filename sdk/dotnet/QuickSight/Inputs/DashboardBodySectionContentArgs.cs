// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardBodySectionContentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The layout configuration of a body section.
        /// </summary>
        [Input("layout")]
        public Input<Inputs.DashboardSectionLayoutConfigurationArgs>? Layout { get; set; }

        public DashboardBodySectionContentArgs()
        {
        }
        public static new DashboardBodySectionContentArgs Empty => new DashboardBodySectionContentArgs();
    }
}
