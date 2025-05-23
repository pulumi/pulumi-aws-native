// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    [OutputType]
    public sealed class PackagingGroupAuthorization
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.
        /// </summary>
        public readonly string CdnIdentifierSecret;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.
        /// </summary>
        public readonly string SecretsRoleArn;

        [OutputConstructor]
        private PackagingGroupAuthorization(
            string cdnIdentifierSecret,

            string secretsRoleArn)
        {
            CdnIdentifierSecret = cdnIdentifierSecret;
            SecretsRoleArn = secretsRoleArn;
        }
    }
}
