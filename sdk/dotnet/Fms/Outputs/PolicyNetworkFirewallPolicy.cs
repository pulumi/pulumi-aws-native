// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fms.Outputs
{

    /// <summary>
    /// Network firewall policy.
    /// </summary>
    [OutputType]
    public sealed class PolicyNetworkFirewallPolicy
    {
        public readonly Pulumi.AwsNative.Fms.PolicyFirewallDeploymentModel FirewallDeploymentModel;

        [OutputConstructor]
        private PolicyNetworkFirewallPolicy(Pulumi.AwsNative.Fms.PolicyFirewallDeploymentModel firewallDeploymentModel)
        {
            FirewallDeploymentModel = firewallDeploymentModel;
        }
    }
}