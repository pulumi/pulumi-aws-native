// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    /// <summary>
    /// An HTTP response header name and its value. CloudFront includes this header in HTTP responses that it sends for requests that match a cache behavior that's associated with this response headers policy.
    /// </summary>
    [OutputType]
    public sealed class ResponseHeadersPolicyCustomHeader
    {
        /// <summary>
        /// The HTTP response header name.
        /// </summary>
        public readonly string Header;
        /// <summary>
        /// A Boolean that determines whether CloudFront overrides a response header with the same name received from the origin with the header specified here.
        /// </summary>
        public readonly bool Override;
        /// <summary>
        /// The value for the HTTP response header.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ResponseHeadersPolicyCustomHeader(
            string header,

            bool @override,

            string value)
        {
            Header = header;
            Override = @override;
            Value = value;
        }
    }
}
