// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// The load balancer details if creating the AWS Verified Access endpoint as load-balancertype.
    /// </summary>
    public sealed class VerifiedAccessEndpointLoadBalancerOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the load balancer.
        /// </summary>
        [Input("loadBalancerArn")]
        public Input<string>? LoadBalancerArn { get; set; }

        /// <summary>
        /// The IP port number.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.VerifiedAccessEndpointPortRangeArgs>? _portRanges;

        /// <summary>
        /// The list of port range.
        /// </summary>
        public InputList<Inputs.VerifiedAccessEndpointPortRangeArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.VerifiedAccessEndpointPortRangeArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// The IP protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The IDs of the subnets.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public VerifiedAccessEndpointLoadBalancerOptionsArgs()
        {
        }
        public static new VerifiedAccessEndpointLoadBalancerOptionsArgs Empty => new VerifiedAccessEndpointLoadBalancerOptionsArgs();
    }
}
