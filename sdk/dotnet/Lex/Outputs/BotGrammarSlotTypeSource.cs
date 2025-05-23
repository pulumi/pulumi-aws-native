// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    [OutputType]
    public sealed class BotGrammarSlotTypeSource
    {
        /// <summary>
        /// The AWS KMS key required to decrypt the contents of the grammar, if any.
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// The name of the Amazon S3 bucket that contains the grammar source.
        /// </summary>
        public readonly string S3BucketName;
        /// <summary>
        /// The path to the grammar in the Amazon S3 bucket.
        /// </summary>
        public readonly string S3ObjectKey;

        [OutputConstructor]
        private BotGrammarSlotTypeSource(
            string? kmsKeyArn,

            string s3BucketName,

            string s3ObjectKey)
        {
            KmsKeyArn = kmsKeyArn;
            S3BucketName = s3BucketName;
            S3ObjectKey = s3ObjectKey;
        }
    }
}
