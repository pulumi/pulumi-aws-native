// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Inputs
{

    /// <summary>
    /// The S3 bucket where artifacts for the pipeline are stored.
    /// </summary>
    public sealed class PipelineArtifactStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service ( AWS KMS) key. If this is undefined, the default key for Amazon S3 is used. To see an example artifact store encryption key field, see the example structure here: [AWS::CodePipeline::Pipeline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html) .
        /// </summary>
        [Input("encryptionKey")]
        public Input<Inputs.PipelineEncryptionKeyArgs>? EncryptionKey { get; set; }

        /// <summary>
        /// The S3 bucket used for storing the artifacts for a pipeline. You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The type of the artifact store, such as S3.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.CodePipeline.PipelineArtifactStoreType> Type { get; set; } = null!;

        public PipelineArtifactStoreArgs()
        {
        }
        public static new PipelineArtifactStoreArgs Empty => new PipelineArtifactStoreArgs();
    }
}
