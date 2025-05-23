// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    [OutputType]
    public sealed class LoggingConfigurationCondition
    {
        /// <summary>
        /// A single action condition.
        /// </summary>
        public readonly Outputs.LoggingConfigurationConditionActionConditionProperties? ActionCondition;
        /// <summary>
        /// A single label name condition.
        /// </summary>
        public readonly Outputs.LoggingConfigurationConditionLabelNameConditionProperties? LabelNameCondition;

        [OutputConstructor]
        private LoggingConfigurationCondition(
            Outputs.LoggingConfigurationConditionActionConditionProperties? actionCondition,

            Outputs.LoggingConfigurationConditionLabelNameConditionProperties? labelNameCondition)
        {
            ActionCondition = actionCondition;
            LabelNameCondition = labelNameCondition;
        }
    }
}
