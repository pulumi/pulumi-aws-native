// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Outputs
{

    [OutputType]
    public sealed class MatchingWorkflowResolutionTechniques
    {
        public readonly Outputs.MatchingWorkflowProviderProperties? ProviderProperties;
        public readonly Pulumi.AwsNative.EntityResolution.MatchingWorkflowResolutionTechniquesResolutionType? ResolutionType;
        public readonly Outputs.MatchingWorkflowRuleBasedProperties? RuleBasedProperties;

        [OutputConstructor]
        private MatchingWorkflowResolutionTechniques(
            Outputs.MatchingWorkflowProviderProperties? providerProperties,

            Pulumi.AwsNative.EntityResolution.MatchingWorkflowResolutionTechniquesResolutionType? resolutionType,

            Outputs.MatchingWorkflowRuleBasedProperties? ruleBasedProperties)
        {
            ProviderProperties = providerProperties;
            ResolutionType = resolutionType;
            RuleBasedProperties = ruleBasedProperties;
        }
    }
}