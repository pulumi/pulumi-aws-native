// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class AutoScalingGroupLaunchTemplateSpecification
    {
        public readonly string? LaunchTemplateId;
        public readonly string? LaunchTemplateName;
        public readonly string Version;

        [OutputConstructor]
        private AutoScalingGroupLaunchTemplateSpecification(
            string? launchTemplateId,

            string? launchTemplateName,

            string version)
        {
            LaunchTemplateId = launchTemplateId;
            LaunchTemplateName = launchTemplateName;
            Version = version;
        }
    }
}