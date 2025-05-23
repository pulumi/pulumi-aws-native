// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// Information about the encryption of the flow.
    /// </summary>
    public sealed class FlowOutputEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
        /// </summary>
        [Input("algorithm")]
        public Input<Pulumi.AwsNative.MediaConnect.FlowOutputEncryptionAlgorithm>? Algorithm { get; set; }

        /// <summary>
        /// The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
        /// </summary>
        [Input("keyType")]
        public Input<Pulumi.AwsNative.MediaConnect.FlowOutputEncryptionKeyType>? KeyType { get; set; }

        /// <summary>
        /// The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        ///  The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.
        /// </summary>
        [Input("secretArn", required: true)]
        public Input<string> SecretArn { get; set; } = null!;

        public FlowOutputEncryptionArgs()
        {
        }
        public static new FlowOutputEncryptionArgs Empty => new FlowOutputEncryptionArgs();
    }
}
