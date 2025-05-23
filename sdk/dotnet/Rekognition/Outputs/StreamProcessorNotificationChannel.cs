// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rekognition.Outputs
{

    /// <summary>
    /// The ARN of the SNS notification channel where events of interests are published, as part of connected home feature.
    /// </summary>
    [OutputType]
    public sealed class StreamProcessorNotificationChannel
    {
        /// <summary>
        /// ARN of the SNS topic.
        /// </summary>
        public readonly string Arn;

        [OutputConstructor]
        private StreamProcessorNotificationChannel(string arn)
        {
            Arn = arn;
        }
    }
}
