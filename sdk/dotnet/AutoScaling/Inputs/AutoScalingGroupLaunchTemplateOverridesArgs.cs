// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Inputs
{

    public sealed class AutoScalingGroupLaunchTemplateOverridesArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceRequirements")]
        public Input<Inputs.AutoScalingGroupInstanceRequirementsArgs>? InstanceRequirements { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("launchTemplateSpecification")]
        public Input<Inputs.AutoScalingGroupLaunchTemplateSpecificationArgs>? LaunchTemplateSpecification { get; set; }

        [Input("weightedCapacity")]
        public Input<string>? WeightedCapacity { get; set; }

        public AutoScalingGroupLaunchTemplateOverridesArgs()
        {
        }
        public static new AutoScalingGroupLaunchTemplateOverridesArgs Empty => new AutoScalingGroupLaunchTemplateOverridesArgs();
    }
}