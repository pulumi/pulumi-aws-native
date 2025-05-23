// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamSnowflakeRetryOptions
    {
        /// <summary>
        /// the time period where Firehose will retry sending data to the chosen HTTP endpoint.
        /// </summary>
        public readonly int? DurationInSeconds;

        [OutputConstructor]
        private DeliveryStreamSnowflakeRetryOptions(int? durationInSeconds)
        {
            DurationInSeconds = durationInSeconds;
        }
    }
}
