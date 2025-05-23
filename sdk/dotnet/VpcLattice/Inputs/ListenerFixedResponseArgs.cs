// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Inputs
{

    public sealed class ListenerFixedResponseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTTP response code. Only `404` and `500` status codes are supported.
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<int> StatusCode { get; set; } = null!;

        public ListenerFixedResponseArgs()
        {
        }
        public static new ListenerFixedResponseArgs Empty => new ListenerFixedResponseArgs();
    }
}
