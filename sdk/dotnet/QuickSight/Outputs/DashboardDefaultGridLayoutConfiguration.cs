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
    public sealed class DashboardDefaultGridLayoutConfiguration
    {
        /// <summary>
        /// Determines the screen canvas size options for a grid layout.
        /// </summary>
        public readonly Outputs.DashboardGridLayoutCanvasSizeOptions CanvasSizeOptions;

        [OutputConstructor]
        private DashboardDefaultGridLayoutConfiguration(Outputs.DashboardGridLayoutCanvasSizeOptions canvasSizeOptions)
        {
            CanvasSizeOptions = canvasSizeOptions;
        }
    }
}
