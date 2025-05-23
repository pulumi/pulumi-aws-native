// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
    /// </summary>
    public sealed class DetectorModelSqsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// You can configure the action payload when you send a message to an Amazon SQS queue.
        /// </summary>
        [Input("payload")]
        public Input<Inputs.DetectorModelPayloadArgs>? Payload { get; set; }

        /// <summary>
        /// The URL of the SQS queue where the data is written.
        /// </summary>
        [Input("queueUrl", required: true)]
        public Input<string> QueueUrl { get; set; } = null!;

        /// <summary>
        /// Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to FALSE.
        /// </summary>
        [Input("useBase64")]
        public Input<bool>? UseBase64 { get; set; }

        public DetectorModelSqsArgs()
        {
        }
        public static new DetectorModelSqsArgs Empty => new DetectorModelSqsArgs();
    }
}
