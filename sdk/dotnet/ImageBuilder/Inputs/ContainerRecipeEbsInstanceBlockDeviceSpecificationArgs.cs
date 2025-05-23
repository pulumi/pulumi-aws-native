// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// Amazon EBS-specific block device mapping specifications. 
    /// </summary>
    public sealed class ContainerRecipeEbsInstanceBlockDeviceSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use to configure delete on termination of the associated device.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        /// <summary>
        /// Use to configure device encryption.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// Use to configure device IOPS.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// Use to configure the KMS key to use when encrypting the device.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The snapshot that defines the device contents.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// For GP3 volumes only - The throughput in MiB/s that the volume supports.
        /// </summary>
        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        /// <summary>
        /// Use to override the device's volume size.
        /// </summary>
        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        /// <summary>
        /// Use to override the device's volume type.
        /// </summary>
        [Input("volumeType")]
        public Input<Pulumi.AwsNative.ImageBuilder.ContainerRecipeEbsInstanceBlockDeviceSpecificationVolumeType>? VolumeType { get; set; }

        public ContainerRecipeEbsInstanceBlockDeviceSpecificationArgs()
        {
        }
        public static new ContainerRecipeEbsInstanceBlockDeviceSpecificationArgs Empty => new ContainerRecipeEbsInstanceBlockDeviceSpecificationArgs();
    }
}
