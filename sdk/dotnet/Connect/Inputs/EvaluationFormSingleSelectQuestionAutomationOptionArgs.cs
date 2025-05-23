// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// The automation options of the single select question.
    /// </summary>
    public sealed class EvaluationFormSingleSelectQuestionAutomationOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The automation option based on a rule category for the single select question.
        /// </summary>
        [Input("ruleCategory", required: true)]
        public Input<Inputs.EvaluationFormSingleSelectQuestionRuleCategoryAutomationArgs> RuleCategory { get; set; } = null!;

        public EvaluationFormSingleSelectQuestionAutomationOptionArgs()
        {
        }
        public static new EvaluationFormSingleSelectQuestionAutomationOptionArgs Empty => new EvaluationFormSingleSelectQuestionAutomationOptionArgs();
    }
}
