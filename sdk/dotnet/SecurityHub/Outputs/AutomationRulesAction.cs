// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Outputs
{

    /// <summary>
    /// One or more actions that ASHlong takes when a finding matches the defined criteria of a rule.
    /// </summary>
    [OutputType]
    public sealed class AutomationRulesAction
    {
        /// <summary>
        /// Specifies that the automation rule action is an update to a finding field.
        /// </summary>
        public readonly Outputs.AutomationRulesFindingFieldsUpdate FindingFieldsUpdate;
        /// <summary>
        /// Specifies the type of action that Security Hub takes when a finding matches the defined criteria of a rule.
        /// </summary>
        public readonly Pulumi.AwsNative.SecurityHub.AutomationRulesActionType Type;

        [OutputConstructor]
        private AutomationRulesAction(
            Outputs.AutomationRulesFindingFieldsUpdate findingFieldsUpdate,

            Pulumi.AwsNative.SecurityHub.AutomationRulesActionType type)
        {
            FindingFieldsUpdate = findingFieldsUpdate;
            Type = type;
        }
    }
}
