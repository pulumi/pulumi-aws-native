// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkspacesInstances
{
    public static class GetVolume
    {
        /// <summary>
        /// Resource Type definition for AWS::WorkspacesInstances::Volume - Manages WorkSpaces Volume resources
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("aws-native:workspacesinstances:getVolume", args ?? new GetVolumeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WorkspacesInstances::Volume - Manages WorkSpaces Volume resources
        /// </summary>
        public static Output<GetVolumeResult> Invoke(GetVolumeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeResult>("aws-native:workspacesinstances:getVolume", args ?? new GetVolumeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WorkspacesInstances::Volume - Manages WorkSpaces Volume resources
        /// </summary>
        public static Output<GetVolumeResult> Invoke(GetVolumeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeResult>("aws-native:workspacesinstances:getVolume", args ?? new GetVolumeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the volume
        /// </summary>
        [Input("volumeId", required: true)]
        public string VolumeId { get; set; } = null!;

        public GetVolumeArgs()
        {
        }
        public static new GetVolumeArgs Empty => new GetVolumeArgs();
    }

    public sealed class GetVolumeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the volume
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public GetVolumeInvokeArgs()
        {
        }
        public static new GetVolumeInvokeArgs Empty => new GetVolumeInvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        /// <summary>
        /// Unique identifier for the volume
        /// </summary>
        public readonly string? VolumeId;

        [OutputConstructor]
        private GetVolumeResult(string? volumeId)
        {
            VolumeId = volumeId;
        }
    }
}
