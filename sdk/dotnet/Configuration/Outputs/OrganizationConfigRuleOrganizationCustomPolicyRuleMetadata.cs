// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    [OutputType]
    public sealed class OrganizationConfigRuleOrganizationCustomPolicyRuleMetadata
    {
        public readonly ImmutableArray<string> DebugLogDeliveryAccounts;
        public readonly string? Description;
        public readonly string? InputParameters;
        public readonly string? MaximumExecutionFrequency;
        public readonly ImmutableArray<string> OrganizationConfigRuleTriggerTypes;
        public readonly string PolicyText;
        public readonly string? ResourceIdScope;
        public readonly ImmutableArray<string> ResourceTypesScope;
        public readonly string Runtime;
        public readonly string? TagKeyScope;
        public readonly string? TagValueScope;

        [OutputConstructor]
        private OrganizationConfigRuleOrganizationCustomPolicyRuleMetadata(
            ImmutableArray<string> debugLogDeliveryAccounts,

            string? description,

            string? inputParameters,

            string? maximumExecutionFrequency,

            ImmutableArray<string> organizationConfigRuleTriggerTypes,

            string policyText,

            string? resourceIdScope,

            ImmutableArray<string> resourceTypesScope,

            string runtime,

            string? tagKeyScope,

            string? tagValueScope)
        {
            DebugLogDeliveryAccounts = debugLogDeliveryAccounts;
            Description = description;
            InputParameters = inputParameters;
            MaximumExecutionFrequency = maximumExecutionFrequency;
            OrganizationConfigRuleTriggerTypes = organizationConfigRuleTriggerTypes;
            PolicyText = policyText;
            ResourceIdScope = resourceIdScope;
            ResourceTypesScope = resourceTypesScope;
            Runtime = runtime;
            TagKeyScope = tagKeyScope;
            TagValueScope = tagValueScope;
        }
    }
}