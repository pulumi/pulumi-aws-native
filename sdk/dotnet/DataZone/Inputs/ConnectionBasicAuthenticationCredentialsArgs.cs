// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// Basic Authentication Credentials
    /// </summary>
    public sealed class ConnectionBasicAuthenticationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public ConnectionBasicAuthenticationCredentialsArgs()
        {
        }
        public static new ConnectionBasicAuthenticationCredentialsArgs Empty => new ConnectionBasicAuthenticationCredentialsArgs();
    }
}
