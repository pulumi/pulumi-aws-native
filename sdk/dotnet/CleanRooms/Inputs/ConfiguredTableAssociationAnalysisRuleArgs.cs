// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class ConfiguredTableAssociationAnalysisRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy of the configured table association analysis rule.
        /// </summary>
        [Input("policy", required: true)]
        public Input<Inputs.ConfiguredTableAssociationAnalysisRulePolicyArgs> Policy { get; set; } = null!;

        /// <summary>
        /// The type of the configured table association analysis rule.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.CleanRooms.ConfiguredTableAssociationAnalysisRuleType> Type { get; set; } = null!;

        public ConfiguredTableAssociationAnalysisRuleArgs()
        {
        }
        public static new ConfiguredTableAssociationAnalysisRuleArgs Empty => new ConfiguredTableAssociationAnalysisRuleArgs();
    }
}