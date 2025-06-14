// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evs.Inputs
{

    public sealed class VcfHostnamesPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudBuilder", required: true)]
        public Input<string> CloudBuilder { get; set; } = null!;

        [Input("nsx", required: true)]
        public Input<string> Nsx { get; set; } = null!;

        [Input("nsxEdge1", required: true)]
        public Input<string> NsxEdge1 { get; set; } = null!;

        [Input("nsxEdge2", required: true)]
        public Input<string> NsxEdge2 { get; set; } = null!;

        [Input("nsxManager1", required: true)]
        public Input<string> NsxManager1 { get; set; } = null!;

        [Input("nsxManager2", required: true)]
        public Input<string> NsxManager2 { get; set; } = null!;

        [Input("nsxManager3", required: true)]
        public Input<string> NsxManager3 { get; set; } = null!;

        [Input("sddcManager", required: true)]
        public Input<string> SddcManager { get; set; } = null!;

        [Input("vCenter", required: true)]
        public Input<string> VCenter { get; set; } = null!;

        public VcfHostnamesPropertiesArgs()
        {
        }
        public static new VcfHostnamesPropertiesArgs Empty => new VcfHostnamesPropertiesArgs();
    }
}
