// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileInforNexusConnectorProfileCredentials
    {
        /// <summary>
        /// The Access Key portion of the credentials.
        /// </summary>
        public readonly string AccessKeyId;
        /// <summary>
        /// The encryption keys used to encrypt data.
        /// </summary>
        public readonly string Datakey;
        /// <summary>
        /// The secret key used to sign requests.
        /// </summary>
        public readonly string SecretAccessKey;
        /// <summary>
        /// The identiﬁer for the user.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private ConnectorProfileInforNexusConnectorProfileCredentials(
            string accessKeyId,

            string datakey,

            string secretAccessKey,

            string userId)
        {
            AccessKeyId = accessKeyId;
            Datakey = datakey;
            SecretAccessKey = secretAccessKey;
            UserId = userId;
        }
    }
}
