// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Inputs
{

    public sealed class RuleActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("forward", required: true)]
        public Input<Inputs.RuleForwardArgs> Forward { get; set; } = null!;

        public RuleActionArgs()
        {
        }
        public static new RuleActionArgs Empty => new RuleActionArgs();
    }
}