// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics
{
    public static class GetDatastore
    {
        /// <summary>
        /// Resource Type definition for AWS::IoTAnalytics::Datastore
        /// </summary>
        public static Task<GetDatastoreResult> InvokeAsync(GetDatastoreArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatastoreResult>("aws-native:iotanalytics:getDatastore", args ?? new GetDatastoreArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoTAnalytics::Datastore
        /// </summary>
        public static Output<GetDatastoreResult> Invoke(GetDatastoreInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreResult>("aws-native:iotanalytics:getDatastore", args ?? new GetDatastoreInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoTAnalytics::Datastore
        /// </summary>
        public static Output<GetDatastoreResult> Invoke(GetDatastoreInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreResult>("aws-native:iotanalytics:getDatastore", args ?? new GetDatastoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatastoreArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the data store.
        /// </summary>
        [Input("datastoreName", required: true)]
        public string DatastoreName { get; set; } = null!;

        public GetDatastoreArgs()
        {
        }
        public static new GetDatastoreArgs Empty => new GetDatastoreArgs();
    }

    public sealed class GetDatastoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the data store.
        /// </summary>
        [Input("datastoreName", required: true)]
        public Input<string> DatastoreName { get; set; } = null!;

        public GetDatastoreInvokeArgs()
        {
        }
        public static new GetDatastoreInvokeArgs Empty => new GetDatastoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatastoreResult
    {
        /// <summary>
        /// Information about the partition dimensions in a data store.
        /// </summary>
        public readonly Outputs.DatastorePartitions? DatastorePartitions;
        /// <summary>
        /// Where data store data is stored.
        /// </summary>
        public readonly Outputs.DatastoreStorage? DatastoreStorage;
        /// <summary>
        /// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
        /// 
        /// The default file format is JSON. You can specify only one format.
        /// 
        /// You can't change the file format after you create the data store.
        /// </summary>
        public readonly Outputs.DatastoreFileFormatConfiguration? FileFormatConfiguration;
        public readonly string? Id;
        /// <summary>
        /// How long, in days, message data is kept for the data store. When `customerManagedS3` storage is selected, this parameter is ignored.
        /// </summary>
        public readonly Outputs.DatastoreRetentionPeriod? RetentionPeriod;
        /// <summary>
        /// Metadata which can be used to manage the data store.
        /// 
        /// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDatastoreResult(
            Outputs.DatastorePartitions? datastorePartitions,

            Outputs.DatastoreStorage? datastoreStorage,

            Outputs.DatastoreFileFormatConfiguration? fileFormatConfiguration,

            string? id,

            Outputs.DatastoreRetentionPeriod? retentionPeriod,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DatastorePartitions = datastorePartitions;
            DatastoreStorage = datastoreStorage;
            FileFormatConfiguration = fileFormatConfiguration;
            Id = id;
            RetentionPeriod = retentionPeriod;
            Tags = tags;
        }
    }
}
