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
    /// Custom request handling.
    /// </summary>
    [OutputType]
    public sealed class WebAclCustomRequestHandling
    {
        /// <summary>
        /// Collection of HTTP headers.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebAclCustomHttpHeader> InsertHeaders;

        [OutputConstructor]
        private WebAclCustomRequestHandling(ImmutableArray<Outputs.WebAclCustomHttpHeader> insertHeaders)
        {
            InsertHeaders = insertHeaders;
        }
    }
}