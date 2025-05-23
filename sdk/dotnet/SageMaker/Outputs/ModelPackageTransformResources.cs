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
    /// Describes the resources, including ML instance types and ML instance count, to use for transform job.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageTransformResources
    {
        /// <summary>
        /// The number of ML compute instances to use in the transform job. For distributed transform jobs, specify a value greater than 1. The default value is 1.
        /// </summary>
        public readonly int InstanceCount;
        /// <summary>
        /// The ML compute instance type for the transform job.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt model data on the storage volume attached to the ML compute instance(s) that run the batch transform job.
        /// </summary>
        public readonly string? VolumeKmsKeyId;

        [OutputConstructor]
        private ModelPackageTransformResources(
            int instanceCount,

            string instanceType,

            string? volumeKmsKeyId)
        {
            InstanceCount = instanceCount;
            InstanceType = instanceType;
            VolumeKmsKeyId = volumeKmsKeyId;
        }
    }
}
