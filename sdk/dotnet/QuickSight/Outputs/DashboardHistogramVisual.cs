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
    public sealed class DashboardHistogramVisual
    {
        public readonly ImmutableArray<Outputs.DashboardVisualCustomAction> Actions;
        public readonly Outputs.DashboardHistogramConfiguration? ChartConfiguration;
        public readonly Outputs.DashboardVisualSubtitleLabelOptions? Subtitle;
        public readonly Outputs.DashboardVisualTitleLabelOptions? Title;
        public readonly string VisualId;

        [OutputConstructor]
        private DashboardHistogramVisual(
            ImmutableArray<Outputs.DashboardVisualCustomAction> actions,

            Outputs.DashboardHistogramConfiguration? chartConfiguration,

            Outputs.DashboardVisualSubtitleLabelOptions? subtitle,

            Outputs.DashboardVisualTitleLabelOptions? title,

            string visualId)
        {
            Actions = actions;
            ChartConfiguration = chartConfiguration;
            Subtitle = subtitle;
            Title = title;
            VisualId = visualId;
        }
    }
}