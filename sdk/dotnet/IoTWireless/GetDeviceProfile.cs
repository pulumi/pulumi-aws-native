// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    public static class GetDeviceProfile
    {
        /// <summary>
        /// Device Profile's resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetDeviceProfileResult> InvokeAsync(GetDeviceProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceProfileResult>("aws-native:iotwireless:getDeviceProfile", args ?? new GetDeviceProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Device Profile's resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetDeviceProfileResult> Invoke(GetDeviceProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceProfileResult>("aws-native:iotwireless:getDeviceProfile", args ?? new GetDeviceProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Device Profile's resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetDeviceProfileResult> Invoke(GetDeviceProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceProfileResult>("aws-native:iotwireless:getDeviceProfile", args ?? new GetDeviceProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Service profile Id. Returned after successful create.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeviceProfileArgs()
        {
        }
        public static new GetDeviceProfileArgs Empty => new GetDeviceProfileArgs();
    }

    public sealed class GetDeviceProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Service profile Id. Returned after successful create.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeviceProfileInvokeArgs()
        {
        }
        public static new GetDeviceProfileInvokeArgs Empty => new GetDeviceProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceProfileResult
    {
        /// <summary>
        /// Service profile Arn. Returned after successful create.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Service profile Id. Returned after successful create.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the device profile.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDeviceProfileResult(
            string? arn,

            string? id,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Id = id;
            Tags = tags;
        }
    }
}
