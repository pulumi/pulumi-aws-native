// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Outputs
{

    [OutputType]
    public sealed class RuleHeaderMatch
    {
        public readonly bool? CaseSensitive;
        public readonly Outputs.RuleHeaderMatchType Match;
        public readonly string Name;

        [OutputConstructor]
        private RuleHeaderMatch(
            bool? caseSensitive,

            Outputs.RuleHeaderMatchType match,

            string name)
        {
            CaseSensitive = caseSensitive;
            Match = match;
            Name = name;
        }
    }
}