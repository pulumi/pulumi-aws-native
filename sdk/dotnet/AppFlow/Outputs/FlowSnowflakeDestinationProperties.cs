// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowSnowflakeDestinationProperties
    {
        /// <summary>
        /// The object key for the destination bucket in which Amazon AppFlow places the files.
        /// </summary>
        public readonly string? BucketPrefix;
        /// <summary>
        /// The settings that determine how Amazon AppFlow handles an error when placing data in the Snowflake destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
        /// </summary>
        public readonly Outputs.FlowErrorHandlingConfig? ErrorHandlingConfig;
        /// <summary>
        /// The intermediate bucket that Amazon AppFlow uses when moving data into Snowflake.
        /// </summary>
        public readonly string IntermediateBucketName;
        /// <summary>
        /// The object specified in the Snowflake flow destination.
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private FlowSnowflakeDestinationProperties(
            string? bucketPrefix,

            Outputs.FlowErrorHandlingConfig? errorHandlingConfig,

            string intermediateBucketName,

            string @object)
        {
            BucketPrefix = bucketPrefix;
            ErrorHandlingConfig = errorHandlingConfig;
            IntermediateBucketName = intermediateBucketName;
            Object = @object;
        }
    }
}
