// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    /// <summary>
    /// S3 data source configuration
    /// </summary>
    public sealed class DataSourceS3DataSourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessControlListConfiguration")]
        public Input<Inputs.DataSourceAccessControlListConfigurationArgs>? AccessControlListConfiguration { get; set; }

        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("documentsMetadataConfiguration")]
        public Input<Inputs.DataSourceDocumentsMetadataConfigurationArgs>? DocumentsMetadataConfiguration { get; set; }

        [Input("exclusionPatterns")]
        private InputList<string>? _exclusionPatterns;
        public InputList<string> ExclusionPatterns
        {
            get => _exclusionPatterns ?? (_exclusionPatterns = new InputList<string>());
            set => _exclusionPatterns = value;
        }

        [Input("inclusionPatterns")]
        private InputList<string>? _inclusionPatterns;
        public InputList<string> InclusionPatterns
        {
            get => _inclusionPatterns ?? (_inclusionPatterns = new InputList<string>());
            set => _inclusionPatterns = value;
        }

        [Input("inclusionPrefixes")]
        private InputList<string>? _inclusionPrefixes;
        public InputList<string> InclusionPrefixes
        {
            get => _inclusionPrefixes ?? (_inclusionPrefixes = new InputList<string>());
            set => _inclusionPrefixes = value;
        }

        public DataSourceS3DataSourceConfigurationArgs()
        {
        }
        public static new DataSourceS3DataSourceConfigurationArgs Empty => new DataSourceS3DataSourceConfigurationArgs();
    }
}