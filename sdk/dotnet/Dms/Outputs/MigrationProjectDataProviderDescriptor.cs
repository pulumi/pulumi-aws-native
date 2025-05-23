// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Outputs
{

    /// <summary>
    /// It is an object that describes Source and Target DataProviders and credentials for connecting to databases that are used in MigrationProject
    /// </summary>
    [OutputType]
    public sealed class MigrationProjectDataProviderDescriptor
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the data provider.
        /// </summary>
        public readonly string? DataProviderArn;
        public readonly string? DataProviderIdentifier;
        /// <summary>
        /// The user-friendly name of the data provider.
        /// </summary>
        public readonly string? DataProviderName;
        /// <summary>
        /// The ARN of the role used to access AWS Secrets Manager.
        /// </summary>
        public readonly string? SecretsManagerAccessRoleArn;
        /// <summary>
        /// The identifier of the AWS Secrets Manager Secret used to store access credentials for the data provider.
        /// </summary>
        public readonly string? SecretsManagerSecretId;

        [OutputConstructor]
        private MigrationProjectDataProviderDescriptor(
            string? dataProviderArn,

            string? dataProviderIdentifier,

            string? dataProviderName,

            string? secretsManagerAccessRoleArn,

            string? secretsManagerSecretId)
        {
            DataProviderArn = dataProviderArn;
            DataProviderIdentifier = dataProviderIdentifier;
            DataProviderName = dataProviderName;
            SecretsManagerAccessRoleArn = secretsManagerAccessRoleArn;
            SecretsManagerSecretId = secretsManagerSecretId;
        }
    }
}
