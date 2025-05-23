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
    public sealed class DeliveryStreamHttpEndpointConfiguration
    {
        /// <summary>
        /// The access key required for Kinesis Firehose to authenticate with the HTTP endpoint selected as the destination.
        /// </summary>
        public readonly string? AccessKey;
        /// <summary>
        /// The name of the HTTP endpoint selected as the destination.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The URL of the HTTP endpoint selected as the destination.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private DeliveryStreamHttpEndpointConfiguration(
            string? accessKey,

            string? name,

            string url)
        {
            AccessKey = accessKey;
            Name = name;
            Url = url;
        }
    }
}
