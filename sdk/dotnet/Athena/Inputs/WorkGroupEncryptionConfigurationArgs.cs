// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Inputs
{

    /// <summary>
    /// If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.
    /// </summary>
    public sealed class WorkGroupEncryptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys ( `SSE_S3` ), server-side encryption with KMS-managed keys ( `SSE_KMS` ), or client-side encryption with KMS-managed keys ( `CSE_KMS` ) is used.
        /// 
        /// If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
        /// </summary>
        [Input("encryptionOption", required: true)]
        public Input<Pulumi.AwsNative.Athena.WorkGroupEncryptionOption> EncryptionOption { get; set; } = null!;

        /// <summary>
        /// For `SSE_KMS` and `CSE_KMS` , this is the KMS key ARN or ID.
        /// </summary>
        [Input("kmsKey")]
        public Input<string>? KmsKey { get; set; }

        public WorkGroupEncryptionConfigurationArgs()
        {
        }
        public static new WorkGroupEncryptionConfigurationArgs Empty => new WorkGroupEncryptionConfigurationArgs();
    }
}
