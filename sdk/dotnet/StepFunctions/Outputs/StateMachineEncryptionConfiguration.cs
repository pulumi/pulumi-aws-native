// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions.Outputs
{

    [OutputType]
    public sealed class StateMachineEncryptionConfiguration
    {
        /// <summary>
        /// Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        /// </summary>
        public readonly int? KmsDataKeyReusePeriodSeconds;
        /// <summary>
        /// An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// Encryption option for a state machine.
        /// </summary>
        public readonly Pulumi.AwsNative.StepFunctions.StateMachineEncryptionConfigurationType Type;

        [OutputConstructor]
        private StateMachineEncryptionConfiguration(
            int? kmsDataKeyReusePeriodSeconds,

            string? kmsKeyId,

            Pulumi.AwsNative.StepFunctions.StateMachineEncryptionConfigurationType type)
        {
            KmsDataKeyReusePeriodSeconds = kmsDataKeyReusePeriodSeconds;
            KmsKeyId = kmsKeyId;
            Type = type;
        }
    }
}
