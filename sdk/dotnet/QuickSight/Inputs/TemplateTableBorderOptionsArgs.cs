// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTableBorderOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        [Input("style")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateTableBorderStyle>? Style { get; set; }

        [Input("thickness")]
        public Input<double>? Thickness { get; set; }

        public TemplateTableBorderOptionsArgs()
        {
        }
        public static new TemplateTableBorderOptionsArgs Empty => new TemplateTableBorderOptionsArgs();
    }
}