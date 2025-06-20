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
    /// The criteria for an analysis rule for an internal access analyzer.
    /// </summary>
    [OutputType]
    public sealed class AnalyzerInternalAccessAnalysisRuleCriteria
    {
        /// <summary>
        /// A list of AWS account IDs to apply to the internal access analysis rule criteria. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers and cannot include the organization owner account.
        /// </summary>
        public readonly ImmutableArray<string> AccountIds;
        /// <summary>
        /// A list of resource ARNs to apply to the internal access analysis rule criteria. The analyzer will only generate findings for resources that match these ARNs.
        /// </summary>
        public readonly ImmutableArray<string> ResourceArns;
        /// <summary>
        /// A list of resource types to apply to the internal access analysis rule criteria. The analyzer will only generate findings for resources of these types.
        /// </summary>
        public readonly ImmutableArray<string> ResourceTypes;

        [OutputConstructor]
        private AnalyzerInternalAccessAnalysisRuleCriteria(
            ImmutableArray<string> accountIds,

            ImmutableArray<string> resourceArns,

            ImmutableArray<string> resourceTypes)
        {
            AccountIds = accountIds;
            ResourceArns = resourceArns;
            ResourceTypes = resourceTypes;
        }
    }
}
