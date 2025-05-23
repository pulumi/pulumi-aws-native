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
    public sealed class AnalysisLoadingAnimation
    {
        /// <summary>
        /// The visibility configuration of `LoadingAnimation` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? Visibility;

        [OutputConstructor]
        private AnalysisLoadingAnimation(Pulumi.AwsNative.QuickSight.AnalysisVisibility? visibility)
        {
            Visibility = visibility;
        }
    }
}
