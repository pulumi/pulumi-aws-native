// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Inputs
{

    public sealed class EndpointSybaseSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        [Input("secretsManagerSecretId")]
        public Input<string>? SecretsManagerSecretId { get; set; }

        public EndpointSybaseSettingsArgs()
        {
        }
        public static new EndpointSybaseSettingsArgs Empty => new EndpointSybaseSettingsArgs();
    }
}