// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    /// <summary>
    /// Specifies settings for the canary deployment in this stage.
    /// </summary>
    public sealed class StageCanarySettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the deployment that the stage points to.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// The percentage (0-100) of traffic diverted to a canary deployment.
        /// </summary>
        [Input("percentTraffic")]
        public Input<double>? PercentTraffic { get; set; }

        /// <summary>
        /// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.
        /// </summary>
        [Input("stageVariableOverrides")]
        public Input<object>? StageVariableOverrides { get; set; }

        /// <summary>
        /// Whether the canary deployment uses the stage cache or not.
        /// </summary>
        [Input("useStageCache")]
        public Input<bool>? UseStageCache { get; set; }

        public StageCanarySettingArgs()
        {
        }
        public static new StageCanarySettingArgs Empty => new StageCanarySettingArgs();
    }
}