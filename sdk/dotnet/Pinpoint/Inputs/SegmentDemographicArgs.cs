// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class SegmentDemographicArgs : global::Pulumi.ResourceArgs
    {
        [Input("appVersion")]
        public Input<Inputs.SegmentSetDimensionArgs>? AppVersion { get; set; }

        [Input("channel")]
        public Input<Inputs.SegmentSetDimensionArgs>? Channel { get; set; }

        [Input("deviceType")]
        public Input<Inputs.SegmentSetDimensionArgs>? DeviceType { get; set; }

        [Input("make")]
        public Input<Inputs.SegmentSetDimensionArgs>? Make { get; set; }

        [Input("model")]
        public Input<Inputs.SegmentSetDimensionArgs>? Model { get; set; }

        [Input("platform")]
        public Input<Inputs.SegmentSetDimensionArgs>? Platform { get; set; }

        public SegmentDemographicArgs()
        {
        }
        public static new SegmentDemographicArgs Empty => new SegmentDemographicArgs();
    }
}