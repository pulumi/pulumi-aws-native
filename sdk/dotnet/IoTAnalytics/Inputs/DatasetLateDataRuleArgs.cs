// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class DatasetLateDataRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The information needed to configure the late data rule.
        /// </summary>
        [Input("ruleConfiguration", required: true)]
        public Input<Inputs.DatasetLateDataRuleConfigurationArgs> RuleConfiguration { get; set; } = null!;

        /// <summary>
        /// The name of the late data rule.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        public DatasetLateDataRuleArgs()
        {
        }
        public static new DatasetLateDataRuleArgs Empty => new DatasetLateDataRuleArgs();
    }
}
