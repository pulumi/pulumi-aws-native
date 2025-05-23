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
    public sealed class ChannelStorage
    {
        /// <summary>
        /// Used to store channel data in an S3 bucket that you manage. If customer managed storage is selected, the `retentionPeriod` parameter is ignored. You can't change the choice of S3 storage after the data store is created.
        /// </summary>
        public readonly Outputs.ChannelCustomerManagedS3? CustomerManagedS3;
        /// <summary>
        /// Used to store channel data in an S3 bucket managed by AWS IoT Analytics . You can't change the choice of S3 storage after the data store is created.
        /// </summary>
        public readonly Outputs.ChannelServiceManagedS3? ServiceManagedS3;

        [OutputConstructor]
        private ChannelStorage(
            Outputs.ChannelCustomerManagedS3? customerManagedS3,

            Outputs.ChannelServiceManagedS3? serviceManagedS3)
        {
            CustomerManagedS3 = customerManagedS3;
            ServiceManagedS3 = serviceManagedS3;
        }
    }
}
