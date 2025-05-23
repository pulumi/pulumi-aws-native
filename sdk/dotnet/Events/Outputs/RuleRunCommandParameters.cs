// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class RuleRunCommandParameters
    {
        /// <summary>
        /// Currently, we support including only one RunCommandTarget block, which specifies either an array of InstanceIds or a tag.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleRunCommandTarget> RunCommandTargets;

        [OutputConstructor]
        private RuleRunCommandParameters(ImmutableArray<Outputs.RuleRunCommandTarget> runCommandTargets)
        {
            RunCommandTargets = runCommandTargets;
        }
    }
}
