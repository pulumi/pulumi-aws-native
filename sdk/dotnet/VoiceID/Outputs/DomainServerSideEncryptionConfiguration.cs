// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VoiceID.Outputs
{

    [OutputType]
    public sealed class DomainServerSideEncryptionConfiguration
    {
        public readonly string KmsKeyId;

        [OutputConstructor]
        private DomainServerSideEncryptionConfiguration(string kmsKeyId)
        {
            KmsKeyId = kmsKeyId;
        }
    }
}