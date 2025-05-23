// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class CapacityProviderAutoScalingGroupProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the Auto Scaling group, or the Auto Scaling group name.
        /// </summary>
        [Input("autoScalingGroupArn", required: true)]
        public Input<string> AutoScalingGroupArn { get; set; } = null!;

        /// <summary>
        /// The managed draining option for the Auto Scaling group capacity provider. When you enable this, Amazon ECS manages and gracefully drains the EC2 container instances that are in the Auto Scaling group capacity provider.
        /// </summary>
        [Input("managedDraining")]
        public Input<Pulumi.AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedDraining>? ManagedDraining { get; set; }

        /// <summary>
        /// The managed scaling settings for the Auto Scaling group capacity provider.
        /// </summary>
        [Input("managedScaling")]
        public Input<Inputs.CapacityProviderManagedScalingArgs>? ManagedScaling { get; set; }

        /// <summary>
        /// The managed termination protection setting to use for the Auto Scaling group capacity provider. This determines whether the Auto Scaling group has managed termination protection. The default is off.
        /// 
        /// &gt; When using managed termination protection, managed scaling must also be used otherwise managed termination protection doesn't work. 
        /// 
        /// When managed termination protection is on, Amazon ECS prevents the Amazon EC2 instances in an Auto Scaling group that contain tasks from being terminated during a scale-in action. The Auto Scaling group and each instance in the Auto Scaling group must have instance protection from scale-in actions on as well. For more information, see [Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the *AWS Auto Scaling User Guide* .
        /// 
        /// When managed termination protection is off, your Amazon EC2 instances aren't protected from termination when the Auto Scaling group scales in.
        /// </summary>
        [Input("managedTerminationProtection")]
        public Input<Pulumi.AwsNative.Ecs.CapacityProviderAutoScalingGroupProviderManagedTerminationProtection>? ManagedTerminationProtection { get; set; }

        public CapacityProviderAutoScalingGroupProviderArgs()
        {
        }
        public static new CapacityProviderAutoScalingGroupProviderArgs Empty => new CapacityProviderAutoScalingGroupProviderArgs();
    }
}
