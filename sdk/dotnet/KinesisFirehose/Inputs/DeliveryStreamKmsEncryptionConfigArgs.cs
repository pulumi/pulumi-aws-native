// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamKmsEncryptionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS KMS encryption key that Amazon S3 uses to encrypt data delivered by the Kinesis Data Firehose stream. The key must belong to the same region as the destination S3 bucket.
        /// </summary>
        [Input("awskmsKeyArn", required: true)]
        public Input<string> AwskmsKeyArn { get; set; } = null!;

        public DeliveryStreamKmsEncryptionConfigArgs()
        {
        }
        public static new DeliveryStreamKmsEncryptionConfigArgs Empty => new DeliveryStreamKmsEncryptionConfigArgs();
    }
}
