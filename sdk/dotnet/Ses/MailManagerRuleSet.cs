// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses
{
    /// <summary>
    /// Definition of AWS::SES::MailManagerRuleSet Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:ses:MailManagerRuleSet")]
    public partial class MailManagerRuleSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the rule set resource.
        /// </summary>
        [Output("ruleSetArn")]
        public Output<string> RuleSetArn { get; private set; } = null!;

        /// <summary>
        /// The identifier of the rule set.
        /// </summary>
        [Output("ruleSetId")]
        public Output<string> RuleSetId { get; private set; } = null!;

        /// <summary>
        /// A user-friendly name for the rule set.
        /// </summary>
        [Output("ruleSetName")]
        public Output<string?> RuleSetName { get; private set; } = null!;

        /// <summary>
        /// Conditional rules that are evaluated for determining actions on email.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.MailManagerRuleSetRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a MailManagerRuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MailManagerRuleSet(string name, MailManagerRuleSetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ses:MailManagerRuleSet", name, args ?? new MailManagerRuleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MailManagerRuleSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ses:MailManagerRuleSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MailManagerRuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MailManagerRuleSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MailManagerRuleSet(name, id, options);
        }
    }

    public sealed class MailManagerRuleSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-friendly name for the rule set.
        /// </summary>
        [Input("ruleSetName")]
        public Input<string>? RuleSetName { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.MailManagerRuleSetRuleArgs>? _rules;

        /// <summary>
        /// Conditional rules that are evaluated for determining actions on email.
        /// </summary>
        public InputList<Inputs.MailManagerRuleSetRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.MailManagerRuleSetRuleArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public MailManagerRuleSetArgs()
        {
        }
        public static new MailManagerRuleSetArgs Empty => new MailManagerRuleSetArgs();
    }
}