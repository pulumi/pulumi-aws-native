// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    /// <summary>
    /// Details of an Amazon MSK cluster.
    /// </summary>
    public sealed class ReplicatorAmazonMskClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of an Amazon MSK cluster.
        /// </summary>
        [Input("mskClusterArn", required: true)]
        public Input<string> MskClusterArn { get; set; } = null!;

        public ReplicatorAmazonMskClusterArgs()
        {
        }
        public static new ReplicatorAmazonMskClusterArgs Empty => new ReplicatorAmazonMskClusterArgs();
    }
}
