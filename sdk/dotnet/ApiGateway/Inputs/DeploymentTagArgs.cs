// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    public sealed class DeploymentTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key name of the tag
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value for the tag
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public DeploymentTagArgs()
        {
        }
        public static new DeploymentTagArgs Empty => new DeploymentTagArgs();
    }
}