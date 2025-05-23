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
    public sealed class AnalysisSpacing
    {
        /// <summary>
        /// String based length that is composed of value and unit
        /// </summary>
        public readonly string? Bottom;
        /// <summary>
        /// String based length that is composed of value and unit
        /// </summary>
        public readonly string? Left;
        /// <summary>
        /// String based length that is composed of value and unit
        /// </summary>
        public readonly string? Right;
        /// <summary>
        /// String based length that is composed of value and unit
        /// </summary>
        public readonly string? Top;

        [OutputConstructor]
        private AnalysisSpacing(
            string? bottom,

            string? left,

            string? right,

            string? top)
        {
            Bottom = bottom;
            Left = left;
            Right = right;
            Top = top;
        }
    }
}
