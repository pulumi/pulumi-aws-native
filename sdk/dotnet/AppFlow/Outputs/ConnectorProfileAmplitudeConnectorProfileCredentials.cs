// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileAmplitudeConnectorProfileCredentials
    {
        /// <summary>
        /// A unique alphanumeric identiﬁer used to authenticate a user, developer, or calling program to your API.
        /// </summary>
        public readonly string ApiKey;
        public readonly string SecretKey;

        [OutputConstructor]
        private ConnectorProfileAmplitudeConnectorProfileCredentials(
            string apiKey,

            string secretKey)
        {
            ApiKey = apiKey;
            SecretKey = secretKey;
        }
    }
}