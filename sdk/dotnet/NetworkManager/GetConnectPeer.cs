// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager
{
    public static class GetConnectPeer
    {
        /// <summary>
        /// AWS::NetworkManager::ConnectPeer Resource Type Definition.
        /// </summary>
        public static Task<GetConnectPeerResult> InvokeAsync(GetConnectPeerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectPeerResult>("aws-native:networkmanager:getConnectPeer", args ?? new GetConnectPeerArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::NetworkManager::ConnectPeer Resource Type Definition.
        /// </summary>
        public static Output<GetConnectPeerResult> Invoke(GetConnectPeerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectPeerResult>("aws-native:networkmanager:getConnectPeer", args ?? new GetConnectPeerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectPeerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Connect peer.
        /// </summary>
        [Input("connectPeerId", required: true)]
        public string ConnectPeerId { get; set; } = null!;

        public GetConnectPeerArgs()
        {
        }
        public static new GetConnectPeerArgs Empty => new GetConnectPeerArgs();
    }

    public sealed class GetConnectPeerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Connect peer.
        /// </summary>
        [Input("connectPeerId", required: true)]
        public Input<string> ConnectPeerId { get; set; } = null!;

        public GetConnectPeerInvokeArgs()
        {
        }
        public static new GetConnectPeerInvokeArgs Empty => new GetConnectPeerInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectPeerResult
    {
        /// <summary>
        /// Configuration of the connect peer.
        /// </summary>
        public readonly Outputs.ConnectPeerConfiguration? Configuration;
        /// <summary>
        /// The ID of the Connect peer.
        /// </summary>
        public readonly string? ConnectPeerId;
        /// <summary>
        /// The ID of the core network.
        /// </summary>
        public readonly string? CoreNetworkId;
        /// <summary>
        /// Connect peer creation time.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The Connect peer Regions where edges are located.
        /// </summary>
        public readonly string? EdgeLocation;
        /// <summary>
        /// State of the connect peer.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectPeerTag> Tags;

        [OutputConstructor]
        private GetConnectPeerResult(
            Outputs.ConnectPeerConfiguration? configuration,

            string? connectPeerId,

            string? coreNetworkId,

            string? createdAt,

            string? edgeLocation,

            string? state,

            ImmutableArray<Outputs.ConnectPeerTag> tags)
        {
            Configuration = configuration;
            ConnectPeerId = connectPeerId;
            CoreNetworkId = coreNetworkId;
            CreatedAt = createdAt;
            EdgeLocation = edgeLocation;
            State = state;
            Tags = tags;
        }
    }
}