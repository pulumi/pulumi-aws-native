// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.
    /// </summary>
    [OutputType]
    public sealed class FunctionSnapStartResponse
    {
        /// <summary>
        /// When set to ``PublishedVersions``, Lambda creates a snapshot of the execution environment when you publish a function version.
        /// </summary>
        public readonly Pulumi.AwsNative.Lambda.FunctionSnapStartResponseApplyOn? ApplyOn;
        /// <summary>
        /// When you provide a [qualified Amazon Resource Name (ARN)](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using), this response element indicates whether SnapStart is activated for the specified function version.
        /// </summary>
        public readonly Pulumi.AwsNative.Lambda.FunctionSnapStartResponseOptimizationStatus? OptimizationStatus;

        [OutputConstructor]
        private FunctionSnapStartResponse(
            Pulumi.AwsNative.Lambda.FunctionSnapStartResponseApplyOn? applyOn,

            Pulumi.AwsNative.Lambda.FunctionSnapStartResponseOptimizationStatus? optimizationStatus)
        {
            ApplyOn = applyOn;
            OptimizationStatus = optimizationStatus;
        }
    }
}
