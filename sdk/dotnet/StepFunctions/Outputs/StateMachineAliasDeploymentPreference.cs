// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions.Outputs
{

    /// <summary>
    /// The settings to enable gradual state machine deployments.
    /// </summary>
    [OutputType]
    public sealed class StateMachineAliasDeploymentPreference
    {
        /// <summary>
        /// A list of CloudWatch alarm names that will be monitored during the deployment. The deployment will fail and rollback if any alarms go into ALARM state.
        /// </summary>
        public readonly ImmutableArray<string> Alarms;
        /// <summary>
        /// The time in minutes between each traffic shifting increment.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// The percentage of traffic to shift to the new version in each increment.
        /// </summary>
        public readonly int? Percentage;
        public readonly string StateMachineVersionArn;
        /// <summary>
        /// The type of deployment to perform.
        /// </summary>
        public readonly Pulumi.AwsNative.StepFunctions.StateMachineAliasDeploymentPreferenceType Type;

        [OutputConstructor]
        private StateMachineAliasDeploymentPreference(
            ImmutableArray<string> alarms,

            int? interval,

            int? percentage,

            string stateMachineVersionArn,

            Pulumi.AwsNative.StepFunctions.StateMachineAliasDeploymentPreferenceType type)
        {
            Alarms = alarms;
            Interval = interval;
            Percentage = percentage;
            StateMachineVersionArn = stateMachineVersionArn;
            Type = type;
        }
    }
}