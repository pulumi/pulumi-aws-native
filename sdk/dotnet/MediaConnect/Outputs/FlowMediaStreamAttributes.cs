// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// Attributes that are related to the media stream.
    /// </summary>
    [OutputType]
    public sealed class FlowMediaStreamAttributes
    {
        /// <summary>
        /// A set of parameters that define the media stream.
        /// </summary>
        public readonly Outputs.FlowFmtp? Fmtp;
        /// <summary>
        /// The audio language, in a format that is recognized by the receiver.
        /// </summary>
        public readonly string? Lang;

        [OutputConstructor]
        private FlowMediaStreamAttributes(
            Outputs.FlowFmtp? fmtp,

            string? lang)
        {
            Fmtp = fmtp;
            Lang = lang;
        }
    }
}