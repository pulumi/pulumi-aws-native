// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Provides a statement the Amazon Lex conveys to the user when the intent is successfully fulfilled.
    /// </summary>
    [OutputType]
    public sealed class BotIntentClosingSetting
    {
        /// <summary>
        /// The response that Amazon Lex sends to the user when the intent is complete.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? ClosingResponse;
        /// <summary>
        /// A list of conditional branches associated with the intent's closing response. These branches are executed when the nextStep attribute is set to EvalutateConditional.
        /// </summary>
        public readonly Outputs.BotConditionalSpecification? Conditional;
        /// <summary>
        /// Specifies whether an intent's closing response is used. When this field is false, the closing response isn't sent to the user. If the active field isn't specified, the default is true.
        /// </summary>
        public readonly bool? IsActive;
        /// <summary>
        /// Specifies the next step that the bot executes after playing the intent's closing response.
        /// </summary>
        public readonly Outputs.BotDialogState? NextStep;

        [OutputConstructor]
        private BotIntentClosingSetting(
            Outputs.BotResponseSpecification? closingResponse,

            Outputs.BotConditionalSpecification? conditional,

            bool? isActive,

            Outputs.BotDialogState? nextStep)
        {
            ClosingResponse = closingResponse;
            Conditional = conditional;
            IsActive = isActive;
            NextStep = nextStep;
        }
    }
}