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
    /// Function definition
    /// </summary>
    [OutputType]
    public sealed class AgentFunction
    {
        /// <summary>
        /// Description of function
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Name for a resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parameters that the agent elicits from the user to fulfill the function.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.AgentParameterDetail>? Parameters;

        [OutputConstructor]
        private AgentFunction(
            string? description,

            string name,

            ImmutableDictionary<string, Outputs.AgentParameterDetail>? parameters)
        {
            Description = description;
            Name = name;
            Parameters = parameters;
        }
    }
}