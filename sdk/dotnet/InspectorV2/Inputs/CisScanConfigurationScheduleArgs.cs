// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Inputs
{

    /// <summary>
    /// Choose a Schedule cadence
    /// </summary>
    public sealed class CisScanConfigurationScheduleArgs : global::Pulumi.ResourceArgs
    {
        [Input("daily")]
        public Input<Inputs.CisScanConfigurationDailyScheduleArgs>? Daily { get; set; }

        [Input("monthly")]
        public Input<Inputs.CisScanConfigurationMonthlyScheduleArgs>? Monthly { get; set; }

        [Input("oneTime")]
        public Input<Inputs.CisScanConfigurationOneTimeScheduleArgs>? OneTime { get; set; }

        [Input("weekly")]
        public Input<Inputs.CisScanConfigurationWeeklyScheduleArgs>? Weekly { get; set; }

        public CisScanConfigurationScheduleArgs()
        {
        }
        public static new CisScanConfigurationScheduleArgs Empty => new CisScanConfigurationScheduleArgs();
    }
}
