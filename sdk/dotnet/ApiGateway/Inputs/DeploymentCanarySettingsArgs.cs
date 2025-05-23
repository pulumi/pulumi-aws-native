// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    /// <summary>
    /// The ``DeploymentCanarySettings`` property type specifies settings for the canary deployment.
    /// </summary>
    public sealed class DeploymentCanarySettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The percentage (0.0-100.0) of traffic routed to the canary deployment.
        /// </summary>
        [Input("percentTraffic")]
        public Input<double>? PercentTraffic { get; set; }

        [Input("stageVariableOverrides")]
        private InputMap<string>? _stageVariableOverrides;

        /// <summary>
        /// A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.
        /// </summary>
        public InputMap<string> StageVariableOverrides
        {
            get => _stageVariableOverrides ?? (_stageVariableOverrides = new InputMap<string>());
            set => _stageVariableOverrides = value;
        }

        /// <summary>
        /// A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.
        /// </summary>
        [Input("useStageCache")]
        public Input<bool>? UseStageCache { get; set; }

        public DeploymentCanarySettingsArgs()
        {
        }
        public static new DeploymentCanarySettingsArgs Empty => new DeploymentCanarySettingsArgs();
    }
}
