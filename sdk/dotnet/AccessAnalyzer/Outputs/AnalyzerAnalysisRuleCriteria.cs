// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AccessAnalyzer.Outputs
{

    /// <summary>
    /// The criteria for an analysis rule for an analyzer.
    /// </summary>
    [OutputType]
    public sealed class AnalyzerAnalysisRuleCriteria
    {
        /// <summary>
        /// A list of AWS account IDs to apply to the analysis rule criteria. The accounts cannot include the organization analyzer owner account. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers.
        /// </summary>
        public readonly ImmutableArray<string> AccountIds;
        /// <summary>
        /// An array of key-value pairs to match for your resources. You can use the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        /// 
        /// For the tag key, you can specify a value that is 1 to 128 characters in length and cannot be prefixed with aws:.
        /// 
        /// For the tag value, you can specify a value that is 0 to 256 characters in length. If the specified tag value is 0 characters, the rule is applied to all principals with the specified tag key.
        /// </summary>
        public readonly ImmutableArray<ImmutableArray<Outputs.AnalyzerTag>> ResourceTags;

        [OutputConstructor]
        private AnalyzerAnalysisRuleCriteria(
            ImmutableArray<string> accountIds,

            ImmutableArray<ImmutableArray<Outputs.AnalyzerTag>> resourceTags)
        {
            AccountIds = accountIds;
            ResourceTags = resourceTags;
        }
    }
}
