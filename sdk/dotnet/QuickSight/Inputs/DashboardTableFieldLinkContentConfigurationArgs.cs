// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTableFieldLinkContentConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("customIconContent")]
        public Input<Inputs.DashboardTableFieldCustomIconContentArgs>? CustomIconContent { get; set; }

        [Input("customTextContent")]
        public Input<Inputs.DashboardTableFieldCustomTextContentArgs>? CustomTextContent { get; set; }

        public DashboardTableFieldLinkContentConfigurationArgs()
        {
        }
        public static new DashboardTableFieldLinkContentConfigurationArgs Empty => new DashboardTableFieldLinkContentConfigurationArgs();
    }
}