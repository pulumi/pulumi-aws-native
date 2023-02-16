// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Includes headers of a web request.
    /// </summary>
    [OutputType]
    public sealed class WebACLHeaders
    {
        public readonly Outputs.WebACLHeaderMatchPattern MatchPattern;
        public readonly Pulumi.AwsNative.WAFv2.WebACLMapMatchScope MatchScope;
        public readonly Pulumi.AwsNative.WAFv2.WebACLOversizeHandling OversizeHandling;

        [OutputConstructor]
        private WebACLHeaders(
            Outputs.WebACLHeaderMatchPattern matchPattern,

            Pulumi.AwsNative.WAFv2.WebACLMapMatchScope matchScope,

            Pulumi.AwsNative.WAFv2.WebACLOversizeHandling oversizeHandling)
        {
            MatchPattern = matchPattern;
            MatchScope = matchScope;
            OversizeHandling = oversizeHandling;
        }
    }
}