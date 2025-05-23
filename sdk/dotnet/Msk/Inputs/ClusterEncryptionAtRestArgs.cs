// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterEncryptionAtRestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon KMS key for encrypting data at rest. If you don't specify a KMS key, MSK creates one for you and uses it.
        /// </summary>
        [Input("dataVolumeKmsKeyId", required: true)]
        public Input<string> DataVolumeKmsKeyId { get; set; } = null!;

        public ClusterEncryptionAtRestArgs()
        {
        }
        public static new ClusterEncryptionAtRestArgs Empty => new ClusterEncryptionAtRestArgs();
    }
}
