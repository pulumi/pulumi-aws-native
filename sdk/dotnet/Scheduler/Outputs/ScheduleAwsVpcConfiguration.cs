// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler.Outputs
{

    /// <summary>
    /// This structure specifies the VPC subnets and security groups for the task, and whether a public IP address is to be used. This structure is relevant only for ECS tasks that use the awsvpc network mode.
    /// </summary>
    [OutputType]
    public sealed class ScheduleAwsVpcConfiguration
    {
        /// <summary>
        /// Specifies whether the task's elastic network interface receives a public IP address. You can specify `ENABLED` only when `LaunchType` in `EcsParameters` is set to `FARGATE` .
        /// </summary>
        public readonly Pulumi.AwsNative.Scheduler.ScheduleAssignPublicIp? AssignPublicIp;
        /// <summary>
        /// Specifies the security groups associated with the task. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// Specifies the subnets associated with the task. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private ScheduleAwsVpcConfiguration(
            Pulumi.AwsNative.Scheduler.ScheduleAssignPublicIp? assignPublicIp,

            ImmutableArray<string> securityGroups,

            ImmutableArray<string> subnets)
        {
            AssignPublicIp = assignPublicIp;
            SecurityGroups = securityGroups;
            Subnets = subnets;
        }
    }
}
