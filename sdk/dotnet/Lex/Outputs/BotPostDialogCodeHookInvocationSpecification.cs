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
    /// Specifies next steps to run after the dialog code hook finishes.
    /// </summary>
    [OutputType]
    public sealed class BotPostDialogCodeHookInvocationSpecification
    {
        /// <summary>
        /// A list of conditional branches to evaluate after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed.
        /// </summary>
        public readonly Outputs.BotConditionalSpecification? FailureConditional;
        /// <summary>
        /// Specifies the next step the bot runs after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed.
        /// </summary>
        public readonly Outputs.BotDialogState? FailureNextStep;
        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? FailureResponse;
        /// <summary>
        /// A list of conditional branches to evaluate after the dialog code hook finishes successfully.
        /// </summary>
        public readonly Outputs.BotConditionalSpecification? SuccessConditional;
        /// <summary>
        /// Specifics the next step the bot runs after the dialog code hook finishes successfully.
        /// </summary>
        public readonly Outputs.BotDialogState? SuccessNextStep;
        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? SuccessResponse;
        /// <summary>
        /// A list of conditional branches to evaluate if the code hook times out.
        /// </summary>
        public readonly Outputs.BotConditionalSpecification? TimeoutConditional;
        /// <summary>
        /// Specifies the next step that the bot runs when the code hook times out.
        /// </summary>
        public readonly Outputs.BotDialogState? TimeoutNextStep;
        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? TimeoutResponse;

        [OutputConstructor]
        private BotPostDialogCodeHookInvocationSpecification(
            Outputs.BotConditionalSpecification? failureConditional,

            Outputs.BotDialogState? failureNextStep,

            Outputs.BotResponseSpecification? failureResponse,

            Outputs.BotConditionalSpecification? successConditional,

            Outputs.BotDialogState? successNextStep,

            Outputs.BotResponseSpecification? successResponse,

            Outputs.BotConditionalSpecification? timeoutConditional,

            Outputs.BotDialogState? timeoutNextStep,

            Outputs.BotResponseSpecification? timeoutResponse)
        {
            FailureConditional = failureConditional;
            FailureNextStep = failureNextStep;
            FailureResponse = failureResponse;
            SuccessConditional = successConditional;
            SuccessNextStep = successNextStep;
            SuccessResponse = successResponse;
            TimeoutConditional = timeoutConditional;
            TimeoutNextStep = timeoutNextStep;
            TimeoutResponse = timeoutResponse;
        }
    }
}