// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Inputs
{

    public sealed class ClusterProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
        /// </summary>
        [Input("keyArn")]
        public Input<string>? KeyArn { get; set; }

        public ClusterProviderArgs()
        {
        }
        public static new ClusterProviderArgs Empty => new ClusterProviderArgs();
    }
}
