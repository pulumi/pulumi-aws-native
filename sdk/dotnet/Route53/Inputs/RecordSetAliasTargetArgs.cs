// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    public sealed class RecordSetAliasTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dNSName", required: true)]
        public Input<string> DNSName { get; set; } = null!;

        [Input("evaluateTargetHealth")]
        public Input<bool>? EvaluateTargetHealth { get; set; }

        [Input("hostedZoneId", required: true)]
        public Input<string> HostedZoneId { get; set; } = null!;

        public RecordSetAliasTargetArgs()
        {
        }
        public static new RecordSetAliasTargetArgs Empty => new RecordSetAliasTargetArgs();
    }
}