// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
        [Input("customerManagedS3Storage")]
        public Input<Inputs.DatastoreCustomerManagedS3StorageArgs>? CustomerManagedS3Storage { get; set; }

        public DatastoreIotSiteWiseMultiLayerStorageArgs()
        {
        }
        public static new DatastoreIotSiteWiseMultiLayerStorageArgs Empty => new DatastoreIotSiteWiseMultiLayerStorageArgs();
    }
}