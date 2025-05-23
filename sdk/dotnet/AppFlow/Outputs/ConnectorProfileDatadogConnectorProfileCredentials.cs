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
    public sealed class ConnectorProfileDatadogConnectorProfileCredentials
    {
        /// <summary>
        /// A unique alphanumeric identiﬁer used to authenticate a user, developer, or calling program to your API.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Application keys, in conjunction with your API key, give you full access to Datadog’s programmatic API. Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
        /// </summary>
        public readonly string ApplicationKey;

        [OutputConstructor]
        private ConnectorProfileDatadogConnectorProfileCredentials(
            string apiKey,

            string applicationKey)
        {
            ApiKey = apiKey;
            ApplicationKey = applicationKey;
        }
    }
}
