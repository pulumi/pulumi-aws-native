// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeFirehoseLogDestination
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Firehose delivery stream to which EventBridge delivers the pipe log records.
        /// </summary>
        public readonly string? DeliveryStreamArn;

        [OutputConstructor]
        private PipeFirehoseLogDestination(string? deliveryStreamArn)
        {
            DeliveryStreamArn = deliveryStreamArn;
        }
    }
}
