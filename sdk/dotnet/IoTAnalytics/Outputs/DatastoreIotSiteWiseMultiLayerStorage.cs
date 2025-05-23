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
    public sealed class DatastoreIotSiteWiseMultiLayerStorage
    {
        /// <summary>
        /// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
        /// </summary>
        public readonly Outputs.DatastoreCustomerManagedS3Storage? CustomerManagedS3Storage;

        [OutputConstructor]
        private DatastoreIotSiteWiseMultiLayerStorage(Outputs.DatastoreCustomerManagedS3Storage? customerManagedS3Storage)
        {
            CustomerManagedS3Storage = customerManagedS3Storage;
        }
    }
}
