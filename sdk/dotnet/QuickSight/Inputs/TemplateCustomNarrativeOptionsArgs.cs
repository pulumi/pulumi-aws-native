// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateCustomNarrativeOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("narrative", required: true)]
        public Input<string> Narrative { get; set; } = null!;

        public TemplateCustomNarrativeOptionsArgs()
        {
        }
        public static new TemplateCustomNarrativeOptionsArgs Empty => new TemplateCustomNarrativeOptionsArgs();
    }
}