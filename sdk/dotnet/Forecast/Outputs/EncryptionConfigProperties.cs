// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Forecast.Outputs
{

    [OutputType]
    public sealed class EncryptionConfigProperties
    {
        public readonly string? KmsKeyArn;
        public readonly string? RoleArn;

        [OutputConstructor]
        private EncryptionConfigProperties(
            string? kmsKeyArn,

            string? roleArn)
        {
            KmsKeyArn = kmsKeyArn;
            RoleArn = roleArn;
        }
    }
}