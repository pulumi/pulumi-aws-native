// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig.Outputs
{

    /// <summary>
    /// Amazon CloudWatch alarm to monitor during the deployment process.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentMonitor
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.
        /// </summary>
        public readonly string AlarmArn;
        /// <summary>
        /// ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor AlarmArn.
        /// </summary>
        public readonly string? AlarmRoleArn;

        [OutputConstructor]
        private EnvironmentMonitor(
            string alarmArn,

            string? alarmRoleArn)
        {
            AlarmArn = alarmArn;
            AlarmRoleArn = alarmRoleArn;
        }
    }
}
