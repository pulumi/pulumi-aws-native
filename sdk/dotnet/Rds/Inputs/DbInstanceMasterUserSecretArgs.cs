// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds.Inputs
{

    public sealed class DbInstanceMasterUserSecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS KMS key identifier that is used to encrypt the secret.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the secret.
        /// </summary>
        [Input("secretArn")]
        public Input<string>? SecretArn { get; set; }

        public DbInstanceMasterUserSecretArgs()
        {
        }
        public static new DbInstanceMasterUserSecretArgs Empty => new DbInstanceMasterUserSecretArgs();
    }
}