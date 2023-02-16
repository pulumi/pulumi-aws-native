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
    /// A set of actions that Amazon Lex should run if none of the other conditions are met.
    /// </summary>
    [OutputType]
    public sealed class BotDefaultConditionalBranch
    {
        /// <summary>
        /// The next step in the conversation.
        /// </summary>
        public readonly Outputs.BotDialogState? NextStep;
        /// <summary>
        /// Specifies a list of message groups that Amazon Lex uses to respond the user input.
        /// </summary>
        public readonly Outputs.BotResponseSpecification? Response;

        [OutputConstructor]
        private BotDefaultConditionalBranch(
            Outputs.BotDialogState? nextStep,

            Outputs.BotResponseSpecification? response)
        {
            NextStep = nextStep;
            Response = response;
        }
    }
}