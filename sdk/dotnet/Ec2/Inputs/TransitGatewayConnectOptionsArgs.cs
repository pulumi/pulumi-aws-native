// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class TransitGatewayConnectOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tunnel protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public TransitGatewayConnectOptionsArgs()
        {
        }
        public static new TransitGatewayConnectOptionsArgs Empty => new TransitGatewayConnectOptionsArgs();
    }
}