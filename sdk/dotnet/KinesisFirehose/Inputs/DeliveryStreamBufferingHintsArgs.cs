// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamBufferingHintsArgs : global::Pulumi.ResourceArgs
    {
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        [Input("sizeInMBs")]
        public Input<int>? SizeInMBs { get; set; }

        public DeliveryStreamBufferingHintsArgs()
        {
        }
        public static new DeliveryStreamBufferingHintsArgs Empty => new DeliveryStreamBufferingHintsArgs();
    }
}