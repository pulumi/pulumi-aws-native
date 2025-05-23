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
    /// Describes the S3 data source.
    /// </summary>
    public sealed class ModelPackageS3DataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 Data Source Type
        /// </summary>
        [Input("s3DataType", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.ModelPackageS3DataSourceS3DataType> S3DataType { get; set; } = null!;

        /// <summary>
        /// Depending on the value specified for the S3DataType, identifies either a key name prefix or a manifest.
        /// </summary>
        [Input("s3Uri", required: true)]
        public Input<string> S3Uri { get; set; } = null!;

        public ModelPackageS3DataSourceArgs()
        {
        }
        public static new ModelPackageS3DataSourceArgs Empty => new ModelPackageS3DataSourceArgs();
    }
}
