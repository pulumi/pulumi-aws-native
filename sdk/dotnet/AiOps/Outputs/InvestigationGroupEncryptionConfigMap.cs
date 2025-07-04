// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AiOps.Outputs
{

    [OutputType]
    public sealed class InvestigationGroupEncryptionConfigMap
    {
        public readonly string? EncryptionConfigurationType;
        public readonly string? KmsKeyId;

        [OutputConstructor]
        private InvestigationGroupEncryptionConfigMap(
            string? encryptionConfigurationType,

            string? kmsKeyId)
        {
            EncryptionConfigurationType = encryptionConfigurationType;
            KmsKeyId = kmsKeyId;
        }
    }
}
