// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardSectionBasedLayoutPaperCanvasSizeOptions
    {
        public readonly Outputs.DashboardSpacing? PaperMargin;
        public readonly Pulumi.AwsNative.QuickSight.DashboardPaperOrientation? PaperOrientation;
        public readonly Pulumi.AwsNative.QuickSight.DashboardPaperSize? PaperSize;

        [OutputConstructor]
        private DashboardSectionBasedLayoutPaperCanvasSizeOptions(
            Outputs.DashboardSpacing? paperMargin,

            Pulumi.AwsNative.QuickSight.DashboardPaperOrientation? paperOrientation,

            Pulumi.AwsNative.QuickSight.DashboardPaperSize? paperSize)
        {
            PaperMargin = paperMargin;
            PaperOrientation = paperOrientation;
            PaperSize = paperSize;
        }
    }
}