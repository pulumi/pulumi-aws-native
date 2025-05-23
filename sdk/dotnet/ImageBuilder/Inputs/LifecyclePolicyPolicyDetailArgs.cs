// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// The policy detail of the lifecycle policy.
    /// </summary>
    public sealed class LifecyclePolicyPolicyDetailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration details for the policy action.
        /// </summary>
        [Input("action", required: true)]
        public Input<Inputs.LifecyclePolicyActionArgs> Action { get; set; } = null!;

        /// <summary>
        /// Additional rules to specify resources that should be exempt from policy actions.
        /// </summary>
        [Input("exclusionRules")]
        public Input<Inputs.LifecyclePolicyExclusionRulesArgs>? ExclusionRules { get; set; }

        /// <summary>
        /// Specifies the resources that the lifecycle policy applies to.
        /// </summary>
        [Input("filter", required: true)]
        public Input<Inputs.LifecyclePolicyFilterArgs> Filter { get; set; } = null!;

        public LifecyclePolicyPolicyDetailArgs()
        {
        }
        public static new LifecyclePolicyPolicyDetailArgs Empty => new LifecyclePolicyPolicyDetailArgs();
    }
}
