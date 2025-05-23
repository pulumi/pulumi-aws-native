// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class IndexRelevance
    {
        public readonly string? Duration;
        public readonly bool? Freshness;
        public readonly int? Importance;
        public readonly Pulumi.AwsNative.Kendra.IndexOrder? RankOrder;
        public readonly ImmutableArray<Outputs.IndexValueImportanceItem> ValueImportanceItems;

        [OutputConstructor]
        private IndexRelevance(
            string? duration,

            bool? freshness,

            int? importance,

            Pulumi.AwsNative.Kendra.IndexOrder? rankOrder,

            ImmutableArray<Outputs.IndexValueImportanceItem> valueImportanceItems)
        {
            Duration = duration;
            Freshness = freshness;
            Importance = importance;
            RankOrder = rankOrder;
            ValueImportanceItems = valueImportanceItems;
        }
    }
}
