// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT1Click
{
    public static class GetDevice
    {
        /// <summary>
        /// Resource Type definition for AWS::IoT1Click::Device
        /// </summary>
        public static Task<GetDeviceResult> InvokeAsync(GetDeviceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceResult>("aws-native:iot1click:getDevice", args ?? new GetDeviceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoT1Click::Device
        /// </summary>
        public static Output<GetDeviceResult> Invoke(GetDeviceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceResult>("aws-native:iot1click:getDevice", args ?? new GetDeviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceArgs : global::Pulumi.InvokeArgs
    {
        [Input("deviceId", required: true)]
        public string DeviceId { get; set; } = null!;

        public GetDeviceArgs()
        {
        }
        public static new GetDeviceArgs Empty => new GetDeviceArgs();
    }

    public sealed class GetDeviceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("deviceId", required: true)]
        public Input<string> DeviceId { get; set; } = null!;

        public GetDeviceInvokeArgs()
        {
        }
        public static new GetDeviceInvokeArgs Empty => new GetDeviceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceResult
    {
        public readonly string? Arn;
        public readonly bool? Enabled;

        [OutputConstructor]
        private GetDeviceResult(
            string? arn,

            bool? enabled)
        {
            Arn = arn;
            Enabled = enabled;
        }
    }
}