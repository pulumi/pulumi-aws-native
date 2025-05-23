// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class ConnectionConnectivityParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameters for EventBridge to use when invoking the resource endpoint.
        /// </summary>
        [Input("resourceParameters", required: true)]
        public Input<Inputs.ConnectionResourceParametersArgs> ResourceParameters { get; set; } = null!;

        public ConnectionConnectivityParametersArgs()
        {
        }
        public static new ConnectionConnectivityParametersArgs Empty => new ConnectionConnectivityParametersArgs();
    }
}
