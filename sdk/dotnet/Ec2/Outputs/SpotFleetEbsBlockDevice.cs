// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class SpotFleetEbsBlockDevice
    {
        public readonly bool? DeleteOnTermination;
        public readonly bool? Encrypted;
        public readonly int? Iops;
        public readonly string? SnapshotId;
        public readonly int? VolumeSize;
        public readonly Pulumi.AwsNative.Ec2.SpotFleetEbsBlockDeviceVolumeType? VolumeType;

        [OutputConstructor]
        private SpotFleetEbsBlockDevice(
            bool? deleteOnTermination,

            bool? encrypted,

            int? iops,

            string? snapshotId,

            int? volumeSize,

            Pulumi.AwsNative.Ec2.SpotFleetEbsBlockDeviceVolumeType? volumeType)
        {
            DeleteOnTermination = deleteOnTermination;
            Encrypted = encrypted;
            Iops = iops;
            SnapshotId = snapshotId;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }
}