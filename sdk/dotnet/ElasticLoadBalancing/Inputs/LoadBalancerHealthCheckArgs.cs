// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancing.Inputs
{

    public sealed class LoadBalancerHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        [Input("healthyThreshold", required: true)]
        public Input<string> HealthyThreshold { get; set; } = null!;

        [Input("interval", required: true)]
        public Input<string> Interval { get; set; } = null!;

        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        [Input("timeout", required: true)]
        public Input<string> Timeout { get; set; } = null!;

        [Input("unhealthyThreshold", required: true)]
        public Input<string> UnhealthyThreshold { get; set; } = null!;

        public LoadBalancerHealthCheckArgs()
        {
        }
        public static new LoadBalancerHealthCheckArgs Empty => new LoadBalancerHealthCheckArgs();
    }
}