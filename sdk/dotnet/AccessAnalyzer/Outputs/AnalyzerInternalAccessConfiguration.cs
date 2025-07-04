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
    /// Specifies the configuration of an internal access analyzer for an AWS organization or account. This configuration determines how the analyzer evaluates internal access within your AWS environment.
    /// </summary>
    [OutputType]
    public sealed class AnalyzerInternalAccessConfiguration
    {
        /// <summary>
        /// Contains information about analysis rules for the internal access analyzer. Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule.
        /// </summary>
        public readonly Outputs.AnalyzerInternalAccessConfigurationInternalAccessAnalysisRuleProperties? InternalAccessAnalysisRule;

        [OutputConstructor]
        private AnalyzerInternalAccessConfiguration(Outputs.AnalyzerInternalAccessConfigurationInternalAccessAnalysisRuleProperties? internalAccessAnalysisRule)
        {
            InternalAccessAnalysisRule = internalAccessAnalysisRule;
        }
    }
}
