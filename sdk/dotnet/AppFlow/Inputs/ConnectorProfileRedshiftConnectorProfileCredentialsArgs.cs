// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileRedshiftConnectorProfileCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password that corresponds to the username.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ConnectorProfileRedshiftConnectorProfileCredentialsArgs()
        {
        }
        public static new ConnectorProfileRedshiftConnectorProfileCredentialsArgs Empty => new ConnectorProfileRedshiftConnectorProfileCredentialsArgs();
    }
}