// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// The action of the policy detail.
    /// </summary>
    public sealed class LifecyclePolicyActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("includeResources")]
        public Input<Inputs.LifecyclePolicyIncludeResourcesArgs>? IncludeResources { get; set; }

        /// <summary>
        /// The action type of the policy detail.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.ImageBuilder.LifecyclePolicyActionType> Type { get; set; } = null!;

        public LifecyclePolicyActionArgs()
        {
        }
        public static new LifecyclePolicyActionArgs Empty => new LifecyclePolicyActionArgs();
    }
}