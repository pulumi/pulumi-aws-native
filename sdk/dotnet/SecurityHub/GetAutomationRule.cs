// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetAutomationRule
    {
        /// <summary>
        /// The ``AWS::SecurityHub::AutomationRule`` resource specifies an automation rule based on input parameters. For more information, see [Automation rules](https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html) in the *User Guide*.
        /// </summary>
        public static Task<GetAutomationRuleResult> InvokeAsync(GetAutomationRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutomationRuleResult>("aws-native:securityhub:getAutomationRule", args ?? new GetAutomationRuleArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SecurityHub::AutomationRule`` resource specifies an automation rule based on input parameters. For more information, see [Automation rules](https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html) in the *User Guide*.
        /// </summary>
        public static Output<GetAutomationRuleResult> Invoke(GetAutomationRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationRuleResult>("aws-native:securityhub:getAutomationRule", args ?? new GetAutomationRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SecurityHub::AutomationRule`` resource specifies an automation rule based on input parameters. For more information, see [Automation rules](https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html) in the *User Guide*.
        /// </summary>
        public static Output<GetAutomationRuleResult> Invoke(GetAutomationRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationRuleResult>("aws-native:securityhub:getAutomationRule", args ?? new GetAutomationRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutomationRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the automation rule that you create. For example, `arn:aws:securityhub:us-east-1:123456789012:automation-rule/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111` .
        /// </summary>
        [Input("ruleArn", required: true)]
        public string RuleArn { get; set; } = null!;

        public GetAutomationRuleArgs()
        {
        }
        public static new GetAutomationRuleArgs Empty => new GetAutomationRuleArgs();
    }

    public sealed class GetAutomationRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the automation rule that you create. For example, `arn:aws:securityhub:us-east-1:123456789012:automation-rule/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111` .
        /// </summary>
        [Input("ruleArn", required: true)]
        public Input<string> RuleArn { get; set; } = null!;

        public GetAutomationRuleInvokeArgs()
        {
        }
        public static new GetAutomationRuleInvokeArgs Empty => new GetAutomationRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutomationRuleResult
    {
        /// <summary>
        /// One or more actions to update finding fields if a finding matches the conditions specified in ``Criteria``.
        /// </summary>
        public readonly ImmutableArray<Outputs.AutomationRulesAction> Actions;
        /// <summary>
        /// A timestamp that indicates when the rule was created.
        /// 
        /// Uses the `date-time` format specified in [RFC 3339 section 5.6, Internet Date/Time Format](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3339#section-5.6) . The value cannot contain spaces. For example, `2020-03-22T13:22:13.933Z` .
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The principal that created the rule. For example, `arn:aws:sts::123456789012:assumed-role/Developer-Role/JaneDoe` .
        /// </summary>
        public readonly string? CreatedBy;
        /// <summary>
        /// A set of [Security Finding Format (ASFF)](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html) finding field attributes and corresponding expected values that ASH uses to filter findings. If a rule is enabled and a finding matches the criteria specified in this parameter, ASH applies the rule action to the finding.
        /// </summary>
        public readonly Outputs.AutomationRulesFindingFilters? Criteria;
        /// <summary>
        /// A description of the rule.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria. This is useful when a finding matches the criteria for multiple rules, and each rule has different actions. If a rule is terminal, Security Hub applies the rule action to a finding that matches the rule criteria and doesn't evaluate other rules for the finding. By default, a rule isn't terminal.
        /// </summary>
        public readonly bool? IsTerminal;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the automation rule that you create. For example, `arn:aws:securityhub:us-east-1:123456789012:automation-rule/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111` .
        /// </summary>
        public readonly string? RuleArn;
        /// <summary>
        /// The name of the rule.
        /// </summary>
        public readonly string? RuleName;
        /// <summary>
        /// An integer ranging from 1 to 1000 that represents the order in which the rule action is applied to findings. Security Hub applies rules with lower values for this parameter first.
        /// </summary>
        public readonly int? RuleOrder;
        /// <summary>
        /// Whether the rule is active after it is created. If this parameter is equal to ``ENABLED``, ASH applies the rule to findings and finding updates after the rule is created.
        /// </summary>
        public readonly Pulumi.AwsNative.SecurityHub.AutomationRuleRuleStatus? RuleStatus;
        /// <summary>
        /// User-defined tags associated with an automation rule.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// A timestamp that indicates when the rule was most recently updated.
        /// 
        /// Uses the `date-time` format specified in [RFC 3339 section 5.6, Internet Date/Time Format](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3339#section-5.6) . The value cannot contain spaces. For example, `2020-03-22T13:22:13.933Z` .
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetAutomationRuleResult(
            ImmutableArray<Outputs.AutomationRulesAction> actions,

            string? createdAt,

            string? createdBy,

            Outputs.AutomationRulesFindingFilters? criteria,

            string? description,

            bool? isTerminal,

            string? ruleArn,

            string? ruleName,

            int? ruleOrder,

            Pulumi.AwsNative.SecurityHub.AutomationRuleRuleStatus? ruleStatus,

            ImmutableDictionary<string, string>? tags,

            string? updatedAt)
        {
            Actions = actions;
            CreatedAt = createdAt;
            CreatedBy = createdBy;
            Criteria = criteria;
            Description = description;
            IsTerminal = isTerminal;
            RuleArn = ruleArn;
            RuleName = ruleName;
            RuleOrder = ruleOrder;
            RuleStatus = ruleStatus;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
