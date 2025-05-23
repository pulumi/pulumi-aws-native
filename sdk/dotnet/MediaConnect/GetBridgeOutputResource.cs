// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect
{
    public static class GetBridgeOutputResource
    {
        /// <summary>
        /// Resource schema for AWS::MediaConnect::BridgeOutput
        /// </summary>
        public static Task<GetBridgeOutputResourceResult> InvokeAsync(GetBridgeOutputResourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBridgeOutputResourceResult>("aws-native:mediaconnect:getBridgeOutputResource", args ?? new GetBridgeOutputResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MediaConnect::BridgeOutput
        /// </summary>
        public static Output<GetBridgeOutputResourceResult> Invoke(GetBridgeOutputResourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBridgeOutputResourceResult>("aws-native:mediaconnect:getBridgeOutputResource", args ?? new GetBridgeOutputResourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MediaConnect::BridgeOutput
        /// </summary>
        public static Output<GetBridgeOutputResourceResult> Invoke(GetBridgeOutputResourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBridgeOutputResourceResult>("aws-native:mediaconnect:getBridgeOutputResource", args ?? new GetBridgeOutputResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBridgeOutputResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Number (ARN) of the bridge.
        /// </summary>
        [Input("bridgeArn", required: true)]
        public string BridgeArn { get; set; } = null!;

        /// <summary>
        /// The network output name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBridgeOutputResourceArgs()
        {
        }
        public static new GetBridgeOutputResourceArgs Empty => new GetBridgeOutputResourceArgs();
    }

    public sealed class GetBridgeOutputResourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Number (ARN) of the bridge.
        /// </summary>
        [Input("bridgeArn", required: true)]
        public Input<string> BridgeArn { get; set; } = null!;

        /// <summary>
        /// The network output name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetBridgeOutputResourceInvokeArgs()
        {
        }
        public static new GetBridgeOutputResourceInvokeArgs Empty => new GetBridgeOutputResourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetBridgeOutputResourceResult
    {
        /// <summary>
        /// The output of the bridge.
        /// </summary>
        public readonly Outputs.BridgeOutputResourceBridgeNetworkOutput? NetworkOutput;

        [OutputConstructor]
        private GetBridgeOutputResourceResult(Outputs.BridgeOutputResourceBridgeNetworkOutput? networkOutput)
        {
            NetworkOutput = networkOutput;
        }
    }
}
