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
    public sealed class AnalysisRangeEndsLabelType
    {
        /// <summary>
        /// The visibility of the range ends label.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? Visibility;

        [OutputConstructor]
        private AnalysisRangeEndsLabelType(Pulumi.AwsNative.QuickSight.AnalysisVisibility? visibility)
        {
            Visibility = visibility;
        }
    }
}
