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
    public sealed class TemplateTableSideBorderOptions
    {
        public readonly Outputs.TemplateTableBorderOptions? Bottom;
        public readonly Outputs.TemplateTableBorderOptions? InnerHorizontal;
        public readonly Outputs.TemplateTableBorderOptions? InnerVertical;
        public readonly Outputs.TemplateTableBorderOptions? Left;
        public readonly Outputs.TemplateTableBorderOptions? Right;
        public readonly Outputs.TemplateTableBorderOptions? Top;

        [OutputConstructor]
        private TemplateTableSideBorderOptions(
            Outputs.TemplateTableBorderOptions? bottom,

            Outputs.TemplateTableBorderOptions? innerHorizontal,

            Outputs.TemplateTableBorderOptions? innerVertical,

            Outputs.TemplateTableBorderOptions? left,

            Outputs.TemplateTableBorderOptions? right,

            Outputs.TemplateTableBorderOptions? top)
        {
            Bottom = bottom;
            InnerHorizontal = innerHorizontal;
            InnerVertical = innerVertical;
            Left = left;
            Right = right;
            Top = top;
        }
    }
}