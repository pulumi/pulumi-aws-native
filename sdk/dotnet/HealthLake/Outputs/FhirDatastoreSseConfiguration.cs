// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.HealthLake.Outputs
{

    /// <summary>
    /// The server-side encryption key configuration for a customer provided encryption key.
    /// </summary>
    [OutputType]
    public sealed class FhirDatastoreSseConfiguration
    {
        public readonly Outputs.FhirDatastoreKmsEncryptionConfig KmsEncryptionConfig;

        [OutputConstructor]
        private FhirDatastoreSseConfiguration(Outputs.FhirDatastoreKmsEncryptionConfig kmsEncryptionConfig)
        {
            KmsEncryptionConfig = kmsEncryptionConfig;
        }
    }
}