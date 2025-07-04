// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetAutomationRuleV2
    {
        /// <summary>
        /// Resource schema for AWS::SecurityHub::AutomationRuleV2
        /// </summary>
        public static Task<GetAutomationRuleV2Result> InvokeAsync(GetAutomationRuleV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutomationRuleV2Result>("aws-native:securityhub:getAutomationRuleV2", args ?? new GetAutomationRuleV2Args(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::SecurityHub::AutomationRuleV2
        /// </summary>
        public static Output<GetAutomationRuleV2Result> Invoke(GetAutomationRuleV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationRuleV2Result>("aws-native:securityhub:getAutomationRuleV2", args ?? new GetAutomationRuleV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::SecurityHub::AutomationRuleV2
        /// </summary>
        public static Output<GetAutomationRuleV2Result> Invoke(GetAutomationRuleV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationRuleV2Result>("aws-native:securityhub:getAutomationRuleV2", args ?? new GetAutomationRuleV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutomationRuleV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the automation rule
        /// </summary>
        [Input("ruleArn", required: true)]
        public string RuleArn { get; set; } = null!;

        public GetAutomationRuleV2Args()
        {
        }
        public static new GetAutomationRuleV2Args Empty => new GetAutomationRuleV2Args();
    }

    public sealed class GetAutomationRuleV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the automation rule
        /// </summary>
        [Input("ruleArn", required: true)]
        public Input<string> RuleArn { get; set; } = null!;

        public GetAutomationRuleV2InvokeArgs()
        {
        }
        public static new GetAutomationRuleV2InvokeArgs Empty => new GetAutomationRuleV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetAutomationRuleV2Result
    {
        /// <summary>
        /// A list of actions to be performed when the rule criteria is met
        /// </summary>
        public readonly ImmutableArray<Outputs.AutomationRuleV2AutomationRulesActionV2> Actions;
        /// <summary>
        /// The timestamp when the V2 automation rule was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The filtering type and configuration of the automation rule.
        /// </summary>
        public readonly Outputs.AutomationRuleV2Criteria? Criteria;
        /// <summary>
        /// A description of the automation rule
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The ARN of the automation rule
        /// </summary>
        public readonly string? RuleArn;
        /// <summary>
        /// The ID of the automation rule
        /// </summary>
        public readonly string? RuleId;
        /// <summary>
        /// The name of the automation rule
        /// </summary>
        public readonly string? RuleName;
        /// <summary>
        /// The value for the rule priority
        /// </summary>
        public readonly double? RuleOrder;
        /// <summary>
        /// The status of the automation rule
        /// </summary>
        public readonly Pulumi.AwsNative.SecurityHub.AutomationRuleV2RuleStatus? RuleStatus;
        /// <summary>
        /// A list of key-value pairs associated with the V2 automation rule.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The timestamp when the V2 automation rule was updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetAutomationRuleV2Result(
            ImmutableArray<Outputs.AutomationRuleV2AutomationRulesActionV2> actions,

            string? createdAt,

            Outputs.AutomationRuleV2Criteria? criteria,

            string? description,

            string? ruleArn,

            string? ruleId,

            string? ruleName,

            double? ruleOrder,

            Pulumi.AwsNative.SecurityHub.AutomationRuleV2RuleStatus? ruleStatus,

            ImmutableDictionary<string, string>? tags,

            string? updatedAt)
        {
            Actions = actions;
            CreatedAt = createdAt;
            Criteria = criteria;
            Description = description;
            RuleArn = ruleArn;
            RuleId = ruleId;
            RuleName = ruleName;
            RuleOrder = ruleOrder;
            RuleStatus = ruleStatus;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
