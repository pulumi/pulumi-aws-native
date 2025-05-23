// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Comprehend.Outputs
{

    [OutputType]
    public sealed class DocumentClassifierDocuments
    {
        /// <summary>
        /// The S3 URI location of the training documents specified in the S3Uri CSV file.
        /// </summary>
        public readonly string S3Uri;
        /// <summary>
        /// The S3 URI location of the test documents included in the TestS3Uri CSV file. This field is not required if you do not specify a test CSV file.
        /// </summary>
        public readonly string? TestS3Uri;

        [OutputConstructor]
        private DocumentClassifierDocuments(
            string s3Uri,

            string? testS3Uri)
        {
            S3Uri = s3Uri;
            TestS3Uri = testS3Uri;
        }
    }
}
