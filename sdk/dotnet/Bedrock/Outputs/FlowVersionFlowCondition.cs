// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Condition branch for a condition node
    /// </summary>
    [OutputType]
    public sealed class FlowVersionFlowCondition
    {
        /// <summary>
        /// Expression for a condition in a flow
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// Name of a condition in a flow
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private FlowVersionFlowCondition(
            string? expression,

            string name)
        {
            Expression = expression;
            Name = name;
        }
    }
}