// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateListControlSelectAllOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? Visibility { get; set; }

        public TemplateListControlSelectAllOptionsArgs()
        {
        }
        public static new TemplateListControlSelectAllOptionsArgs Empty => new TemplateListControlSelectAllOptionsArgs();
    }
}