// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class ConfiguredTableAnalysisRule
    {
        public readonly Outputs.ConfiguredTableAnalysisRulePolicy Policy;
        public readonly Pulumi.AwsNative.CleanRooms.ConfiguredTableAnalysisRuleType Type;

        [OutputConstructor]
        private ConfiguredTableAnalysisRule(
            Outputs.ConfiguredTableAnalysisRulePolicy policy,

            Pulumi.AwsNative.CleanRooms.ConfiguredTableAnalysisRuleType type)
        {
            Policy = policy;
            Type = type;
        }
    }
}