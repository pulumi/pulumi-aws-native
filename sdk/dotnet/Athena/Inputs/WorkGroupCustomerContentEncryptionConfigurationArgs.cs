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
    /// Indicates the KMS key for encrypting notebook content.
    /// </summary>
    public sealed class WorkGroupCustomerContentEncryptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The customer managed KMS key that is used to encrypt the user's data stores in Athena.
        /// </summary>
        [Input("kmsKey", required: true)]
        public Input<string> KmsKey { get; set; } = null!;

        public WorkGroupCustomerContentEncryptionConfigurationArgs()
        {
        }
        public static new WorkGroupCustomerContentEncryptionConfigurationArgs Empty => new WorkGroupCustomerContentEncryptionConfigurationArgs();
    }
}
