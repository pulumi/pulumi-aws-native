// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    [OutputType]
    public sealed class AiAgentOrCondition0Properties
    {
        public readonly ImmutableArray<Outputs.AiAgentTagCondition> AndConditions;

        [OutputConstructor]
        private AiAgentOrCondition0Properties(ImmutableArray<Outputs.AiAgentTagCondition> andConditions)
        {
            AndConditions = andConditions;
        }
    }
}
