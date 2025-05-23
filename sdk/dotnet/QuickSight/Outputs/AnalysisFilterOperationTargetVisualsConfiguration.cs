// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisFilterOperationTargetVisualsConfiguration
    {
        /// <summary>
        /// The configuration of the same-sheet target visuals that you want to be filtered.
        /// </summary>
        public readonly Outputs.AnalysisSameSheetTargetVisualConfiguration? SameSheetTargetVisualConfiguration;

        [OutputConstructor]
        private AnalysisFilterOperationTargetVisualsConfiguration(Outputs.AnalysisSameSheetTargetVisualConfiguration? sameSheetTargetVisualConfiguration)
        {
            SameSheetTargetVisualConfiguration = sameSheetTargetVisualConfiguration;
        }
    }
}
