// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Specifies the location of ML model data to deploy during endpoint creation.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageModelDataSource
    {
        public readonly Outputs.ModelPackageS3ModelDataSource? S3DataSource;

        [OutputConstructor]
        private ModelPackageModelDataSource(Outputs.ModelPackageS3ModelDataSource? s3DataSource)
        {
            S3DataSource = s3DataSource;
        }
    }
}