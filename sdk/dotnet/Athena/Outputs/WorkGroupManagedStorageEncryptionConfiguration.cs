// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Outputs
{

    /// <summary>
    /// Indicates the encryption configuration for Athena Managed Storage. If not setting this field, Managed Storage will encrypt the query results with Athena's encryption key
    /// </summary>
    [OutputType]
    public sealed class WorkGroupManagedStorageEncryptionConfiguration
    {
        public readonly string? KmsKey;

        [OutputConstructor]
        private WorkGroupManagedStorageEncryptionConfiguration(string? kmsKey)
        {
            KmsKey = kmsKey;
        }
    }
}
