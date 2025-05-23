// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IvsChat.Outputs
{

    /// <summary>
    /// Configuration information for optional review of messages.
    /// </summary>
    [OutputType]
    public sealed class RoomMessageReviewHandler
    {
        /// <summary>
        /// Specifies the fallback behavior if the handler does not return a valid response, encounters an error, or times out.
        /// </summary>
        public readonly Pulumi.AwsNative.IvsChat.RoomMessageReviewHandlerFallbackResult? FallbackResult;
        /// <summary>
        /// Identifier of the message review handler.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private RoomMessageReviewHandler(
            Pulumi.AwsNative.IvsChat.RoomMessageReviewHandlerFallbackResult? fallbackResult,

            string? uri)
        {
            FallbackResult = fallbackResult;
            Uri = uri;
        }
    }
}
