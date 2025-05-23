// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Custom response.
    /// </summary>
    [OutputType]
    public sealed class WebAclCustomResponse
    {
        /// <summary>
        /// Custom response body key.
        /// </summary>
        public readonly string? CustomResponseBodyKey;
        /// <summary>
        /// The HTTP status code to return to the client.
        /// 
        /// For a list of status codes that you can use in your custom responses, see [Supported status codes for custom response](https://docs.aws.amazon.com/waf/latest/developerguide/customizing-the-response-status-codes.html) in the *AWS WAF Developer Guide* .
        /// </summary>
        public readonly int ResponseCode;
        /// <summary>
        /// Collection of HTTP headers.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebAclCustomHttpHeader> ResponseHeaders;

        [OutputConstructor]
        private WebAclCustomResponse(
            string? customResponseBodyKey,

            int responseCode,

            ImmutableArray<Outputs.WebAclCustomHttpHeader> responseHeaders)
        {
            CustomResponseBodyKey = customResponseBodyKey;
            ResponseCode = responseCode;
            ResponseHeaders = responseHeaders;
        }
    }
}
