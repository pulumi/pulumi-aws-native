// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// Provides the path to the S3 bucket that contains the user context filtering files for the data source. For the format of the file, see [Access control for S3 data sources](https://docs.aws.amazon.com/kendra/latest/dg/s3-acl.html) .
        /// </summary>
        [Input("accessControlListConfiguration")]
        public Input<Inputs.DataSourceAccessControlListConfigurationArgs>? AccessControlListConfiguration { get; set; }

        /// <summary>
        /// The name of the bucket that contains the documents.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Specifies document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes. Each metadata file contains metadata about a single document.
        /// </summary>
        [Input("documentsMetadataConfiguration")]
        public Input<Inputs.DataSourceDocumentsMetadataConfigurationArgs>? DocumentsMetadataConfiguration { get; set; }

        [Input("exclusionPatterns")]
        private InputList<string>? _exclusionPatterns;

        /// <summary>
        /// A list of glob patterns (patterns that can expand a wildcard pattern into a list of path names that match the given pattern) for certain file names and file types to exclude from your index. If a document matches both an inclusion and exclusion prefix or pattern, the exclusion prefix takes precendence and the document is not indexed. Examples of glob patterns include:
        /// 
        /// - */myapp/config/** —All files inside config directory.
        /// - ***/*.png* —All .png files in all directories.
        /// - ***/*.{png, ico, md}* —All .png, .ico or .md files in all directories.
        /// - */myapp/src/**/*.ts* —All .ts files inside src directory (and all its subdirectories).
        /// - ***/!(*.module).ts* —All .ts files but not .module.ts
        /// - **.png , *.jpg* —All PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
        /// - **internal** —All files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
        /// - ***/*internal** —All internal-related files in a directory and its subdirectories.
        /// 
        /// For more examples, see [Use of Exclude and Include Filters](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.
        /// </summary>
        public InputList<string> ExclusionPatterns
        {
            get => _exclusionPatterns ?? (_exclusionPatterns = new InputList<string>());
            set => _exclusionPatterns = value;
        }

        [Input("inclusionPatterns")]
        private InputList<string>? _inclusionPatterns;

        /// <summary>
        /// A list of glob patterns (patterns that can expand a wildcard pattern into a list of path names that match the given pattern) for certain file names and file types to include in your index. If a document matches both an inclusion and exclusion prefix or pattern, the exclusion prefix takes precendence and the document is not indexed. Examples of glob patterns include:
        /// 
        /// - */myapp/config/** —All files inside config directory.
        /// - ***/*.png* —All .png files in all directories.
        /// - ***/*.{png, ico, md}* —All .png, .ico or .md files in all directories.
        /// - */myapp/src/**/*.ts* —All .ts files inside src directory (and all its subdirectories).
        /// - ***/!(*.module).ts* —All .ts files but not .module.ts
        /// - **.png , *.jpg* —All PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
        /// - **internal** —All files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
        /// - ***/*internal** —All internal-related files in a directory and its subdirectories.
        /// 
        /// For more examples, see [Use of Exclude and Include Filters](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.
        /// </summary>
        public InputList<string> InclusionPatterns
        {
            get => _inclusionPatterns ?? (_inclusionPatterns = new InputList<string>());
            set => _inclusionPatterns = value;
        }

        [Input("inclusionPrefixes")]
        private InputList<string>? _inclusionPrefixes;

        /// <summary>
        /// A list of S3 prefixes for the documents that should be included in the index.
        /// </summary>
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
