// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Inputs
{

    public sealed class DetectorRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The rule ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The time when the event type was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The detector for which the rule is associated.
        /// </summary>
        [Input("detectorId")]
        public Input<string>? DetectorId { get; set; }

        /// <summary>
        /// The rule expression. A rule expression captures the business logic. For more information, see [Rule language reference](https://docs.aws.amazon.com/frauddetector/latest/ug/rule-language-reference.html) .
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        /// <summary>
        /// The rule language.
        /// 
        /// Valid Value: DETECTORPL
        /// </summary>
        [Input("language")]
        public Input<Pulumi.AwsNative.FraudDetector.DetectorRuleLanguage>? Language { get; set; }

        /// <summary>
        /// The time when the event type was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        [Input("outcomes")]
        private InputList<Inputs.DetectorOutcomeArgs>? _outcomes;

        /// <summary>
        /// The rule outcome.
        /// </summary>
        public InputList<Inputs.DetectorOutcomeArgs> Outcomes
        {
            get => _outcomes ?? (_outcomes = new InputList<Inputs.DetectorOutcomeArgs>());
            set => _outcomes = value;
        }

        /// <summary>
        /// The rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// The rule version.
        /// </summary>
        [Input("ruleVersion")]
        public Input<string>? RuleVersion { get; set; }

        [Input("tags")]
        private InputList<Inputs.DetectorTagArgs>? _tags;

        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        public InputList<Inputs.DetectorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DetectorTagArgs>());
            set => _tags = value;
        }

        public DetectorRuleArgs()
        {
        }
        public static new DetectorRuleArgs Empty => new DetectorRuleArgs();
    }
}
