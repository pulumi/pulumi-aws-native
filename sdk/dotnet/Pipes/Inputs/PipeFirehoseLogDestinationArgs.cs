// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeFirehoseLogDestinationArgs : global::Pulumi.ResourceArgs
    {
        [Input("deliveryStreamArn")]
        public Input<string>? DeliveryStreamArn { get; set; }

        public PipeFirehoseLogDestinationArgs()
        {
        }
        public static new PipeFirehoseLogDestinationArgs Empty => new PipeFirehoseLogDestinationArgs();
    }
}