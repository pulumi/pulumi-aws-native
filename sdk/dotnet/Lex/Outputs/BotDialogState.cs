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
    /// The current state of the conversation with the user.
    /// </summary>
    [OutputType]
    public sealed class BotDialogState
    {
        /// <summary>
        /// Defines the action that the bot executes at runtime when the conversation reaches this step.
        /// </summary>
        public readonly Outputs.BotDialogAction? DialogAction;
        /// <summary>
        /// Override settings to configure the intent state.
        /// </summary>
        public readonly Outputs.BotIntentOverride? Intent;
        /// <summary>
        /// List of session attributes to be applied when the conversation reaches this step.
        /// </summary>
        public readonly ImmutableArray<Outputs.BotSessionAttribute> SessionAttributes;

        [OutputConstructor]
        private BotDialogState(
            Outputs.BotDialogAction? dialogAction,

            Outputs.BotIntentOverride? intent,

            ImmutableArray<Outputs.BotSessionAttribute> sessionAttributes)
        {
            DialogAction = dialogAction;
            Intent = intent;
            SessionAttributes = sessionAttributes;
        }
    }
}