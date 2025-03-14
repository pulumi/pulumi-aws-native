// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Provides a list of conditional branches. Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
    /// </summary>
    [OutputType]
    public sealed class BotConditionalSpecification
    {
        /// <summary>
        /// A list of conditional branches. A conditional branch is made up of a condition, a response and a next step. The response and next step are executed when the condition is true.
        /// </summary>
        public readonly ImmutableArray<Outputs.BotConditionalBranch> ConditionalBranches;
        /// <summary>
        /// The conditional branch that should be followed when the conditions for other branches are not satisfied. A conditional branch is made up of a condition, a response and a next step.
        /// </summary>
        public readonly Outputs.BotDefaultConditionalBranch DefaultBranch;
        /// <summary>
        /// Determines whether a conditional branch is active. When active is false, the conditions are not evaluated.
        /// </summary>
        public readonly bool IsActive;

        [OutputConstructor]
        private BotConditionalSpecification(
            ImmutableArray<Outputs.BotConditionalBranch> conditionalBranches,

            Outputs.BotDefaultConditionalBranch defaultBranch,

            bool isActive)
        {
            ConditionalBranches = conditionalBranches;
            DefaultBranch = defaultBranch;
            IsActive = isActive;
        }
    }
}
