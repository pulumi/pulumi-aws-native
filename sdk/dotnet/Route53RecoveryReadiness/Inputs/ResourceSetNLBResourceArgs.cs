// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryReadiness.Inputs
{

    /// <summary>
    /// The Network Load Balancer resource that a DNS target resource points to.
    /// </summary>
    public sealed class ResourceSetNLBResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Network Load Balancer resource Amazon Resource Name (ARN).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        public ResourceSetNLBResourceArgs()
        {
        }
        public static new ResourceSetNLBResourceArgs Empty => new ResourceSetNLBResourceArgs();
    }
}