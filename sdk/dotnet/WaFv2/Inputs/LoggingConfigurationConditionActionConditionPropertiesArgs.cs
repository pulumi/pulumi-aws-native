// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// A single action condition.
    /// </summary>
    public sealed class LoggingConfigurationConditionActionConditionPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
        /// </summary>
        [Input("action", required: true)]
        public Input<Pulumi.AwsNative.WaFv2.LoggingConfigurationConditionActionConditionPropertiesAction> Action { get; set; } = null!;

        public LoggingConfigurationConditionActionConditionPropertiesArgs()
        {
        }
        public static new LoggingConfigurationConditionActionConditionPropertiesArgs Empty => new LoggingConfigurationConditionActionConditionPropertiesArgs();
    }
}
