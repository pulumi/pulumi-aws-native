// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class ConfiguredTableAssociationAnalysisRulePolicyV12Properties
    {
        public readonly Outputs.ConfiguredTableAssociationAnalysisRuleCustom Custom;

        [OutputConstructor]
        private ConfiguredTableAssociationAnalysisRulePolicyV12Properties(Outputs.ConfiguredTableAssociationAnalysisRuleCustom custom)
        {
            Custom = custom;
        }
    }
}