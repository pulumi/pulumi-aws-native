// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake.Inputs
{

    /// <summary>
    /// Provides replication details of Amazon Security Lake object.
    /// </summary>
    public sealed class DataLakeReplicationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// Specifies one or more centralized rollup Regions. The AWS Region specified in the region parameter of the `CreateDataLake` or `UpdateDataLake` operations contributes data to the rollup Region or Regions specified in this parameter.
        /// 
        /// Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different Regions or within the same Region as the source bucket.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public DataLakeReplicationConfigurationArgs()
        {
        }
        public static new DataLakeReplicationConfigurationArgs Empty => new DataLakeReplicationConfigurationArgs();
    }
}
