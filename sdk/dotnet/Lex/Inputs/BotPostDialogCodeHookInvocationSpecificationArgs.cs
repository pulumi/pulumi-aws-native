// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// Specifies next steps to run after the dialog code hook finishes.
    /// </summary>
    public sealed class BotPostDialogCodeHookInvocationSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of conditional branches to evaluate after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed.
        /// </summary>
        [Input("failureConditional")]
        public Input<Inputs.BotConditionalSpecificationArgs>? FailureConditional { get; set; }

        /// <summary>
        /// Specifies the next step the bot runs after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed.
        /// </summary>
        [Input("failureNextStep")]
        public Input<Inputs.BotDialogStateArgs>? FailureNextStep { get; set; }

        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        [Input("failureResponse")]
        public Input<Inputs.BotResponseSpecificationArgs>? FailureResponse { get; set; }

        /// <summary>
        /// A list of conditional branches to evaluate after the dialog code hook finishes successfully.
        /// </summary>
        [Input("successConditional")]
        public Input<Inputs.BotConditionalSpecificationArgs>? SuccessConditional { get; set; }

        /// <summary>
        /// Specifics the next step the bot runs after the dialog code hook finishes successfully.
        /// </summary>
        [Input("successNextStep")]
        public Input<Inputs.BotDialogStateArgs>? SuccessNextStep { get; set; }

        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        [Input("successResponse")]
        public Input<Inputs.BotResponseSpecificationArgs>? SuccessResponse { get; set; }

        /// <summary>
        /// A list of conditional branches to evaluate if the code hook times out.
        /// </summary>
        [Input("timeoutConditional")]
        public Input<Inputs.BotConditionalSpecificationArgs>? TimeoutConditional { get; set; }

        /// <summary>
        /// Specifies the next step that the bot runs when the code hook times out.
        /// </summary>
        [Input("timeoutNextStep")]
        public Input<Inputs.BotDialogStateArgs>? TimeoutNextStep { get; set; }

        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        [Input("timeoutResponse")]
        public Input<Inputs.BotResponseSpecificationArgs>? TimeoutResponse { get; set; }

        public BotPostDialogCodeHookInvocationSpecificationArgs()
        {
        }
        public static new BotPostDialogCodeHookInvocationSpecificationArgs Empty => new BotPostDialogCodeHookInvocationSpecificationArgs();
    }
}