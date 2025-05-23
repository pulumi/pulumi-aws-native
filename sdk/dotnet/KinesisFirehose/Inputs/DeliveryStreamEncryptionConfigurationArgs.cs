// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamEncryptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Key Management Service ( AWS KMS) encryption key that Amazon S3 uses to encrypt your data.
        /// </summary>
        [Input("kmsEncryptionConfig")]
        public Input<Inputs.DeliveryStreamKmsEncryptionConfigArgs>? KmsEncryptionConfig { get; set; }

        /// <summary>
        /// Disables encryption. For valid values, see the `NoEncryptionConfig` content for the [EncryptionConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_EncryptionConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
        /// </summary>
        [Input("noEncryptionConfig")]
        public Input<Pulumi.AwsNative.KinesisFirehose.DeliveryStreamEncryptionConfigurationNoEncryptionConfig>? NoEncryptionConfig { get; set; }

        public DeliveryStreamEncryptionConfigurationArgs()
        {
        }
        public static new DeliveryStreamEncryptionConfigurationArgs Empty => new DeliveryStreamEncryptionConfigurationArgs();
    }
}
