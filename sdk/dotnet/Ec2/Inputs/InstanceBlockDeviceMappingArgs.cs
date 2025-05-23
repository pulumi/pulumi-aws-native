// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class InstanceBlockDeviceMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name (for example, /dev/sdh or xvdh).
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// Parameters used to automatically set up EBS volumes when the instance is launched.
        /// </summary>
        [Input("ebs")]
        public Input<Inputs.InstanceEbsArgs>? Ebs { get; set; }

        /// <summary>
        /// To omit the device from the block device mapping, specify an empty string.
        /// 
        /// &gt; After the instance is running, modifying this parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
        /// </summary>
        [Input("noDevice")]
        public Input<object>? NoDevice { get; set; }

        /// <summary>
        /// The virtual device name ( `ephemeral` N). The name must be in the form `ephemeral` *X* where *X* is a number starting from zero (0). For example, an instance type with 2 available instance store volumes can specify mappings for `ephemeral0` and `ephemeral1` . The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.
        /// 
        /// NVMe instance store volumes are automatically enumerated and assigned a device name. Including them in your block device mapping has no effect.
        /// 
        /// *Constraints* : For M3 instances, you must specify instance store volumes in the block device mapping for the instance. When you launch an M3 instance, we ignore any instance store volumes specified in the block device mapping for the AMI.
        /// 
        /// &gt; After the instance is running, modifying this parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
        /// </summary>
        [Input("virtualName")]
        public Input<string>? VirtualName { get; set; }

        public InstanceBlockDeviceMappingArgs()
        {
        }
        public static new InstanceBlockDeviceMappingArgs Empty => new InstanceBlockDeviceMappingArgs();
    }
}
