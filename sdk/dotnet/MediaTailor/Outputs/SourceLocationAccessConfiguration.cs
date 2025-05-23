// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Outputs
{

    /// <summary>
    /// &lt;p&gt;Access configuration parameters.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class SourceLocationAccessConfiguration
    {
        /// <summary>
        /// The type of authentication used to access content from `HttpConfiguration::BaseUrl` on your source location. Accepted value: `S3_SIGV4` .
        /// 
        /// `S3_SIGV4` - AWS Signature Version 4 authentication for Amazon S3 hosted virtual-style access. If your source location base URL is an Amazon S3 bucket, MediaTailor can use AWS Signature Version 4 (SigV4) authentication to access the bucket where your source content is stored. Your MediaTailor source location baseURL must follow the S3 virtual hosted-style request URL format. For example, https://bucket-name.s3.Region.amazonaws.com/key-name.
        /// 
        /// Before you can use `S3_SIGV4` , you must meet these requirements:
        /// 
        /// • You must allow MediaTailor to access your S3 bucket by granting mediatailor.amazonaws.com principal access in IAM. For information about configuring access in IAM, see Access management in the IAM User Guide.
        /// 
        /// • The mediatailor.amazonaws.com service principal must have permissions to read all top level manifests referenced by the VodSource packaging configurations.
        /// 
        /// • The caller of the API must have s3:GetObject IAM permissions to read all top level manifests referenced by your MediaTailor VodSource packaging configurations.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaTailor.SourceLocationAccessType? AccessType;
        /// <summary>
        /// AWS Secrets Manager access token configuration parameters.
        /// </summary>
        public readonly Outputs.SourceLocationSecretsManagerAccessTokenConfiguration? SecretsManagerAccessTokenConfiguration;

        [OutputConstructor]
        private SourceLocationAccessConfiguration(
            Pulumi.AwsNative.MediaTailor.SourceLocationAccessType? accessType,

            Outputs.SourceLocationSecretsManagerAccessTokenConfiguration? secretsManagerAccessTokenConfiguration)
        {
            AccessType = accessType;
            SecretsManagerAccessTokenConfiguration = secretsManagerAccessTokenConfiguration;
        }
    }
}
