// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dlm.Inputs
{

    public sealed class LifecyclePolicyArchiveRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("retainRule", required: true)]
        public Input<Inputs.LifecyclePolicyArchiveRetainRuleArgs> RetainRule { get; set; } = null!;

        public LifecyclePolicyArchiveRuleArgs()
        {
        }
        public static new LifecyclePolicyArchiveRuleArgs Empty => new LifecyclePolicyArchiveRuleArgs();
    }
}