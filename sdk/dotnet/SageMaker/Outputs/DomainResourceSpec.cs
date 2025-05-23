// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class DomainResourceSpec
    {
        /// <summary>
        /// The instance type that the image version runs on.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.DomainResourceSpecInstanceType? InstanceType;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
        /// </summary>
        public readonly string? LifecycleConfigArn;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
        /// </summary>
        public readonly string? SageMakerImageArn;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image version created on the instance.
        /// </summary>
        public readonly string? SageMakerImageVersionArn;

        [OutputConstructor]
        private DomainResourceSpec(
            Pulumi.AwsNative.SageMaker.DomainResourceSpecInstanceType? instanceType,

            string? lifecycleConfigArn,

            string? sageMakerImageArn,

            string? sageMakerImageVersionArn)
        {
            InstanceType = instanceType;
            LifecycleConfigArn = lifecycleConfigArn;
            SageMakerImageArn = sageMakerImageArn;
            SageMakerImageVersionArn = sageMakerImageVersionArn;
        }
    }
}
