// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The first port in the Traffic Mirror port range.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// The last port in the Traffic Mirror port range.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs()
        {
        }
        public static new TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs Empty => new TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs();
    }
}
