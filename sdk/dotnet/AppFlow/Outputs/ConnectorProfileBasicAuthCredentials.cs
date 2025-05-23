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
    public sealed class ConnectorProfileBasicAuthCredentials
    {
        /// <summary>
        /// The password to use to connect to a resource.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The username to use to connect to a resource.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ConnectorProfileBasicAuthCredentials(
            string password,

            string username)
        {
            Password = password;
            Username = username;
        }
    }
}
