// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class DomainResourceSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance type that the image version runs on.
        /// </summary>
        [Input("instanceType")]
        public Input<Pulumi.AwsNative.SageMaker.DomainResourceSpecInstanceType>? InstanceType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
        /// </summary>
        [Input("lifecycleConfigArn")]
        public Input<string>? LifecycleConfigArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
        /// </summary>
        [Input("sageMakerImageArn")]
        public Input<string>? SageMakerImageArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the image version created on the instance.
        /// </summary>
        [Input("sageMakerImageVersionArn")]
        public Input<string>? SageMakerImageVersionArn { get; set; }

        public DomainResourceSpecArgs()
        {
        }
        public static new DomainResourceSpecArgs Empty => new DomainResourceSpecArgs();
    }
}