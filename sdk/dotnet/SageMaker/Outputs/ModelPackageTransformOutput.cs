// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Describes the results of a transform job.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageTransformOutput
    {
        /// <summary>
        /// The MIME type used to specify the output data. Amazon SageMaker uses the MIME type with each http call to transfer data from the transform job.
        /// </summary>
        public readonly string? Accept;
        /// <summary>
        /// Defines how to assemble the results of the transform job as a single S3 object.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.ModelPackageTransformOutputAssembleWith? AssembleWith;
        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// The Amazon S3 path where you want Amazon SageMaker to store the results of the transform job.
        /// </summary>
        public readonly string S3OutputPath;

        [OutputConstructor]
        private ModelPackageTransformOutput(
            string? accept,

            Pulumi.AwsNative.SageMaker.ModelPackageTransformOutputAssembleWith? assembleWith,

            string? kmsKeyId,

            string s3OutputPath)
        {
            Accept = accept;
            AssembleWith = assembleWith;
            KmsKeyId = kmsKeyId;
            S3OutputPath = s3OutputPath;
        }
    }
}
