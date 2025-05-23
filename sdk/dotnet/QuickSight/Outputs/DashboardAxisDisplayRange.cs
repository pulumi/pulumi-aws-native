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
    public sealed class DashboardAxisDisplayRange
    {
        /// <summary>
        /// The data-driven setup of an axis display range.
        /// </summary>
        public readonly Outputs.DashboardAxisDisplayDataDrivenRange? DataDriven;
        /// <summary>
        /// The minimum and maximum setup of an axis display range.
        /// </summary>
        public readonly Outputs.DashboardAxisDisplayMinMaxRange? MinMax;

        [OutputConstructor]
        private DashboardAxisDisplayRange(
            Outputs.DashboardAxisDisplayDataDrivenRange? dataDriven,

            Outputs.DashboardAxisDisplayMinMaxRange? minMax)
        {
            DataDriven = dataDriven;
            MinMax = minMax;
        }
    }
}
