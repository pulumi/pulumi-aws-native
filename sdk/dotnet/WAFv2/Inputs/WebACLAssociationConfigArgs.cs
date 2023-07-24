// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// AssociationConfig for body inspection
    /// </summary>
    public sealed class WebACLAssociationConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("requestBody")]
        public Input<Inputs.WebACLRequestBodyArgs>? RequestBody { get; set; }

        public WebACLAssociationConfigArgs()
        {
        }
        public static new WebACLAssociationConfigArgs Empty => new WebACLAssociationConfigArgs();
    }
}