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
    public sealed class DashboardYAxisOptions
    {
        /// <summary>
        /// The Y axis type to be used in the chart.
        /// 
        /// If you choose `PRIMARY_Y_AXIS` , the primary Y Axis is located on the leftmost vertical axis of the chart.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardSingleYAxisOption YAxis;

        [OutputConstructor]
        private DashboardYAxisOptions(Pulumi.AwsNative.QuickSight.DashboardSingleYAxisOption yAxis)
        {
            YAxis = yAxis;
        }
    }
}
