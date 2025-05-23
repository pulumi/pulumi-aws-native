// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// An endpoint for ingesting source content for a Channel.
    /// </summary>
    [OutputType]
    public sealed class ChannelIngestEndpoint
    {
        /// <summary>
        /// The system generated unique identifier for the IngestEndpoint
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The system generated password for ingest authentication.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The ingest URL to which the source stream should be sent.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// The system generated username for ingest authentication.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ChannelIngestEndpoint(
            string id,

            string password,

            string url,

            string username)
        {
            Id = id;
            Password = password;
            Url = url;
            Username = username;
        }
    }
}
