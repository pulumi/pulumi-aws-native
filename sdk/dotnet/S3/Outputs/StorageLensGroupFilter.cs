// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Sets the Storage Lens Group filter.
    /// </summary>
    [OutputType]
    public sealed class StorageLensGroupFilter
    {
        public readonly Outputs.StorageLensGroupAnd? And;
        public readonly ImmutableArray<string> MatchAnyPrefix;
        public readonly ImmutableArray<string> MatchAnySuffix;
        public readonly ImmutableArray<Outputs.StorageLensGroupTag> MatchAnyTag;
        public readonly Outputs.StorageLensGroupMatchObjectAge? MatchObjectAge;
        public readonly Outputs.StorageLensGroupMatchObjectSize? MatchObjectSize;
        public readonly Outputs.StorageLensGroupOr? Or;

        [OutputConstructor]
        private StorageLensGroupFilter(
            Outputs.StorageLensGroupAnd? and,

            ImmutableArray<string> matchAnyPrefix,

            ImmutableArray<string> matchAnySuffix,

            ImmutableArray<Outputs.StorageLensGroupTag> matchAnyTag,

            Outputs.StorageLensGroupMatchObjectAge? matchObjectAge,

            Outputs.StorageLensGroupMatchObjectSize? matchObjectSize,

            Outputs.StorageLensGroupOr? or)
        {
            And = and;
            MatchAnyPrefix = matchAnyPrefix;
            MatchAnySuffix = matchAnySuffix;
            MatchAnyTag = matchAnyTag;
            MatchObjectAge = matchObjectAge;
            MatchObjectSize = matchObjectSize;
            Or = or;
        }
    }
}