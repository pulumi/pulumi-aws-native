// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Outputs
{

    [OutputType]
    public sealed class RuleMatch
    {
        /// <summary>
        /// The HTTP criteria that a rule must match.
        /// </summary>
        public readonly Outputs.RuleHttpMatch HttpMatch;

        [OutputConstructor]
        private RuleMatch(Outputs.RuleHttpMatch httpMatch)
        {
            HttpMatch = httpMatch;
        }
    }
}
