// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeEcsInferenceAcceleratorOverrideArgs : global::Pulumi.ResourceArgs
    {
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        public PipeEcsInferenceAcceleratorOverrideArgs()
        {
        }
        public static new PipeEcsInferenceAcceleratorOverrideArgs Empty => new PipeEcsInferenceAcceleratorOverrideArgs();
    }
}