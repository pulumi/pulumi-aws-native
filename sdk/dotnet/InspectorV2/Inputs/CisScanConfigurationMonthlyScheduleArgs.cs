// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Inputs
{

    public sealed class CisScanConfigurationMonthlyScheduleArgs : global::Pulumi.ResourceArgs
    {
        [Input("day", required: true)]
        public Input<Pulumi.AwsNative.InspectorV2.CisScanConfigurationDay> Day { get; set; } = null!;

        [Input("startTime", required: true)]
        public Input<Inputs.CisScanConfigurationTimeArgs> StartTime { get; set; } = null!;

        public CisScanConfigurationMonthlyScheduleArgs()
        {
        }
        public static new CisScanConfigurationMonthlyScheduleArgs Empty => new CisScanConfigurationMonthlyScheduleArgs();
    }
}