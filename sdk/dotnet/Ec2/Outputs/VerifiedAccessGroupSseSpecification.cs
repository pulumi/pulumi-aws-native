// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// The configuration options for customer provided KMS encryption.
    /// </summary>
    [OutputType]
    public sealed class VerifiedAccessGroupSseSpecification
    {
        /// <summary>
        /// Whether to encrypt the policy with the provided key or disable encryption
        /// </summary>
        public readonly bool? CustomerManagedKeyEnabled;
        /// <summary>
        /// KMS Key Arn used to encrypt the group policy
        /// </summary>
        public readonly string? KmsKeyArn;

        [OutputConstructor]
        private VerifiedAccessGroupSseSpecification(
            bool? customerManagedKeyEnabled,

            string? kmsKeyArn)
        {
            CustomerManagedKeyEnabled = customerManagedKeyEnabled;
            KmsKeyArn = kmsKeyArn;
        }
    }
}