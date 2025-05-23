// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotConditionalSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditionalBranches", required: true)]
        private InputList<Inputs.BotConditionalBranchArgs>? _conditionalBranches;

        /// <summary>
        /// A list of conditional branches. A conditional branch is made up of a condition, a response and a next step. The response and next step are executed when the condition is true.
        /// </summary>
        public InputList<Inputs.BotConditionalBranchArgs> ConditionalBranches
        {
            get => _conditionalBranches ?? (_conditionalBranches = new InputList<Inputs.BotConditionalBranchArgs>());
            set => _conditionalBranches = value;
        }

        /// <summary>
        /// The conditional branch that should be followed when the conditions for other branches are not satisfied. A conditional branch is made up of a condition, a response and a next step.
        /// </summary>
        [Input("defaultBranch", required: true)]
        public Input<Inputs.BotDefaultConditionalBranchArgs> DefaultBranch { get; set; } = null!;

        /// <summary>
        /// Determines whether a conditional branch is active. When `IsActive` is false, the conditions are not evaluated.
        /// </summary>
        [Input("isActive", required: true)]
        public Input<bool> IsActive { get; set; } = null!;

        public BotConditionalSpecificationArgs()
        {
        }
        public static new BotConditionalSpecificationArgs Empty => new BotConditionalSpecificationArgs();
    }
}
