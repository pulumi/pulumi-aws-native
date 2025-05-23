// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Describes the input source of a transform job and the way the transform job consumes it.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageDataSource
    {
        /// <summary>
        /// The S3 location of the data source that is associated with a channel.
        /// </summary>
        public readonly Outputs.ModelPackageS3DataSource S3DataSource;

        [OutputConstructor]
        private ModelPackageDataSource(Outputs.ModelPackageS3DataSource s3DataSource)
        {
            S3DataSource = s3DataSource;
        }
    }
}
