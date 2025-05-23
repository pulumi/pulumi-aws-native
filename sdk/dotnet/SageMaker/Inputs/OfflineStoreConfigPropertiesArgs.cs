// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The configuration of an `OfflineStore` .
    /// </summary>
    public sealed class OfflineStoreConfigPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The meta data of the Glue table that is autogenerated when an `OfflineStore` is created.
        /// </summary>
        [Input("dataCatalogConfig")]
        public Input<Inputs.FeatureGroupDataCatalogConfigArgs>? DataCatalogConfig { get; set; }

        /// <summary>
        /// Set to `True` to disable the automatic creation of an AWS Glue table when configuring an `OfflineStore` . If set to `False` , Feature Store will name the `OfflineStore` Glue table following [Athena's naming recommendations](https://docs.aws.amazon.com/athena/latest/ug/tables-databases-columns-names.html) .
        /// 
        /// The default value is `False` .
        /// </summary>
        [Input("disableGlueTableCreation")]
        public Input<bool>? DisableGlueTableCreation { get; set; }

        /// <summary>
        /// The Amazon Simple Storage (Amazon S3) location of `OfflineStore` .
        /// </summary>
        [Input("s3StorageConfig", required: true)]
        public Input<Inputs.FeatureGroupS3StorageConfigArgs> S3StorageConfig { get; set; } = null!;

        /// <summary>
        /// Format for the offline store table. Supported formats are Glue (Default) and [Apache Iceberg](https://docs.aws.amazon.com/https://iceberg.apache.org/) .
        /// </summary>
        [Input("tableFormat")]
        public Input<Pulumi.AwsNative.SageMaker.FeatureGroupTableFormat>? TableFormat { get; set; }

        public OfflineStoreConfigPropertiesArgs()
        {
        }
        public static new OfflineStoreConfigPropertiesArgs Empty => new OfflineStoreConfigPropertiesArgs();
    }
}
