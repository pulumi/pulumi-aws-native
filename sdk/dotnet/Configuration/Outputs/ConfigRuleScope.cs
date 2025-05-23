// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    /// <summary>
    /// Defines which resources trigger an evaluation for an CC rule. The scope can include one or more resource types, a combination of a tag key and value, or a combination of one resource type and one resource ID. Specify a scope to constrain which resources trigger an evaluation for a rule. Otherwise, evaluations for the rule are triggered when any resource in your recording group changes in configuration.
    /// </summary>
    [OutputType]
    public sealed class ConfigRuleScope
    {
        /// <summary>
        /// The ID of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for ``ComplianceResourceTypes``.
        /// </summary>
        public readonly string? ComplianceResourceId;
        /// <summary>
        /// The resource types of only those AWS resources that you want to trigger an evaluation for the rule. You can only specify one type if you also specify a resource ID for ``ComplianceResourceId``.
        /// </summary>
        public readonly ImmutableArray<string> ComplianceResourceTypes;
        /// <summary>
        /// The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.
        /// </summary>
        public readonly string? TagKey;
        /// <summary>
        /// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule. If you specify a value for ``TagValue``, you must also specify a value for ``TagKey``.
        /// </summary>
        public readonly string? TagValue;

        [OutputConstructor]
        private ConfigRuleScope(
            string? complianceResourceId,

            ImmutableArray<string> complianceResourceTypes,

            string? tagKey,

            string? tagValue)
        {
            ComplianceResourceId = complianceResourceId;
            ComplianceResourceTypes = complianceResourceTypes;
            TagKey = tagKey;
            TagValue = tagValue;
        }
    }
}
