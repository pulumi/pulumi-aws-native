// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses
{
    public static class GetMailManagerRuleSet
    {
        /// <summary>
        /// Definition of AWS::SES::MailManagerRuleSet Resource Type
        /// </summary>
        public static Task<GetMailManagerRuleSetResult> InvokeAsync(GetMailManagerRuleSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMailManagerRuleSetResult>("aws-native:ses:getMailManagerRuleSet", args ?? new GetMailManagerRuleSetArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::SES::MailManagerRuleSet Resource Type
        /// </summary>
        public static Output<GetMailManagerRuleSetResult> Invoke(GetMailManagerRuleSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMailManagerRuleSetResult>("aws-native:ses:getMailManagerRuleSet", args ?? new GetMailManagerRuleSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMailManagerRuleSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the rule set.
        /// </summary>
        [Input("ruleSetId", required: true)]
        public string RuleSetId { get; set; } = null!;

        public GetMailManagerRuleSetArgs()
        {
        }
        public static new GetMailManagerRuleSetArgs Empty => new GetMailManagerRuleSetArgs();
    }

    public sealed class GetMailManagerRuleSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the rule set.
        /// </summary>
        [Input("ruleSetId", required: true)]
        public Input<string> RuleSetId { get; set; } = null!;

        public GetMailManagerRuleSetInvokeArgs()
        {
        }
        public static new GetMailManagerRuleSetInvokeArgs Empty => new GetMailManagerRuleSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetMailManagerRuleSetResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the rule set resource.
        /// </summary>
        public readonly string? RuleSetArn;
        /// <summary>
        /// The identifier of the rule set.
        /// </summary>
        public readonly string? RuleSetId;
        /// <summary>
        /// A user-friendly name for the rule set.
        /// </summary>
        public readonly string? RuleSetName;
        /// <summary>
        /// Conditional rules that are evaluated for determining actions on email.
        /// </summary>
        public readonly ImmutableArray<Outputs.MailManagerRuleSetRule> Rules;
        /// <summary>
        /// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetMailManagerRuleSetResult(
            string? ruleSetArn,

            string? ruleSetId,

            string? ruleSetName,

            ImmutableArray<Outputs.MailManagerRuleSetRule> rules,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            RuleSetArn = ruleSetArn;
            RuleSetId = ruleSetId;
            RuleSetName = ruleSetName;
            Rules = rules;
            Tags = tags;
        }
    }
}