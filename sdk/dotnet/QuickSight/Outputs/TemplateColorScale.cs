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
    public sealed class TemplateColorScale
    {
        public readonly Pulumi.AwsNative.QuickSight.TemplateColorFillType ColorFillType;
        public readonly ImmutableArray<Outputs.TemplateDataColor> Colors;
        public readonly Outputs.TemplateDataColor? NullValueColor;

        [OutputConstructor]
        private TemplateColorScale(
            Pulumi.AwsNative.QuickSight.TemplateColorFillType colorFillType,

            ImmutableArray<Outputs.TemplateDataColor> colors,

            Outputs.TemplateDataColor? nullValueColor)
        {
            ColorFillType = colorFillType;
            Colors = colors;
            NullValueColor = nullValueColor;
        }
    }
}