// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    [OutputType]
    public sealed class DataSourceAuthorizationConfig
    {
        /// <summary>
        /// The authorization type that the HTTP endpoint requires.
        /// </summary>
        public readonly string AuthorizationType;
        /// <summary>
        /// The AWS Identity and Access Management settings.
        /// </summary>
        public readonly Outputs.DataSourceAwsIamConfig? AwsIamConfig;

        [OutputConstructor]
        private DataSourceAuthorizationConfig(
            string authorizationType,

            Outputs.DataSourceAwsIamConfig? awsIamConfig)
        {
            AuthorizationType = authorizationType;
            AwsIamConfig = awsIamConfig;
        }
    }
}
