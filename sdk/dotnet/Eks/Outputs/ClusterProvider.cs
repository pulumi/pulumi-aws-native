// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Outputs
{

    [OutputType]
    public sealed class ClusterProvider
    {
        /// <summary>
        /// Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
        /// </summary>
        public readonly string? KeyArn;

        [OutputConstructor]
        private ClusterProvider(string? keyArn)
        {
            KeyArn = keyArn;
        }
    }
}