// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileBasicAuthCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password to use to connect to a resource.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username to use to connect to a resource.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ConnectorProfileBasicAuthCredentialsArgs()
        {
        }
        public static new ConnectorProfileBasicAuthCredentialsArgs Empty => new ConnectorProfileBasicAuthCredentialsArgs();
    }
}
