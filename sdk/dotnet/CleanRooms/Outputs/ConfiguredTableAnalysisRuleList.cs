// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class ConfiguredTableAnalysisRuleList
    {
        public readonly ImmutableArray<Pulumi.AwsNative.CleanRooms.ConfiguredTableJoinOperator> AllowedJoinOperators;
        public readonly ImmutableArray<string> JoinColumns;
        public readonly ImmutableArray<string> ListColumns;

        [OutputConstructor]
        private ConfiguredTableAnalysisRuleList(
            ImmutableArray<Pulumi.AwsNative.CleanRooms.ConfiguredTableJoinOperator> allowedJoinOperators,

            ImmutableArray<string> joinColumns,

            ImmutableArray<string> listColumns)
        {
            AllowedJoinOperators = allowedJoinOperators;
            JoinColumns = joinColumns;
            ListColumns = listColumns;
        }
    }
}