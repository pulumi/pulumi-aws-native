// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupActionDefinition
    {
        /// <summary>
        /// Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. This setting defines a CloudWatch dimension value to be published.
        /// 
        /// You can pair this custom action with any of the standard stateless rule actions. For example, you could pair this in a rule action with the standard action that forwards the packet for stateful inspection. Then, when a packet matches the rule, Network Firewall publishes metrics for the packet and forwards it.
        /// </summary>
        public readonly Outputs.RuleGroupPublishMetricAction? PublishMetricAction;

        [OutputConstructor]
        private RuleGroupActionDefinition(Outputs.RuleGroupPublishMetricAction? publishMetricAction)
        {
            PublishMetricAction = publishMetricAction;
        }
    }
}
