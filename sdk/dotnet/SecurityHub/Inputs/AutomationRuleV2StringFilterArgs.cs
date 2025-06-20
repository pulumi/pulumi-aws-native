// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Inputs
{

    /// <summary>
    /// A string filter for filtering findings
    /// </summary>
    public sealed class AutomationRuleV2StringFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition to apply to a string value when filtering findings
        /// </summary>
        [Input("comparison", required: true)]
        public Input<Pulumi.AwsNative.SecurityHub.AutomationRuleV2StringFilterComparison> Comparison { get; set; } = null!;

        /// <summary>
        /// The string filter value
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public AutomationRuleV2StringFilterArgs()
        {
        }
        public static new AutomationRuleV2StringFilterArgs Empty => new AutomationRuleV2StringFilterArgs();
    }
}
