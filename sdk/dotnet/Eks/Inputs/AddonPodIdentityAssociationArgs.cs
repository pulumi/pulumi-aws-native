// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Inputs
{

    /// <summary>
    /// A pod identity to associate with an add-on.
    /// </summary>
    public sealed class AddonPodIdentityAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM role ARN that the pod identity association is created for.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The Kubernetes service account that the pod identity association is created for.
        /// </summary>
        [Input("serviceAccount", required: true)]
        public Input<string> ServiceAccount { get; set; } = null!;

        public AddonPodIdentityAssociationArgs()
        {
        }
        public static new AddonPodIdentityAssociationArgs Empty => new AddonPodIdentityAssociationArgs();
    }
}