// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    public sealed class RuleGroupIpSetReferenceStatementArgs : global::Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("ipSetForwardedIpConfig")]
        public Input<Inputs.RuleGroupIpSetForwardedIpConfigurationArgs>? IpSetForwardedIpConfig { get; set; }

        public RuleGroupIpSetReferenceStatementArgs()
        {
        }
        public static new RuleGroupIpSetReferenceStatementArgs Empty => new RuleGroupIpSetReferenceStatementArgs();
    }
}