// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAxisLabelOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("applyTo")]
        public Input<Inputs.TemplateAxisLabelReferenceOptionsArgs>? ApplyTo { get; set; }

        [Input("customLabel")]
        public Input<string>? CustomLabel { get; set; }

        [Input("fontConfiguration")]
        public Input<Inputs.TemplateFontConfigurationArgs>? FontConfiguration { get; set; }

        public TemplateAxisLabelOptionsArgs()
        {
        }
        public static new TemplateAxisLabelOptionsArgs Empty => new TemplateAxisLabelOptionsArgs();
    }
}