// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions.Outputs
{

    [OutputType]
    public sealed class ActivityEncryptionConfiguration
    {
        public readonly int? KmsDataKeyReusePeriodSeconds;
        public readonly string? KmsKeyId;
        public readonly Pulumi.AwsNative.StepFunctions.ActivityEncryptionConfigurationType Type;

        [OutputConstructor]
        private ActivityEncryptionConfiguration(
            int? kmsDataKeyReusePeriodSeconds,

            string? kmsKeyId,

            Pulumi.AwsNative.StepFunctions.ActivityEncryptionConfigurationType type)
        {
            KmsDataKeyReusePeriodSeconds = kmsDataKeyReusePeriodSeconds;
            KmsKeyId = kmsKeyId;
            Type = type;
        }
    }
}