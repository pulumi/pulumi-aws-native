// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// The body of a web request. This immediately follows the request headers.
    /// </summary>
    [OutputType]
    public sealed class WebAclBody
    {
        public readonly Pulumi.AwsNative.WaFv2.WebAclOversizeHandling? OversizeHandling;

        [OutputConstructor]
        private WebAclBody(Pulumi.AwsNative.WaFv2.WebAclOversizeHandling? oversizeHandling)
        {
            OversizeHandling = oversizeHandling;
        }
    }
}