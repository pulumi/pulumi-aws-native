// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterFirehoseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
        /// </summary>
        [Input("deliveryStream")]
        public Input<string>? DeliveryStream { get; set; }

        /// <summary>
        /// Specifies whether broker logs get send to the specified Kinesis Data Firehose delivery stream.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public ClusterFirehoseArgs()
        {
        }
        public static new ClusterFirehoseArgs Empty => new ClusterFirehoseArgs();
    }
}
