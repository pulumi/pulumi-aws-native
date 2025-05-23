// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    [OutputType]
    public sealed class BotSlotCaptureSetting
    {
        /// <summary>
        /// A list of conditional branches to evaluate after the slot value is captured.
        /// </summary>
        public readonly Outputs.BotConditionalSpecification? CaptureConditional;
        /// <summary>
        /// Specifies the next step that the bot runs when the slot value is captured before the code hook times out.
        /// </summary>
        public readonly Outputs.BotDialogState? CaptureNextStep;
        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? CaptureResponse;
        /// <summary>
        /// Code hook called after Amazon Lex successfully captures a slot value.
        /// </summary>
        public readonly Outputs.BotDialogCodeHookInvocationSetting? CodeHook;
        /// <summary>
        /// Code hook called when Amazon Lex doesn't capture a slot value.
        /// </summary>
        public readonly Outputs.BotElicitationCodeHookInvocationSetting? ElicitationCodeHook;
        /// <summary>
        /// A list of conditional branches to evaluate when the slot value isn't captured.
        /// </summary>
        public readonly Outputs.BotConditionalSpecification? FailureConditional;
        /// <summary>
        /// Specifies the next step that the bot runs when the slot value code is not recognized.
        /// </summary>
        public readonly Outputs.BotDialogState? FailureNextStep;
        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input when the slot fails to be captured.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? FailureResponse;

        [OutputConstructor]
        private BotSlotCaptureSetting(
            Outputs.BotConditionalSpecification? captureConditional,

            Outputs.BotDialogState? captureNextStep,

            Outputs.BotResponseSpecification? captureResponse,

            Outputs.BotDialogCodeHookInvocationSetting? codeHook,

            Outputs.BotElicitationCodeHookInvocationSetting? elicitationCodeHook,

            Outputs.BotConditionalSpecification? failureConditional,

            Outputs.BotDialogState? failureNextStep,

            Outputs.BotResponseSpecification? failureResponse)
        {
            CaptureConditional = captureConditional;
            CaptureNextStep = captureNextStep;
            CaptureResponse = captureResponse;
            CodeHook = codeHook;
            ElicitationCodeHook = elicitationCodeHook;
            FailureConditional = failureConditional;
            FailureNextStep = failureNextStep;
            FailureResponse = failureResponse;
        }
    }
}
