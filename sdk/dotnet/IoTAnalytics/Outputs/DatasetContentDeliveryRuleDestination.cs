// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatasetContentDeliveryRuleDestination
    {
        /// <summary>
        /// Configuration information for delivery of dataset contents to AWS IoT Events .
        /// </summary>
        public readonly Outputs.DatasetIotEventsDestinationConfiguration? IotEventsDestinationConfiguration;
        /// <summary>
        /// Configuration information for delivery of dataset contents to Amazon S3.
        /// </summary>
        public readonly Outputs.DatasetS3DestinationConfiguration? S3DestinationConfiguration;

        [OutputConstructor]
        private DatasetContentDeliveryRuleDestination(
            Outputs.DatasetIotEventsDestinationConfiguration? iotEventsDestinationConfiguration,

            Outputs.DatasetS3DestinationConfiguration? s3DestinationConfiguration)
        {
            IotEventsDestinationConfiguration = iotEventsDestinationConfiguration;
            S3DestinationConfiguration = s3DestinationConfiguration;
        }
    }
}
