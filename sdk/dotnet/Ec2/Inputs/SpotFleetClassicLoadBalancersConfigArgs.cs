// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class SpotFleetClassicLoadBalancersConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("classicLoadBalancers", required: true)]
        private InputList<Inputs.SpotFleetClassicLoadBalancerArgs>? _classicLoadBalancers;

        /// <summary>
        /// One or more Classic Load Balancers.
        /// </summary>
        public InputList<Inputs.SpotFleetClassicLoadBalancerArgs> ClassicLoadBalancers
        {
            get => _classicLoadBalancers ?? (_classicLoadBalancers = new InputList<Inputs.SpotFleetClassicLoadBalancerArgs>());
            set => _classicLoadBalancers = value;
        }

        public SpotFleetClassicLoadBalancersConfigArgs()
        {
        }
        public static new SpotFleetClassicLoadBalancersConfigArgs Empty => new SpotFleetClassicLoadBalancersConfigArgs();
    }
}
