// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileSapoDataConnectorProfileCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("basicAuthCredentials")]
        public Input<Inputs.ConnectorProfileBasicAuthCredentialsArgs>? BasicAuthCredentials { get; set; }

        [Input("oAuthCredentials")]
        public Input<Inputs.ConnectorProfileSapoDataConnectorProfileCredentialsOAuthCredentialsPropertiesArgs>? OAuthCredentials { get; set; }

        public ConnectorProfileSapoDataConnectorProfileCredentialsArgs()
        {
        }
        public static new ConnectorProfileSapoDataConnectorProfileCredentialsArgs Empty => new ConnectorProfileSapoDataConnectorProfileCredentialsArgs();
    }
}