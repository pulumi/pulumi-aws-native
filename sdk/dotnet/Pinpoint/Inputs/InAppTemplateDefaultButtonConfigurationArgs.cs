// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class InAppTemplateDefaultButtonConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("backgroundColor")]
        public Input<string>? BackgroundColor { get; set; }

        [Input("borderRadius")]
        public Input<int>? BorderRadius { get; set; }

        [Input("buttonAction")]
        public Input<Pulumi.AwsNative.Pinpoint.InAppTemplateButtonAction>? ButtonAction { get; set; }

        [Input("link")]
        public Input<string>? Link { get; set; }

        [Input("text")]
        public Input<string>? Text { get; set; }

        [Input("textColor")]
        public Input<string>? TextColor { get; set; }

        public InAppTemplateDefaultButtonConfigurationArgs()
        {
        }
        public static new InAppTemplateDefaultButtonConfigurationArgs Empty => new InAppTemplateDefaultButtonConfigurationArgs();
    }
}