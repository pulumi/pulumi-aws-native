// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    /// <summary>
    /// Resource schema for AWS::SecurityHub::AutomationRuleV2
    /// </summary>
    [AwsNativeResourceType("aws-native:securityhub:AutomationRuleV2")]
    public partial class AutomationRuleV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of actions to be performed when the rule criteria is met
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.AutomationRuleV2AutomationRulesActionV2>> Actions { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the V2 automation rule was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The filtering type and configuration of the automation rule.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.AutomationRuleV2Criteria> Criteria { get; private set; } = null!;

        /// <summary>
        /// A description of the automation rule
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The ARN of the automation rule
        /// </summary>
        [Output("ruleArn")]
        public Output<string> RuleArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the automation rule
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// The name of the automation rule
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;

        /// <summary>
        /// The value for the rule priority
        /// </summary>
        [Output("ruleOrder")]
        public Output<double> RuleOrder { get; private set; } = null!;

        /// <summary>
        /// The status of the automation rule
        /// </summary>
        [Output("ruleStatus")]
        public Output<Pulumi.AwsNative.SecurityHub.AutomationRuleV2RuleStatus?> RuleStatus { get; private set; } = null!;

        /// <summary>
        /// A list of key-value pairs associated with the V2 automation rule.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the V2 automation rule was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a AutomationRuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutomationRuleV2(string name, AutomationRuleV2Args args, CustomResourceOptions? options = null)
            : base("aws-native:securityhub:AutomationRuleV2", name, args ?? new AutomationRuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private AutomationRuleV2(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:securityhub:AutomationRuleV2", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AutomationRuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutomationRuleV2 Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AutomationRuleV2(name, id, options);
        }
    }

    public sealed class AutomationRuleV2Args : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.AutomationRuleV2AutomationRulesActionV2Args>? _actions;

        /// <summary>
        /// A list of actions to be performed when the rule criteria is met
        /// </summary>
        public InputList<Inputs.AutomationRuleV2AutomationRulesActionV2Args> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.AutomationRuleV2AutomationRulesActionV2Args>());
            set => _actions = value;
        }

        /// <summary>
        /// The filtering type and configuration of the automation rule.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.AutomationRuleV2CriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// A description of the automation rule
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The name of the automation rule
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        /// <summary>
        /// The value for the rule priority
        /// </summary>
        [Input("ruleOrder", required: true)]
        public Input<double> RuleOrder { get; set; } = null!;

        /// <summary>
        /// The status of the automation rule
        /// </summary>
        [Input("ruleStatus")]
        public Input<Pulumi.AwsNative.SecurityHub.AutomationRuleV2RuleStatus>? RuleStatus { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A list of key-value pairs associated with the V2 automation rule.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AutomationRuleV2Args()
        {
        }
        public static new AutomationRuleV2Args Empty => new AutomationRuleV2Args();
    }
}
