// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Inputs
{

    public sealed class IntegrationTlsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("serverNameToVerify")]
        public Input<string>? ServerNameToVerify { get; set; }

        public IntegrationTlsConfigArgs()
        {
        }
        public static new IntegrationTlsConfigArgs Empty => new IntegrationTlsConfigArgs();
    }
}