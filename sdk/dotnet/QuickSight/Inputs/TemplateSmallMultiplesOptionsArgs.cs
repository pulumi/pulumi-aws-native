// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSmallMultiplesOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxVisibleColumns")]
        public Input<double>? MaxVisibleColumns { get; set; }

        [Input("maxVisibleRows")]
        public Input<double>? MaxVisibleRows { get; set; }

        [Input("panelConfiguration")]
        public Input<Inputs.TemplatePanelConfigurationArgs>? PanelConfiguration { get; set; }

        public TemplateSmallMultiplesOptionsArgs()
        {
        }
        public static new TemplateSmallMultiplesOptionsArgs Empty => new TemplateSmallMultiplesOptionsArgs();
    }
}