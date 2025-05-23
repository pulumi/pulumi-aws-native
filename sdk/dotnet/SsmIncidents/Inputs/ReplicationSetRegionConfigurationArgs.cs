// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Inputs
{

    /// <summary>
    /// The ReplicationSet regional configuration.
    /// </summary>
    public sealed class ReplicationSetRegionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Key Management Service key ID or Key Alias to use to encrypt your replication set.
        /// </summary>
        [Input("sseKmsKeyId", required: true)]
        public Input<string> SseKmsKeyId { get; set; } = null!;

        public ReplicationSetRegionConfigurationArgs()
        {
        }
        public static new ReplicationSetRegionConfigurationArgs Empty => new ReplicationSetRegionConfigurationArgs();
    }
}
