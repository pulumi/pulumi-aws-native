// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class ConfiguredTableAnalysisRulePolicyV11PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregation", required: true)]
        public Input<Inputs.ConfiguredTableAnalysisRuleAggregationArgs> Aggregation { get; set; } = null!;

        public ConfiguredTableAnalysisRulePolicyV11PropertiesArgs()
        {
        }
        public static new ConfiguredTableAnalysisRulePolicyV11PropertiesArgs Empty => new ConfiguredTableAnalysisRulePolicyV11PropertiesArgs();
    }
}
