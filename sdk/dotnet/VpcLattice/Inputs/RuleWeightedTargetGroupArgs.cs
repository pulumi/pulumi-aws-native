// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Inputs
{

    public sealed class RuleWeightedTargetGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetGroupIdentifier", required: true)]
        public Input<string> TargetGroupIdentifier { get; set; } = null!;

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public RuleWeightedTargetGroupArgs()
        {
        }
        public static new RuleWeightedTargetGroupArgs Empty => new RuleWeightedTargetGroupArgs();
    }
}