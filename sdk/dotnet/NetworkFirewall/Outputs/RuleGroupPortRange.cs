// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupPortRange
    {
        public readonly int FromPort;
        public readonly int ToPort;

        [OutputConstructor]
        private RuleGroupPortRange(
            int fromPort,

            int toPort)
        {
            FromPort = fromPort;
            ToPort = toPort;
        }
    }
}