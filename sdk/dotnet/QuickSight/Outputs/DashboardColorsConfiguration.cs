// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardColorsConfiguration
    {
        public readonly ImmutableArray<Outputs.DashboardCustomColor> CustomColors;

        [OutputConstructor]
        private DashboardColorsConfiguration(ImmutableArray<Outputs.DashboardCustomColor> customColors)
        {
            CustomColors = customColors;
        }
    }
}