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
    public sealed class DashboardShapeConditionalFormat
    {
        /// <summary>
        /// The conditional formatting for the shape background color of a filled map visual.
        /// </summary>
        public readonly Outputs.DashboardConditionalFormattingColor BackgroundColor;

        [OutputConstructor]
        private DashboardShapeConditionalFormat(Outputs.DashboardConditionalFormattingColor backgroundColor)
        {
            BackgroundColor = backgroundColor;
        }
    }
}
