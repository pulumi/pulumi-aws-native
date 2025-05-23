// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    /// <summary>
    /// ``VCpuCountRequest`` is a property of the ``InstanceRequirements`` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum number of vCPUs for an instance type.
    /// </summary>
    [OutputType]
    public sealed class AutoScalingGroupVCpuCountRequest
    {
        /// <summary>
        /// The maximum number of vCPUs.
        /// </summary>
        public readonly int? Max;
        /// <summary>
        /// The minimum number of vCPUs.
        /// </summary>
        public readonly int? Min;

        [OutputConstructor]
        private AutoScalingGroupVCpuCountRequest(
            int? max,

            int? min)
        {
            Max = max;
            Min = min;
        }
    }
}
