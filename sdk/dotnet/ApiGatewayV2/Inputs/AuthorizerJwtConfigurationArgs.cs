// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Inputs
{

    public sealed class AuthorizerJwtConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("audience")]
        private InputList<string>? _audience;
        public InputList<string> Audience
        {
            get => _audience ?? (_audience = new InputList<string>());
            set => _audience = value;
        }

        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        public AuthorizerJwtConfigurationArgs()
        {
        }
        public static new AuthorizerJwtConfigurationArgs Empty => new AuthorizerJwtConfigurationArgs();
    }
}