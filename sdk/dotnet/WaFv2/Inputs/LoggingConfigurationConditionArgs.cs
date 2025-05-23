// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    public sealed class LoggingConfigurationConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A single action condition.
        /// </summary>
        [Input("actionCondition")]
        public Input<Inputs.LoggingConfigurationConditionActionConditionPropertiesArgs>? ActionCondition { get; set; }

        /// <summary>
        /// A single label name condition.
        /// </summary>
        [Input("labelNameCondition")]
        public Input<Inputs.LoggingConfigurationConditionLabelNameConditionPropertiesArgs>? LabelNameCondition { get; set; }

        public LoggingConfigurationConditionArgs()
        {
        }
        public static new LoggingConfigurationConditionArgs Empty => new LoggingConfigurationConditionArgs();
    }
}
