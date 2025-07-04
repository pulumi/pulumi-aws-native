// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Information about the automation configuration in numeric questions.
    /// </summary>
    [OutputType]
    public sealed class EvaluationFormNumericQuestionAutomation
    {
        public readonly object? AnswerSource;
        /// <summary>
        /// The property value of the automation.
        /// </summary>
        public readonly Outputs.EvaluationFormNumericQuestionPropertyValueAutomation? PropertyValue;

        [OutputConstructor]
        private EvaluationFormNumericQuestionAutomation(
            object? answerSource,

            Outputs.EvaluationFormNumericQuestionPropertyValueAutomation? propertyValue)
        {
            AnswerSource = answerSource;
            PropertyValue = propertyValue;
        }
    }
}
