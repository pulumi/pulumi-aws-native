// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class DatastoreIotSiteWiseMultiLayerStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
        /// </summary>
        [Input("customerManagedS3Storage")]
        public Input<Inputs.DatastoreCustomerManagedS3StorageArgs>? CustomerManagedS3Storage { get; set; }

        public DatastoreIotSiteWiseMultiLayerStorageArgs()
        {
        }
        public static new DatastoreIotSiteWiseMultiLayerStorageArgs Empty => new DatastoreIotSiteWiseMultiLayerStorageArgs();
    }
}
