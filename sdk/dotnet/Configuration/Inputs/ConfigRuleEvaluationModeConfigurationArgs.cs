// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    /// <summary>
    /// Evaluation mode for the AWS Config rule
    /// </summary>
    public sealed class ConfigRuleEvaluationModeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode of evaluation of AWS Config rule
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public ConfigRuleEvaluationModeConfigurationArgs()
        {
        }
        public static new ConfigRuleEvaluationModeConfigurationArgs Empty => new ConfigRuleEvaluationModeConfigurationArgs();
    }
}