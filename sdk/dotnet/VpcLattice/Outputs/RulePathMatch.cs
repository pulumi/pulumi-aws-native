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
    public sealed class RulePathMatch
    {
        public readonly bool? CaseSensitive;
        public readonly Outputs.RulePathMatchType Match;

        [OutputConstructor]
        private RulePathMatch(
            bool? caseSensitive,

            Outputs.RulePathMatchType match)
        {
            CaseSensitive = caseSensitive;
            Match = match;
        }
    }
}