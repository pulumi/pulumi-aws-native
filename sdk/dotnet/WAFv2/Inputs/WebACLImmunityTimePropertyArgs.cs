// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    public sealed class WebACLImmunityTimePropertyArgs : global::Pulumi.ResourceArgs
    {
        [Input("immunityTime", required: true)]
        public Input<int> ImmunityTime { get; set; } = null!;

        public WebACLImmunityTimePropertyArgs()
        {
        }
        public static new WebACLImmunityTimePropertyArgs Empty => new WebACLImmunityTimePropertyArgs();
    }
}